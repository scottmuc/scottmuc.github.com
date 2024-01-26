---
title: 'Growing an Open Source Project: The Pester Story'
date: 2014-08-07T17:50:58+02:00
tags:
    - pester
aliases:
    - /growing-an-open-source-project-the-pester-story

---


[![test image](https://scottmuc.com/images/blog/github-pester.png)][pester]

### In the Beginning

In 2010 I was on a project that focused on build and release software for a bunch of .Net projects (and a few Java
projects). Our "glue" was written in PowerShell because of the wonderful remoting capabilities and it's integration with
Windows automation. Our code was untested and our code base was growing over time because our tools proved to be useful and
we were getting more and more feedback. Thus, the code required change more frequently. Without the safety net of tests
for the most basic logic, our feedback loop was slow as it required us to run builds and deploys to adequately cover our
bases.

It was January 1st, 2011. I was a touch hungover from the previous night and was chilling in my apartment in Calgary.
The difficulty in changing our PowerShell code was bothering me enough that I got a bunch of coffee from
[Cafe Beano][beano] and attempted to make an `rspec` like testing tool for PowerShell. The next day I was able to make a [commit][first-commit]
with actual working code! A few days later I was able to [add tests][poweryaml-tests] to PowerYaml. That's how
[Pester][pester] was born.

[first-commit]: https://github.com/pester/Pester/commit/a1d6a0e01f58375175ed090647ab8245a049f1a6
[poweryaml-tests]: https://github.com/scottmuc/PowerYaml/commits/master?page=1
[pester]: https://github.com/pester/Pester
[beano]: http://www.yelp.com/biz/caffe-beano-calgary

In the next couple months I integrated Pester into our project. It was a real superficial integration as we didn't have
too much time to retroactively add a lot of tests. So for me, my consumption of Pester was a few months in early 2011!

### The Early Days (2011 - 2012)

Over the course of the year, the tool seemed to be adopted by a few. A [friend in Vancouver][martin] apparently had begun using it
in one of his projects. He's also the first [contributor][martins-commits] to the project! It wasn't until the next year
that [Rob Reynolds (aka ferventcoder)][ferventcoder] used Pester in the [Chocolatey][chocolatey] project. That was my
first hint that there's potential in this project.

In 2012 there wasn't much going on until [Matt Wrock][matt] added mocking support and contributed a whole lot of fixes.
It was then that Pester got a new core maintainer. It was very cool to meet him during my [career break][career-break]
in Seattle and give him a proper thank you for the work that's done. I believe he also wrote the [earliest post][matts-post]
on Pester. I then noticed a comprehensive post on the [endjin.com][endjins-posts] blog website.

Seeing other people use Pester was a massively ego boosting thing. I'm afraid I'm one of those people in the software
community that does feed off of this kind of thing. For some reason it made me feel validated as a professional in the
software community. Despite the pleasure I got from Pester being associated with me, I decided to move Pester from my
personal github user to a Pester organization. I'm really happy I did that, because it made it feel more like a
community project than a personal one. A [google group][pester-group] was eventually created too as to not clutter
Github with conversations in the issue tracker.

Later on I would find out that it's been used in several [ThoughtWorks][thoughtworks] projects! These use cases weren't
even initiated by me. Someone wanted to test their PowerShell code, found Pester, and started using it. They would find
out I also worked for ThoughtWorks and would ping me on our internal chat and ask me questions.

[martins-commits]: https://github.com/pester/Pester/commits/master?author=mrtns
[martin]: http://mrtn.nrd.io/
[thoughtworks]: http://www.thoughtworks.com/
[ferventcoder]: https://twitter.com/ferventcoder
[chocolatey]: http://chocolatey.org/
[matt]: https://twitter.com/mwrockx
[matts-post]: http://www.mattwrock.com/post/2012/11/15/Unit-Testing-Powershell-and-Hello-Pester.aspx
[endjins-posts]: http://blogs.endjin.com/?s=pester
[career-break]: /blog/categories/career-break/
[pester-group]: https://groups.google.com/forum/#!forum/pester

### The Year of Change (2013)

While I was living in India I had several months with some extra time that I used to update Pester. Throughout the year
I spent quite a bit of time maintaining the project. There definitely was a motive for this activity. I knew that
projects that have activity in the codebase are usually favoured by other devs. I still haven't used Pester on a project
since early 2011! There's no functionality that I could possibly need out of Pester but I carved out time to take care
of a few niggling design flaws and close some issues. This resulted in a [major release][major-release] which is pretty
much the last time I really did anything to the code. It was great to see another [quality blog post][quality-post] by
Johan Leino appear on the Web.

[major-release]: /powershell-pester-2-and-1-dot-2-released/
[quality-post]: http://johanleino.wordpress.com/2013/09/13/pester-unit-testing-for-powershell/

### The Year of Transition (2014)

[![test image](https://scottmuc.com/images/blog/pester-mention.png)](https://twitter.com/LogicalDiagram/status/461638955725631488)
2014 marked a large surge in Pester interest. During the PowerShell Summit it was rumoured that the creator of
PowerShell, [Jeffrey Snover][jsnover], mentioned Pester in one of his talks! Heck, I was even [interviewed][interview] for the
PowerScripting Podcast!

Another contributor by the name of [Jakub Jareš][jakub] was very active and soon became a core member of the Pester team. During
my travels on my [career break][career-break] I got to meet up with him in Prague! He also wrote a
[couple][jakubs-articles] great articles too.

Now it looks like there's another one joining in too. Welcome [Dave][dave] to the Pester team!

Looking at the logs of the Beta branch, it's wonderful to see my commit percentage diminishing. I've now only committed
30% of the code base!

[jsnover]: http://www.jsnover.com/blog/
[jakub]: https://twitter.com/nohwnd
[dave]: https://twitter.com/MSH_Dave
[interview]: http://powershell.org/wp/2014/03/23/episode-262-powerscripting-podcast-scott-muc-on-testing-with-pester/
[jakubs-articles]: http://www.powershellmagazine.com/author/jjakub/

### Lessons Learned

Though not a big project, I can say I learned a lot maintaining this open source project. There's something very
different than working with a colocated team and random contributors around the world. Colocated teams develop unwritten
rules (a culture) and there's less debate over small issues. I didn't realize how much brace alignment, trailing
whitespace, tab stops, and other petty things would annoy me. What's harder is that I didn't want to complain about it
to the community afraid that I would come across as an OCD control freak. My biggest mistake with Pester was thinking
too much on how it reflected upon me, and not what is good for the community. If I had lived and breathed Windows
projects over the last few years, I think that would have been different. I'm happy that Pester has a few excellent
maintainers now and that I've participated in an open source project from first commit to stepping away.

Another lesson was that empowering would be contributors was an investment well worth doing. It's hard to say, but I
feel that doing this is what's enabled me to lessen my involvement over time.

This talk also resonated with me a bit (the last 4-5 minutes):

{{< youtube UIDb6VBO9os >}}

### Stepping Down

The adventure of watching a project grow and have a life on its own has been a wonderful experience. Looking at the
project history and correspondance it's pretty easy to see that I'm not nearly as active. That's really thanks to Jakub
and Dave doing nearly all the maintenance.

That being said, I'm "officially" stepping down and retiring from Pester. This will be done by unsubcribing from any
e-mail notifications (the google group, and the github project).

![test image](https://scottmuc.comhttp://pesterbdd.com/images/Pester.png)
I would like to shout a big thank you to everyone who made this project so wonderful. Especially to [Manoj
Mahalingam][manoj] who did some great work in the early days and really helped me get this project off the ground.
Another shout out to [Stefan Vant][vant] for the excellent logo!

[manoj]: https://twitter.com/manojlds
[vant]: http://vant.ca/


