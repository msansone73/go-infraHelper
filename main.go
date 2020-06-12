package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/msansone73/go-infraHelper/Parametros"
)

func main()  {

	fmt.Println("starting service...")
	fmt.Println("USER=", Parametros.GetParametro("user"))

	r := gin.Default()
	setRotas(r)

	err := r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if (err!=nil){
		fmt.Println("Erro = ", err)
	}
	
}