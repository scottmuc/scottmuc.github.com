---
layout: post
title: "Devops Days 2012 in Summary"
date: 2012-06-29 17:08
comments: true
categories: 
---
[DevopsDays](http://devopsdays.org/) is a juxtaposition in the world of IT. Most
tech conferences focus on tools and processes, whereas Devops Days focuses on
culture and the delivery of value. This year's conference was no exception and
here is a summary of what I experienced:

## Day 1

#### The Audubon Society for Partial Failures

The first [presentation](http://cl.ly/2t1J0H1M3b2O092a1v1S) was done by [Cliff Moon][cliff-moon-twitter]. He argues
that there are many common patterns in the IT space and it would be valuable
to have a repository of knowledge that we can visit to read about the types of
errors that can happen. If we could build a taxonomy for the partial failures
that exist we may better understand our systems and explain when full-blown
outages occur.

He stresses sharing as the ultimate method for this. Not just sharing ideas and
hypotheses but also the data that lead to the hypothesis. This allows for peer
review and corrections to be made if analysis of the data could have been
misinterpreted.

What I really enjoyed was the spirit of Cliff's talk. The influence for the title
was a man who was obsessed with categorizing birds, John James Adubon. What
really hit home was the idea that there's always someone on your team that is
somewhat of a shaman when it comes to resolving production issues. It might
simply be because that person understands the conventions of the system at
work. Creating this taxonomy could help in revealing the shaman's "secrets" and
allow for everyone to weave a little magic.

#### Continuous Improvement

[Noah Sussman][noah-sussman-twitter] is a test architect at Etsy and he
[presented][continous-improvement-slides] on how Continuous Deployment and
Continuous Delivery change how quality assurance needs to be done.

I believe this to be a very important topic because this is going to be
a drastic change for most organisations. Large companies are enamoured with
startup culture and ideas, but aren't seeing what's required of them to change
to this model. A change in testing strategy is a significant barrier to entry
for most large companies. Even if the devs and ops teams are collaborating,
they can't deliver anything unless it gets passed the quality gates.

#### Vagrant 2.0

I got a chance to meet [Mitchell Hashimoto][mitchell-hashimoto-twitter],
creator of a tool that I've been enjoying quite a bit lately,
[Vagrant][vagrant-website]. There was an excellent discussion on how people
are using vagrant in their organisation. What was also exciting was hearing
what's coming out next:

* run on any VM software
* run any guest OS (that includes OS X)
* distributed as a package, not as a gem
* Builder (essentially merging veewee into vagrant)

During the session I also learned about the concept of baking an image vs
frying one (and everyone knows I love analogies). To bake an image is to use
an image as is, no configuration management is needed. This doesn't mean
configuration management isn't used to bake that image though which I thought
was an important detail. Frying an image is to start with a common base and
always reach an end state via configuration management. Frying is the approach
I favour because I like to execute my chef recipes as often as possible. I
also learned about different ways of baking images and still having the same
level of confidence in configuration management scripting.

## Day 2

#### Gauntlt and bridging the gap between security and everyone

[James Wickett][james-wickett-twitter], a security consultant, discussed how security is feeling left
out in this crazy world of Continuous Integration, Testing, and Delivery.
What is cool, is that he's doing something about it.

In typical security processes, an organisation has someone come in and
perform an assessment of the security of the applications being developed.
The artifact is usually a 100 page PDF describing everything that's wrong
with the system being developed. Instead of a document, he suggests delivering
a test suite that can be executed continuously. Hence [gauntlt][guantlt-website] is born; a tool
to execute security tests against a running application.

I chatted with James briefly and somehow managed to volunteer to join their
team. I am no security expert, but I believe there's a gap in security knowledge
in the software industry. Hopefully by integrating gauntlt into future projects
I can provide some feedback to make the tool more useful.

#### Continuous Delivery at Ancestry.com

I posted my thoughts on a similar presentation at [Chef Conf][chef-conf-post] and again John
delivered an excellent presentation on the values of culture and delivery focus
at their organisation.

Again I got a chance to chat with Dan Gilmer who implemented a lot of tech
around their pipeline and am looking forward to collaborating with him on
open sourcing his robust solution to dealing with Windows configuration
management using PowerShell and Chef. Hopefully we'll have some of that
knowledge in a project in github soon.

#### Automating Application Configuration Deployment

[Dan Nemec][dan-nemec-twitter] delivered a presentation that is very near and dear to my heart.
Application configuration management can be messy if left to people who do not
have visibility into all the systems that are in play. He brought a system that
required 400+ configuration files to a 2 file system. The 400+ files still
existed but by using templates, and simple configuration hierarchy in chef
he was able to make application configuration management much simpler and
easier to scale.

This hit home for me because I've implemented my own turnkey implementation
of this on my last few projects. Those that have worked with me know I refer
to these as environment yaml files. These files are a manifest of the topology
of an environment and all its configuration.

You can read more about Dan's thoughts on this on his [blog](http://blog.geeksgonemad.com/2012/05/automating-application-configuration.html).

#### Where Does SmallOps Fit In

This session was very relevant to me as someone who has been working on a small
team for the last few months. [Matthew Kocher][matthew-kocher-twitter]
highlights the simple solutions which is why I really enjoy hearing him discuss
his opinion.

During this session not much of a consensus was reached. I got the feeling that
most of the people in the room didn't quite know what to do with SmallOps. How
do you manage 20 servers and not complicate the infrastructure
management?

#### Devops as a role on projects

This topic, I love in principle, but do a horrible job in practice. There's
no such thing as a "DevOps Engineer", it's just a mindset. This is where
the cultural message is most important. DevOps is about transforming
organisations to remove silos. Jamie who brought up the open space likened
a software delivery team to a sports team (again, I love the analogies). With
sports you still have specialists, but they work has a cohesive unit. In
football, there's no reason a defense person can't score a goal. If you have
a team where offensive players begin blaming defense for their lack of success
you then have a disfunctional team.

#### What Am I Going To Do Now?

It really hit home to me after this Devops Days how I am doing a lot of things
wrong in the mission to bridge these gaps.

I'm going to:

* Measure all things.
* Use monitoring tools (I've been a cowboy ops person for too long)
* Keep my engagements short to force knowledge transfer early


[cliff-moon-twitter]: https://twitter.com/moonpolysoft
[noah-sussman-twitter]: https://twitter.com/noahsussman
[dan-nemec-twitter]: https://twitter.com/dcnem
[james-wickett-twitter]: https://twitter.com/wickett
[mitchell-hashimoto-twitter]: https://twitter.com/mitchellh
[matthew-kocher-twitter]: https://twitter.com/mkocher
[vagrant-website]: http://vagrantup.com/
[guantlt-website]: https://github.com/thegauntlet/gauntlt
[chef-conf-post]: /chefconf-2012-in-summary/
[continous-improvement-slides]: http://www.slideshare.net/noahsussman/continuous-improvement-devops-day-mountain-view-2012
