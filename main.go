package main

import (
	"fmt"
	"projects/go-infraHelper/parametro"

	"github.com/gin-gonic/gin"
)

func main()  {

	fmt.Println("starting service...")

	fmt.Println("USER=", parametro.GetParametro("user"))

	r := gin.Default()
	setRotas(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	
}