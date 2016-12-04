---
layout: post
title: "OSX Workstation Setup Automation"
date: 2016-10-11 15:45:48 +0100
comments: true
categories:
---

**DISCLAIMER**

**Try out at your own trisk. This setup is not meant to be copy+paste reusable. It's about keeping MY
workstation under source control, and this strategy may destroy the machine you wish to attempt this on.**

## Objective

IT infrastructure has gone done a long journey of automating all of the things. We have automated deployments, production
environments, and regression testing. We can make an API request to Amazon and get a server instances in minutes.

The developer workstation is another story. Often, the developer workstation is a finely crafted machine. Just like
production, your workstation should be treated like a [hotel room][phil-cattle-vs-pets] rather than an owned residence.

[phil-cattle-vs-pets]: https://twitter.com/pcalcado/status/759218156493795329

## My Workstation

Why bother with this for my personal machine? It serves as a good disaster recovery method. Using this automation for
the last several years (and several laptops + multiple laptop repaves) has given me a lot of confidence that I know what
is required to get my machine ready.

## Existing Tools

There are plenty of [tools](#existing-tools) that help with this process. What I am going to describe is yet another
approach, but I hope the simplicity of it will be attractive to those who have struggled trying to get other frameworks
to work.

## Starting From Scratch

Technically you cannot start from scratch because you need at least an existing OSX machine to create the installer USB
stick. To do this, I followed the instructions on Brooke Kuhlmann's [workstation automation][bkuhlmann-setup]\*
repository. Though, I'll probably copy this information into my own repository for reasons I'll explain later.

When following this process, I reformat the machines disk with an encrypted volume and always create my new machine with
a different username (to ensure my automation doesn't make too many assumptions on pathes).

\* *Similar steps worked just recently for macOS Sierra*

[bkuhlmann-setup]: https://github.com/bkuhlmann/osx#os-x-el-capitan

## There's No Place Like Home

{% img center /images/blog/gh-dotfiles.png %}

I've found that many people have [dotfile repositories][gh-dotfiles] that get installed by their scripting. This adds
a level of complexity and absraction that is not really necessary. Cloning straight into home removes a lot of moving parts.
All of the sudden all the abstractions and scripts that go into setting up your home directory disappear. 
If you can successfully transform home into a repository you have a simple and easy way of maintaining dot files,
and scripts to support your day-to-day activities. You've essentially already completed the first
step in getting your machine setup. Also, because it's a repository, when you delete a file, it goes away.
It's inherently an idempotent system!

One type of complexity that comes with using re-usable components is that there's almost always a shim that allows you
add your own customisations. For my personal machine, this is quite unneccessary and I want to minimize the number of
mechanisms used to actually configure my machine.

Rather than adding a `bash` plugin manager, I favour adding scripts to `~/bin` (which is added to `PATH`). This keeps
the dependency on `bash` quite small and hence leaves the `.bashrc` pretty lean.

So this is why [my workstation automation][osx-homedir] is cloned straight into `$HOME` on a fresh machine:

```
# Assuming my keys are loaded using the process described in the next section
sudo xcodebuild -license
git init
git remote add origin git@github.com:scottmuc/osx-homedir.git
git fetch --all
git checkout master
bin/coalesce_this_machine
```

[osx-homedir]: https://github.com/scottmuc/osx-homedir/
[gh-dotfiles]: https://github.com/search?utf8=%E2%9C%93&q=dotfiles

## Awkward OSX Settings

My previous iteration of this automation was 100% `sprout` based. If you look [closely][sprout-cmd] you will see that this
automation still invokes `sprout`. I've gone the route of using `sprout` as a collection of OSX configurations. I could
maintain some of these as shell scripts, but I feel ok having this additional complexity and added dependency.

[sprout-cmd]: https://github.com/scottmuc/osx-homedir/blob/9e9c6bee3e24b481c06ba0cfaffdb0e3b6ac93ff/bin/coalesce_this_machine#L14-L21

## What Could Possibly Go Wrong

Workstation automation can be extremely flaky and that's due to the nature of it fetching dependencies from all sorts
of locations. This drives home the reason for needing to run it frequently. I can pretty much guarantee that the code
that worked a month ago will not work today. Packages change name, or download locations change. This can result in a
lot of mistrust of the automation. This leads to the following...

## Duplication is OK

Choosing when to vendor something is an important decision to make around construction your setup. The more you vendor,
the more reproducable your automation will be. The more you vendor, the harder it is to keep up to date with changes.

Something that took me a while to come to terms with is the amount of duplication that will exist. There does not exist
one true workstation to rule them all, but it is important to figure out what level of "sameness" is needed for your
development needs. I have my personal workstation automation, but I also work with a team-wide system for workstation
automation. I believe the team level is the largest group you can realistically manage without creating something too
complex of manage. Your automation is going to be a somewhat handcrafted piece of art, but it's going to be a simple piece
of art (hopefully).

The [OSX installer][bkuhlmann-setup] steps is something I'd be willing to copy and paste into my own repository. The `vim`
[pathogen][pathogen] plugin manager is vendored as are the plugins (albeit as submodules, which means they could still
fail).

[pathogen]: https://github.com/scottmuc/osx-homedir/blob/master/.vim/autoload/pathogen.vim

## Summary

In the end, I have a setup I'm relatively happy with. It turns out that my implementation looks oddly similar to a
colleague of mine.

<hr />

## Appendix

1. <span id="existing-tools">**Existing tools**</span>
    * [Sprout][sprout-wrap] - Intially this was called `pivotal_workstation` but transformed into a slightly more
      configurable and modular system. It uses `chef` under the covers and is quite comprehensive.
    * [Boxen][boxen] - Created by GitHub, this tool uses `puppet` under the covers.
    * [osxc][osxc] - Not sure where it was born from, but it uses `ansible` underneath.

    All three of these strive for the same thing, but use different engines. It probably makes sense to use whichever one
    uses the Configuration Management tool you're most comfortable with. The problem though, is that these tools may add a
    bit more complexity than is really necessary.

    [Sam Gibson][sams-homepage] wrote a [great article][sams-setup] about this using [Babushka][babushka]. He believed the
    above styles are too complex and strived to find a simpler solution. I like it, but I felt that it could go a bit
    further.

    [Gerhard Lazu's][gerhards-setup] workstation setup. He replaced **sprout** with a collection to
    shell scripts that I may be tempted to migrate towards. But one thing at a time; that's an important principle around
    working on this kind of stuff. There are a lot of moving parts and it'll only prove itself out over time.

[sams-setup]: https://www.thoughtworks.com/p2magazine/issue08/babushka/
[sams-homepage]: http://www.ifdown.net/
[sprout-wrap]: https://github.com/pivotal-sprout/sprout-wrap
[boxen]: https://github.com/boxen/boxen
[osxc]: https://osxc.github.io/
[babushka]: https://babushka.me/
[gerhards-setup]: https://github.com/gerhard/setup

1. Since this way of setting up your machine depends on being able to clone git repositories as a first step, it means that
  you're going to need to have your keys. For this, I follow Tammer Saleh's excellent post on
  [building an encrypted usb drive for your keys][tammer-usb].

  Most other secrets goes into 1Password. At work we use LastPass and are able to integrate it into our scripting by
  running commands like `eval "$(lpass show 'Awesome/aws-secrets' --notes)"`. This way secrets are never recorded on
  disk in their unencryped form. I try to keep my memorized passwords to a minimum this way.

[tammer-usb]: http://tammersaleh.com/posts/building-an-encrypted-usb-drive-for-your-ssh-keys-in-os-x/

1. For fun, here's the complete output of the latest run:

    ```
    ~ ? coalesce_this_machine
    ~/workspace/sprout-wrap ~
    Rubygems 2.0.14.1 is not threadsafe, so your gems will be installed one at a time. Upgrade to Rubygems 2.1.0 or higher to enable parallel gem installation.
    Using rake 10.5.0
    Using awesome_print 1.6.1
    Using builder 3.2.2
    Using mixlib-config 2.2.1
    Using mixlib-shellout 2.2.5
    Using libyajl2 1.2.0
    Using hashie 2.1.2
    Using mixlib-log 1.6.0
    Using rack 1.6.4
    Using uuidtools 2.1.5
    Using diff-lcs 1.2.5
    Using erubis 2.7.0
    Using highline 1.7.8
    Using rspec-support 3.4.1
    Using mixlib-cli 1.5.0
    Using net-ssh 2.9.2
    Using ffi 1.9.10
    Using ipaddress 0.8.2
    Using systemu 2.6.5
    Using wmi-lite 1.0.0
    Using plist 3.1.0
    Using proxifier 1.0.3
    Using coderay 1.1.0
    Using method_source 0.8.2
    Using slop 3.6.0
    Using multi_json 1.11.2
    Using net-telnet 0.1.1
    Using sfl 2.2
    Using syslog-logger 1.6.8
    Using mini_portile2 2.0.0
    Using rufus-lru 1.0.5
    Using polyglot 0.3.5
    Using yajl-ruby 1.2.1
    Using thor 0.19.1
    Using minitar 0.5.4
    Using bundler 1.13.2
    Using chef-config 12.6.0
    Using ffi-yajl 2.2.3
    Using rspec-core 3.4.1
    Using rspec-expectations 3.4.0
    Using rspec-mocks 3.4.1
    Using net-ssh-gateway 1.2.0
    Using net-scp 1.2.1
    Using pry 0.10.3
    Using gherkin 2.12.2
    Using nokogiri 1.6.7.2
    Using treetop 1.6.3
    Using librarian 0.1.2
    Using chef-zero 4.4.2
    Using ohai 8.8.1
    Using rspec_junit_formatter 0.2.3
    Using rspec-its 1.2.0
    Using mixlib-authentication 1.4.0
    Using rspec 3.4.0
    Using net-ssh-multi 1.2.1
    Using specinfra 2.50.3
    Using foodcritic 6.0.0
    Using serverspec 2.29.1
    Using chef 12.6.0
    Using librarian-chef 0.0.4
    Using soloist 1.0.3
    Bundle complete! 3 Gemfile dependencies, 61 gems now installed.
    Use `bundle show [gemname]` to see where a bundled gem is installed.
    Installing build-essential (2.2.4)
    Installing dmg (2.3.0)
    Installing homebrew (2.0.5)
    Installing osx (0.1.0)
    Installing sprout-base (0.3.0)
    Installing sprout-osx-settings (0.1.0)
    Installing sprout-osx-apps (0.1.0)
    Password:
    [2016-10-11T22:06:06+01:00] INFO: Forking chef instance to converge...
    Starting Chef Client, version 12.6.0
    [2016-10-11T22:06:06+01:00] INFO: *** Chef 12.6.0 ***
    [2016-10-11T22:06:06+01:00] INFO: Chef-client pid: 14253
    [2016-10-11T22:06:20+01:00] INFO: Setting the run_list to ["sprout-base", "sprout-osx-settings::defaults_fast_key_repeat_rate", "sprout-osx-settings::function_keys", "sprout-osx-settings::screensaver", "sprout-osx-settings::set_menubar_clock_format", "sprout-osx-settings::dock_preferences"] from CLI options
    [2016-10-11T22:06:20+01:00] INFO: Run List is [recipe[sprout-base], recipe[sprout-osx-settings::defaults_fast_key_repeat_rate], recipe[sprout-osx-settings::function_keys], recipe[sprout-osx-settings::screensaver], recipe[sprout-osx-settings::set_menubar_clock_format], recipe[sprout-osx-settings::dock_preferences]]
    [2016-10-11T22:06:20+01:00] INFO: Run List expands to [sprout-base, sprout-osx-settings::defaults_fast_key_repeat_rate, sprout-osx-settings::function_keys, sprout-osx-settings::screensaver, sprout-osx-settings::set_menubar_clock_format, sprout-osx-settings::dock_preferences]
    [2016-10-11T22:06:20+01:00] INFO: Starting Chef Run for Meriadocs-MacBook-Air.local
    [2016-10-11T22:06:20+01:00] INFO: Running start handlers
    [2016-10-11T22:06:20+01:00] INFO: Start handlers complete.
    Compiling Cookbooks...
    [2016-10-11T22:06:20+01:00] INFO: /Users/merry
    [2016-10-11T22:06:20+01:00] WARN: Cloning resource attributes for osx_defaults[adjusts dock size to ] from prior resource (CHEF-3694)
    [2016-10-11T22:06:20+01:00] WARN: Previous osx_defaults[adjusts dock size to ]: /Users/merry/workspace/sprout-wrap/cookbooks/sprout-osx-settings/recipes/dock_preferences.rb:24:in `from_file'
    [2016-10-11T22:06:20+01:00] WARN: Current  osx_defaults[adjusts dock size to ]: /Users/merry/workspace/sprout-wrap/cookbooks/sprout-osx-settings/recipes/dock_preferences.rb:31:in `from_file'
    Converging 17 resources
    Recipe: sprout-base::var_chef_cache
      * directory[/var/chef/cache] action create[2016-10-11T22:06:20+01:00] INFO: Processing directory[/var/chef/cache] action create (sprout-base::var_chef_cache line 3)
     (up to date)
    Recipe: sprout-osx-settings::defaults_fast_key_repeat_rate
      * osx_defaults[set key repeat rate] action write[2016-10-11T22:06:20+01:00] INFO: Processing osx_defaults[set key repeat rate] action write (sprout-osx-settings::defaults_fast_key_repeat_rate line 1)
     (up to date)
      * execute[set key repeat rate - /Users/merry/Library/Preferences/.GlobalPreferences - KeyRepeat] action run[2016-10-11T22:06:20+01:00] INFO: Processing execute[set key repeat rate - /Users/merry/Library/Preferences/.GlobalPreferences - KeyRepeat] action run (/Users/merry/workspace/sprout-wrap/cookbooks/osx/providers/defaults.rb line 4)
    [2016-10-11T22:06:20+01:00] INFO: Processing execute[Guard resource] action run (dynamically defined)
    1
    [2016-10-11T22:06:20+01:00] INFO: execute[Guard resource] ran successfully
     (skipped due to not_if)
      * osx_defaults[set initial key repeat delay] action write[2016-10-11T22:06:20+01:00] INFO: Processing osx_defaults[set initial key repeat delay] action write (sprout-osx-settings::defaults_fast_key_repeat_rate line 8)
     (up to date)
      * execute[set initial key repeat delay - /Users/merry/Library/Preferences/.GlobalPreferences - InitialKeyRepeat] action run[2016-10-11T22:06:20+01:00] INFO: Processing execute[set initial key repeat delay - /Users/merry/Library/Preferences/.GlobalPreferences - InitialKeyRepeat] action run (/Users/merry/workspace/sprout-wrap/cookbooks/osx/providers/defaults.rb line 4)
    [2016-10-11T22:06:20+01:00] INFO: Processing execute[Guard resource] action run (dynamically defined)
    15
    [2016-10-11T22:06:20+01:00] INFO: execute[Guard resource] ran successfully
     (skipped due to not_if)
    Recipe: sprout-osx-settings::function_keys
      * osx_defaults[Turn on function-keys-work-as-function keys] action write[2016-10-11T22:06:20+01:00] INFO: Processing osx_defaults[Turn on function-keys-work-as-function keys] action write (sprout-osx-settings::function_keys line 5)
     (up to date)
      * execute[Turn on function-keys-work-as-function keys - /Users/merry/Library/Preferences/.GlobalPreferences - com.apple.keyboard.fnState] action run[2016-10-11T22:06:20+01:00] INFO: Processing execute[Turn on function-keys-work-as-function keys - /Users/merry/Library/Preferences/.GlobalPreferences - com.apple.keyboard.fnState] action run (/Users/merry/workspace/sprout-wrap/cookbooks/osx/providers/defaults.rb line 4)
    [2016-10-11T22:06:20+01:00] INFO: Processing execute[Guard resource] action run (dynamically defined)
    [2016-10-11T22:06:20+01:00] INFO: execute[Turn on function-keys-work-as-function keys - /Users/merry/Library/Preferences/.GlobalPreferences - com.apple.keyboard.fnState] ran successfully

        - execute defaults write /Users/merry/Library/Preferences/.GlobalPreferences com.apple.keyboard.fnState -boolean true
      * ruby_block[Fix Function Keys] action run[2016-10-11T22:06:20+01:00] INFO: Processing ruby_block[Fix Function Keys] action run (sprout-osx-settings::function_keys line 13)
    merry             275   0.0  0.5  2611920  40944   ??  S    Sat06pm   7:11.60 /System/Library/CoreServices/SystemUIServer.app/Contents/MacOS/SystemUIServer
    157:242: execution error: access for assistive devices is NOT enabled! (This is not an error, just a warning) (-2700)
    [2016-10-11T22:06:21+01:00] INFO: ruby_block[Fix Function Keys] called

        - execute the ruby block Fix Function Keys
    Recipe: sprout-osx-settings::screensaver
      * osx_defaults[ask for password when screen is locked] action write[2016-10-11T22:06:21+01:00] INFO: Processing osx_defaults[ask for password when screen is locked] action write (sprout-osx-settings::screensaver line 1)
     (up to date)
      * execute[ask for password when screen is locked - com.apple.screensaver - askForPassword] action run[2016-10-11T22:06:21+01:00] INFO: Processing execute[ask for password when screen is locked - com.apple.screensaver - askForPassword] action run (/Users/merry/workspace/sprout-wrap/cookbooks/osx/providers/defaults.rb line 4)
    [2016-10-11T22:06:21+01:00] INFO: Processing execute[Guard resource] action run (dynamically defined)
    1
    [2016-10-11T22:06:21+01:00] INFO: execute[Guard resource] ran successfully
     (skipped due to not_if)
      * osx_defaults[wait 60 seconds between screensaver & lock] action write[2016-10-11T22:06:21+01:00] INFO: Processing osx_defaults[wait 60 seconds between screensaver & lock] action write (sprout-osx-settings::screensaver line 7)
     (up to date)
      * execute[wait 60 seconds between screensaver & lock - com.apple.screensaver - askForPasswordDelay] action run[2016-10-11T22:06:21+01:00] INFO: Processing execute[wait 60 seconds between screensaver & lock - com.apple.screensaver - askForPasswordDelay] action run (/Users/merry/workspace/sprout-wrap/cookbooks/osx/providers/defaults.rb line 4)
    [2016-10-11T22:06:21+01:00] INFO: Processing execute[Guard resource] action run (dynamically defined)
    [2016-10-11T22:06:21+01:00] INFO: execute[wait 60 seconds between screensaver & lock - com.apple.screensaver - askForPasswordDelay] ran successfully

        - execute defaults write com.apple.screensaver askForPasswordDelay -float 60.0
      * osx_defaults[set screensaver timeout] action write[2016-10-11T22:06:21+01:00] INFO: Processing osx_defaults[set screensaver timeout] action write (sprout-osx-settings::screensaver line 19)
     (up to date)
      * execute[set screensaver timeout - ByHost/com.apple.screensaver.E30179FA-16E3-5588-9F5E-127D843E0F27 - idleTime] action run[2016-10-11T22:06:21+01:00] INFO: Processing execute[set screensaver timeout - ByHost/com.apple.screensaver.E30179FA-16E3-5588-9F5E-127D843E0F27 - idleTime] action run (/Users/merry/workspace/sprout-wrap/cookbooks/osx/providers/defaults.rb line 4)
    [2016-10-11T22:06:21+01:00] INFO: Processing execute[Guard resource] action run (dynamically defined)
    600
    [2016-10-11T22:06:21+01:00] INFO: execute[Guard resource] ran successfully
     (skipped due to not_if)
      * execute[set display, disk and computer sleep times] action run[2016-10-11T22:06:21+01:00] INFO: Processing execute[set display, disk and computer sleep times] action run (sprout-osx-settings::screensaver line 26)
    [2016-10-11T22:06:21+01:00] INFO: execute[set display, disk and computer sleep times] ran successfully

        - execute pmset -a displaysleep 20 disksleep 15 sleep 0
    Recipe: sprout-osx-settings::set_menubar_clock_format
      * osx_defaults[turn on date & seconds for menubar clock] action write[2016-10-11T22:06:21+01:00] INFO: Processing osx_defaults[turn on date & seconds for menubar clock] action write (sprout-osx-settings::set_menubar_clock_format line 2)
     (up to date)
      * execute[turn on date & seconds for menubar clock - com.apple.menuextra.clock - DateFormat] action run[2016-10-11T22:06:21+01:00] INFO: Processing execute[turn on date & seconds for menubar clock - com.apple.menuextra.clock - DateFormat] action run (/Users/merry/workspace/sprout-wrap/cookbooks/osx/providers/defaults.rb line 4)
    [2016-10-11T22:06:21+01:00] INFO: Processing execute[Guard resource] action run (dynamically defined)
    MMM d  HH:mm:ss
    [2016-10-11T22:06:21+01:00] INFO: execute[Guard resource] ran successfully
     (skipped due to not_if)
    Recipe: sprout-osx-settings::dock_preferences
      * osx_defaults[set dock to be on left] action write[2016-10-11T22:06:21+01:00] INFO: Processing osx_defaults[set dock to be on left] action write (sprout-osx-settings::dock_preferences line 3)
     (up to date)
      * execute[set dock to be on left - com.apple.dock - orientation] action run[2016-10-11T22:06:21+01:00] INFO: Processing execute[set dock to be on left - com.apple.dock - orientation] action run (/Users/merry/workspace/sprout-wrap/cookbooks/osx/providers/defaults.rb line 4)
    [2016-10-11T22:06:21+01:00] INFO: Processing execute[Guard resource] action run (dynamically defined)
    left
    [2016-10-11T22:06:21+01:00] INFO: execute[Guard resource] ran successfully
     (skipped due to not_if)
      * osx_defaults[set dock autohide to ] action write[2016-10-11T22:06:21+01:00] INFO: Processing osx_defaults[set dock autohide to ] action write (sprout-osx-settings::dock_preferences line 10)
     (up to date)
      * execute[set dock autohide to  - com.apple.dock - autohide] action run[2016-10-11T22:06:21+01:00] INFO: Processing execute[set dock autohide to  - com.apple.dock - autohide] action run (/Users/merry/workspace/sprout-wrap/cookbooks/osx/providers/defaults.rb line 4)
    [2016-10-11T22:06:21+01:00] INFO: Processing execute[Guard resource] action run (dynamically defined)
    [2016-10-11T22:06:21+01:00] INFO: execute[set dock autohide to  - com.apple.dock - autohide] ran successfully

        - execute defaults write com.apple.dock autohide -boolean true
      * osx_defaults[remove persistent apps from the dock] action write[2016-10-11T22:06:21+01:00] INFO: Processing osx_defaults[remove persistent apps from the dock] action write (sprout-osx-settings::dock_preferences line 17)
     (up to date)
      * execute[remove persistent apps from the dock - com.apple.dock - persistent-apps] action run[2016-10-11T22:06:21+01:00] INFO: Processing execute[remove persistent apps from the dock - com.apple.dock - persistent-apps] action run (/Users/merry/workspace/sprout-wrap/cookbooks/osx/providers/defaults.rb line 4)
    [2016-10-11T22:06:21+01:00] INFO: Processing execute[Guard resource] action run (dynamically defined)
    [2016-10-11T22:06:21+01:00] INFO: execute[remove persistent apps from the dock - com.apple.dock - persistent-apps] ran successfully

        - execute defaults write com.apple.dock persistent-apps -array
      * osx_defaults[adjusts dock size to ] action write[2016-10-11T22:06:21+01:00] INFO: Processing osx_defaults[adjusts dock size to ] action write (sprout-osx-settings::dock_preferences line 24)
     (skipped due to only_if)
      * osx_defaults[adjusts dock size to ] action write[2016-10-11T22:06:21+01:00] INFO: Processing osx_defaults[adjusts dock size to ] action write (sprout-osx-settings::dock_preferences line 31)
     (skipped due to only_if)
      * osx_defaults[toggle dock magnification on/off] action write[2016-10-11T22:06:21+01:00] INFO: Processing osx_defaults[toggle dock magnification on/off] action write (sprout-osx-settings::dock_preferences line 38)
     (skipped due to not_if)
      * execute[relaunch dock] action run[2016-10-11T22:06:21+01:00] INFO: Processing execute[relaunch dock] action run (sprout-osx-settings::dock_preferences line 45)
    [2016-10-11T22:06:21+01:00] INFO: execute[relaunch dock] ran successfully

        - execute killall Dock
    [2016-10-11T22:06:21+01:00] INFO: Chef Run complete in 1.275448 seconds
    [2016-10-11T22:06:21+01:00] INFO: Skipping removal of unused files from the cache

    Running handlers:
    [2016-10-11T22:06:21+01:00] INFO: Running report handlers
    Running handlers complete
    [2016-10-11T22:06:21+01:00] INFO: Report handlers complete
    Chef Client finished, 7/27 resources updated in 15 seconds
    ~
    Already up-to-date.
    Succeeded in installing golang
    Succeeded in installing jq
    Succeeded in installing mp3blaster
    Succeeded in installing pstree
    Succeeded in installing tree
    Succeeded in installing 1password
    Succeeded in installing dropbox
    Succeeded in installing evernote
    Succeeded in installing flycut
    Succeeded in installing google-chrome
    Succeeded in installing iterm2
    Succeeded in installing selfcontrol
    Succeeded in installing shiftit
    Succeeded in installing skype
    Succeeded in installing tunnelbear
    Succeeded in installing vagrant
    Succeeded in installing vlc
    Succeeded in installing virtualbox

    Success: 18 Fail: 0
    Last software update was 3 days ago
    Since softwareupdate ran recently, not going to do anything
    remove /Users/merry/.softwareupdate_indicator to force an update
    ```

*This post took 7 pomodoros to complete (more than enough time to setup a laptop from scratch)*
