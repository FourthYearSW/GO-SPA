// Author: Andrej Lavrinovic & Will Hogan & Christy Madden
// Date: November 07th, 2016
// Adapted from: https://go-macaron.com

package main

import "gopkg.in/macaron.v1"

func main() {
	m := macaron.Classic()
	m.Use(macaron.Renderer())

	m.Get("/", func(ctx *macaron.Context) {
		//ctx.Data["Name"] = "christy"

		// Give a status code and the partial html file you want to serve
		ctx.HTML(200, "header")
		ctx.HTML(200, "body")
		ctx.HTML(200, "footer")

	})

	// Start Web Server on port 8080
	m.Run(8090)
} // End main
