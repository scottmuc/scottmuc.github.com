---
title: Book Review - A Mind for Numbers by Barbara Oakley
date: 2014-10-26T13:35:53+05:30
tags:
    - book
    - review
aliases:
    - /book-review-a-mind-for-numbers-by-barbara-oakley

---

[{{< figure src="book.jpg" >}}][book]

Over the last 4 weeks I've been taking the [Coursera][coursera] course [Learning how to Learn][course-homepage] on the
recommendation of my friend [Srini][srini-twitter]. The 2nd assignment in the course is to create a presentation or some
form of media to present an understanding of the topics. An accompaniment to the course is the book,
[A Mind for Numbers][book] by [Barbara Oakley][barbara] (the primary course instructor), which is a short read but very
insightful and full of references. I'm thoroughly enjoying the course and love the concept around the assignment. This
blog is my way of trying to synthesize ideas that are flying around in my head, which is exactly the purpose of the
assignment.

[coursera]: https://www.coursera.org/
[course-homepage]: https://www.coursera.org/course/learning
[srini-twitter]: https://twitter.com/sragu
[barbara]: http://www.barbaraoakley.com/

Over my education and career, I feel that I have developed some pretty good habits. The book has confirmed some of these
habits and given me an understanding around why they work. It's also able to point out some flaws that I tend to
down-play, but now realize that they are habits I need to deal with.

[book]: http://www.penguin.com/book/a-mind-for-numbers-by-barbara-oakley/9780399165245

## Core Concepts

The material in the book (and the course) were broken up into easy to consume chunks (see what I did there?).
I can't really cover all of it here but these are the 3 concepts that really stuck with me.

### Chunking

Chunks can be thought of as an abstraction around a collection of specific thoughts. When I speak to another developer
and I say that this code *makes a web request*, that is a chunk of knowledge that I've developed. I could break that
phrase down and say: establish a TCP connection with the remote server, using the HTTP protocol send a GET request
message, the server will receive this and route it to an appropriate controller, the application code will send an HTTP
response, the client will take the results and render the HTML to the browser. Woah, that was a lot of detail, and
that's one example of why it's important to chunk knowledge.

{{< imglink src="chunking.png" href="chunking.png" width="800x" >}}

It wasn't mentioned in the book, but a lot of what I read on chunking looked like reaching the **competent** level on the
[Dreyfus Model of Skill Acquisition][dreyfus-model]. I'm now curious if **diffuse mode** is an essential part of
becoming **proficent** or an **expert** in a particular topic.

This section of the book has made me realize some flaws in the way I learn and I feel I fall into the trap called
**illusion of competence**. This can be caused by [poor learning practices][10-rules] (below), and I know I am guilty of many of
these. Speaking of the Dreyfus model, **Erik Dietrich** wrote an article about the [Advanced Beginner][advanced-beginner],
which is another great way of explaining the **illusion of competence**.

> * Passive rereading
> * Letting highlights overwhelm you
> * Merely glancing at a problem's solution and thinking you know how to do it
> * Waiting until the last minute to study
> * Repeatedly solving problems of the same type that you already know how to solve
> * Letting study sessions with friends turn into chat sessions
> * Neglecting to read the textbook before you start working problems
> * Not checking with your instructors or classmates to clear up points of confusion
> * Thinking you can learn deeply when you are being constantly distracted
> * Not getting enough sleep
>
> Barbara Oakley -- [Learning How (Not) to Learn][hntl]

[hntl]: http://blog.coursera.org/post/93424900982/learning-how-not-to-learn

I've been lucky to accidentally avoid some of these habits. As a computer savvy person for the last 25 years, I still
prefer taking my notes on paper. I've always stressed sleep too. I'm definitely guilty of passive rereading, but I'm
getting much better now. Solving similar problems is another one I need to catch myself doing.

[10-rules]: http://blog.coursera.org/post/93424900982/learning-how-not-to-learn
[dreyfus-model]: https://en.wikipedia.org/wiki/Dreyfus_model_of_skill_acquisition
[advanced-beginner]: http://www.daedtech.com/how-developers-stop-learning-rise-of-the-expert-beginner

### Diffuse and Focused Modes of Thinking

**Focused Mode** - When the brain is in this mode, it's busy building up foundations of material for access later. I
like to think of it as going to the gym. You're exercising ideas and building up their strength. If you don't practice
them deliberately, they can lose strength. If you don't focus on concepts with discipline, you don't effectively
chunk the concepts. When you get up from your task and realize that a lot more time has gone by than you had expected,
then you know you've been in the focused mode.

**Diffuse Mode** - This is where synthesis of multiple ideas (or chunks) happens. Our brain is good at continuing effort, solving
problems, while our bodies are performing activities like going for a walk, running, socializing, and cooking. This is
why when struggling with a problem late at night, sleeping on it, usually results in some new insights in the morning.
You know when you have yesterdays problem figured out in the morning during your shower? That's the diffuse mode kicking in.
Using the workout analogy, this is like using your in-shape body to enjoy life: traveling, playing with your children,
fixing up your garden.

{{< figure src="diffuse-focused.jpg" width=400 >}}

