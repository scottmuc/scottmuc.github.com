---
title: Setting up Overtone
date: 2013-01-27T21:11:00Z
tags: []
aliases:
    - /setting-up-overtone

---

Every now and then I like to sit in a coffee shop on a Sunday and focus on writing some code. This weekend I had the pleasure of spending this geeky time with a colleague, [Chris Ford](http://literateprogrammer.blogspot.in/)([twitter](https://twitter.com/ctford)).

Chris has gotten me interested into an open source toolkit for creating sounds and making music called [Overtone](https://github.com/overtone/overtone). Here's a talk about demonstrates functional composition using clojure through music and overtone.

{{< youtube Mfsnlbd-4xQ >}}

That talk made me want to play with overtone. Here's the diary of my workstation setup.

## Leiningen

![test image](https://scottmuc.comhttp://leiningen.org/img/leiningen.jpg) Thinking I know what I'm doing I say "So, Overtone is in Clojure, right? So I'll first install that.". To which he replied "No, you need to install [Leiningen](http://leiningen.org/).". What?! I want to install clojure, why do I need this oddly named tool? Already my world has been turned upside down.

Leiningen is a package manager for Clojure. I can't say I understand completely what it's doing, but when I did a `sudo port install leiningen` I saw Maven was being installed. Chris calmed me down and said that `lein` hides the pain so I don't have to worry about it.

Sure enough, the install works. I run `lein new mucmusic` and a simple project structure is setup for me. I updated the `project.clj` to refer to the 1.4.0 version of Clojure.

I then run `lein repl` and see the package manager at work. It downloads Clojure for me and I'm in a repl. I try the most basic thing I know `(+ 1 1)` and it works! I can only imagine what's going on behind the scenes. I find it interesting that a package management tool is responsible for obtaining the language you're going to develop in. It makes a whole lot of sense actually. It's sort of like [virtualenv](http://www.virtualenv.org/en/latest/), but not quite.

## Clojure

![test image](https://scottmuc.comhttp://clojure.org/space/showimage/clojure-icon.gif) Since I'm still coming to grips on writing code in Clojure I can't comment on the language, but what I really like so far is how accessible the documentation is. Simply evaluating `(doc +)` will print the documentation for '+' function  and `(source +)` prints out the source code. The REPL made it very easy to get started.

One thing that tripped me up was using a namespace I had defined. I learned the the ' character is used to indicate the following text is a symbol and not to be evaluated. Also, the first project I created had an underscore in it. In order to use that namespace I had to use a hyphen when using it.

## Overtone

![test image](https://scottmuc.comhttp://overtone.github.com/media/logo.png) Happy that I can write and evaluate Clojure code it was time to setup Overtone. This part was really nice. All I did was update the `project.clj` file and added the overtone dependency. The next time I ran `lein repl` it downloaded overtone and set everything up for me.

In the repl I loaded overtone via `(use 'overtone.live)` and here's the output:

    reply.eval-modes.nrepl=> (use 'overtone.live)
    --> Loading Overtone...
    --> Booting internal SuperCollider server...
    --> Connecting to internal SuperCollider server...
    --> Connection established

        _____                 __
       / __  /_  _____  _____/ /_____  ____  ___
      / / / / | / / _ \/ ___/ __/ __ \/ __ \/ _ \
     / /_/ /| |/ /  __/ /  / /_/ /_/ / / / /  __/
     \____/ |___/\___/_/   \__/\____/_/ /_/\___/

        Collaborative Programmable Music. v0.8

    Hello ThoughtWorks, just take a moment to pause and focus your creative powers...

Woah! Booting internal SuperCollider server? Is Clojure so powerful that I can run my own atom smasher?
Turns out the answer is unknown on that, but what `overtone.live` does is smartly start up
[SuperCollider](http://supercollider.sourceforge.net/) which is "an environment and programming language
for real time audio synthesis and algorithmic composition". Ahh! So overtone needs this to actually
create audio.

I continue to follow the [README](https://github.com/overtone/overtone) and make my first sound.

    (demo (sin-osc))

This code rewards me with the beautiful sound that is a sine wave. Curious as to what `demo` and `sin-osc`
do I load up the docs in the repl.

    overtone.live/demo
    ([& body])
    Macro
      Listen to an anonymous synth definition for a fixed period of time.
      Useful for experimentation.  If the root node is not an out ugen, then
      it will add one automatically.  You can specify a timeout in seconds
      as the first argument otherwise it defaults to *demo-time* ms. See
      #'run for a version of demo that does not add an out ugen.

      (demo (sin-osc 440))      ;=> plays a sine wave for *demo-time* ms
      (demo 0.5 (sin-osc 440))  ;=> plays a sine wave for half a second
    nil
    user=> (doc sin-osc)
    
