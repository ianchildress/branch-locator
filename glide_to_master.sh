#!/usr/bin/env bash
BRANCH=$1
git checkout $1 && git pull
echo "Changing all glide.yaml repositories from $1 to master."
sed -i "s,$1,master,g" glide.yaml
git add .
git commit -m "Changing all glide.yaml repositories from $1 to master."