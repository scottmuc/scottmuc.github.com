---
layout: post
title: "Manager Suffixes Are a Code Smell"
date: 2008-06-30 16:31
comments: true
permalink: /blog/development/manager-suffixes-are-a-code-smell/
categories: 
---

After attending [Nothing But .Net Bootcamp](/blog/development/i-survived-jp-s-nothing-but-dot-net-boot-camp/) I noticed none of the classes that were written in the course had a *Manager* suffix. I've sort had a eureka moment and have a good idea on why this is.

<!-- more -->

JP heavily emphasized the [Single Responsibility Principle](http://en.wikipedia.org/wiki/Single_responsibility_principle) throughout the course. If a class is only has a single responsibility it will be pretty difficult to attach a Manager suffix to the class name.

So often have have I grouped together methods into some sort of "Manager" class and have always felt guilty about it. Having a Manager class almost corresponds to the [Broken Windows story](http://www.artima.com/intv/fixit.html). You start thinking that since you have the Manager class it won't hurt to add another method to it, which is simply giving it more responsibilities.

I think a good real world exercise for me to do is to [refactor](http://en.wikipedia.org/wiki/Refactoring) one of my Manager classes into a set of classes that perform the same operations in a much more cohesive way.

