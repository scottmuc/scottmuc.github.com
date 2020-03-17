---
layout: post
title: "Adopting PsGet for PowerShell Module Management"
date: 2011-08-03 11:18
permalink: /blog/development/adopting-psget-for-powershell-module-management/
comments: true
categories: 
---

For a while now I've been waiting for something that can be as close to ruby gems as possible for PowerShell. For a while I thought Nuget was the way but it's tight coupling to Visual Studio made it feel like it wasn't quite the right fit for PowerShell.

<!-- more -->

Mike Chaliy has made probably what could be called the simplest thing that could possibly work, and it works well! It's called PsGet (not to be confused with another project called PsGet). Take a look at the homepage and bask in its simplicity! Installation is copying/pasting the bootstrap code. Then install modules via name or absolute url to a powershell module.

Take a look at my post about Pester for an example of setting it up.

I have a couple features that I would like to bake into it. Hopefully I won't add too much complexity to it as I do really appreciate it's simple approach.

* Ability to deploy a specific version (using git tags) eg: https://github.com/scottmuc/Pester/archives/1.0.3
* Something similar to the concept of gemsets in rvm for easy switching between module versions

