--- 
layout: post
title: "Migrated to Octopress!"
date: 2012-01-15 08:45
comments: true
categories: 
---

It's funny how many posts there are from [people who have migrated to Octopress](https://www.google.com/search?q=migrating+to+octopress). What would motivate someone to post their migration experience? I think I understand why. It's such a shift in thinking that it's worth talking about. My migration experience was so much pleasant that I feel that I must let everyone know.

<!-- more -->

So what's my motivation for moving to a new blogging platform? Was what I was using before broken? Not
exactly, but I didn't enjoy the blogging experience that much because I found the experience to not be a
lot of fun. I want to start writing more and feel that removing any barriers to this activity will get more
content out there. I also didn't like having to rely on a server side application server. For a website that changes around 30 times a year, a database backed content management system was overkill. Also, the website was hosted on a friends server that I had no backups for, and didn't want to spend much effort creating a backup/restoration strategy.

After some searching, I found [jekyll](http://jekyllrb.com/), wich I love for its simplicity. I was going to use bare bones jekyll but then I was directed to [Octopress](http://octopress.org/) which is a framework to make working
with jekyll a lot easier.

Since I was on Windows my install was a bit different than the documentation. Luckily I have created a tool
to make [ruby installation simple on Windows](http://scottmuc.com/blog/development/simplifying-ruby-installation-in-windows/). Here's everything I did to go from nothing to having a working
blog hosted on github!

1. Installed [yari](https://github.com/scottmuc/yari), and got my environment using ruby 1.9.2
2. Followed the octopress [setup instructions](http://octopress.org/docs/setup/)
3. Followed the octopress [deployment instructions](http://octopress.org/docs/deploying/github/)

I was blown away when I saw my blog appear when I went to [scottmuc.github.com](http://scottmuc.github.com/).

### Migration

Unfortunately my previous blog engine's export functionality is broken so migration had to be manual. Every post
is going to be copy and pasted and re-linkified in markdown syntax. I wanted to get this new version up and 
running first so I made a few compromises.

1. I decided not to migrate every post. Some of my old content doesn't really fit with the rest of the content
so I decided to abandon the posts. I'm keeping the [old blog](http://old.scottmuc.com/) up and running so that people can still find the posts there if they really want them.
2. Keeping links working is **extremely** important to me. Octopress provides the ability to set the permalink on
an individual post which made this element of the migration super easy. Every blog post has "front matter" which is a little bit of YAML meta data. Adding a permalink value will make the post accessible from that URL path: <pre>permalink: /foo/bar/post</pre>
3. I went live with an incomplete migration. I only have about a third of my posts migrated. I did this so I can 
get this live and sort out issues right away before everything is migrated. Also, it's a way to motivate me to get
everything into the new blog ASAP.
4. Added a custom [404](http://scottmuc.com/non-existent-page/) to guide those who want to see non-migrated content to the old blog.
5. Added "scottmuc.com" to my CNAME file. I followed the instructions [here](https://github.com/blog/315-cname-support-for-github-pages). Note that this is enabled for non-paying users now! I use [gandi.net](http://gandi.net/) for my DNS and updating the A record and CNAME (for www) was a snap.


### Other Stuff

When setting up the blog I had an online discussion about how url design. I decided to leave dates out of my url
and to include them in the title slug only when appropriate. I came to the conclusion that url slug design is
extremely important and doesn't always have to be word for word related to the post title. Because Octopress is
so easy to configure, the configuration setting for this is simply this value in \_config.yml:

<pre>permalink: /:title/</pre>

I'm now using other 3rd party services like feedburner, disqus, and gist.github.com. With all of these things 
taken care of for me, authoring a post is a call to *rake new_post["post title"]*, edit in vim, and preview
using *rake preview* to view the changes before I deploy them.

The one thing I would like to change is learn how to create a custom theme. It's quite common to see a lot of
octopress based blogs use the same default theme that I'm using now.

### Summary

Octopress is awesome! I love the workflow for authoring new content. I much prefer markdown over using some
WYSIWIG editor over a web connection. The ability to do everything local means there's significantly less
effort for me. I believe Octopress is an excellent example of how an excellent deployment mechanism can shift
how work can be done. If pushing brand new code to production is that easy, it removes the need for complex
server side code.

Lastly, if you would like to see what a raw blog post looks like, take a look [here](https://github.com/scottmuc/scottmuc.github.com/tree/source/source/_posts).
