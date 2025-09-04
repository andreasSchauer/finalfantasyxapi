#!/bin/bash
set -e

if [ $# -eq 0 ]; then
    echo "Usage: ./commit.sh \"commit message\""
    exit 1
fi


MESSAGE="$1"

git add .
git commit -m "$MESSAGE"
git push origin main

echo "Successfully committed and pushed: $MESSAGE"