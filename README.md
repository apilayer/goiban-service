goiban-service
==============

Implements a basic REST Web-service for validating IBAN account numbers in GO. Uses the logic from http://github.com/fourcube/goiban. Deployed at http://openiban.com .

Running the service
-------

You have to install go >= 1.2, setup your GOPATH and install a MySQL server.
Goiban requires a database called 'goiban'.

```
go get github.com/fourcube/goiban-data-loader
cd $GOPATH/src/github.com/fourcube/goiban-data-loader
go build

# apply schema to database
mysql -uroot -proot goiban < schema/1.sql

# load data
./goiban-data-loader bundesbank root:root@/goiban?charset=utf8

go get github.com/fourcube/goiban-service
cd $GOPATH/src/github.com/fourcube/goiban-service
go build
./goiban-service 8080 root:root@/goiban?charset=utf8
```

Using metrics
-------

You can use [KeenIO](http://keen.io) metrics if you want. Just use the following
command when starting goiban-service:

```
./goiban-service 8080 root:root@/goiban?charset=utf8 <ENV> <ProjectID> <WriteAPIKey>
```

The metrics will be logged in a collection named after the <ENV> parameter. You should not use "Test" or "Live" because those environments enable the HTML frontend.

In-memory metrics are automatically available under http://<host>:<port>/metrics.



The MIT License (MIT)
------
Copyright (c) 2015 Chris Grieger

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
