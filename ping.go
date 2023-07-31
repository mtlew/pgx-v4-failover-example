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

func (p *ping) run() {
	var row test
	var err error

	for {
		row, err = p.storage.first()

		if err != nil {
			log.Println(fmt.Sprintf("Background: Error, %s!", err.Error()))
			continue
		}
		log.Println(fmt.Sprintf("Background: ID: %s, Title: %s", row.ID, row.Title))

		time.Sleep(10 * time.Millisecond)
	}
}
