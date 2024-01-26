---
title: What Makes Something a "Devops" Tool?
date: 2014-12-05T09:43:52+05:30
tags:
    - devops
aliases:
    - /what-makes-something-a-devops-tool

---

The term "devops" has only existed for 5 years and the trend is to label your tool as a devops tool in order for it to
gain traction. When I look at the landscape of tools, most look like Ops tools, just relabeled.

The reason this has been on my mind is that myself and [Subhas][subhas] released a website called
[Devops Bookmarks][devops-bookmarks] to aggregate all the tools in the landscape. Using a [Github Pull Request][ghpr]
for contributions results in reviewing a lot of tools and after a short amount of time I found myself second-guessing
the need for such a website. It felt that we were diluting the devops term.

Here's a couple that I think got it wrong:

* http://blog.xebialabs.com/2014/04/30/taxonomy/
* http://newrelic.com/devops/toolset

I'm greatly concerned the Devops Bookmarks is going to feel the same unless context is added to support the inclusion of
the tools. It's why I made a [devops links post](/valuable-links-devops/) on my blog so that I could explain a little
bit why something is included.

[subhas]: twitter.com/rdsubhas
[devops-bookmarks]: http://www.devopsbookmarks.com/
[ghpr]: https://github.com/devopsbookmarks/devopsbookmarks.com/blob/master/CONTRIBUTING.md

### Going To the [Gemba][gemba]

Honestly, [Patrick Debois's][patrick-blog] blog is a must read to get adequate context around devops. He has an
excellent [introductory post][intro-post] (written almost 5 years ago) and he's made an excellent slide deck pointing
out the same concepts I'm struggling with right now.

<div style="width: 450px; margin: 0 auto 0 auto;">
<iframe src="//www.slideshare.net/slideshow/embed_code/9631120" width="425" height="355" frameborder="0" marginwidth="0"
marginheight="0" scrolling="no" style="border:1px solid #CCC; border-width:1px; margin-bottom:5px; max-width: 100%;"
allowfullscreen> </iframe> <div style="margin-bottom:5px"> <strong> <a
href="//www.slideshare.net/jedi4ever/devops-tools-fools-and-other-smart-things" title="Devops Tools Fools and Other
smart things" target="_blank">Devops Tools Fools and Other smart things</a> </strong> from <strong><a
href="//www.slideshare.net/jedi4ever" target="_blank">Patrick Debois</a></strong> </div>
</div>

In that presentation the point of articulating desired *Behaviours* is very clear. To me a devops tool is something that
promotes good behaviour if it's executed well. The desired behaviours are cleary represented by
[John Willis'][john-willis] excellent blog post titled: [What Devops Means to Me][devops-to-me].

[gemba]: https://en.wikipedia.org/wiki/Gemba
[patrick-blog]: http://www.jedi.be/blog/
[intro-post]: http://www.jedi.be/blog/2010/02/12/what-is-this-devops-thing-anyway/

[john-willis]: https://twitter.com/botchagalupe
[devops-to-me]: https://www.chef.io/blog/2010/07/16/what-devops-means-to-me/

**Examples**

* **Chef** - promotes the behaviour of automating the configuration of machines. This reduces the tribal knowledge and
  provides a collaborative environment to work on this automation (because it's now code instead of a Word doc).
* **Vagrant** - enables the ability to create production like environments anywhere. All roles in the software lifecycle
  work with consistent infrastructure, thus can share their improvements and feedback.
* **Gauntlt** - security testing for humans. Using Cucumber, security tests can be easily read by everyone and the tests
  can be executed in a continuous integration pipeline.

**Not Sure About These**

* **Gitorious** - is it a devops tool because having source control is necessary?
* **Static Analysis** - is a python static analysis tool really a devops tool? What about something like Puppet Lint?

Are some tools purely dev, and some purely ops? Everything lies on a spectrum and I guess what I'm wondering is if
devops tools are simply the ones that are in the middle of that spectrum.

### Moving Forward

I'm torn, but I think a feature in [Devops Bookmarks][devops-bookmarks] should be a requirement to explain why the tools
is a devops tool, and what behaviours does it try to support. If no explanation is provided, then it cannot be listed.

Currently the website is broken down by *Topics*, but that could change to *Behaviours* and the tools can be tagged with
the behaviours that it supports.

Consider this a request for comments.

*This post took 3 pomodoros to complete*






