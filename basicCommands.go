package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	)


func pong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func cmd_welcome(c *gin.Context) {
	firstname := c.DefaultQuery("firstname", "Guest")
	lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")

	c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
}