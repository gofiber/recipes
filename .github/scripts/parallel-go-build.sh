#!/usr/bin/env bash
# parallel-go-build.sh
# Recursively find go.mod files and run `go build ./...` in each module directory.
# Optional: run with `-i` to execute `go mod tidy` and `go mod vendor` per module before build.
# - truncates build-errors.log at start
# - appends failures immediately, annotated with source line + 1-line context
# - runs in parallel (jobs = CPU cores)
# - atomic appends (mkdir lock)
# - handles Ctrl+C (SIGINT) and SIGTERM: kills children and cleans up
# - colored terminal output (OK = green, FAIL = red, WARN = yellow)
set -euo pipefail
IFS=$'\n\t'

JOBS="$(getconf _NPROCESSORS_ONLN 2>/dev/null || echo 4)"
LOGFILE="build-errors.log"
INPLACE_SYNC=0

usage() {
  cat <<'EOF'
Usage: parallel-go-build.sh [-i] [-h]

Options:
  -i    Run `go mod tidy` and `go mod vendor` in each module before build.
  -h    Show this help.
EOF
}

while getopts ":ih" opt; do
  case "$opt" in
    i) INPLACE_SYNC=1 ;;
    h)
      usage
      exit 0
      ;;
    \?)
      printf 'Unknown option: -%s\n' "$OPTARG" >&2
      usage >&2
      exit 2
      ;;
  esac
done
shift $((OPTIND - 1))

# Colors (only if stdout is a tty)
if [ -t 1 ]; then
  GREEN=$'\e[32m'
  RED=$'\e[31m'
  YELLOW=$'\e[33m'
  RESET=$'\e[0m'
else
  GREEN=""
  RED=""
  YELLOW=""
  RESET=""
fi

# truncate central logfile at start
: > "$LOGFILE"

# find all module directories (skip common vendor trees)
mapfile -t MODULE_DIRS < <(
  find . \( -path "./.git" -o -path "./vendor" -o -path "./node_modules" \) -prune -o \
    -type f -name 'go.mod' -print0 |
  xargs -0 -n1 dirname |
  sort -u
)

if [ "${#MODULE_DIRS[@]}" -eq 0 ]; then
  printf '%s\n' "No go.mod files found. Nothing to do."
  exit 0
fi

# create absolute temp dir
TMPDIR="$(mktemp -d 2>/dev/null || mktemp -d /tmp/build-logs.XXXXXX)"
if [ -z "$TMPDIR" ] || [ ! -d "$TMPDIR" ]; then
  printf '%s\n' "Failed to create temporary directory" >&2
  exit 2
fi

# ensure cleanup on exit (will try multiple times)
cleanup_tmpdir() {
  local attempts=0
  while [ "$attempts" -lt 5 ]; do
    if rm -rf "$TMPDIR" 2>/dev/null; then
      break
    fi
    attempts=$((attempts+1))
    sleep 0.1
  done
  if [ -d "$TMPDIR" ]; then
    printf '%s\n' "WARNING: could not fully remove temporary dir: $TMPDIR"
  fi
}

# signal handling: kill children, wait, cleanup, exit
pids=()
on_interrupt() {
  printf '%b\n' "${YELLOW}Received interrupt. Killing background jobs...${RESET}"
  # kill all tracked background pids
  for pid in "${pids[@]:-}"; do
    if kill -0 "$pid" 2>/dev/null; then
      kill "$pid" 2>/dev/null || true
    fi
  done
  # give them a moment, then force
  sleep 0.1
  for pid in "${pids[@]:-}"; do
    if kill -0 "$pid" 2>/dev/null; then
      kill -9 "$pid" 2>/dev/null || true
    fi
  done
  cleanup_tmpdir
  printf '%b\n' "${YELLOW}Aborted by user.${RESET}"
  exit 130
}
trap on_interrupt INT TERM

# helper: sanitize a dir to a filename-safe token
sanitize_name() {
  local d="$1"
  d="${d#./}"
  d="${d//\//__}"
  d="${d// /_}"
  printf '%s' "${d//[^A-Za-z0-9._-]/_}"
}

