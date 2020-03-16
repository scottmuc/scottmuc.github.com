---
layout: post
title: "Book Review - Learning Continuous Integration with TeamCity by Manoj Mahalingam"
date: 2014-10-13 17:34:30 +0530
comments: true
categories:
  - book
  - review
---

**Disclaimer** - I personally know Manoj and have worked with him in the past. I was asked to review his book,
[Learning Continuous Integration with TeamCity][book], by [Packt Publishing][packt-publishing].

[packt-publishing]: https://www.packtpub.com/
[book]: http://goo.gl/5wUy3v

[{% img left https://www.packtpub.com/sites/default/files/9518OT%5f0.jpg 200 %}][book]
As someone who had spent many years consulting around the benefits of CI and implementing build and release strategies,
I was interested in how the minutiae would be brought up in a book.

The book begins by going through the creation and configuration of a similar build process for 3 specific tech stacks:
Java, .Net, and Ruby. At first it feels rather repetitive but the repetition actually helps reinforce the importance of
consistency. With all tech stacks, the procedure is pretty much the same: `Checkout, Build, Test, Package, Deploy, Run
UI Tests`. This pattern really stands out after going through chapters 4, 5 and 6 and builds up a foundation for the
reader so that the rest of the material makes sense.

After setting the stage with these build configurations, he then dives into content that's a bit more TeamCity
specific. This area was much more interesting to me as I'm already experienced with build and release concepts, but it
was nice to see TeamCity specific features explained in a little bit more depth.

One piece that felt a little out of place was the topic on Continuous Delivery. It's really hard to introduce a concept
that has a massive book dedicated to the subject. Not sure if its inclusion will help out the beginner that this book
appears to be targeting.

What I enjoyed the most was raising operational concerns around running a CI server system. It was likely
deliberately left out, but I would have liked to see a mention of automated configuration management. There's something
incredibly powerful around automation around your automation.

It became clear to me that this book is meant for the absolute beginner. Manoj hints at better ways of implementing
his example configurations (eg: using wrapper scripts over the built in runners) but doesn't distract the reader
with these possibly subjective tangents. I would recommend this book if you've never worked with any CI tool and like
to have a guidebook as you try and set it up yourself. Despite being over 250 pages, the book is packed with
screenshots so that it's absolutely clear and won't take too long to go through the whole book. If you're already
experienced with some tools the principles around CI and good build hygiene, this book is useful to skim to see what's
new with TeamCity and some context of the usage of TeamCity specific features.

You can find more information about Manoj at his [twitter account][twitter], on [stackoverflow][stackoverflow], or his
[website][website].

[twitter]: https://twitter.com/manojlds
[website]: http://stacktoheap.com/
[stackoverflow]: http://stackoverflow.com/users/526535/manojlds

