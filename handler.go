package main

import (
	"fmt"
	"net/http"
)

type handler struct {
	storage *storage
}

func newHandler(storage *storage) *handler {
	return &handler{storage: storage}
}

func (h *handler) init() *handler {
	http.HandleFunc("/", h.index)
	http.HandleFunc("/first", h.first)
	http.HandleFunc("/list", h.list)
	return h
}

func (h *handler) run() error {
	return http.ListenAndServe(":8080", nil)
}

func (h *handler) first(w http.ResponseWriter, _ *http.Request) {
	row, err := h.storage.first()
	if err != nil {
		fmt.Fprintf(w, "Error, %s!", err.Error())
		return
	}
	fmt.Fprintf(w, "ID: %s, Title: %s\n", row.ID, row.Title)
}

func (h *handler) list(w http.ResponseWriter, _ *http.Request) {
	rows, err := h.storage.list()
	if err != nil {
		fmt.Fprintf(w, "Error, %s!", err.Error())
		return
	}
	for _, row := range rows {
		fmt.Fprintf(w, "ID: %s, Title: %s\n", row.ID, row.Title)
	}
}

func (h *handler) index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
