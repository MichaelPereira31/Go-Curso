package main

import (
	"os"
	"text/template"
)

type Curso struct {
	Nome        string
	CargaHoraria int
}

func main() {
	curso := Curso{
		Nome:         "Go Avançado",
		CargaHoraria: 40,
	}
	t := template.Must(template.New("CursoTemplate").Parse("Curso: {{.Nome}}, Carga Horária: {{.CargaHoraria}} horas"))
	err := t.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}

}