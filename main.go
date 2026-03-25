package main

import (
	"log"

	"golang.org/x/sys/unix"
)

func main() {
	log.Println(unix.Getpid)
}
