package main

import (
	"fmt"
	"net/http"

	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

func hello(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, Goji")
}

func main() {
	goji.Get("/", hello)
	goji.Serve()
}
