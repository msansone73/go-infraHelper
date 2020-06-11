package main

import (
	//"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)


func cmd_index(c *gin.Context) {

	response := "Bem vindo, marcio"

	//c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(response))
	c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
		"title": response,
	})

}

func cmd_about(c *gin.Context) {
	firstname := c.DefaultQuery("firstname", "Guest")
	lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")

	c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
}