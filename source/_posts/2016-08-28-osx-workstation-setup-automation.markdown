---
layout: post
title: "OSX Workstation Setup Automation"
date: 2016-08-28 15:45:48 +0100
comments: true
categories:
---

## Objective

IT infrastructure has gone done a long journey of automating all of the things. We have automate deployments, production
environments, and regression testing. We can make an API request to Amazon and get a server instances in minutes. In
general, you would get what you asked for.

The developer workstation is another story. Often, the developer workstation is a finely crafted machine. Just like
production, your workstation should be treated like a [hotel room][phil-cattle-vs-pets] rather than a home. The problems
end to manifest as the following:

TODO insert screen grab of twitter thread

* Onboarding a new employee takes too long. How long is too long? For me a new developer should be able to go from an
  unboxed machine to a successful build in less than a day.
* Lot's of cruft builds over time leaving to things working (or breaking) without much indication as to why. Things like
  out of date tools like `vagrant` or `docker`.
* The machine starts off well, but as changes are introduced, newer machines never receive these changes.

[phil-cattle-vs-pets]: https://twitter.com/pcalcado/status/759218156493795329

## My Workstation

Some of what I'm discussing here is in a professional work context, but I find this automation valuable for my personal
machines too. Being able to recreate my machine on a regular basis gives me faith that I can replicate this environment
if needed. I also don't worry about backups either because I just distribute my persistent files on cloud services.

## Existing Tools

There are plenty of tools that help with this process.

* [Sprout][sprout-wrap] - Intially this was called `pivotal_workstation` but transformed into a slightly more
  configurable and modular system. It uses `chef` under the covers and is quite comprehensive.
* [Boxen][boxen] - Created by GitHub, this tool uses `puppet` under the covers.
* [osxc][osxc] - Not sure where it was born from, but it uses `ansible` underneath.

All three of these strive for the same thing, but use different engines. It probably makes sense to use whichever one
uses the Configuration Management tool you're most comfortable with. The problem though, is that these tools may add a
bit more complexity than is really necessary.

[Sam Gibson][sams-setup] wrote a great article about this. He believed the above styles are too complex and strived to
find a simpler solution.

[sams-setup]: https://www.thoughtworks.com/p2magazine/issue08/babushka/
[sprout-wrap]: https://github.com/pivotal-sprout/sprout-wrap
[boxen]: https://github.com/boxen/boxen
[osxc]: https://osxc.github.io/

## Starting From Scratch

Technically you cannot start from scratch because you need at least an existing OSX machine to create the installer USE
stick. To do this, I followed the instructions on Brooke Kuhlmann's [workstation automation][bkuhlmann-setup]
repository. Though, I'll probably copy this information into my own repository for reasons I'll explain later.

[bkuhlmann-setup]: https://github.com/bkuhlmann/osx#os-x-el-capitan

## There's No Place Like Home

While I like where he went with it, but it missed a few things that I liked that **sprout**
offered. I did have a full **sprout** way of setting up my machine that was inspired by
[Gary Bernhardt's][bernhardt-setup] where his home directory is a repository.

I've found that many people have `dotfile` repositories that get installed by their scripting. This adds a level of
complexity and absraction that is not really necessary. Cloning straight into home removes a lot of moving parts.
All of the sudden all the abstractions and scripts that go into setting up your home directory disappear. 
If you can successfully transform home into a repository you have a simple and easy way of maintaining dot files,
and scripts to support your day-to-day activities. You've essentially already completed the first
step in getting your machine setup. Also, because it's a repository, when you delete a file, it goes away.
It's inherently an idempotent system!

[bernhardt-setup]: https://github.com/garybernhardt/dotfiles

## Keys and Secrets Management

Since this way of setting up your machine depends on being able to clone git repositories as a first step, it means that
you're going to need to have your keys. For this, I follow Tammer Saleh's excellent post on
[building an encrypted usb drive for your keys][tammer-usb].

Every other secret goes into 1Password. At work we use LastPass and are able to integrate it into our scripting by
running commands like `eval "$(lpass show 'Awesome/aws-secrets' --notes)"`. This way secrets are never recorded on disk in
their unencryped form.

[tammer-usb]: http://tammersaleh.com/posts/building-an-encrypted-usb-drive-for-your-ssh-keys-in-os-x/

## What Could Possibly Go Wrong

Workstation automation can be extremely flaky and that's due to the nature of it fetching dependencies from all sorts
of locations. This drives home the reason for needing to run this frequently. I can pretty much guarantee that the code
that worked a month ago will work today. Packages change name, or download locations change. This can result in a lot of
mistrust of the automation. This leads to the following...

## Duplication is OK

Choosing when to vendor something is an important decision to make around construction your setup. The more you vendor,
the more reproducable your automation will be. The more you vendor, the harder it is to keep up to date with changes.

Something that took me a while to come to terms with is the amount of duplication in the world. There does not exist one
true workstation to rule them all, but it is important to figure out what level of "sameness" is needed for your
development needs. I have my personal workstation automation, but I also work with a team-wide system for workstation
automation. I believe the team level is the largest group you can realistically manage without creating something too
complex of manage. Your automation is going to be a somewhat handcrafted piece of art, but it's going to be a simple piece
of art (hopefully). 

## Summary

In the end, I have a setup I'm relatively happy with. It turns out that my implementation looks oddly similar to a
colleague of mine's, [Gerard Lazu's][gerhards-setup] workstation setup. He replaced **sprout** with a collection to
shell scripts that I may be tempted to migrate towards. But one thing at a time; that's an important principle around
working on this kind of stuff. There are a lot of moving parts and it'll only prove itself out over time.

[gerhards-setup]: https://github.com/gerhard/setup

*This post took 4 pomodoros to complete*
