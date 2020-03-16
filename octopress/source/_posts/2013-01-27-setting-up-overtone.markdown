---
layout: post
title: "Setting up Overtone"
date: 2013-01-27 21:11
comments: true
categories:
---
Every now and then I like to sit in a coffee shop on a Sunday and focus on writing some code. This weekend I had the pleasure of spending this geeky time with a colleague, [Chris Ford](http://literateprogrammer.blogspot.in/)([twitter](https://twitter.com/ctford)).

Chris has gotten me interested into an open source toolkit for creating sounds and making music called [Overtone](https://github.com/overtone/overtone). Here's a talk about demonstrates functional composition using clojure through music and overtone.

{% youtube Mfsnlbd-4xQ %}

That talk made me want to play with overtone. Here's the diary of my workstation setup.

## Leiningen

{% img right http://leiningen.org/img/leiningen.jpg %} Thinking I know what I'm doing I say "So, Overtone is in Clojure, right? So I'll first install that.". To which he replied "No, you need to install [Leiningen](http://leiningen.org/).". What?! I want to install clojure, why do I need this oddly named tool? Already my world has been turned upside down.

Leiningen is a package manager for Clojure. I can't say I understand completely what it's doing, but when I did a `sudo port install leiningen` I saw Maven was being installed. Chris calmed me down and said that `lein` hides the pain so I don't have to worry about it.

Sure enough, the install works. I run `lein new mucmusic` and a simple project structure is setup for me. I updated the `project.clj` to refer to the 1.4.0 version of Clojure.

I then run `lein repl` and see the package manager at work. It downloads Clojure for me and I'm in a repl. I try the most basic thing I know `(+ 1 1)` and it works! I can only imagine what's going on behind the scenes. I find it interesting that a package management tool is responsible for obtaining the language you're going to develop in. It makes a whole lot of sense actually. It's sort of like [virtualenv](http://www.virtualenv.org/en/latest/), but not quite.

## Clojure

{% img left http://clojure.org/space/showimage/clojure-icon.gif %} Since I'm still coming to grips on writing code in Clojure I can't comment on the language, but what I really like so far is how accessible the documentation is. Simply evaluating `(doc +)` will print the documentation for '+' function  and `(source +)` prints out the source code. The REPL made it very easy to get started.

One thing that tripped me up was using a namespace I had defined. I learned the the ' character is used to indicate the following text is a symbol and not to be evaluated. Also, the first project I created had an underscore in it. In order to use that namespace I had to use a hyphen when using it.

## Overtone

{% img right http://overtone.github.com/media/logo.png %} Happy that I can write and evaluate Clojure code it was time to setup Overtone. This part was really nice. All I did was update the `project.clj` file and added the overtone dependency. The next time I ran `lein repl` it downloaded overtone and set everything up for me.

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
    -------------------------
    overtone.live/sin-osc
    ([freq phase])

      Sine table lookup oscillator 

      [freq 440.0, phase 0.0]

      freq  - Frequency in Hertz 
      phase - Phase offset or modulator in radians 

      Outputs a sine wave with values oscillating between -1 and 1 similar to 
      osc except that the table has already been fixed as a sine table of 8192 
      entries.

      Sine waves are often used for creating sub-basses or are mixed with 
      other waveforms to add extra body or bottom end to a sound. They contain 
      no harmonics and consist entirely of the fundamental frequency. This 
      means that they're not suitable for subtractive synthesis i.e. 
      passing through filters such as a hpf or lpf. However, they are useful 
      for additive synthesis i.e. adding multiple sine waves together at 
      different frequencies, amplitudes and phase to create new timbres. 

      Categories: Generators -> Deterministic
      Rates: [ :ar, :kr ]
      Default rate: :ar

No wonder people love programming in clojure! This idea of the documentation living in the
code is wonderful. I'm generally used to googling for a function name, but this makes that step
unnecessary.

## VimClojure

Chris, like myself, also uses [Vim](http://www.vim.org/). So first thing I wanted to do is get some syntax highlighting. Chris points me to [vimclojure-easy](https://github.com/daveray/vimclojure-easy), but I didn't want to replace my vim configuration. I already maintain my vim in a tidy and versioned way. Looking at the this repo gave me the direction I needed to set this up. I cloned [VimClojure](https://github.com/vim-scripts/VimClojure) into my `~/.vim/bundle` directory (and removed its .git directory). I launch vim on a clojure file and I have syntax highlighting!

## lein-tarsier and nailgun

Thinking I was complete I tried some of the vim shortcuts to evaluate some clojure from within vim. Then I realized I'm
not quite done (feels like my visit to the [FRRO](http://indianfrro.gov.in/frro/). This is now where the JVM comes in.
Clojure is a JIT compiled language. In order to compile, it needs a JVM. Starting up a JVM all the time is a very
slow and painful process.

Lein-tarsier (tarsier is the type of creature on the [O'Reilly Vim book](http://shop.oreilly.com/product/9780596529833.do)
is a lein plugin that starts up a VimClojure server. To install this, I just added it to my lein profile. It didn't work
at first, but that was because MacPorts installs version 1.6 of lein. Nailgun was some C code that you have to download a compile and make reference to it in your vimrc. I'm sort of lost on how all these components work together.

A quick `lein upgrade` got me to version 2.0. Then an execution of `lein vimclojure` and the server is up and running.

## Summary

After installing a bunch of oddly named things I was able to play the pieces from Chris' talk in vim.

There were some details that I left out (and forgot). I don't think I would have gotten this far without help from Chris
and the considerable documentation that's in the README files and wikis. Part of me also thinks that it shouldn't
require this much work. As a side project, I'm going to look into pulling all of these steps into a Vagrant
configuration management setup. I think it would be nice to go `vagrant up overtone` and have a working overtone
development environment.



