# Deployment Details
A single page application built with the Go programming language

Application is showing current article retrieved from www.guardian.com and allows the users to discuss the topic.
To use the application user must to follow next steps:

1. Establish GO environment on local machine
2. Include go packages assosiated to the application
3. Download application
4. Run

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

```
go get <package address>
```

for all dependencies.

###Download Application

To install application:

1. Using command prompt
  * cd %GOPATH%\src - for WINDOWS
  * cd $GOPATH/src  - for LINUX
2. You need to clone it from remote repository - https://github.com/FourthYearSW/GO-SPA

```
git clone https://github.com/FourthYearSW/GO-SPA
```

###Run Application

To run application you need to go to the application root directory

```
cd %GOPATH%\src\GO-SPA  #for WINDOWS
cd $GOPATH/src/GO-SPA   #for LINUX
```

then command

```
go run main.go
```

In the browser's address bar type

```
127.0.0.1:8080
```

and discus the article with other participants. 
