---
layout: post
title: "The Most Over-Engineered Thing I Have Made"
date: 2013-06-25 08:13
comments: true
categories:
---
The other day I overheard some interns boasting about their complicated shell setup. They felt pretty good about
using tmux, zsh, and aliasing the heck out of everything. One of them had to re-map copy and paste because they messed
that up, but they felt awesome when they figured out a clever work around.

We've all been there at some point. It's something that you grow out of over time. At least I hope so. Here's a story of
the most complicated thing I've made and I'll start with a network diagram of my basement suite circa 2003:

[{% img center /images/blog/complicated-network.png 600 %}][network-diagram]

## My Contribution to Over-Engineering

Hmm, after looking at that network diagram again, I'm now thinking about how complicated I've made a lot of things!
Anyways, the most complicated thing I've made was a music playing setup for my home. Doesn't sound like it should
be complicated, right?

Here are some of the requirements:

* Rebooting my laptop doesn't stop the music. I was quad booting (that's another story) and would switch frequently
* Control can be from any computer (I did some NAT translation so I could control it from outside the house too)
* The machine playing the music had to be quiet (PCs those days were much louder)

### Architecture

> The result was an mp3 playing system that required 5 computers in order to work.

[{% img right /images/blog/complicated-rack.jpg 300 %}][rack] First there was the **file server** that contained
all my music. It would expose files to the network via Nfs and Samba. I was pretty meticulous when it came to
organizing my files. Every artist had a directory, and every album was a sub-directory. In case you're wondering,
I organized artist names as *Last Name, First Name*. All the mp3s were created by ripping my personal CD collection.

Second was the **music player** box. This was an underclocked machine with no hard drive. This way I was able to run it
without any fans except for the power supply. It used PXE boot load the operating system and would Nfs mount it's filesystem.
The only thing running on it was [music player daemon][mpd] which played mp3s that were also Nfs mounted.

So since I required PXE boot, I setup a **network config** server. It was my local DNS, DHCP, and TFTP host. I
configured DHCP to respond to the MAC address of the music player box and provide it a TFTP location to fetch its kernel.
I had a DNS zone setup that handled internal .muc-central.com records (eg: jukebox.muc-central.com).

Those 3 servers were sufficient to have a working music playing system, but it didn't have a user interface. This
required the addition the **web server** which ran Apache and PHP to host [PhpMp][phpmp] that connected to the music
player boxes `mpd` service.

Now with a UI, you can start playing songs, and update a playlist from any machine in the house. The 5th machine is the
**client**, which was my laptop at the time. This was before the era of smart phones, but I imagine with some tweaks, a
mobile web browser would have worked just fine too.

## The Result

[{% img left /images/blog/complicated-livingroom.jpg 515 %}][livingroom]
[{% img left /images/blog/complicated-ui.jpg 515 %}][desktop-screenshot]

<span style="color: #eaeaea;">hack to break the floats</span>

The machine that plays music is in a closet near the entrance into my apartment. Then I setup some shelves to
put my [Altec Langsing ACS48][acs48] speakers on top of. The volume on the speakers was set to max so the volume could be
controlled through the software.

The web interface wasn't fancy but I still prefer it to many of the applications I've used recently. It's a basic PHP
app called [PhpMp][phpmp]. What I liked about it the most was the playlist was easy to modify. My friends would browse
the music collection and click a link to add a song to the playlist. The only picture I can find is an old desktop
screenshot where you can see the playlist in the bottom left of the screen.

## Lesson Learned

Although I would never set something up like that again, I'm glad I did it. It was an excellent learning experience and
it actually was one of the best digital music setups I had. It also was the gateway to me starting to write technical
[articles][first-article] online 10 years ago! While listening to the interns, it came to me that this is simply a right
of passage. I don't think I would be where I am today if I did the simplest thing to solve all my problems. One thing it
did make me think about is that I probably should contrive complex solutions as a way to practice my understanding
of of complex systems.

Things I learned by doing this:

* PXE booting
* Diskless computing
* TFTP
* Non trivial DHCP

[mpd]: http://www.musicpd.org/ "Music Player Daemon Website"
[acs48]: http://www.amazon.com/Altec-Lansing-ACS48-Computer-Speakers/product-reviews/B00000JBJF "ACS48 Reviews"
[phpmp]: http://mpd.wikia.com/wiki/Client:PhpMp "PHP Music Player"
[first-article]: http://muc-central.com/articles/index.php?file=FreeBSD%20Jukebox "My First Online Article"
[network-diagram]: /images/blog/complicated-network.png "Network Diagram"
[rack]: /images/blog/complicated-rack.jpg "Server Rack"
[livingroom]: /images/blog/complicated-livingroom.jpg "Living Room"
[desktop-screenshot]: /images/blog/complicated-ui.jpg "Screenshot"
