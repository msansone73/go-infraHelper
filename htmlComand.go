package main

import (
	_ "encoding/json"
	_ "fmt"
	"net/http"
	"projects/go-infraHelper/repository"
	"github.com/gin-gonic/gin"
)

func cmd_home(c *gin.Context)  {
	projetos := repositorio.FindAll()

	response := "Bem vindo, marcio"
	c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
		"cumprimento": response,
		"Projetos":    projetos,
	})
}

func cmd_about(c *gin.Context) {
	firstname := c.DefaultQuery("firstname", "Guest")
	lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")

	c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
}
