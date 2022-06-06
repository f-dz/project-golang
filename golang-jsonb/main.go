package main

import (
	"golang-jsonb/controller"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()
	url := "/apiumur"
	urlWithName := url + `/{name}`

	r.Get("/people", controller.GetAllData)
	r.Post(url, controller.CreateData)
	r.Get(urlWithName, controller.GetData)
	r.Put(urlWithName, controller.UpdateData)
	r.Delete(urlWithName, controller.DeleteData)

	http.ListenAndServe(":8080", r)
}
