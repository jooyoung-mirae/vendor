package main

import (
	"log"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/sys/unix"
)

func main() {
	log.Println(unix.Getpid)
}
