package main

import (
	"flag"
	"fmt"
)

var background = flag.String("bg", "#000000", "backgrouind color")

func init() {
	flag.Parse()
}

func draw() {
	fmt.Printf("drawing with background %s\n", *background)
}

func main() {
	draw()
}
