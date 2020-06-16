package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/msansone73/go-infraHelper/Repositorio"
)

func cmd_home(c *gin.Context) {
	projetos := Repositorio.FindAll()

	response := "Bem vindo, marcio"
	c.HTML(http.StatusOK, "main/index.tmpl", gin.H{
		"cumprimento": response,
		"Projetos":    projetos,
	})
}

func cmd_about(c *gin.Context) {

	response := "Bem vindo, marcio"
	//c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	c.HTML(http.StatusOK, "main/about.tmpl", gin.H{
		"cumprimento": response,
	})
}
