package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var people = []Person{
	{Name: "Tom", Age: 12},
	{Name: "Jerry", Age: 13},
}

func main() {
	var url = "/apiumur"
	r := chi.NewRouter()

	r.Get(url+"/people", getAllData)
	r.Get(url, getData)
	r.Post(url, createData)

	http.ListenAndServe(":8080", r)
}

func getAllData(w http.ResponseWriter, r *http.Request) {
	output, err := json.MarshalIndent(people, "", "   ")
	if err != nil {
		log.Fatal(err)
	}
	w.Write([]byte(output))
}

func getData(w http.ResponseWriter, r *http.Request) {
	var person Person

	name := r.URL.Query().Get("name")
	person.Name = strings.Replace(name, `"`, "", -1)

	age := r.URL.Query().Get("age")
	person.Age, _ = strconv.Atoi(age)

	output, err := json.MarshalIndent(person, "", "   ")
	if err != nil {
		log.Fatal(err)
	}
	w.Write([]byte(output))
}

func createData(w http.ResponseWriter, r *http.Request) {
	var person Person
	json.NewDecoder(r.Body).Decode(&person)
	people = append(people, person)

	output, err := json.MarshalIndent(person, "", "   ")
	if err != nil {
		log.Fatal(err)
	}
	w.Write([]byte(output))
}
