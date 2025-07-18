package main

import (
	"html/template"
	"net/http"
	"strings"
)

type Curso struct {
	Nome        string
	CargaHoraria int
}

type Cursos []Curso

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func main() {
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.New("content.html")
		t.Funcs(template.FuncMap{
			"ToUpper": ToUpper,
		})
		t = template.Must(t.ParseFiles(templates...))
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