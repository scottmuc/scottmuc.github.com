---
layout: post
title: "Pester - PowerShell BDD Style Testing For The System Administrator"
date: 2011-05-11 20:05
permalink: /blog/development/pester-bdd-for-the-system-administrator/
comments: true
categories: 
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
---------- ----                      ----------------
Script     PsGet                     {Get-PsGetModuleInfo, Install-Module}
Script     Pester                    {It, Describe, New-Fixture, Invoke-Pester...}

</pre>

Pester includes a helper function called Create-Fixture. Calling the function with no args looks like the following:

<pre>
SMUC-PC {C:\d\IDeploy} Create-Fixture
invalid usage, please specify (path, name)
eg: .\Create-Fixture -Path Foo -Name Bar
creates .\Foo\Bar.ps1 and .\Foo.Bar.Tests.ps1
</pre>

Create-Fixture wants to know what the path of your function is going to be and what to call it. I personally like having my tests next to what I'm testing so Create-Fixture sort of enforces this convention. Note that Pester can be used without ever using this function. I just never remember how to create my fixtures so this saves me a bit of copying and pasting.

Armed with a quick way to create fixtures Michael runs his Create-Fixture to scaffold his feature. He decides he wants it to be called Ensure-AspNetDebugIsFalse and places it in the Deploy\Functions directory of his project.

<pre>
SMUC-PC {C:\d\IDeploy} Create-Fixture Deploy\Functions Ensure-AspNetDebugIsFalse
Creating => Deploy\Functions\Ensure-AspNetDebugIsFalse.ps1
Creating => Deploy\Functions\Ensure-AspNetDebugIsFalse.Tests.ps1
</pre>

Wanting to see some red he runs the tests by running Invoke-Pester which loads all files that match \*.Tests.ps1 recursively in the current directory.

<pre>
SMUCS-PC {C:\d\IDeploy} Invoke-Pester
Executing all tests in C:\dev\IDeploy\Deploy
Describing Ensure-AspNetDebugIsFalse
does something useful
Tests completed
Passed: 0 Failed: 1
</pre>

As you can see Pester by default makes a failing test. Now it's time for Michael to update the test to make it more meaningful. Here's his test file after he's setup his expectations.

Wow, it failed! Why is that? By default Pester will generate a fixture that is silly and won't ever pass. So what should we do with this broken function? As it stands now it's totally empty. Let's update our specification (aka Test) and do something useful.

So Michael has written a test that didn't require a lot of code, but is actually doing a few cool things. I'll try to go line by line and explain what's going on:

Ensure-AspNetDebugIsFalse.Tests.ps1 contents:

<pre>
$pwd = Split-Path -Parent $MyInvocation.MyCommand.Path
$sut = (Split-Path -Leaf $MyInvocation.MyCommand.Path).Replace(".Tests.", ".")
. "$pwd\$sut"
. "$pwd\..\..\Pester.1.0.1\tools\Pester.ps1"
 
Describe "Ensure-AspNetDebugIsFalse" {
 
    Setup -File "inetpub\wwwroot\testsite\web.config" `
                "&lt;configuration&gt;&lt;system.web&gt;&lt;compilation debug='true' /&gt;&lt;/system.web&gt;&lt;/configuration&gt;"
 
    It "switches debug attribute to false for a web.config in a given website path" {
        Ensure-AspNetDebugIsFalse "$TestDrive\inetpub\wwwroot\testsite"
 
        [xml] $xml = Get-Content "$TestDrive\inetpub\wwwroot\testsite\web.config"
        $xml.configuration."system.web".compilation.debug.should.be("false")
    }
}
</pre>

Ensure-AspNetDebugIsFalse.ps1 contents:

    function Ensure-AspNetDebugIsFalse($websitePath) {
     
    }

Here are some details on some of the lines in the test.

1. We obtain the directory of the test script because we need to base all other paths off of it

2. Here's how we ensure our code/test conventions. This means that the code in Foo.Tests.ps1 will automatically include Foo.ps1 (the code under test) in line 3

4. Create-Fixture automagically resolves the Pester path no matter what version of Pester you have. This directory will remain static. I might implement an Upgrade-Fixture script so that as you update your Pester version you can update your tests as well.

8. Pester has the ability to do filesystem based setups. This line is creating a file in isolation in what Pester calls the $TestDrive. The $TestDrive is disposed of after every Describe context. This allows you to perform filesystem related tasks without having to maintain separate test files. This line has created a file called web.config in the inetpub\wwwroot\testsite path in the $TestDrive and the 2nd argument is the content of that file.

12. We execute the code we want to test

15. We assert that the debug attribute has been set to false.

Let's make this bad boy pass!
Here's the function fleshed out and the test passes!

<pre>
function Ensure-AspNetDebugIsFalse($websitePath) {
    $webConfigPath = "$websitePath\web.config"
 
    [xml] $webConfig = Get-Content $webConfigPath
    $webConfig.configuration."system.web".compilation.debug = "false"
    $webConfig.Save($webConfigPath)
}
</pre>

and the execution of the test:

<pre>
SMUC-PC {C:\d\IDeploy} Invoke-Pester
Executing all tests in C:\dev\IDeploy\Deploy
Describing Ensure-AspNetDebugIsFalse
switches debug attribute to false for a web.config in a given website path
Tests completed
Passed: 1 Failed: 0
</pre>

Now Michael has an extra tool at his disposal to call when he's performing deployments. If only he had taken this approach when he wrote his money transfer code. He wouldn't have missed all those decimal places!

Hope this helps get you running with Pester. There's a lot more I would like to add but I'm really only adding stuff as I see the need for it. Pester was driven out by my desire to refactor another Powershell library I wrote called PowerYaml. Take a look at the project page to see how I took the untested code, wrapped tests around it, then refactored.

Feedback is greatly appreciate and I hope this helps your teams make reliable Powershell scripts.

