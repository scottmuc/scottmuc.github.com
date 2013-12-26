---
layout: post
title: "Unit Testing Domain Persistence With NDbUnit"
date: 2009-03-09 07:51
comments: true
permalink: /blog/development/unit-testing-domain-persistence-with-ndbunit-nhibernate-and-sqlite/
categories: 
---

Ever since I've begun using NHibernate the number one thing that's caused me a lot of headaches is learning how to properly map my domain objects. Even the most basic mappings I wrote had bugs simple because I overlooked trivial items. This made me realize that I could save a lot of time and hassle if I unit tested by data access layer.

Ever since I've begun using NHibernate the number one thing that's caused me a lot of headaches is learning how to properly map my domain objects. Even the most basic mappings I wrote had bugs simple because I overlooked trivial items. This made me realize that I could save a lot of time and hassle if I unit tested by data access layer.

<!-- more -->

I'm a neophyte when it comes to unit testing, but over the last six months I've started to understand the benefits of testing and where it fits in with my development cycle. In order to test my mappings I needed a simple way of placing dummy data into my database. I also wanted to use SQLite as my test database implementation for simplicity, and speed reasons. When a new developer jumps on board, I think it would be cool if they can run the database unit test suite without having to setup SQL Server, or any other database application server. My research led me to a project called NDbUnit. It allows a simple way to populate a database with data from an Xml DataSet. This is exactly what I wanted, but it did not support SQLite. I took the plunge and implemented a SQLite provider and even got it added to their codebase! [I've already blogged](/blog/development/when-1-does-not-equal-1-a-debugging-tale/) about that tangent here.

With tools in hand I started creating an integration test suite for the CBC Radio 3 website codebase. The website is currently undergoing a re-write where I am migrating from my hand written data access layer to an NHibernate based DAL (Data Access Layer). Oh how do I wish I wrote it with NHibernate from the beginning :-). The current database has 92 tables and I've only just begun mapping them to proper Domain objects. So far I've written about 20 Domain objects and I've hit enough walls with the mapping files. I realized that I needed a testing foundation to assert the persistence of my domain objects.

Enough about my objectives, here's how I plan on making a database testing suite that doesn't get bogged down by the size of the domain or the number of repositories I have.

I've attached an NDbUnit demo code base that should be fully runnable within Visual Studio if you care to try this out for yourself. Note that I've used a lot of thirdparty tools to assist with this setup:

* Castle.Windsor
* Rhino.Commons.Binsor
* Castle.Facilities.NHibernateIntegration
* NHibernate

The first step in setting all this up is my first test which does not have any assertions. It's existence is purely to have a simple way of creating the SQLite database file with the schema populated from the Domain.

{% gist 1551024 CreateDatabase.cs %}

To elaborate briefly, the first line gets the Windsor container from a Boo config file. Next, it obtains the NHibernate Configuration object from the container, then exports the schema. There's a lot of magic being done in these few lines of code and I would like to thank all those who worked on these projects for making my life a lot easier. My TestConfiguration class is just a wrapper around some hard coded string:

{% gist 1551024 TestConfiguration.cs %}

The test setup is inspired by Chris Canal's post on "Blistering Fast Integration Tests with NHibernate and SQLIte". My Boo config files are nearly identical to his, so I'll save the bits and not paste them here.

Next I create a RepositoryTestBase to encapsulate the construction of my container and expose the ISessionManager to all my tests.

{% gist 1551024 RepositoryTestBase.cs %}

So that's pretty much all the underlying infrastructure for my database tests so far. Since I have a lot of tables to test, and NDbUnit requires that I model them using Xml DataSet schema definitions I didn't want to put them all in one file. I eventually came to the realization that I could segregate my XSD files to contain only the tables that I needed to test certain Aggregate Roots (although I barely know if I'm using that term correctly). When testing my IArtistRepository, I didn't want to have to put Blog tables in my XSD and vice-versa. I eventually came up with a directory structure to deal with separating the concerns of my database tests. Each repository would get its own directory/namespace for all the classes that it's concerned with.

I ended up designing the system so I have three repositories that returned a BlogPost, User, and Profile. So in my integration test suite I have directories for { Blog, Profiles, Users } to group repository functionality.

In each repository directory I create a custom test context class as a base class for all other fixtures in the directory. This test context will be used to expose the repository that's under test, and to populate the dummy database using NDbUnit. Here's an example of the UserRepositoryTestContext

{% gist 1551024 UserRepositoryTestContext.cs %}

The XSD for this section contains a dataset that for the Users, Roles, and UserRoles tables and the dummy data looks something like the following:

{% img /images/blog/UsersDS.jpg %}

{% gist 1551024 Users.xml %}

Now I can write succinct unit tests leveraging all the infrastructure that I've put in place:

The User domain object is quite simple and that's why I also have tests for a simple BlogPost class that has a child collection of Comments. Using a similar structure I can write painless tests like the following:

{% gist 1551024 UserPersistenceTests.cs %}

With this kind of segregation between the repository tests I don't have to worry about creating XSD DataSets for my entire database in one file. I can pick and choose the tables I want to test directly using NDbUnit. For my Blog tests, I didn't want to worry about User persistence even though it's mapped as a property on my BlogPost class. By leveraging lazy loading I don't need to worry that I don't have a valid User object in the database because my tests never access the User property, therefore I don't get any exceptions when the NHibernate generated proxy attempts to obtain that entity from the session/database.

That concludes my one week analysis of using NDbUnit to testing my NHibernate persistence layer. Please leave a comment if you have suggestions on this setup. So far this setup hasn't caused me too much grief, but it's only been in use for a few days. On the surface it seems clean to me, but take it all with a grain of salt.

