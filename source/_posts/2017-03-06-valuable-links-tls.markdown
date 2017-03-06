---
layout: post
title: "Valuable Links - TLS"
date: 2017-03-06 21:40:48 +0000
comments: true
categories:
  - links
  - tls
---
This is embarassing to admit but I have to relearn TLS on a continuous basis. Unless I'm working on something that's TLS
heavy I will forget everything over a short amount of time.

## SSL vs TLS

The fact that I still think [SSL][ssl-link] over [TLS][tls-link] is likely a result of my first exposure to dealing with
these sorts of things was in a multi-tenant Windows web hosting environment. I seriously can't believe I was able to do
the things I did back then without having one iota of understanding what was going on.

I only learned that the difference between *SSL* and *TLS* is so simple just a couple years ago. It boils down to this:

* Are you expecting a secure connection because the port you're connecting to is by definition intended to be a secure
  port? Then we're talking about **SSL**.
* Do you start connecting and switch to a secure mode based on the feedback coming back over the "wire"? Then we're
  talking about **TLS**.

Yup, it's as simple as that. Big thanks to the [SSL vs TLS - What's the Difference?][best-ssl-vs-tls] post for helping
me finally understand this!

[best-ssl-vs-tls]: https://luxsci.com/blog/ssl-versus-tls-whats-the-difference.html

[ssl-link]: https://en.wikipedia.org/wiki/Transport_Layer_Security
[tls-link]: https://en.wikipedia.org/wiki/Transport_Layer_Security

## What Goes On When You Use HTTPS?

Nearly everytime I need to setup [HTTPS][https-link] I find that I need to reread the
[First Few Milliseconds of an HTTPS Connection][https-link] article. It's a long read but if you spend a day following
the links and focusing on understanding that article, you'll tackle any *HTTPS* related task with ease. You'll just
forget it all after a few days.

[https-link]: https://en.wikipedia.org/wiki/HTTPS
[first-few-seconds]: http://www.moserware.com/2009/06/first-few-milliseconds-of-https.html

## What Are Certificates?



## I Keep Hearing About Exploits

The following video by [Moxie Marlinspike][moxie-website] is told at a nice slow pace on how security can break down.
It's a good reminder to be humble when dealing with these complex security tools. It really helps if you understand the
basics of certificates and the protocols that apply.

{% youtube ibF36Yyeehw %}

[moxie-website]: https://moxie.org/

## Summary

Hope these help. It's inspired by all the other folks that also couldn't tell me the difference between SSL and TLS.
It's easy to feel ashamed about not knowing these technologies that are the cornerstone to our industry.

*This post took 1 pomodoro to complete*


