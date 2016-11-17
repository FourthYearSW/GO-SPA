// Author: Andrej Lavrinovic & Will Hogan & Christy Madden
// Date: November 07th, 2016
// Adapted from: https://go-macaron.com

package main

import (
	"net/http"

	"gopkg.in/macaron.v1"
)

func main() {

	m := macaron.Classic()
	m.Use(macaron.Renderer())

	// Basic Hello world
	m.Get("/hello", func() string {
		return "Hello world!"
	})

	m.Get("/hello", func(resp http.ResponseWriter, req *http.Request) {
		// resp and req are injected by Macaron
		resp.WriteHeader(200) // HTTP 200
	})

	// Basic Hello world
	m.Get("/helloandrej", func() string {
		return "Hello Andrej"
	})

	// Using the template from template folder
	m.Get("/helloagain", func(ctx *macaron.Context) {
		ctx.Data["Name"] = "Person"
		ctx.HTML(200, "hello") // 200 is the response code.
	})

	m.Run(8080) // Start Web Server on port 8080

}
