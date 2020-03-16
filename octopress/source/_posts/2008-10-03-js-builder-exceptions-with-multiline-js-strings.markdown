---
layout: post
title: "JS Builder Exceptions With Multiline JS Strings"
date: 2008-10-03 16:32
comments: true
permalink: /blog/development/js-builder-exceptions-with-multiline-js-strings/
categories: 
---
I'm creating an automated build process and one of the items in the build is to build an external resource which is a java script library. The javascript is contained in multiple files and I want to concatenate them into one script. I found a very useful tool called [JS Builder](http://code.google.com/p/js-builder/) which has a command line executable which I can use to automate the construction of this javascript file.

<!-- more -->

Setup was very simple. I created a JS Build project in my build directory. Inside it I reference the .js files I want to include/minimize/concatenate. I select the folder containing the scripts and attempt to run the build.

My first attempt results in the following error:

An unhandled exception occurred while building:
<pre>
System.Exception: Error: JSMIN unterminated string literal: 10
 at JSBuild.JSMin.action(Int32 d)
 at JSBuild.JSMin.jsmin()
 at JSBuild.JSMin.MinifyToString(String src)
 at JSBuild.JavascriptFile.get_Minified()
 at JSBuild.JavascriptFile.MinifyTo(String target)
 at JSBuild.ProjectBuilder.Build(Project project, String outputDir)
 at JSBuild.ProjectBuilder.Build(Project project)
 at JSBuild.MainForm.Build()
</pre>
I tried to google the error but could not find anything. To troubleshoot I added 1 javascript file at a time and found that there were a small group of files that caused this error. It turns out that JS Build chokes on multiline strings that use the '\' character to have the string declared on multiple lines in the source code.

Not a big deal, I did multiline string concatenation to build the strings instead (probably implement a string builder in the future). After that, everything built fine. During the research I came across an [interesting article](http://kazoolist.blogspot.com/2008/06/multi-line-strings-in-javascript.html) regarding the validitiy of the syntax of strings using the '\' character.

Now I have a simple NAnt task that executes JS Build that constructs a single javascript file for a complex javascript audio player for the future CBC Radio 3 website (kudos to [Phil Rabin](http://www.rabin.ca/) for the code). Next step is to figure out how to deal with thirdparty javascript libraries that the project depends on.