# annotate a module's temp log and append to central logfile atomically
annotate_and_append() {
  local src_log="$1"   # absolute path to per-module temp log
  local module_dir="$2"
  local lockdir="$TMPDIR/.lock"

  # create annotated temp file
  local annotated
  annotated="$(mktemp "$TMPDIR/annotated.XXXXXX")" || {
    # fallback: append raw log
    until mkdir "$lockdir" 2>/dev/null; do sleep 0.01; done
    printf '==== %s ====\n' "$module_dir" >> "$LOGFILE"
    cat "$src_log" >> "$LOGFILE"
    rm -f "$src_log" || true
    rmdir "$lockdir" 2>/dev/null || true
    return
  }

  # process lines: if matches file.go:LINE[:COL] annotate with source snippet
  while IFS= read -r line || [ -n "$line" ]; do
    # regex: capture path ending with .go, line number, optional column, rest
    if [[ $line =~ ^([^:]+\.go):([0-9]+):?([0-9]*)[:[:space:]]*(.*)$ ]]; then
      local fp="${BASH_REMATCH[1]}"
      local ln="${BASH_REMATCH[2]}"
      local col="${BASH_REMATCH[3]}"
      local rest="${BASH_REMATCH[4]}"
      local candidate=""

      # try to resolve file path: as-is, relative to module_dir, or relative to repo root
      if [ -f "$fp" ]; then
        candidate="$fp"
      elif [ -f "$module_dir/$fp" ]; then
        candidate="$module_dir/$fp"
      elif [ -f "./$fp" ]; then
        candidate="./$fp"
      fi

      if [ -n "$candidate" ]; then
        # pick context: line-1 .. line+1
        local start=$(( ln > 1 ? ln - 1 : 1 ))
        local end=$(( ln + 1 ))
        printf '%s\n' "---- source: $candidate:$ln ----" >> "$annotated"
        # print numbered context lines
        awk -v s="$start" -v e="$end" 'NR>=s && NR<=e { printf("%6d  %s\n", NR, $0) }' "$candidate" >> "$annotated"
        printf '%s\n\n' "Error: $line" >> "$annotated"
      else
        printf '%s\n' "---- (source not found) $line ----" >> "$annotated"
      fi
    else
      # plain copy line
      printf '%s\n' "$line" >> "$annotated"
    fi
  done < "$src_log"

  # append annotated file to central logfile under lock
  until mkdir "$lockdir" 2>/dev/null; do
    sleep 0.01
  done
  {
    printf '==== %s ====\n' "$module_dir"
    cat "$annotated"
    printf '\n\n'
  } >> "$LOGFILE"
  rm -f "$annotated" "$src_log" || true
  rmdir "$lockdir" 2>/dev/null || true
}

# run build for one module
run_build() {
  local module_dir="$1"
  local safe
  safe="$(sanitize_name "$module_dir")"
  local mod_log="$TMPDIR/$safe.log"
  mkdir -p "$(dirname "$mod_log")"

  # run inside subshell to avoid changing caller's cwd
  if ( cd "$module_dir" 2>/dev/null && go build ./... >/dev/null 2>"$mod_log" ); then
    printf '%b\n' "${GREEN}OK:   ${RESET}$module_dir"
    rm -f "$mod_log" >/dev/null 2>&1 || true
    return 0
  else
    # ensure we captured something
    if [ ! -s "$mod_log" ]; then
      ( cd "$module_dir" 2>/dev/null && go build ./... >"$mod_log" 2>&1 ) || true
    fi
    printf '%b\n' "${RED}FAIL: ${RESET}$module_dir (appending to $LOGFILE)"
    annotate_and_append "$mod_log" "$module_dir"
    return 1
  fi
}

# main launcher: spawn jobs, throttle to JOBS, wait properly
fail_count=0

if [ "$INPLACE_SYNC" -eq 1 ]; then
  printf '%b\n' "${YELLOW}Running module sync (go mod tidy && go mod vendor) before build...${RESET}"
  for md in "${MODULE_DIRS[@]}"; do
    safe="$(sanitize_name "$md")"
    init_log="$TMPDIR/${safe}.init.log"
    if ( cd "$md" 2>/dev/null && go mod tidy && go mod vendor ) >"$init_log" 2>&1; then
      printf '%b\n' "${GREEN}SYNC: ${RESET}$md"
      rm -f "$init_log" >/dev/null 2>&1 || true
    else
      printf '%b\n' "${RED}SYNC FAIL: ${RESET}$md (appending to $LOGFILE)"
      annotate_and_append "$init_log" "$md"
      fail_count=$((fail_count+1))
    fi
  done
fi

for md in "${MODULE_DIRS[@]}"; do
  run_build "$md" &
  pids+=( "$!" )

  # if we reached jobs, wait for the oldest launched
  if [ "${#pids[@]}" -ge "$JOBS" ]; then
    if wait "${pids[0]}"; then
      :  # success
    else
      fail_count=$((fail_count+1))
    fi
    # drop first pid
    pids=( "${pids[@]:1}" )
  fi
done

# wait remaining background jobs
for pid in "${pids[@]:-}"; do
  if wait "$pid"; then :; else fail_count=$((fail_count+1)); fi
done

# final cleanup and status
cleanup_tmpdir

if [ "$fail_count" -gt 0 ]; then
  printf '\n%b\n' "${RED}Done. $fail_count module(s) failed. See $LOGFILE for details (annotated snippets included).${RESET}"
  exit 1
else
  printf '\n%b\n' "${GREEN}Done. All modules built successfully.${RESET}"
  # remove logfile if empty
  if [ -f "$LOGFILE" ] && [ ! -s "$LOGFILE" ]; then
    rm -f "$LOGFILE"
  fi
  exit 0
fi
