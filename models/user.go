// adopted from https://github.com/swhite24/go-rest-tutorial/blob/master/models/user.go
package models

type (
	// User represents the structure of our resource
	User struct {
		Name   string        `json:"name" bson:"name"`
		Gender string        `json:"gender" bson:"gender"`
		Age    int           `json:"age" bson:"age"`
	}
)
