#!/usr/bin/env bash

rm -rf public
if command -v nix 2>/dev/null; then
  nix --extra-experimental-features "nix-command flakes" run nixpkgs#hugo -- --gc
else
  hugo --gc
fi
rm -rf public/1
