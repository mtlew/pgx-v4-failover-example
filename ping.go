package main

import (
	"fmt"
	"log"
	"time"
)

type ping struct {
	storage *storage
}

func newPing(storage *storage) *ping {
	return &ping{storage: storage}
}

func (p *ping) run(pid int) {
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
