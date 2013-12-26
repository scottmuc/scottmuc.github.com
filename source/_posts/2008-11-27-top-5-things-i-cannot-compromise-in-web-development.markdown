---
layout: post
title: "Top 5 Things I Cannot Compromise In Web Development"
date: 2008-11-27 16:33
comments: true
permalink: /blog/development/top-5-things-i-can-t-compromise-on-in-web-development/
categories: 
---
Over the past few months I've realized that I'm a stubborn mule when it comes to certain aspects of web development. After some internal reflection I've come up with my top 5 things I strongly feel that I won't compromise on in the web development process. These items are just for the presentation layer, I'm loaded with issues deeper into the development stack.

<!-- more -->

*1 - URLs must be as beautiful as possible*

I really don't need to say much about this because [Mike Schinkel](http://mikeschinkel.com/blog/welldesignedurlsarebeautiful/) has already said it all. We've been contemplating the use of [SWFAddress](http://www.asual.com/swfaddress/) as a method of deep linking because the new [Radio 3 website](http://radio3.cbc.ca/) is wrapped in a FrameSet therefore you don't see the URL change in the Address bar of a web browser. I know how well deep linking solves this problem, but I still have a hard time looking at a URL with the hash (#). I personally would never use it for a personal website because that 1 character of noise is so distracting to me.

Secondly, I believe urls should be guessable and hackable. [Stackoverlow](http://stackoverflow.com/) has some clean urls but they aren't perfect. They are somewhat hackable, but what's with have the database key in the url? this makes urls impossible to guess. I really wish my profile URL was: http://stackoverflow.com/users/scott-muc instead.

*2 - Web pages must be constructed with semantic markup with no display decoration*

Seperating the design from the markup can be difficult but when done correctly it is a thing of beauty. The introduction of [micro-formats](http://microformats.org/) has made me see a huge potential in the markup as data and not about display. I'm extremely impressed by [Corkd](http://content.corkd.com/) because you can actually use their website markup as a database for their wine reviews because of their usage of the [hreview](http://microformats.org/wiki/hreview) microformat.

Along with that I believe the markup should contain zero inline CSS and javascript. Easier said than done, but I simply believe that CSS should belong in the stylesheet files (or in the page header in the worst case). Inline boot strapping javascript I find ok, but using inline javascript to handle click events is really messy.

*3 - No browser specific CSS and/or hacks*

The first time I learned about the [Star Hack](http://css-discuss.incutio.com/wiki/Star_Html_Hack) I was in heaven. I finally felt I could make IE 6 work. Unfortunately this hack came with the cost of making CSS maintenance difficult. It takes more time, but having a design that requires the minimal amount of IE 6 only hacks the better. Or better yet, don't do any IE 6 hacks and just use [progressive enhancement](http://www.alistapart.com/articles/progressiveenhancementwithcss) and don't fret about IE 6 looking bad in some situations. It's simply not worth the time because we all have better things to do than support a nearly 10 year old web browser.

*4 - No flash or objects on the page*

Flash creates a load delay that's an order of magnitude slower than a regular text only page load. I do like flash and a lot of the features it provides, but I think the best usage of flash is to only load it when the user wishes to use a feature that requires flash. The embedable music players and the embedded youtube videos are the worst offenders in my humble opinion. These can easily be handled by loading of the flash on click events.

*5 - Don't let your server side application framework take over your markup*

The majority of my web experience is with Asp.Net and I really don't like how some of the controls create horrible markup. I'm a big fan of the Repeater and ListView controls simply because I have 100% control over the markup they create. I can't say I've ever used GridView in a production web site.

