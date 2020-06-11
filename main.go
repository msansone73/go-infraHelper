package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	)

func main()  {

	fmt.Println("starting service...")
	r := gin.Default()
	setRotas(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	
}