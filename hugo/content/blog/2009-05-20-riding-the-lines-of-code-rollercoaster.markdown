---
layout: post
title: "Riding The Lines-Of-Code Rollercoaster"
date: 2009-05-20 08:52
comments: true
permalink: /blog/development/riding-the-lines-of-code-rollercoaster/
categories: 
---

Lines of code has always been a poor guide in measuring developer productivity. I'm currently reading [Scott Rosenberg's](http://en.wikipedia.org/wiki/Scott_Rosenberg_\(journalist\)) [Dreaming In Code](http://www.dreamingincode.com/) and came across and interesting quote.

>   "Atkinson [ed: Bill] had just completed rewriting a portion of the Quickdraw code, making it more efficient and faster. The new version was 2000 lines of code shorter than the old one. What to report? He wrote the number -2000" -- **Rosenberg**

I couldn't help but smile when I read that. I have felt that some of my best coding successes are when I eliminate lots of code. This probably pushes Test Driven Development (TDD) as a great practice even further. Lately I've been a bit stricter by doing TDD. This means that I'm writing a lot more code, but as I get my tests passing I start refactoring and remove much of the code that's written. For every couple hundred lines of code I write, I eventually remove 50-80% of it as I glean knowledge from all the code I previously wrote.

It appears that LOC Rollercoaster development is an artifact of TDD and it's a great ride!

