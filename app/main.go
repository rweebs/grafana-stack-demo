package main

import (
	"log"
	"time"
)

func main() {
	for {
		log.Println("Helllo, This is from the application start up.")
		time.Sleep(1 * time.Second)
	}
}
