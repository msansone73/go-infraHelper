package main

import (
	"fmt"
	"projects/go-infraHelper/parametro"
	"github.com/gin-gonic/autotls"
	//"golang.org/x/crypto/acme/autocert"
	"github.com/gin-gonic/gin"
	"log"
)

func main()  {

	fmt.Println("starting service...")

	fmt.Println("USER=", parametro.GetParametro("user"))

	r := gin.Default()
	setRotas(r)

	// m := autocert.Manager{
	// 	Prompt:     autocert.AcceptTOS,
	// 	HostPolicy: autocert.HostWhitelist("msansone.com.br"),
	// 	Cache:      autocert.DirCache("/var/www/.cache"),
	// }

	//log.Fatal(autotls.RunWithManager(r, &m))
	log.Fatal(autotls.Run(r, "msansone.com.br"))

	//r.Run(":443") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	
}