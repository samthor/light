package main

import (
	"github.com/samthor/light"

	"log"
)

func main() {
	light.Add(light.Task{
		Color: &light.Red,
	})

	err := light.Update()
	log.Fatal(err)
}
