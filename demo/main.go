package main

import (
	"github.com/samthor/light"

	"log"
)

func main() {
	light.Add(light.Task{
		Color: &light.Red,
	})

	color, err := light.Update()
	log.Printf("set color: %+v", color)
	log.Fatal(err)
}
