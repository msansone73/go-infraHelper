package main

import (
	"encoding/json"
	_ "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Student struct {
	Id    int
	Name  string
	Idade int
}

const jsondata = `[
  {
    "id": 1,
	"name": "Mary",
	"idade": 25
	
  },
  {
    "id": 2,
	"name": "John",
	"idade": 40
  }
]`

func cmd_index(c *gin.Context) {

	var students []Student
	if err := json.Unmarshal([]byte(jsondata), &students); err != nil {
		panic(err)
	}
	response := "Bem vindo, marcio"
	c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
		"cumprimento": response,
		"Students":    students,
	})

}

func cmd_about(c *gin.Context) {
	firstname := c.DefaultQuery("firstname", "Guest")
	lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")

	c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
}
