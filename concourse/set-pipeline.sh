#!/usr/bin/env bash

set -euo pipefail

fly -t concourse.scottmuc.com \
  set-pipeline \
  --pipeline www.scottmuc.com \
  --config ./cd-pipeline.yml
