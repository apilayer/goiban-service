goiban-service
==============

## Openiban is going to be revived, stay tuned!

--- 

Implements a basic REST Web-service for validating IBAN account numbers in GO. Uses the logic from http://github.com/fourcube/goiban. Deployed at http://openiban.com .

# Running the service

## Via `go get`:

```bash
$ go get -u github.com/fourcube/goiban-service
# Launch the service listening on port 8080 and serve static content
$ $GOPATH/bin/goiban-service -port 8080 -w 
```

## Download a binary package:

A list of all releases is available [here](https://github.com/fourcube/goiban-service/releases).

```bash
# Make sure to choose the correct operating system and architecture
$ curl -Lo goiban-service.tar.gz "https://github.com/fourcube/goiban-service/releases/download/v1.0.2/goiban-service-1.0.2-linux-386.tar.gz"
$ tar -xzf goiban-service.tar.gz
$ cd goiban-service
# Launch the service listening on port 8080, using the bank data from ./data and serving
# the web interface from ./static
$ ./goiban-service -dataPath ./data -staticPath ./static -port 8080 -w
```

## Via Docker

https://hub.docker.com/r/fourcube/openiban/

```

$ docker run --name openiban -d -p8080:8080 fourcube/openiban
$ curl localhost:8080/validate/DE89370400440532013000
```

You will see something like:

```json
{
  "valid": true,
  "messages": [],
  "iban": "DE89370400440532013000",
  "bankData": {
    "bankCode": "",
    "name": ""
  },
  "checkResults": {}
}
```

## Via Kubernetes

You can deploy above docker file in your kubernetes cluster by using the configuration files container in [./k8s](./k8s).

Run the following command to apply the deployment and service (`NodePort service` running on port 32111)

```bash
kubectl apply -f ./k8s
```

# Building the service (MySQL)

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
$ ./goiban-service -port 8080 -dbURL root:root@/goiban?charset=utf8
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


Client Libraries
------

Name                                                                                         | Language
-------------------------------------------------------------------------------------------- | ---------------------
[openiban.js](https://github.com/fourcube/openiban.js) :star:                                | JavaScript/TypeScript
[OpenIBAN](https://github.com/nathanIL/openiban)                                             | Python
[OpenIban Connector](http://store.shopware.com/webch77589500739/openiban.com-connector.html) | PHP (Plugin for Shopware)

:star: officially supported


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
