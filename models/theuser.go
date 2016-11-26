// adopted from https://github.com/swhite24/go-rest-tutorial/blob/master/models/user.go
package models

import "gopkg.in/mgo.v2/bson"

type (
	// User represents the structure of our resource
	Theuser struct {
		Id     bson.ObjectId `json:"id" bson:"_id"`
		Name     string `json:"name" bson:"name"`
		Email    string `json:"email" bson:"email"`
		Password string `json:"password" bson:"password"`
	}
)