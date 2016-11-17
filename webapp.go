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

	// Get the template page
	m.Get("/", func(ctx *macaron.Context) {
		ctx.Data["Name"] = "Andrej"
		ctx.HTML(200, "hello") // 200 is the response code.
	})

	// Req.body will recieve the post
	m.Get("/hello", func(resp http.ResponseWriter, req *http.Request) {
		// resp and req are injected by Macaron
		resp.WriteHeader(200) // HTTP 200

	})

	// Start Web Server on port 8080
	m.Run(8080)

}
