---
title: "Installing Windows Home Server On Old Hardware"
date: "2008-04-27T23:08:00Z"
---
For the last 7 years I've been running a [FreeBSD](http://www.freebsd.org/ "FreeBSD") server at home for all my file serving needs. It's always been solid and never caused problems. It was the gateway into Unix knowledge for me and I loved the whole experience. A fellow 8-ball competitor revealed that he's a sales agent for Microsoft and kindly gave me an OEM copy of [Windows Home Server](https://en.wikipedia.org/wiki/Windows_Home_Server "Windows Home Server") (WHS)! I reluctantly decided to retire the ol BSD box and try out WHS.


The first thing I'll discuss is the hardware. The BSD fileserver did not have high hardware requirements so it never got upgraded often. Here's what I had to work with:

* Tyan Trinity 400 Motherboard
* Pentium III 866MHz
* 256MB Ram
* 3 Maxtor Hardrives (250GB, 80GB, 40GB)
* Promise Ultra 66 Controller
* Netgear NIC

My first stumbling block was that I did not have a DVD drive. A few months ago I did a purge of almost all my computers (over a dozen) so my spare parts closet was a bit sparse. Fortunately my kind neighbor who lives about me loaned me her HP box and I was able to pull the DVD drive from that and use it.

With the box in DVD bootable condition I was able to get the installer going until I ran into my second bump. It complained that [ACPI](http://en.wikipedia.org/wiki/Advanced_Configuration_and_Power_Interface) support was not found. Fortunately that was a simple BIOS setting. I did a lot of messing around with the BIOS a while back and was unsure what changes what could complicate matters so I reverted the BIOS back to its defaults and I got past that point without further issues.

The next block was the installer would not continue with less than 512MB of RAM. The Tyan motherboard only has 3 DIMM slots and I couldn't find any sticks larger than 128MB! Fortunately I remembered that my old dev server had 512MB in it so I was able to scoop a couple 256MB DIMMs from it.

A quick reboot later and I now discover that the installer does not have drivers for my Promise card. Instead of trying to find drivers I decided to ditch the card. Sure I'll only have ATA 33 hardrive speeds but who cares since my network will be more of a bottleneck anyways and I'm the only user too. Hopefully this decision doesn't come and haunt me later.

![Windows Home Server install](https://lh3.googleusercontent.com/lfxvx0vLbC5clQKfaLh1tvOdca0G8LvMag6Y4FW3ti8fPBih2OTV9sxtL4FHlT7uHpjHyInoxO4FksIiArIomLv3ckQ3xmIKU6JXja12X49dDVgl_UY9ok3Fg2CEvK48buAINzJzjOvArBsexS6YBjgrPv0PX1pvytia0v0ehVSk0dNWM8M8B4xtrX1Ey8yn9FPyWSxFIYqZ2Q9ec-VrKzg0N2Ci8scRMQi6mdj2Jqa5FZXdJJB5K3TdgIIno7d62SgP3OHNRI3_lPu-GaJb82hKQciCOhh7nGomX6hZaxDcoGCFLRyi7EEhtghy94MWJEt4XDhMZLWPWDpcKwgWw9_6w0woUt0jeFq1312CPuBBHwEgMGFh5O5bcgrXxaGqSaRkphRzCeVHbYrYoz5PLg1YuAahVd_SVQaC5lGG1Wqzm0wNrV05pNAgFy_su1pZuZIZ2vYr8Eo5Q8v8DzUiXujkFsYSfQM4rGSC58W1tWAiSSnGnEewfk9SP0D1uUW-yIT5ntZ1tu7ukbo0KTO1gEyATJvGYiIWpko_4zc4sCumPCgGHeCjJ2AaeHYIcaSbTxdLNBgGhcoxfwcin-6H0T_vCox9m0TEj2m7DNQx6vPNZi7kW-ogO4z0BacT7n-BLEl4uZMbGAGbqxtcP2jO5O-ppKGPYE4lAkjtGABQeZW8Vdidq4Fa3NOhk8X_uaslQqhOFrAyvXpnE5eWMT4otHM2MsvjT8lfNPWffc9X6OQWKD6cpXyBcj0=w400-h300-no)

Finally I am at a point where the installer gets going! Figuring that this was going to take a while I take a break and watch [Adam Scott](http://www.pgatour.com/golfers/024502/adam-scott/) win the Byron Nelson Championship. Go Adam Scott!

A few hours later it's now booting into Windows Server 2003 and it's getting ready to install the Home Server component. For some reason the final portion of the install keeps on crashing without much of an error message. Not too sure what's going on I try to 'Windows Key' my way back to the desktop to see what's going on. Unfortunately I can't recall where I saw this, but I noticed a message displaying a 'getting updates' kind of message. I got into the Device Manager as fast as I could and noticed the *Netgear NIC* is not being detected. I found my usual standby *3com 3c905b NIC* in the closet and that worked like a charm and the crashes stopped.

Everything then went smoothly! *39* Windows Updates had to be installed but that's to be expected. I have the box back on my rack and am currently copying over the gigs of content back to the server. All in all it wasn't a bad install experience but I got quite lucky with a lot of my troubleshooting. I'm sure all of those problems would have been averted had I worked with modern hardware.

Anyways, that's my tale of installing WHS on older hardware. I hope this encourages people to recycle their old computers for this task rather than throwing them away. I always love getting the most out of old computer hardware that people think is useless. Heck, to me a PIII 866 still seems like a pretty decent machine!

![My rack](https://lh3.googleusercontent.com/GA0USczzaCv2G1b76Sq7TlGBTaQIWdfGkCIK1s62phw0Vxmt7zVxrpveCWvDcxzYuGego6-F6S8VdRU3bJWCqrpnfK2PHCt0OGcbZPPrj2sg9bsW2yw-CQlmvrabDM79PP-qIkdiJUO7qOCRTUcUdA7KNPcSJzkKtWBS2yMkM_nonweY9bn_wRA4O0SFr3k5k6QmYLUtu0S6FFfPWQ2xoV4FgW1-MyTFVdv17FcGGJ0O2y3wEbyBkn47NT7wcFm8f7NfgbwuKYj5yo9j9xF0qH5s-gHj1Jq8vaXrgbNr5CDoVe99HlnXVCukKCJGgfDNKKd1CtdbvRaZ3NB02UqlJsq43e2Ar0YDUpeOddejtj--WOqyB1LEokHd5bV4lOev-QZXRqku-XZa6lnsPKXwN1D4WVH6GEokPGkzLXXpI-t-XbmdVPPboZ0Nd_RqXcghCU8RFvsW3xj8OME0Vvg1l40yTWRiUmK2a2wKjFqt2sYRe0JFFv9jQupmrbZSN29S6eK4KDnEHgCqFYjujOS0gyXrOMxe-lzx3vnrNvWWs8FmPewj_h7KuVlcNNcjdxHnudiyUDsWHJs0JFz0RJkiIXCnuLvTura49ec9wzf-ndBKaR_Rt-wbr1mRylCE5jk3de30HRYOSZBD6zAIk3p0z45WWGqgXEbdzc8GkZJ7FJNXdf7Wv96tHxVxpfv0hQEZOCjA3Mof158UJDCX0Z86vVZlFbSzlaBq4uICDEmRhsGQBDv7ivixe_M=w350-h522-no)
