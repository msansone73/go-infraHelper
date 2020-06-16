package main 

import (
	"github.com/gin-gonic/gin"
	)

	func setRotas(r *gin.Engine){

	r.Static("/main", "../templates/main")
	r.LoadHTMLGlob("templates/**/*")

	//info
	r.GET("/", cmd_home)
	r.GET("/about", cmd_about)

	// test
	r.GET("/ping", pong)

	//
	r.GET("/ps", cmd_ps)
	r.GET("/ls", cmd_ls)
	r.GET("/welcome", cmd_welcome)	
	
}