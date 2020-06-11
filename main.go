package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"bytes"
	"log"
	"os/exec"	
	)

func main()  {

	fmt.Println("starting service...")

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/ps", func(c *gin.Context) {

		cmd := exec.Command("ps")
		var out bytes.Buffer		
		cmd.Stdout = &out
	
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
		c.String(http.StatusOK, fmt.Sprintf((out.String())))
	})

	r.GET("/ls", func(c *gin.Context) {

		cmd := exec.Command("ls")
		var out bytes.Buffer		
		cmd.Stdout = &out
	
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
		c.String(http.StatusOK, fmt.Sprintf((out.String())))
	})

	r.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")

		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})	

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	
}