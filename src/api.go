package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/ant0ine/go-json-rest/rest"
)

type Routing struct {
	sync.RWMutex
}

type Students struct {
	StudentList []Student `json:"studentList"`
}

type Student struct {
	ID                   string `json:"id"`
	Name                 string `json:"name"`
	LastName             string `json:"lastName"`
	BasicEnglish1        string `json:"basicEnglish1"`
	Calculus1            string `json:"calculus1"`
	ComputerGraphics     string `json:"computerGraphics"`
	SoftwareArchitecture string `json:"softwareArchitecture"`
	DatabaseSystem       string `json:"databaseSystem"`
}

func (rt *Routing) GetStudentInfo(w rest.ResponseWriter, r *rest.Request) {
	url := "http://localhost:8882/api/student"
	response, err := http.Get(url)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer response.Body.Close()

	studentList := Students{}

	json.NewDecoder(response.Body).Decode(&studentList)
	w.WriteJson(studentList)
}

func main() {
	routing := Routing{}

	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Get("/api/student", routing.GetStudentInfo),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}
