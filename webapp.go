// Author: Andrej Lavrinovic & Will Hogan
// Date: November 07th, 2016
// Adapted from: https://go-macaron.com

package main

import "gopkg.in/macaron.v1"

func main() {
  m := macaron.Classic()
  m.Use(macaron.Renderer())
  
  // Edit accordingly....
  m.Get("/hello", func() string {
    return "Hello from Macaron!"
  })
  m.Run(8080) // Start Web Server on port 8080
}