---
layout: post
title: "When 1 Does Not Equal 1"
date: 2009-02-27 07:41
comments: true
permalink: /blog/development/when-1-does-not-equal-1-a-debugging-tale/
categories: 
---

I'm starting to get comfortable with [NHibernate](http://nhforge.org/Default.aspx) but I've been [cowboy coding](http://en.wikipedia.org/wiki/Cowboy_coding) it without any testing framework to let me know if my mappings are really doing what I think they are doing. I'm starting to get tired of manually interacting with my applications to see if things work so I've begun the course of starting database integration testing. For the most part, it's all about testing my mapping files.

<!-- more -->

Because these tests are meant to be automated, I wanted to be as easy as possible to get the test framework into a consistent state. I found tool called [NDbUnit](http://code.google.com/p/ndbunit/) and it allowed me to dump data stored in Xml straight to your database. It was looking great until I found out it didn't support [SQLite](http://www.sqlite.org/) which is the database engine I want to use for the testing because of speed reasons. At least the tool is extensible and I took some time to write a SQLite portion to the codebase.

Anyways, enough about what lead me to this post. During the development of SQLite integration I came across a weird bug. A feature of NDbUnit is to allow a refresh of the data, so you can merge/update/insert new data without having to completely remove everything from the test database. One of the parts of this code is to compare the data set from the xml file to the data in the database. To find out if an **UPDATE** needs to be done versus and **INSERT** is by the return value of this method:

<pre>
protected bool IsPrimaryKeyValueEqual(DataRow dataRow1, DataRow dataRow2, DataColumn[] primaryKey)
{
    if (primaryKey.Length == 0)
    {
        return false;
    }

    for (int i = 0; i < primaryKey.Length; ++i)
    {
        DataColumn dataColumn = primaryKey[i];
        // Primary key column value is not equal.
        if (!dataRow1[dataColumn.ColumnName].Equals(dataRow2[dataColumn.ColumnName]))
        {
            return false;
        }
    }

    return true;
}
</pre>

It simply goes through all the keys of the table, and if they are all equal, then we've found a tuple that needs to be updated. Seems pretty straightforward, and not a method that I would presume would cause any problems. Strangely enough, here is where some strangeness started to occur. This method would always return false! I hooked up the debugger and inspected the rows and everything seemed normal. Until what seemed like the 30th time; I noticed something. The row from the XML had 32 bit IDs while the row from the SQLite database had 64 bit IDs.

<img src="/images/blog/debugger_tail_32bit.png" />

and here's what that DataRow object was being compared to:

<img src="/images/blog/debugger_tail_32bit.png" />

I didn't want jump to conclusions as to whether or not I found my problem so I wrote a test. Sure enough, the following test fails (and for good reason).

<pre>
[Test]
public void One_should_equal_one()
{
    object _32bit1 = (Int32) 1;
    object _64bit1 = (Int64) 1;

    bool result = _64bit1.Equals(_32bit1);
    
    Assert.IsTrue(result);
}
</pre>

Since SQLite always uses 64 bits for its integers I simply made a copy of the XSD file (which defines how the XML data relates to the database schema) and replaced the xs:int id declarations with xs:long data type.

The End.

