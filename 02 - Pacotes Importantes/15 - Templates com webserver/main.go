package main

import (
	"net/http"
	"text/template"
)

type Curso struct {
	Nome        string
	CargaHoraria int
}

type Cursos []Curso

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.New("template.html").ParseFiles("template.html"))
		err := t.Execute(w, Cursos{
			{"GO", 40},
			{"Java", 30},
			{"C#", 20},
		})
		if err != nil {
			http.Error(w, "Erro ao processar o template", http.StatusInternalServerError)
			return
		}
	})
	http.ListenAndServe(":8080", nil)

}