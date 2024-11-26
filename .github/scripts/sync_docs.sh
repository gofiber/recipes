#!/usr/bin/env bash
set -e

# Some env variables
BRANCH="main"
REPO_URL="github.com/gofiber/docs.git"
AUTHOR_EMAIL="github-actions[bot]@users.noreply.github.com"
AUTHOR_USERNAME="github-actions[bot]"
REPO_DIR="recipes"
COMMIT_URL="https://github.com/gofiber/recipes"

# Set commit author
git config --global user.email "${AUTHOR_EMAIL}"
git config --global user.name "${AUTHOR_USERNAME}"

git clone https://${TOKEN}@${REPO_URL} fiber-docs

latest_commit=$(git rev-parse --short HEAD)

# remove all files in the docs directory
rm -rf $ROOT/../fiberDocs/docs/${REPO_DIR}/*

for f in $(find -E . -type f -iregex '.*\.(md|png|jpe?g|gif|bmp|svg|webp)$' -not -path "./(fiberDocs)/*" -not -path "*/vendor/*" -not -path "*/.github/*" -not -path "*/.*"); do
  log_output=$(git log --oneline "${BRANCH}" HEAD~1..HEAD --name-status -- "${f}")

    if [[ $log_output != "" || ! -f "fiber-docs/docs/${REPO_DIR}/$f" ]]; then
      mkdir -p fiber-docs/docs/${REPO_DIR}/$(dirname $f)
      cp "${f}" fiber-docs/docs/${REPO_DIR}/$f
  fi
done


# Push changes
cd fiber-docs/ || true
git add .

git commit -m "Add docs from ${COMMIT_URL}/commit/${latest_commit}"

MAX_RETRIES=5
DELAY=5
retry=0

while ((retry < MAX_RETRIES))
do
    git push https://${TOKEN}@${REPO_URL} && break
    retry=$((retry + 1))
    git pull --rebase
    sleep $DELAY
done

if ((retry == MAX_RETRIES))
then
    echo "Failed to push after $MAX_RETRIES attempts. Exiting with 1."
    exit 1
fi