What's important to understand is that both of these modes can't operate at the same time. I [even wrote][flow-post]
about this a while ago and this book has given me the knowledge to better understand what I was observing when I wrote
that post. As a developer, the state of flow (deep focused mode) has its appeal, but breaking that up so that you can
piece together your new knowledge is equally important.

[flow-post]: http://scottmuc.com/going-against-the-flow/

Sometimes you'll hear someone say that they are a great multi-tasker. This usually results in an inability to focus!
If someone is a multi-tasker, then they are taking time away from the **focused mode** of operating, and thus fall into
the many ways [not to learn][10-rules].

### Procrastination and Habits

The topic of procrastination was very informative. The explanation of how dopamine effects long term beneficial actions
([reference][dopamine-reference]) is insightful and can be used to influence how we develop habits. The book is full of
ideas on how prevent procrastination and here's my list of tactics:

* **Zero Notifications** - this is an easy one for me as I already try to live a notification free life. My phone is
  almost always on silent unless I'm *expecting* a call or message. All my applications on my phone and computer have
  notifications turned off. It's very hard to think of any reason why any notification would need immediate action.
* **[Pomodoro Technique][pomodoro]** - I've been a fan of this process for a long time. Procrastination doesn't stand
  a chance when the timer is on. I like to use [pomodoro one][pomodoro-one] on my laptop. There are 4 components of
  developing a habit, and the first one is the **cue**. Starting the timer triggers me to go into a focused mode.
* **[Self Control][self-control]** (pictured) - This application does a machine level block of the hostnames that you specify for a
  specific time period. It runs outside of the web-browser so it'll block applications that talk to these services and
  even killing the process or rebooting your machine won't get the hostname working again. I will use this every time I'm doing a pomodoro.
  
	{{< imglink src="self-control.png" href="self-control.png" width="400x" >}}
* **Process over Product** - Unless the task is very very small, the time allocated to a single pomodoro should be
  process based. As of right now, I have on my task list the following: "2 pomodoros of coursera project". This keeps
  the focus on the activity and now the anxiety around the outcome. I don't know how long it's going to take to complete
  this post, but I know if I chip away at it with time-boxes of focus it'll be completed before I know it.

I like reading about other people's techniques and thought [James Clear][james-clear] has a good
[blog post][james-clear-post] and a good list. On my reading list is [Daily Rituals][daily-rituals] which looks very
interesting. One thing that seems clear, is that many productive people have developed a system and it's important to
discover my own.

[daily-rituals]: http://masoncurrey.com/Daily-Rituals
[james-clear]: http://jamesclear.com/
[james-clear-post]: http://jamesclear.com/multipliers
[pomodoro]: http://pomodorotechnique.com/
[pomodoro-one]: http://lifehacker.com/pomodoro-one-is-a-free-lightweight-pomodoro-timer-for-1626504270
[self-control]: http://selfcontrolapp.com/
[dopamine-reference]: http://www.psychologytoday.com/blog/intrinsic-motivation-and-magical-unicorns/201206/procrastination-and-dopamine-receptor-density

### Summary

I began reading the book before taking the course and immediately benefited from a process that is described in the
book. The book on its own is a fantastic resource and is an excellent compliment to the course. I really think taking the
course at the same time instilled a lot of the information far better than if I had only just read the book. I was
surprised when I finished the book because my Kindle reported that I was only 65% complete when I was at the end. Well
that's because the remainder is a massive collection of references of all the research that's been done on all of the
topics! It shows that the content is truly backed by science which is refreshing given all the anecdotal type of
information that's published on the Internet.

Reading the book and taking the course was very meta; as I learned techniques from the
book, I applied them to the course and vice versa. I get the feeling this is going to be a book that I read on a yearly
basis (spaced repetition) to continuously improve these concepts. This post could have been a lot longer, but I'd rather
encourage you to read the book yourself and practice the exercises at the end of each chapter.

#### Other Resources and References

* [Pragmatic Thinking and Learning][prag-thinking] which is geared more towards software developers.
* [HealthyHacker Podcast][healthyhacker] which talks a lot about memorization techniques and developing positive habits.
* [Better Explained][better-explained] Kalid Azad does an exceptional job of explaining things. In an interview in the
  course he describes his method of ADEPT (analogy, diagram, example, plain English, technical).
* [Slow Tech][slow-tech], a talk about distraction and how to be focused. After the course, the points in this video
  appear well connected with the material.
* [John Cleese on 5 Factors to Make Your Life More Creative][john-cleese]. A comedic genius explains how he uses time
  boxing to focus on his craft.

1 "diffuse vs focused mode" picture http://tdlc.ucsd.edu/educators/educators%5Fwebinar%5Foakley5%F031213.html<br />
2 "chunking" picture p.56 A Mind For Numbers, Barbara Oakley<br />

[better-explained]: http://betterexplained.com/
[prag-thinking]: https://pragprog.com/book/ahptl/pragmatic-thinking-and-learning
[healthyhacker]: http://www.healthyhacker.com/
[slow-tech]: https://www.youtube.com/watch?v=EzpX0TLKS9Q
[john-cleese]: http://www.brainpickings.org/2012/04/12/john-cleese-on-creativity-1991/


