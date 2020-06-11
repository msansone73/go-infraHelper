package repositorio

import(
	"encoding/json"
	"fmt"
	"io/ioutil"
	"projects/go-infraHelper/Model"
)

type Projetos struct {
	Item 		[]Model.Projeto
}

func FindAll() []Model.Projeto{ 

	data := Projetos{}	
	//file, _ := ioutil.ReadFile("/Users/msansone/go/src/projects/go-infraHelper/projetos.json")
 
	file, _ := ioutil.ReadFile("./projetos.json")

	err := json.Unmarshal([]byte(file), &data)
	if (err!=nil) {
		fmt.Println("ERRO = ",err)
	}
	return data.Item

}