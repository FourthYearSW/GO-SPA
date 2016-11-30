// adopted from https://github.com/swhite24/go-rest-tutorial/blob/master/models/user.go
package models

//import "time"

// Adapted from https://github.com/swhite24/go-rest-tutorial/blob/master/models/user.go
type (
	// Comment is a struct which holds details about the articles and can be marshalled into json and bson
	Comment struct {
		ID      string `json:"id" bson:"_id"`
		Comment string `json:"comment" bson:"comment"`
	}
)