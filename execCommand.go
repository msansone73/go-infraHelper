package main


import (
	"bytes"
	"log"
	"net/http"
	"os/exec"
	"strings"

	"github.com/gin-gonic/gin"
)


func cmd_ps(c *gin.Context) {

	cmd := exec.Command("ps")
	var out bytes.Buffer		
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	//c.String(http.StatusOK, fmt.Sprintf((out.String())))
	response := "Bem vindo, marcio"
	resultado:= strings.TrimSpace(out.String())
	c.HTML(http.StatusOK, "main/list_ls.tmpl", gin.H{
		"cumprimento": response,
		"resultado": resultado,
	})	
}

func cmd_ls(c *gin.Context) {

	cmd := exec.Command("ls")
	var out bytes.Buffer		
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	//c.String(http.StatusOK, fmt.Sprintf((out.String())))

	response := "Bem vindo, marcio"
	resultado:= strings.TrimSpace(out.String())
	c.HTML(http.StatusOK, "main/list_ls.tmpl", gin.H{
		"cumprimento": response,
		"resultado": resultado,
	})
}