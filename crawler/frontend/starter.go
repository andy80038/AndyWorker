package main

import (
	"net/http"

	"github.com/andy80038/AndyWorker/crawler/frontend/controller"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("crawler/frontend/view")))

	http.Handle("/search", controller.CreateSearchResultHandler(
		"crawler/frontend/view/template.html",
	))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
