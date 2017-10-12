goiban-service
==============

Implements a basic REST Web-service for validating IBAN account numbers in GO. Uses the logic from http://github.com/fourcube/goiban. Deployed at http://openiban.com .

Running the service
-------

You have to install go >= 1.8, setup your GOPATH and install a MySQL server.
Goiban requires a database called 'goiban'. The following commands assume a 
MySQL database running on `localhost:3306` with database `goiban` and 
user `root` with password `root`.

```
$ go get -u github.com/fourcube/goiban-data-loader
$ cd $GOPATH/src/github.com/fourcube/goiban-data-loader
$ DATABASE_URL="root:root@tcp(localhost:3306)/goiban?charset=utf8" make migrate

# load data
$ go build
$ ./goiban-data-loader bundesbank root:root@/goiban?charset=utf8

$ go get github.com/fourcube/goiban-service
$ cd $GOPATH/src/github.com/fourcube/goiban-service
$ go build
$ ./goiban-service 8080 root:root@/goiban?charset=utf8
```

To create a build without the metrics support (e.g if you run on go < 1.8) run:

```
$ go build -tags no_metrics
```

MySQL development instance
-------
To quickly run a MySQL database inside a docker container you can use
the following command:

`docker run -d --name openiban-mysql -p3306:3306 -e MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=goiban mysql`


The MIT License (MIT)
------
Copyright (c) 2013-2017 Chris Grieger

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
