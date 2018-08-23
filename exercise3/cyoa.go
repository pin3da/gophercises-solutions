package main

import (
	"flag"
	"net/http"

	"./story"
)

func main() {
	mux := http.NewServeMux()

	filename := flag.String("filename", "gopher.json", "path to the chapters")
	flag.Parse()

	adventure, err := story.LoadArcsFromFile(*filename)

	if err != nil {
		panic(err)
	}

	mux.Handle("/", story.NewHandler(adventure))
	http.ListenAndServe(":8080", mux)
}
