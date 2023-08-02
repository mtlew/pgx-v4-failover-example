package main

import (
	"fmt"
	"log"
	"time"
)

type daemon struct {
	storage *storage
}

func newDaemon(storage *storage) *daemon {
	return &daemon{storage: storage}
}

func (p *daemon) run(pid int) {
	var first test
	var err error

	for {
		first, err = p.storage.first()
		if err != nil {
			log.Println(fmt.Sprintf("Background (%d): err: %s", pid, err))
		} else {
			log.Println(fmt.Sprintf("Background (%d): ID: %s, Title: %s", pid, first.ID, first.Title))
		}
		time.Sleep(time.Millisecond)
	}
}
