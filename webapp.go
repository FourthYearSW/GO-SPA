// Author: Andrej Lavrinovic & Will Hogan
// Date: November 07th, 2016
// Adapted from: https://go-macaron.com

package main

import "gopkg.in/macaron.v1"

func main() {
	m := macaron.Classic()
	m.Use(macaron.Renderer())

	// Basic Hello world
	m.Get("/hello", func() string {
		return "Hello world!"
	})

	// Using the template from template folder
	m.Get("/helloagain", func(ctx *macaron.Context) {
		ctx.Data["Name"] = "Person"
		ctx.HTML(200, "hello") // 200 is the response code.
	})

	m.Run(8080) // Start Web Server on port 8080
}
