package controller

import (
	"context"
	"encoding/json"
	"golang-jsonb/helper"
	"golang-jsonb/model"
	"golang-jsonb/model/entity"
	rep "golang-jsonb/model/repository"
	"net/http"

	"github.com/go-chi/chi"
)

func GetAllData(w http.ResponseWriter, r *http.Request) {
	db := model.GetConnection()
	defer db.Close()

	ctx := context.Background()

	personRepository := rep.NewPersonImplementation(db)
	result, err := personRepository.GetAllData(ctx)

	var output []byte
	if err != nil {
		w.WriteHeader(400)
		output = helper.WriteJSONMessage("Failed get all data")
		w.Write([]byte(output))
	} else {
		output = helper.WriteJSONPeople(result)
		w.Write([]byte(output))
	}

}

func GetData(w http.ResponseWriter, r *http.Request) {
	db := model.GetConnection()
	defer db.Close()

	ctx := context.Background()

	name := chi.URLParam(r, "name")

	personRepository := rep.NewPersonImplementation(db)
	result, err := personRepository.GetData(ctx, name)

	var output []byte
	if err != nil {
		output = helper.WriteJSONMessage("Failed get data with name: " + name)
		w.Write([]byte(output))
	} else {
		if result.Name != "" {
			output := helper.WriteJSONPerson(result)
			w.Write([]byte(output))
		} else {
			w.WriteHeader(404)
			output := helper.WriteJSONMessage("Name " + name + " not found")
			w.Write([]byte(output))
		}
	}
}

func CreateData(w http.ResponseWriter, r *http.Request) {
	db := model.GetConnection()
	defer db.Close()

	ctx := context.Background()

	var person entity.Person
	json.NewDecoder(r.Body).Decode(&person)

	personRepository := rep.NewPersonImplementation(db)

	existingData, err := personRepository.GetData(ctx, person.Name)

	var output []byte
	if err == nil {
		if existingData.Name == "" {
			result, err := personRepository.CreateData(ctx, person)
			if err != nil {
				w.WriteHeader(400)
				output = helper.WriteJSONMessage("Failed create data")
				w.Write([]byte(output))
			} else {
				w.WriteHeader(201)
				output = helper.WriteJSONPerson(result)
				w.Write([]byte(output))
			}
		} else {
			w.WriteHeader(400)
			output = helper.WriteJSONMessage("Name " + person.Name + " already exist")
			w.Write([]byte(output))
		}
	} else {
		w.WriteHeader(400)
		output = helper.WriteJSONMessage("Failed create data")
		w.Write([]byte(output))
	}
}

func UpdateData(w http.ResponseWriter, r *http.Request) {
	db := model.GetConnection()
	defer db.Close()

	ctx := context.Background()

	var person entity.Person
	json.NewDecoder(r.Body).Decode(&person)
	name := chi.URLParam(r, "name")

	personRepository := rep.NewPersonImplementation(db)
	result, err := personRepository.UpdateData(ctx, name, person.Age)

	var output []byte
	if err != nil {
		w.WriteHeader(400)
		output = helper.WriteJSONMessage("Failed update data")
		w.Write([]byte(output))
	} else {
		output := helper.WriteJSONPerson(result)
		w.Write([]byte(output))
	}
}

func DeleteData(w http.ResponseWriter, r *http.Request) {
	db := model.GetConnection()
	defer db.Close()

	ctx := context.Background()

	name := chi.URLParam(r, "name")

	personRepository := rep.NewPersonImplementation(db)
	err := personRepository.DeleteData(ctx, name)

	var output []byte
	if err != nil {
		w.WriteHeader(404)
		output = helper.WriteJSONMessage("Failed delete data")
		w.Write([]byte(output))
	} else {
		output = helper.WriteJSONMessage("Success delete data")
		w.Write([]byte(output))
	}
}
