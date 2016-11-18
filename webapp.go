// Author: Will Hogan
// Date: 18th Nov, 2016
// Countdown Timer Adapted from https://play.golang.org/p/jl5VwaurB5

package main

import (
	"flag"
	"fmt"
	"time"

	"gopkg.in/macaron.v1"
)

func main() {
	m := macaron.Classic()
	m.Use(macaron.Renderer())

	m.Get("/name", func(ctx *macaron.Context) {
		ctx.Data["Name"] = "Will"
		ctx.HTML(200, "hello") // 200 is the response code.
	})

	// For testing purposes, uses the 'timer.tmpl'
	m.Get("/timer", func(ctx *macaron.Context) {
		ctx.Data["timer"] = "Countdown in progress..."
		countdown() //  Trying to figure out how this will work with html page
		ctx.HTML(200, "timer")
	})
	m.Run(8080)
}

// Countdown Timer Adapted from https://play.golang.org/p/jl5VwaurB5
func countdown() {
	// Create a pointer named duration
	// See http://golang.org/pkg/flag/
	var duration = flag.Duration("duration", 5*time.Second, "set timer duration, for example 60s, 20m, .5h, or 1d")
	flag.Parse()

	// See http://golang.org/pkg/time/#Time.Format
	const layout = "15:04:05"
	t := time.Now()
	fmt.Println(t.Format(layout))

	timer := time.NewTimer(*duration)

	<-timer.C
	fmt.Println("1 Hour is up (obvious shorter test time of 5 seconds set for testing), need to reset and make new call to API...")
	t = time.Now()
	fmt.Println(t.Format(layout))
}
