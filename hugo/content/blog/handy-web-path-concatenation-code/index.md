---
title: Handy Web Path Concatenation Code
date: 2008-12-03T16:33:00Z
tags: []
---

Lot of web developement code is spent concatenating path snippets and I always hated having to deal with slashes. I wanted a function that was consistent with its return type, and very loose in its parameter requirements. I ended up with the following:

*Test First*:

{{< highlight csharp "linenos=table" >}}
using bddunit.core;
using CBC.Radio3.Commons.Specs;
using Observation = MbUnit.Framework.TestAttribute;
using SUT = CBC.Radio3.Commons.Web.WebPath;

namespace CBC.Radio3.Commons.Web
{
    [Concern(typeof (WebPath))]
    public class when_WebPath_Combines_a_bunch_of_strings : ContextSpecification
    {
        private string result;

        protected override void establish_context()
        {
            
        }

        protected override void because()
        {
            result = SUT.Combine("/slashesonbothsides/", "//doubleslash//", "rightslash/", "/leftslash", "file.jpg");
        }

        [Observation]
        public void should_return_an_absolute_web_path()
        {
            result.should_be_equal_to("/slashesonbothsides/doubleslash/rightslash/leftslash/file.jpg");
        }
    }
}
{{< / highlight >}}

The [BDD](http://en.wikipedia.org/wiki/Behavior_Driven_Development) style of testing is inspired (and pretty much copied) from JP Boodhoo's Nothing but .Net class.

*Here's the code*:

{{< highlight csharp "linenos=table" >}}
using System.Text;

namespace CBC.Radio3.Commons.Web
{
    public class WebPath
    {
        /// <summary>
        /// Takes a parameter list of strings and returns the strings concatenated with a
        /// preceding '/' character. If the string has a trailing/preceding slash it will
        /// be removed. Any slashes in the middle of the string will remain.
        /// </summary>
        /// <param name="paths"></param>
        /// <returns>Contatenation of paths seperated by directory seperator</returns>
        public static string Combine(params string[] paths)
        {
            var sb = new StringBuilder();

            foreach (var path in paths)
                sb.AppendFormat("/{0}", path.TrimEnd('/').TrimStart('/'));

            return sb.ToString();
        }        
    }
}
{{< / highlight >}}

That was my first go at it. I have a feeling I might generalize the class to be a delimited string builder. Which would be nice because then the method won't be static. Anyways, just thought I would share something that I think is concise but will be used a lot.

