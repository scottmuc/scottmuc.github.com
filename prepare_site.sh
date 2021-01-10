#!/usr/bin/env bash

set -e

if [ ! -d ./master ]; then
  git clone -b master --depth 1 git@github.com:scottmuc/scottmuc.github.com.git master
fi

# prepate hugo
pushd hugo
hugo # generates hugo/public
popd

# prepare octopress
pushd octopress
./docker-bundle-exec.sh rake generate # generates octopress/public
popd

# Bring them together!

rsync -av octopress/public/ master/
rsync -av hugo/public/ master/hugo

# verify changes and commit and push!
