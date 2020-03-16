---
layout: post
title: "Examples of how production like environments can catch bugs early"
date: 2012-06-08 17:03
comments: true
categories:
---
On a previous project, I witnessed a few interesting build breakages that we caught in [CI](http://wikipedia.org/Continuous_Integration)
that could have easily gotten passed our automated quality gates. These issues would have hurt us in downstream
environments.

It's commonly referenced in the [Continuous Delivery](http://continuousdelivery.com/) book that development
environments should be as close to production like as possible. This might not
be 100% possible but you can come close. What can come that close is your
system integration test environment. On this project we had a typical
web application architecture: web server, database server, and a load balancer
to support multiple web servers. Here are a couple integration failures that we experienced:

## Issue 1

A route in our application returned the version of the application and the
hostname of the machine. It wasn't used by the application at all, but it was
used in environments that had load balancers because it is what they used to
detect the health of the application. Someone deleted this route and
deployments ceased to be successful because the application was considered
down by the load balancer so it returned [HTTP 503](http://www.w3.org/Protocols/rfc2616/rfc2616-sec10.html).

Had our [system integration](http://en.wikipedia.org/wiki/System_integration_testing) environment simply used the hostname of the
application server, this would not have been caught until it made it to
downstream environments. It was easy to fix and helped transfer knowledge to
the team. It was also much less stressful to catch here than when a product
owner is pushing something to UAT and things all of the sudden break.

## Issue 2

A change to some application code used unconfigured database connection
code. Default behaviour of the code is to connect to localhost when issuing
queries against the database. This worked fine locally, but because the
system integration environment seperated the database server from the web
application server, this unconfigured code was caught before it made it to
any other downstream environment.

## Isn't That Overboard?

If I was managing environments manually, I would say yes. Since the environments
are nearly 100% configured by [Chef](http://opscode.com/chef) and our deployment code is automated, I don't think
having the environment set up this way is going too far. It caught what could be
extremely stressful defects to debug in production much earlier such that that our fixes
could be done quickly and with confidence.

I predict that in the future, tools like [vagrant](http://vagrantup.com/) will help us setup these
architectures on our local machines so that we can catch these issues before
they even get checked into version control. Imagine developing software in
a way that the mismatch between dev and prod becomes such a non-issue that
it's no longer a worry? I look forward to helping make that world possible.
