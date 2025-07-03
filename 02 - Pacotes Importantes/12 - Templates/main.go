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

	tmp := template.New("CursoTemplate")
	tmp, _ = tmp.Parse("Curso: {{.Nome}}, Carga Horária: {{.CargaHoraria}} horas")
	err := tmp.Execute(os.Stdout, curso)

	if err != nil {
		panic(err)
	}

}