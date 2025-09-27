package main

import (
	"log"
	"time"
)

func main() {
	for {
		log.Println("Helllo, World!")
		time.Sleep(1 * time.Second)
	}
}
