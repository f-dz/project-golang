package main

import (
	"strconv"
	"net/http"
	"github.com/gin-gonic/gin"
)

type Person struct {
	Name string `json:"name"`
	Age int `json:"age"`
}


func main() {
	var url = "/apiumur"
	r:= gin.Default()

	r.GET(url + "/people", getAllData)
	r.GET(url, getData)

	r.Run()
}

func getAllData(c *gin.Context) {
	var people = []Person{
		{Name: "Tom", Age: 22},
		{Name: "Jerry", Age: 23},
	}
	c.IndentedJSON(http.StatusOK, gin.H{"data": people})
}

func getData(c *gin.Context) {
	var person Person
	person.Name = c.Query("name")
	age := c.Query("age")
	person.Age, _ = strconv.Atoi(age)

	c.IndentedJSON(http.StatusOK, gin.H{
		"name": person.Name,
		"age": person.Age,
	})
}