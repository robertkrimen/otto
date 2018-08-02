#!/bin/sh
git add .
echo "git commit:"
read commit
git commit -m "$commit"
git push -u origin master