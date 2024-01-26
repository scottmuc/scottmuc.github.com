---
title: Testing Powershell Code That Talks to CLR Objects
date: 2014-06-05T17:39:41-07:00
tags:
    - pester
aliases:
    - /testing-powershell-code-that-talks-to-clr-objects

---

Several months ago someone posted an excellent question on the [Pester Group][pester-group] about a certain problem area
in testing PowerShell with [Pester][pester].  He brings up a relevant example of the difficulties that can occur when writing tests
is done afterwards. It also reveals a missing feature of Pester.

First let me display the code that needed tests:

{{< highlight powershell "linenos=table" >}}
Add-Type -Path (Join-Path -Path $psScriptRoot "Renci.SshNet.dll")

function Get-FromSFtp {
    $fileList=$null
    try {
        $sftp=New-Object Renci.SshNet.SftpClient("127.0.0.1","test","test")
        $sftp.Connect()
        if($sftp.IsConnected) {
            $sftp.ChangeDirectory("/test")
            $fileList=$sftp.ListDirectory(".")
        }
        return $fileList
    } catch [Exception] {
        Write-Host $_.Exception.ToString()
    } finally {
        if(($sftp -ne $null) -and ($sftp.IsConnected)) {
            $sftp.Disconnect()
        }
    }
}
{{< / highlight >}}

Here are a couple reasons why this code is difficult to test:

* Line 4 returns a CLR object that actually does SFTP
* Many parameters are hard coded

### Start With Integration

I started this exercise with the intention of not changing a single character of the code submitted. Here's the story of
how the changes took place (I had a git repo of this, but the files got corrupted). First thing to do was to make the
existing code pass a test as is. In order to do that, I needed an integration environment for it to talk to. So I used
[Vagrant][vagrant] to spin up a local linux box. I manually configured the box such that the credentials in the code
work. Here's the Vagrantfile

{{< highlight ruby "linenos=table" >}}
# -*- mode: ruby -*-
# vi: set ft=ruby :

# Vagrantfile API/syntax version. Don't touch unless you know what you're doing!
VAGRANTFILE_API_VERSION = "2"

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|
  config.vm.box = "precise64"
  config.vm.box_url = "http://files.vagrantup.com/precise64.box"
  config.vm.network "forwarded_port", guest: 22, host: 22
end
{{< / highlight >}}

and here's the passing test. I'm happy with it because I could change the code or the environment and cause the test to
fail, thus giving me confidence that the test is good enough to tell if me I broke anything.

```powershell
Describe "Get-FromSFtp Integration" {
    It "returns the 2 default dir entries" {
        Get-FromSFtp | ForEach-Object { $_.Name } | Should Be ".", ".."
    }
}
```

Happy that I have a feedback loop going now, I begin to start testing the code in isolation. I start by writing the
contexts and specifications:

### Follow Up With Isolated Tests

```powershell
Describe "Get-FromSFtp" {
    Context "when connected" {
        It "attempts to connect to an sftp server" { }
        It "changes to the directory /test" { }
        It "gets the contents listing of the directory" { }
        It "disconnects from the sftp server" { }
    }

    Context "when not connected" {
        It "returns 0 files" { }
    }

    Context "when an error happens" {
        It "prints out the error to the user"
    }
}
```

I was stumped for a bit because I realized Pester didn't have the power necessary to do this easily. It didn't mean it
couldn't be tested though. The Mocking functionality came to the rescue as I was able to mock `New-Object` and decouple
my tests from the `Renci.SshNet` library. The hard part was hand-rolling the mock object to be used in the tests.

One by one, I made each spec pass. The beginning of each `Context` was the setup, and each `It` asserted something from
after that setup and action had been taken. The nice thing that resulted was I could break a specific piece of the
production code, and the specific test would tell me that it's broken at that place (eg: change the directory the
production code goes to).

Here's the complete test file for the code:

{{< highlight powershell "linenos=table" >}}
$here = Split-Path -Parent $MyInvocation.MyCommand.Path
$sut = (Split-Path -Leaf $MyInvocation.MyCommand.Path).Replace(".Tests.", ".")
. "$here\$sut"

$mockSFtpCode = @"
public class MockSFtp {
    public bool ConnectCalled { get; private set; }
    public bool DisconnectCalled { get; private set; }
    public bool IsConnected { get; set; }
    public string ChangeDirectoryParam { get; private set; }
    public string ListDirectoryParam { get; private set; }

    public void Connect() {
        this.ConnectCalled = true;
    }

    public void Disconnect() {
        this.DisconnectCalled = true;
    }

    public void ChangeDirectory(string directory) {
        this.ChangeDirectoryParam = directory;
    }

    public string[] ListDirectory(string path) {
        this.ListDirectoryParam = path;
        return new string[] { "test1" };
    }
}
"@

$throwingMockSFtpCode = @"
public class ThrowingMockSFtp {
    public void Connect() {
        throw new System.Exception();
    }
}
"@

Add-Type -TypeDefinition $mockSFtpCode
Add-Type -TypeDefinition $throwingMockSFtpCode

Describe "Get-FromSFtp" {
    Context "when connected" {
        $mockSFtp = New-Object MockSFtp
        $mockSFtp.IsConnected = $true
        Mock New-Object -MockWith { $mockSFtp }
        $files = Get-FromSFtp

        It "attempts to connect to an sftp server" {
            $mockSFtp.ConnectCalled | Should Be $true
        }

        It "changes to the directory /test" {
            $mockSFtp.ChangeDirectoryParam | Should Be '/test'
        }

        It "gets the content listing of the directory" {
            $mockSFtp.ListDirectoryParam | Should Be '.'
            $files | Should Be @("test1")
        }

        It "disconnects from the sftp server" {
            $mockSFtp.DisconnectCalled | Should Be $true
        }
    }

    Context "when not connected" {
        $mockSFtp = New-Object MockSFtp
        $mockSFtp.IsConnected = $false
        Mock New-Object -MockWith { $mockSFtp }
        $files = Get-FromSFtp

        It "returns 0 files" {
            $files | Should BeNullOrEmpty
        }
    }

    Context "when an exception happens" {
        $mockSFtp = New-Object ThrowingMockSFtp
        Mock New-Object -MockWith { $mockSFtp }

        It "prints out the error to the user" {
            # I do not think this is a good test, but until
            # mocking of Write-Host can be done, this is all
            # that we can really do.
            { Get-FromSFtp } | Should Not Throw
        }
    }
}

Describe "Get-FromSFtp Integration" {
    It "returns the 2 default dir entries" {
        Get-FromSFtp | ForEach-Object { $_.Name } | Should Be ".", ".."
    }
}
{{< / highlight >}}

### Summary

I can't say this is the best test code out there. I think testing PowerShell is in a different place than testing C# or
other OO languages. It seems like the best you can do is test the function logic and test the composition of it all in
an integrated way. The first refactoring change I would likely do to this function is make the connection a parameter.
That way I wouldn't need to mock `New-Object`, but passing a mock object is still an issue. Perhaps Pester needs a mock
object builder of some kind, though writing mocks by hand wasn't too difficult either. It's something I would expect a
testing framework to be able to do for me though.

[pester-group]: https://groups.google.com/forum/#!topic/pester/Y8UzSiNlcSE
[vagrant]: http://www.vagrantup.com/
[pester]: https://github.com/pester/Pester


