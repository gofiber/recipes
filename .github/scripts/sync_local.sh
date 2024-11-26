#!/usr/bin/env bash
set -e

if [ "$#" -eq 1 ]; then
    REPO_DIR="$1"
else
    REPO_DIR="recipes"  # default value
fi

if [[ ! "$REPO_DIR" =~ ^[a-zA-Z0-9_-]+$ ]]; then
    echo "Error: REPO_DIR must contain only alphanumeric characters, underscores, and hyphens" >&2
    exit 1
fi

# determine root repo directory
ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd -P)"

# remove all files in the docs directory
rm -rf $ROOT/../fiberDocs/docs/${REPO_DIR}/*

# Find and copy relevant files
find . \
  -type f \
  \( -iname '*.md' -o -iname '*.png' -o -iname '*.jpg' -o -iname '*.jpeg' -o -iname '*.gif' -o -iname '*.bmp' -o -iname '*.svg' -o -iname '*.webp' \) \
  -not -path "./fiber-docs/*" \
  -not -path "*/vendor/*" \
  -not -path "*/.github/*" \
  -not -path "*/.*" |
while IFS= read -r f; do
  echo "Copying $f"
  mkdir -p $ROOT/../fiberDocs/docs/${REPO_DIR}/$(dirname "$f")
  cp "$f" $ROOT/../fiberDocs/docs/${REPO_DIR}/$f
done
