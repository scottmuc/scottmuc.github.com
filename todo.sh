#!/usr/bin/env bash

echo Copy podcasts_opml.xml
echo Migrate static pages, like the about page
echo Copy twinery boat tour
echo Sort out RSS feed
echo Create post page-bundle archetype
echo Fix Twitter inline references
echo The following images should be sorted:
grep -R "test image" hugo/content/blog
