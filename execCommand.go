package main


import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"bytes"
	"log"
	"os/exec"	
	)


func cmd_ps(c *gin.Context) {

	cmd := exec.Command("ps")
	var out bytes.Buffer		
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	c.String(http.StatusOK, fmt.Sprintf((out.String())))
}

func cmd_ls(c *gin.Context) {

	cmd := exec.Command("ls")
	var out bytes.Buffer		
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	c.String(http.StatusOK, fmt.Sprintf((out.String())))
}