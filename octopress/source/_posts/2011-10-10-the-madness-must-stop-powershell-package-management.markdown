---
layout: post
title: "The Madness Must Stop! - PowerShell Package Management"
date: 2011-10-10 10:57
permalink: /blog/development/the-madness-must-stop-powershell-package-management/
comments: true
categories: 
---

As a publisher for a few Powershell modules I'm getting frustrated by the lack of standarization in the community on how to support versioning, deployment and conventions within Powershell module distribution. All I can find is rather adhoc mechanisms of sharing poorly written code with the "use at your own risk" kinds of disclaimers. 

<!-- more -->

I keep on looking at the ruby community with their amazing tools like rvm, ruby gems, and bundler to make ruby scripting highly testable and repeatable. The formalization also means that you don't see people writing ruby code with random imports and magical includes that only work in very specific scenarios. By creating a community around gems and some basic standards, it's forced developers to factor their ruby code accordingly and modularize when necessary.

So far I've found the following package management tools for Windows:

* Nuget - conversations online have lead me to understand that nuget isn't meant for powershell
* OpenWrap - still under evaluation
* Chocolatey (by Rob Reynolds) - wrapper around nuget. Seems to be about config management like Puppet and Chef.
* Chewie (by Eric Ridgeway) - bundler for nuget. Could have some input on how to manage PS module packages. Potential to be the front end for managing suites of modules.
* Web Platform Installer - not even going to go there
* PS-Get (by Andrew Nurse) - built around nuget but Andrew has the MS connection to make it popular.
* PS-Get (by Mike Chaliy) - my favourite but needs some growth in the community. A showcase for simplest thing that can possibly work.
* Nugit - almost forgot about this late entry. Don't know anything about it yet.

Here's my wishlist (that none of the above can fulfill):

* Ability to deploy multiple versions of the same package
* Different levels of isolation (System | Profile | Project ) levels
* A management tool so that projects can declare versions and packages that they need along with dependency chaining (ala bundler and ivy)
* Packages can be stored in many different locations (Nuget.org, github, some url pattern)

History has shown that pretty much all languages need their own package management system (which can then be wrapped by a system level package manager). I vote that some of us in the community have a round-table and discuss how to move forward with this. There's so much to learn from what already exists to get Windows out of the scripting ghetto.

