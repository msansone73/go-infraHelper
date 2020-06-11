package Model

// Projeto do cliente
type Projeto struct {
    Id   		int
	Nome 		string
	Descricao	string
	Git			string
	Path		string
}

// construtor
func NewProjeto(id int,
		nome string, 
		descricao string, 
		git string,
		path string) *Projeto {
	p:= Projeto{Id:id}
	p.Nome=nome
	p.Descricao=descricao
	p.Git=git
	p.Path=path
	return &p
}