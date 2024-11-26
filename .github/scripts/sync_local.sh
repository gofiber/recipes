#!/usr/bin/env bash
set -e

REPO_DIR="recipes"

# determine root repo directory
ROOT=$(cd "$(dirname "$0")/.." && dirname "$(pwd -P)")

# remove all files in the docs directory
rm -rf $ROOT/../fiberDocs/docs/${REPO_DIR}/*

for f in $(find -E . -type f -iregex '.*\.(md|png|jpe?g|gif|bmp|svg|webp)$' -not -path "./(fiberDocs)/*" -not -path "*/vendor/*" -not -path "*/.github/*" -not -path "*/.*"); do
  echo "Copying $f"
    mkdir -p $ROOT/../fiberDocs/docs/${REPO_DIR}/$(dirname $f)
    cp "${f}" $ROOT/../fiberDocs/docs/${REPO_DIR}/$f
done
