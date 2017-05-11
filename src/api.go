package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/ant0ine/go-json-rest/rest"
)

type Routing struct {
	sync.RWMutex
}

func (rt *Routing) Index(w rest.ResponseWriter, r *rest.Request) {
	w.WriteJson(map[string]string{
		"id":   "1",
		"name": "Golang",
	})
}

func main() {
	routing := Routing{}

	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Get("/", routing.Index),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}
