package main

import (
	"github.com/kataras/iris"
	"github.com/valyala/fasthttp"
	"GO-SPA/controllers"
	"gopkg.in/mgo.v2"
)

func main() {
	// Get Guardian API Instance
	gapi := controllers.NewGuardianAPI()
	us := controllers.NewUserController(getSession())

	api := iris.New()
	api.Get("/", gapi.Search)

	// Create User
	api.Get("/user", us.Newuser)

	api.Build()
	fsrv := &fasthttp.Server{Handler: api.Router}
	fsrv.ListenAndServe(":9999")
}

// addopted from https://github.com/swhite24/go-rest-tutorial/blob/master/server.go
// getSession creates a new mongo session and panics if connection error occurs
func getSession() *mgo.Session {
	// Connect to our local mongo
	s, err := mgo.Dial("localhost")

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}

	// Deliver session
	return s
}