---
layout: post
title: "Valuable Links - TLS"
date: 2017-03-06 21:40:48 +0000
comments: true
categories:
  - links
  - tls
---
This is embarassing to admit but I have to re-learn TLS on a continuous basis. This post is my cheatsheet to remembering
all the itty bitty details, which I'll forgot shortly after I do whatever TLS related task I have on my plate.

## What's The Difference Between SSL And TLS?

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

## What Are Certificates?

When it comes to certifcates, I've yet to find a resource that does a good high level explanation. For the time being
I've found [SSL/TLS Certificates Beginners' Tutorial][cert-tut] the best resource because it does explain the `openssl`
command in a bit of detail and links to even more resources. Having a basic grasp of [Certificate Authorities][ca-link],
and the [chain of trust][chain-of-trust] (not a 90's Christian metal band) is another piece
of the puzzle to understand well because it's where things can often go wrong and result in cryptic errors
(like the browser warnings that we all ignore).

My simplified understanding is that a certificate certifies hat the one (machine) with the certificate is the entity
that the certificate says it's certifying. I believe this is true for server and client certificates, but I am not 100% sure.

[chain-of-trust]: https://en.wikipedia.org/wiki/Chain_of_trust
[ca-link]: https://en.wikipedia.org/wiki/Certificate_authority
[cert-tut]: https://blog.talpor.com/2015/07/ssltls-certificates-beginners-tutorial/

## What Goes On When You Use HTTPS?

Nearly everytime I need to setup [HTTPS][https-link] I find that I need to reread the
[First Few Milliseconds of an HTTPS Connection][https-link] article. It's a long read but if you spend a day following
the links and focusing on understanding that article, you'll tackle any *HTTPS* related task with ease. You'll just
forget it all after a few days.

[https-link]: https://en.wikipedia.org/wiki/HTTPS
[first-few-seconds]: http://www.moserware.com/2009/06/first-few-milliseconds-of-https.html

## How Do Things Go Wrong?

The following video by [Moxie Marlinspike][moxie-website] is told at a nice slow pace on how certificates can break down.
It's a good reminder to be humble when dealing with these complex security tools. I tend to rewatch this video after
reading the above articles. If I can follow the video, then I've chunked enough knowledge to move on.

{% youtube ibF36Yyeehw %}

[moxie-website]: https://moxie.org/

## Summary

Hope these help. It's inspired by all the other folks that also couldn't tell me the difference between SSL and TLS.
It's easy to feel ashamed about not knowing these technologies, security is a very hard domain. I've kept my
explanations to a minimum because I don't feel confident I can provide correct information, nor am I able to explain it
simply. When in doubt, I'll look to the resources here. Hopefully this will give you a good starting ground if you've
been asked to do anything related to **SSL/TLS** and **certificates**.

*This post took 2 pomodoros to complete*
