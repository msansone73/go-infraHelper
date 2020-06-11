package main 

import (
	"github.com/gin-gonic/gin"
	)

	func setRotas(r *gin.Engine){

	r.LoadHTMLGlob("templates/**/*")

	r.GET("/", cmd_index)
	r.GET("/about", cmd_about)

	r.GET("/ping", pong)
	r.GET("/ps", cmd_ps)
	r.GET("/ls", cmd_ls)
	r.GET("/welcome", cmd_welcome)	
	
}