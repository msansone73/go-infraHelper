package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func main() {
	//cmd := exec.Command("ps")
	cmd := exec.Command("ls", "/Users/msansone/go/src/projects/infraHelper")
    var out bytes.Buffer
	var saida bytes.Buffer
	
	cmd.Stdout = &out
	cmd.Stderr = &saida

    err := cmd.Run()
    if err != nil {
        log.Fatal(err)
	}
	fmt.Println("saida:")
	fmt.Printf("%q\n", out.String() )
	fmt.Printf("%q\n", saida.String() )
	fmt.Println("fim.")
}




