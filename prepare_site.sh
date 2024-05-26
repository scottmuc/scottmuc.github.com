#!/usr/bin/env bash

set -e

# the master branch is currently the method of deployment to github pages
if [ ! -d ./live ]; then
  git clone -b live --depth 1 git@github.com:scottmuc/scottmuc.github.com.git live
fi

# prepate hugo
pushd hugo
./build.sh # generates hugo/public
popd

rsync -av hugo/public/ live
# verify changes and commit and push!
