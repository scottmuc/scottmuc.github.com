---
title: It Is Good To Laugh At Your Old Code
date: 2008-11-10T16:33:00Z
tags: []
---

One of the things I love about my job is the ability to see myself grow as a developer. I'm glad I can laugh at my old code because if I can't then I'm definitely not improving my skillset.

Here's some code from around 3 years ago:

{{< highlight csharp "linenos=table" >}}
using System;
using System.Web;
using System.Configuration;

using System.Text;
using System.Text.RegularExpressions;
using CBC.Radio3.Core.Web;

namespace CBC.Radio3.Core
{
    public class UrlRewriteHandler : IHttpHandler, System.Web.SessionState.IRequiresSessionState
    {
        // Override the ProcessRequest method.
          public void ProcessRequest(HttpContext context) 
          {
            string path = (context.Request.QueryString["url"] == null) 
                            ? string.Empty 
                            : Convert.ToString(context.Request.QueryString["url"]);
            
            string sendToUrl = path;
                        
            RewriterRuleCollection rules = RewriterConfiguration.GetConfig().Rules;
            context.Trace.Write("ModuleRewriter", "Url Requested " + path);

            for(int i = 0; i < rules.Count; i++)
            {
                string lookFor = rules[i].LookFor;
                Regex re = new Regex(lookFor, RegexOptions.IgnoreCase);

                if (re.IsMatch(path))
                {
                    sendToUrl = SiteInfo.ApplicationPath + re.Replace(path, rules[i].SendTo);
                    context.Trace.Write("ModuleRewriter", "Rewriting URL to " + sendToUrl);
                    
                    string referer               = context.Request.Headers["Referer"] ?? string.Empty;
                    string permalink             = path;
                    RewriterRuleType ruleType     = UrlRewriteHandler.GetRewriterRuleType(path); 
                    string siteDomain            = SiteInfo.Hostname;
                    
                    if ( !IsBotRequest(context) && !referer.Contains( siteDomain ))
                    {
                        PermalinkLogger.Instance().Log(permalink, ruleType, referer);
                    }
                    
                    context.Response.Redirect(sendToUrl);
                }
            }
            
            // We haven't found a match so return a 404
            context.Response.Status        = "404 Not Found";
            context.Response.StatusCode = 404;
          }
 

          private static bool IsBotRequest(HttpContext context)
          {
              return context.Request.Browser.Crawler;
          }

          public bool IsReusable { get { return true; } }
          
          
        public static RewriterRuleType GetRewriterRuleType(string permalink)
        {
            RewriterRuleCollection rules = RewriterConfiguration.GetConfig().Rules;

            for(int i = 0; i < rules.Count; i++)
            {
                string lookFor = rules[i].LookFor;
                Regex re = new Regex(lookFor, RegexOptions.IgnoreCase);

                if (re.IsMatch(permalink))
                {
                    return rules[i].RuleType;
                }
            }
            
            return RewriterRuleType.Undefined;
        } 
        
        public static string GetRuleDataXml(string permalink)
        {
            RewriterRuleCollection rules = RewriterConfiguration.GetConfig().Rules;

            for(int i = 0; i < rules.Count; i++)
            {
                string lookFor = rules[i].LookFor;
                Regex re = new Regex(lookFor, RegexOptions.IgnoreCase);

                if (re.IsMatch(permalink))
                {
                    string xmlData = re.Replace(permalink, rules[i].NodeForm);
                    return xmlData;
                }
            }
            
            return string.Empty;        
        }        
   }
}
{{< / highlight >}}

Here's what the code I used to replace it:

{{< highlight csharp "linenos=table" >}}
using System.Web;
using CBC.Radio3.Permalinks.Common;

namespace CBC.Radio3.Permalinks.HttpHandlers
{
    public class PermalinkProcessingHandler : IHttpHandler
    {
        private readonly IPermalinkParser permalinkParser;

        public PermalinkProcessingHandler()
            : this(new PermalinkParser())
        {
            
        }

        public PermalinkProcessingHandler(IPermalinkParser permalinkParser)
        {
            this.permalinkParser = permalinkParser;
        }

        public void ProcessRequest(HttpContext context) 
        {
            string permalink = context.Request.QueryString["url"] ?? string.Empty;

            if (!string.IsNullOrEmpty(permalink) && permalinkParser.CanParse(permalink))
            {
                string redirectUrl = permalinkParser.GetRedirectUrlFrom(permalink);
                context.Response.Redirect(redirectUrl);
            }

            context.Response.Status        = "404 Not Found";
            context.Response.StatusCode = 404;
        }

        public bool IsReusable { get { return true; } }
    }
}
{{< / highlight >}}


