---
layout: post
title: "Intentional Programming aka Calling Your Shots"
date: 2012-02-21 00:00
comments: true
categories: 
---
I just read this [blog post](http://playswithfire.com/blog/2012/02/19/you-are-not-ruthless-enough/) and it's inspired me to finally post
about something that's been in the back of my mind for while now.

Back in the day I used to play 8-ball on a fairly regular basis. I would play with a regular group of friends who would all work on
improving our game. One of the players was nearly a professional and provided me one training tip that helps me everyday. He would ask me
to declare every variable of my shot. Not just what ball and pocket I'm targeting, but what trajectory the cue ball will take after contact
and how it will carom off of every rail, and to predict where the cue ball will end up. By doing this it revealed so much more of the game
to me. As a novice, your primary concern is potting balls. But to be good, you need be in control of everything on the table.

<!-- more -->

I like to apply this to software development. While watching Kent Beck's [TDD screencast](http://pragprog.com/screencasts/v-kbtdd/test-driven-development)
I heard him label this kind of thing as calling yours shots. I absolutely love that term for it and can't believe I never thought of naming it that 
before! Ultimately, calling your shots is like this:

> **When changing part of the system, predict how the the behaviour of the system has changed.**

So if it's code under test, predict the test result. If you predict it's going to fail, declare exactly why it's going to fail. When making
code change, declare your intent of what you want to happen as a result of that change. Before I went into software development I worked as
a system administrator and this routine became valuable. Saying things like "Because I have update this registry entry, a file named blah.log
should be in C:\foo\logs". If that doesn't happen, then you have an opportunity to learn more on how the system works.

I had an experience recently where I wasn't able to call my shots. I was working on some CSS positioning and nearly every change I made to
that poor little SASS file was a shot in the dark. This was a clear indicator that I didn't know what I was doing and that some deeper
understanding was necessary to be at all effective.

This is the paragraph from the [blog post](http://playswithfire.com/blog/2012/02/19/you-are-not-ruthless-enough/) that really hit home
for me.

> Every time you throw in something that seems to work, find out why it works. If you don’t understand why it works then the code is 
> just magic and you’re letting yourself off the hook. That code will become enshrined in your version control system for later 
> generations (read: you, three months from now) to discover, wonder at, puzzle over, and leave it because no one understands what it does.

When it comes to software and computers, I have little patience for magic. This probably slows me down initially but as the system grows
I feel confident that I understand everything quite well.

On the other side of the coin, playing around and exploring is also necessary. This is the realm of the spike. I had an experience with a colleague
recently where I kept on insisting that they called their shot. They announce that they couldn't because of a lack of context and understanding
and that they were simply exploring. It was a good reminder that exploration is extremely important and that we don't spike things out
often enough. Our spikes end up being our production code which is part of what the linked to article is trying to get at.

