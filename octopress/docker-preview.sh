#!/usr/bin/env bash

docker run --rm -d -p 4000:4000 -v $PWD:/app -w /app octopress bundle exec rake preview
