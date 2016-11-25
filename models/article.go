package models


// Adapted from https://github.com/swhite24/go-rest-tutorial/blob/master/models/user.go
type (
	// Article is a struct which holds details about the articles and can be marshalled into json and bson
	Article struct {
		ID   string `json:"id" bson:"_id"`
		Name string `json:"name" bson:"name"`
		URL  string `json:"url" bson:"url"`
	}
)