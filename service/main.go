package main

import (
	"log"
	"os"
)

func main() {
	log.Println("Received args:")
	log.Println(os.Args[1:])
}
