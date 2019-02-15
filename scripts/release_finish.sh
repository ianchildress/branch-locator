#!/usr/bin/env bash
BRANCH = $1
git checkout $1 && git pull
sed -i "s,release/24.1,master,g" glide.yaml
git add .
git commit -m "pointing glide.yaml to master"
git checkout master && git pull
git checkout develop && git pull
git checkout release/24.1
git flow release finish && git checkout master && git push && git checkout develop && git push
