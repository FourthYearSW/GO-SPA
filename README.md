# GO-SPA
A single page application built with the Go programming language

Application is showing current article retrieved from www.guardian.com and allows the users to discuss the topic.
To use the application user must to follow next steps:

1. Establish GO environment on local machine
2. Include go packages assosiated to the application
3. Download application
4. Build application
5. Use it.

###Establish Go environment on local machine

If you don't have the Go language installed on your machine you need download from [here](https://golang.org/dl/) and install it using the followin [instructions](https://golang.org/doc/install#install).

Very important to configure environment variables for GO. Instructions is [here](https://golang.org/doc/install)

###Include go packages assosiated to the application

Packages need to be installed:

 * github.com/guardian/gocapiclient
 * github.com/kataras/iris
 * github.com/valyala/fasthttp
 * gopkg.in/mgo.v2

To include all packages for GO-SPA you need command

```go get <package address>```

for all dependencies.

