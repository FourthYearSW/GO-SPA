// adopted from https://github.com/swhite24/go-rest-tutorial/blob/master/models/user.go
package controllers

import (
	"gopkg.in/mgo.v2"
	"fmt"
	"github.com/kataras/iris"
)

type (
	// UserController represents the controller for operating on the User resource
	UserController struct {
		session *mgo.Session
	}
)

/*
// NewUserController provides a reference to a UserController with provided mongo session
func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

// CreateUser creates a new user resource
func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Param) {
	// Stub an user to be populated from the body
	u := models.User{}

	// Populate the user data
	json.NewDecoder(r.Body).Decode(&u)

	// Add an Id
	u.Id = bson.NewObjectId()

	// Write the user to mongo
	uc.session.DB("gospa").C("users").Insert(u)

	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(u)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", uj)
}
*/

func (us UserController) CreateUser(ctx *iris.Context){
	fmt.Println("Connection to ther mongo db and createding records")
	ctx.Write("Connection to ther mongo db and createding records")

	//u := models.User{}
}