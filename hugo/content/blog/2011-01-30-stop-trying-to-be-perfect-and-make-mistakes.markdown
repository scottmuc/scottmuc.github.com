---
layout: post
title: "Stop Trying to be Perfect and Make Mistakes!"
date: 2011-01-30 21:30
permalink: /blog/development/stop-trying-to-be-perfect-and-make-mistakes/
comments: true
categories: 
---

Think about how you get good at anything. What do you remember spending most of your time on? Most likely you spent the majority of your time messing things up and getting it wrong. Why are kids so quick at picking up new skills, especially dangerous ones? It's because kids are fearless experimenters. I'm currently learning how to skate and ski and my ability to learn is hindered by my concern of hurting myself; potentially preventing me from working and earning a living. I want to try and correlate the idea of failure as a healthy activity in the road to expertise in the world of operations in the software world.

<!-- more -->

Operations people have two numbers they care about. The number 9 and how many of them they have (99.999% uptime is commonly known as five nines). This is a great metric for a system that doesn't change, but in a world where new requirements are constantly conjured up and need to be live yesterday, the number of 9's you have becomes harder to keep stable. A new operational metric needs to be created to handle this new kind of ever changing system.

Historically operations people are stereotyped as grumpy curmudgeons who will only let you deploy something in the 1 minute downtime window that they provide for you on the 3rd full moon after each solstice on a Sunday night at 2:29am, and you better get it done before the backups fire up or else you're screwed! Developers on the other hand are eager to get their newly written code live so that people can see the fruits of their labour. It's pretty obvious that these mentalities are in complete opposition.

This clash of values can be lessened if operators adopted a new bragging metric. The **Mean Time to Recovery (MTTR)**. An operations person that can fix a problem fast is much stronger than someone without that kind of debugging ability. To get ops to fix problems that means they need to have strong knowledge of the applications they are deploying and strong relationships with the teams that developed them. The hard part about this metric is that it needs to be exercised. You can't measure your MTTR unless things break. 

Another great metric was created by the Lean way of thinking where minimization of completed features (an inventory) waiting to go live is important because that backlog is considered wasteful. This is still a great system wide view of shortening the time from concept to market, but I want to focus a little on the mind of an operations person. 

One of the problems with software development is a general sense of hubris when it comes to the stability and functionality of an application. Stats like the number of unit tests and code coverage are often leaned on with the implicit understanding that it means the code is working. That's when concepts like [Mutation Testing](http://en.wikipedia.org/wiki/Mutation_testing) are great for truly determining the value and confidence that the test suite brings. Tools like [Jester perform](http://jester.sourceforge.net/) a little bytecode magic behind the scenes to determine how brittle your tests are.

Operations can suffer from the same kind illogical confidence. How are the 9's measured? I've often seen uptime measured by how long heartbeat.html is still up, but how confident would that make you believe that the entire system can still process an order? A server that's up and running serving up broken software should not count as being "up". This is where a common agreement of the term "done" is important in bridging the gap between software development and operations. 

It's too bad that ops can't perform activities like TDD where the developer intentionally fails a test in order to drive out the code to implement a feature.

At a few conferences I've met a few people who deploy to production rapidly in massive systems ([IMVU](http://engineering.imvu.com/2010/04/09/imvus-approach-to-integrating-quality-assurance-with-continuous-deployment/), [Facebook](http://glinden.blogspot.com/2009/11/continuous-deployment-at-facebook.html)). They don't know 100% the code they are deploying won't break in production. By deploying to partitions they can use system wide health metrics to determine whether or not things are working. Are orders still being processed in the partition that the new code was deployed to?

In my previous work experience I deployed code to production all the time (and broke things). Every time I broke something was a learning experience, but I didn't stress about it because I knew how to fix things reasonably fast. I didn't realize at the time how valuable that was. Periodic failure was welcome because it re-assured us that we knew how the system worked. Now of course mission critical applications have much different sensitivities to failure. I'm not talking about software that launches the space shuttle here.

In conclusion I would like to see how we can allow operations to fail without such a negative impact. If a deployment fails because someone forgot to chmod a file with +x, that's a great lesson learned and hopefully someone will update the deployment script. After all, we're all human and we make mistakes. By recording the operational MTTR we can grow a team that updates production frequently and fixes errors fast, rather than a team performs massive batch releases to production and slowly figures out the errors. So let's figure out how we can help our ops people fear downtime a bit less and empower them with the ability fix the software we build. On the flip side lets write software that is easily deployable and doesn't require pages of instructions to make live so that operations can feel more comfortable tweaking and redeploying with confidence.

I don't propose any solutions to this mindset but would love to hear some feedback. Would measuring MTTR be important in an operations group?

