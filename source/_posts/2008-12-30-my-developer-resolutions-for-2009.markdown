---
layout: post
title: "My Developer Resolutions For 2009"
date: 2008-12-30 00:11
comments: true
permalink: /blog/development/my-developer-resolutions-for-2009/
categories: 
---

I've never made any New Years resolutions before. This year I'm going to break that trend simply to create a checklist of tasks to keep me on track. I easily get distracted, but having to check things off a list is a sure way for me to stay focused.

Not wanting to over do it, I made goals that are easily attainable. Hopefully upon completion I'll get more ambitious the next time around.

<!-- more -->

### Development Goals:

Lately I've been spending a lot of time thinking about "the build". Strangely enough, I get a lot of satisfaction out of a useful build environment. When I can checkout some code and be able to build and run a local instance of an application with a single command line execution, I feel great sense of value from the work put into my build scripts. These thoughts have heavily influenced my technical aspirations for 2009.

#### Learn A New Build System

I Love [NAnt](http://nant.sourceforge.net/), but I'm really starting to hate XML. It's not that I want to ditch NAnt because I know it well, and it does the job. It's just that I think it's could practice to see how the same thing can be done in other ways.

* <span style="text-decoration: line-through;">Boobs</span> [Bake](http://code.google.com/p/boo-build-system/) - Written by the great [Ayende](http://ayende.com/) and uses the Boo language as its syntax.
* [PSake](https://github.com/JamesKovacs/psake) - Written by [James Kovacs](http://jameskovacs.com/) and is written in PowerShell. I've yet to get into PowerShell and a build system seems like a natural place to start learning it. The one downside is that I don't believe it'll work in a Linux environment (or any environment without PowerShell for that matter).
* [Rake](http://rake.rubyforge.org/) - Ruby is another language I would like to learn and build scripts is yet another applicable task for it. Rake has a fairly large user base and it's cross platform so it's a strong contender.

I don't really need to learn any of these considering I've accomplished quite a bit with NAnt (You can compile and deploy the [CBC Radio 3](http://radio3.cbc.ca/) site in one command). Still, if I can get away from the XML build files tasks would be more concise and easier to extend.

#### Implement a Continuous Integration Setup

I've mentioned how much I love the one line build execution. What I hate is when that one line build execution breaks! I don't know a whole lot about CI, but just the ability to have an automated build health check is reason enough for me to implement CI. Automated testing, and deployment are just bonuses. The only one I have some knowledge about is [CruiseControl.Net](http://www.cruisecontrolnet.org/) written by the folks at ThoughtWorks. If it's good enough for [Martin Fowler's](http://martinfowler.com/) team, it's good enough for me.

#### Change Version Control System

[Subversion](http://subversion.apache.org/) has changed my life. I wish I had used it (or any versioning system) when in University. In fact, that's the first thing I would recommend any new Comp-Sci student to learn is some sort of version system. It'll save you so much hassle in the long run, and you'll be versioning ninja once you get out into the workforce.

I've been reading about distributed version systems and like the idea of decoupling my repository from a server. I recently found a [great article](http://www.smashingmagazine.com/2008/09/18/the-top-7-open-source-version-control-systems/) comparing other systems and there are two other systems that I want to try:

* [Git](http://git-scm.com/) - Written by Linus Torvalds, this distributed versioning system is quite powerful and has lots of mainstream support. I've seen discussions where some say Git is a bit too difficult, but once understood is extremely powerful.
* [Bazaar](http://bazaar-vcs.org/) - A distributed version system with a focus on usability. The one interesting note on their website is that their code base has over 10,000 unit tests! Whether or not I use it, I'm sure the code base would be an interesting read in how to write python applications.

### Life Organization:

#### Bookmarks

I still have great pain in maintaining my browser bookmarks. I use Delicious, and love the FireFox integration, but now that doesn't matter now that I'm using Chrome now. I'm still brainstorming ideas and may post later about how I'm going to move forward.

#### Calendar

One of my favourite features of my phone is how it syncs with my work Outlook calendar. The problem is that I'm not at work all the time (wait, that's a problem?). I want an ubiquitous calendar. When I refer to a calendar, I mean the one calendar to rule them all and in darkness... oops, I mean... When I say I'm doing something on a particular date and time, I want my announcement cascaded to all calendar applications I use. Notice I didn't say how I announce an appointment. I don't want to limit myself to adding events in Outlook, my phone, or my Google apps calendar. So I'm on the lookout a calendaring strategy that allows me to add events to my "abstract" calendar. Good luck with that, eh?

