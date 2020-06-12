package Parametro

import (
    "fmt"
    "os"
    "gopkg.in/ini.v1"
)

const parFile = "./go-infraHelper.ini";

func GetParametro(chave string ) string {

	cfg, err := ini.Load(parFile)
    if err != nil {
        fmt.Printf("Fail to read file: %v", err)
        os.Exit(1)
	}
	
	return cfg.Section("").Key(chave).String();

}