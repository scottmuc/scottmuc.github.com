---
title: Pester - PowerShell BDD Style Testing For The System Administrator
date: 2011-05-11T20:05:00Z
tags:
    - pester
aliases:
    - /pester-powershell-bdd-style-testing-for-the-system-administrator

---


Hi there and welcome to my demo of Pester, a BDD style testing framework for Powershell. The creation of Pester came out of the desire to test some build/deployment infrastructure we were creating for a project. We wrote nearly all the code without tests and it came to bite us in the end. I wanted to find a way ensure these problems didn't happen again as well as provide some code coverage to give new entrants to the codebase some confidence that they won't break everything.

<!-- more -->

Inside Pester there's already a trivial Calculator example but it's not really the best way to demonstrate the app. Pester itself is tested using Pester. In fact it's being tested by the version of Pester that's under test (although Martin pointed out some uncovered areas). Still, probably not the best way to show how it all works.

What I'm going to go through here is the beginning of what could possibly be a real world scenario for a person assigned to write deployment code.

## Here's the story:

Initech has had issues where their .Net web applications have been deployed on production servers with the debug compilation flag set to true. This has made production support people irritable because they now manually tweak the web.config every single time they do a deploy. Michael Bolton has decided he's going to automate this step but wants to right it test first. He doesn't want to repeat the debacle of his previous attempt at being clever.

First step is to setup a project. Mike (as he now likes be known as) decides to call his project of tools IDeploy and creates a folder of that name where he does is development. 

Setting up Pester is simply a matter of following the instructions at PsGet (author Mike Chaliy has tons of great PowerShell modules)then running Install-Pester. Then running Import-Module Pester anytime you open up a PowerShell session where you want to use Pester. The console should look like the following:

<pre>
SMUC-PC {C:\d\IDeploy} (new-object Net.WebClient).DownloadString("http://bit.ly/GetPsGet") | iex
Downloading PsGet from https://github.com/chaliy/psget/raw/master/PsGet/PsGet.psm1
PsGet is installed and ready to use
USAGE:
    import-module PsGet
    install-module https://github.com/chaliy/psurl/raw/master/PsUrl/PsUrl.psm1

For more details:
    get-help install-module
Or visit http://psget.net
SMUC-PC {C:\d\IDeploy} import-module PsGet
SMUC-PC {C:\d\IDeploy} install-module Pester
Module Pester was successfully installed.
SMUC-PC {C:\d\IDeploy} import-module Pester
SMUC-PC {C:\d\IDeploy} Get-Module

ModuleType Name                      ExportedCommands

