#!/usr/bin/env bash
BRANCH=$1
FILE=glide.yaml

# check if glide.yam. exists
if [ ! -f $FILE ]; then
  echo "$FILE not found."
  exit
fi

# checkout the specified branch
if ! git checkout $BRANCH; then
    echo "branch '$BRANCH' not found, exiting."
    exit 1
fi

# pull latest
if ! git pull; then
    echo "failed to pull branch '$BRANCH', exiting."
    exit 1
fi

# check if we need to make any changes
if ! grep -Fxq "$BRANCH" $FILE; then
  echo "'$BRANCH' not found in $FILE. No work to do, exiting."
  exit
fi

# replace the specified branch with master
echo "changing all glide.yaml repositories from $1 to master"
sed -i "s,$BRANCH,master,g" glide.yaml

# commit
git add .
git commit -m "replaced branch dependency from '$1' to 'master' in the glide.yaml file"