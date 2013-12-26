---
layout: post
title: "Handy Web Path Concatenation Code"
date: 2008-12-03 16:33
comments: true
permalink: /blog/development/handy-web-path-concatenation-code/
categories: 
---
Lot of web developement code is spent concatenating path snippets and I always hated having to deal with slashes. I wanted a function that was consistent with its return type, and very loose in its parameter requirements. I ended up with the following:

<!-- more -->

*Test First*:

{% gist 1545719 WebPathSpecs.cs %} 

The [BDD](http://en.wikipedia.org/wiki/Behavior_Driven_Development) style of testing is inspired (and pretty much copied) from JP Boodhoo's Nothing but .Net class.

*Here's the code*:

{% gist 1545719 WebPath.cs %}

That was my first go at it. I have a feeling I might generalize the class to be a delimited string builder. Which would be nice because then the method won't be static. Anyways, just thought I would share something that I think is concise but will be used a lot.
