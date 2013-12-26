---
layout: post
title: "My Developer Resolutions for 2013"
date: 2012-12-31 10:45
comments: true
categories:
---

2012 started off interesting because I was on on an iOS/ruby project. Nearly everything was new to me, but I liked how
my past experiences with other technology stacks were still relevant. Even though I didn't know the languages very well
the patterns I've seen before still applied.

The last quarter of the year was spent teaching new software developers in India, which has been one of the most
rewarding experiences of my career! It has definitely influenced my goals for 2013.

<!-- more -->

## Reflections on 2012 and carry over from 2011:

* Learn a server provisioning automation tool (2011) - I may have failed this one big time in 2011 but 2012 was a great
  year for this. I spent much of the year getting to know Chef. I feel I can check this one off now.
* Stop reading comments (2012) - Can't say I've stopped reading comments but I'm more conscious of when I'm wasting my
  time and tend to stop immediately. I started following a [useful twitter account](https://twitter.com/AvoidComments)
  which provides constant reminders that reading comments can be a huge waste of time.
* No client work when away from client (2012) - I twisted this one into a conditional success. I still like to hack in
  my spare time but I hack on my personal projects and stay away from doing extra billable hours when I should be having
  ME time. I've been much better at saying NO to excess outings and have focused on being healthy.
* Complete [YDeliver](https://github.com/manojlds/YDeliver/) (2012) - This one is a big failure. I barely touched
  YDeliver all year. Apparently it's been used in a few ThoughtWorks projects by [Manoj](https://github.com/manojlds) and he now
  has ownership of the repository. I've come to the realization that OSS doesn't come out of a vacuum. YDeliver was
  going strong when I still had fresh memories of that previous project, but as time went on, the relevance (to me) of
  YDeliver was diminishing. The takeaway for is me to maintain the momentum if it's there.

## Development Goals for 2013:

While 2012 was a great mix of fun ops tooling learning and consulting, I want to get back to my roots (pun?) and do more low
level tech related things. At the same time, I want to improve on my teaching skills which will mean facilitating more learning
workshops and presenting more frequently.

### Read [7 Languages in 7 Weeks](http://pragprog.com/book/btlang/seven-languages-in-seven-weeks)

Much of my effort is spent making many puzzle pieces fit together. Being isolated to a
[REPL](http://en.wikipedia.org/wiki/Read%E2%80%93eval%E2%80%93print_loop) and a text editor would be
pleasant change. Also, I haven't dived deep into a language lately. All my ruby learning has been simply by having to
write ruby code for work. I can't say I know the language, but I've managed to figure things out. I miss having a little
bit deeper knowledge of a language like I used to have with C#.

My action plan is to purchase the book and do 1 language a month. So by July this goal will be accomplished.

### More Presentations

Being a trainer at ThoughtWorks University taught me that I have a lot to learn about presenting. [Sumeet
Moghe](http://www.learninggeneralist.com/), who was the previous lead trainer, showed excellent examples of what makes a
good presentation. It was interesting, because I always felt those points internally, but he expresses them externally
in a digestable way that actually inspires me to make excellent presentations.

My colleague [Tim Brown](http://twitter.com/tpbrown) will be happy about this goal because he's been nagging me for a
while about this. Some topics I've been pondering so far are:

* Using [Vagrant](http://vagrantup.com/) to provide local infrastructure (aka never install a DB server on your machine again)
* Application configuration patterns with a non-trivial number of environments
* PowerShell BDD with [Pester](https://github.com/pester/Pester/)

The action plan is to choose 1 of these topics and present it 3 times. As well as the presentations, I will design 3
different workshops and faciliate each one 3 times.

### Expand my Ops Toolbelt

I really love ops and I think this is the career path I want to continue on. 2012 was great for learning
[Chef](http://www.opscode.com/chef/), so this year I want to learn [Puppet](http://puppetlabs.com/). Another area
that I would like to improve upon my knowledge is the domain of remoting.

My action plan is to duplicate any Vagrant server setups that use Chef to use Puppet as well. This will be great for
a side-by-side comparison. So currently that would mean the following repositories need a puppet provisioner:

* [vagrant-postgresql](https://github.com/scottmuc/vagrant-postgresql)
* [vagrant-rails](https://github.com/scottmuc/vagrant_rails)
* [gordon](https://github.com/scottmuc/gordon)

For remoting I plan on writing 10 blog posts on tools like ssh, winRM. I want to dive into the mechanics of
authentication and networking to better understand technologies like kerberos, certificates, and gssapi which are
crucial for making remoting work on a larger scale.

### Summary

The last few years I've learned a lot about managing my goals. I'm doing my best to keep my goals
[SMART](http://en.wikipedia.org/wiki/SMART_criteria) so that I can actually achieve them. We'll see how things turn
out 365 days from now!

