#!/usr/bin/env bash

set -e

# the master branch is currently the method of deployment to github pages
if [ ! -d ./master ]; then
  git clone -b master --depth 1 git@github.com:scottmuc/scottmuc.github.com.git master
fi

# prepate hugo
pushd hugo
./build.sh # generates hugo/public
popd

rsync -av hugo/public/ master
# verify changes and commit and push!
