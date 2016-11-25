// adopted from https://github.com/swhite24/go-rest-tutorial/blob/master/models/user.go
package controllers

import (
	"gopkg.in/mgo.v2"
	"GO-SPA/models"
	"github.com/kataras/iris"
	"log"
)

type (
	// UserController represents the controller for operating on the User resource
	UserController struct {
		session *mgo.Session
	}
)

// NewUserController provides a reference to a UserController with provided mongo session
func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

// adobted from https://labix.org/mgo and modified
func (us UserController) Newuser(ctx *iris.Context){

	u := models.User{"Will Hogan", "Man", 12} // thios is an emulator of data

	// path to the database.collection if there is no collection
	// with particular name then the collection created
	// same with database -> it is created if database
	// with the same name does not exist
	c := us.session.DB("gospa").C("users")

	// mgo package => inserts the data into the mongodb
	err := c.Insert(u)
	if err != nil {log.Fatal(err)}

	ctx.Next()
}