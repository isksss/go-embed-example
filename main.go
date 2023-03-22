package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

//go:embed public/*
var FS embed.FS

func main() {
	public, err := fs.Sub(FS, "public")
	if err != nil {
		panic(err)
	}

	http.Handle("/",
		http.FileServer((http.FS(public))),
	)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
