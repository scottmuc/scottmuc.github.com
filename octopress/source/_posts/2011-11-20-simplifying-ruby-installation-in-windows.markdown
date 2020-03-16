---
layout: post
title: "Simplifying Ruby Installation in Windows"
date: 2011-11-20 10:54
permalink: /blog/development/simplifying-ruby-installation-in-windows/
comments: true
categories: 
---

The one thing that has always annoyed me with installing ruby on Windows is that at some point in time I always need to setup the DevKit because I want to install gems that require native extensions. Also, I want the ability to switch between different versions of ruby.

<!-- more -->

### Introducing Yari

yari is a tool (implemented in Powershell) I've made to simplify both of these operations. Right now it's hard coded to only install the 2 ruby versions I care about right now.

* 1.8.7 so I can play with [Puppet](http://projects.puppetlabs.com/projects/1/wiki/Puppet_Windows)
* 1.9.2 to be bleeding edge

**Installation** is described in the readme on the github page. Once installed you can simply run yari from anywhere in the command prompt. Also, you can pass a -InstallMachine argument to yari and it will permanently configure your PATH for the ruby version selected.

So far this tool has made my ruby management in Windows a lot simpler and hopefully it can help you out too. If you find it useful but can use a few changes please submit issues to the projects github page. Or better yet, fork it and submit a pull request.

After installing, here's what you'll be able to do (and installing things that require native extensions will work):
<pre>
C:\Users\smuc> ruby -v
'ruby' is not recognized as an internal or external command,
operable program or batch file.

C:\user\smuc> yari 1.9.2

C:\user\smuc> ruby -v
ruby 1.9.2p290 (2011-07-09) [i386-mingw32]

C:\user\smuc> yari 1.8.7

C:\user\smuc> ruby -v
ruby 1.8.7 (2011-02-18 patchlevel 334) [i386-mingw32]

C:\user\smuc>
</pre>
