#!/usr/bin/env bash

set -e

# Script assumes the stage is created via
#   git clone -b master --depth 1 git@github.com:scottmuc/scottmuc.github.com.git master

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
