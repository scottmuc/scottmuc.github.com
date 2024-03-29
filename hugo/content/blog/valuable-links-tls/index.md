---
title: Valuable Links - TLS
date: 2017-03-06T21:40:48Z
tags:
    - links
    - tls
aliases:
    - /valuable-links-tls

---

This is embarassing to admit but I have to re-learn Transport Layer Security (TLS) on a continuous basis. This post
is my cheatsheet to remembering all the itty bitty details, which I'll forgot shortly after I do whatever TLS related
task I have on my plate.

What's TLS you ask? Take 9 minutes to listen to this fantastic episode on [Public key cryptography][podcast-episode]
which will prime (hah!) you up for the rest of the post.

[podcast-episode]: http://www.bbc.co.uk/programmes/p04vqrwy

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

[Everything You Ever Wanted to Know SSL...][everything-ssl] is another great post that dives deep into the low level of
what happens in creating a secure connection.

[https-link]: https://en.wikipedia.org/wiki/HTTPS
[first-few-seconds]: http://www.moserware.com/2009/06/first-few-milliseconds-of-https.html
[everything-ssl]: http://www.robinhowlett.com/blog/2016/01/05/everything-you-ever-wanted-to-know-about-ssl-but-were-afraid-to-ask/

## Hostname Verification

[Will Sargent][will-sargent] wrote an excellent series on [Testing Hostname Verification][testing-hostname-verification]
which goes into the guts of many of the parts discussed in this post.

[will-sargent]: "https://twitter.com/will_sargent"
[testing-hostname-verification]: https://tersesystems.com/blog/2014/03/31/testing-hostname-verification/

## How Do Things Go Wrong?

The following video by [Moxie Marlinspike][moxie-website] is told at a nice slow pace on how certificates can break down.
It's a good reminder to be humble when dealing with these complex security tools. I tend to rewatch this video after
reading the above articles. If I can follow the video, then I've chunked enough knowledge to move on.

{{< youtube ibF36Yyeehw >}}

[moxie-website]: https://moxie.org/

## Summary

Hope these help. It's inspired by all the other folks that also couldn't tell me the difference between SSL and TLS.
It's easy to feel ashamed about not knowing these technologies, security is a very hard domain. I've kept my
explanations to a minimum because I don't feel confident I can provide correct information, nor am I able to explain it
simply. When in doubt, I'll look to the resources here. Hopefully this will give you a good starting ground if you've
been asked to do anything related to **SSL/TLS** and **certificates**.

**Update** I was contacted by [Matt Banner][matt-banner] who asked me to update this article with some content from his
website. I took a look at it and it's a [really good guide][infographic] for updating from `http` to `https` (pop quiz, is
this SSL or TLS?) plus it has a really good infographic.

[matt-banner]: https://twitter.com/BlastYourBlog
[infographic]: https://www.onblastblog.com/http-to-https/

**Update 2** Found a couple more excellent resources on the subject. [A Tour of TLS][tour-tls] is probably the clearest
video I've watched on the subject. It clarified some of the more obscure pieces of the protocol like OCSP and Stapling
in such a way that I finally got it! [Everything PKI][everything-pki] brings pretty much of what's in the Tour of
TLS video in the frame of a system operator (e.g.: through `openssl` commands).

[everything-pki]: https://smallstep.com/blog/everything-pki.html
[tour-tls]: https://www.youtube.com/watch?v=yzz3bcnWf7M

*This post took 2 pomodoros to complete*

