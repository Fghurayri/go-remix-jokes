package page

import (
	"html/template"
	"net/http"
)

type Page struct {
	PageFilePath string
}

func NewPage(path string) *Page {
	return &Page{
		PageFilePath: path,
	}
}

func (p *Page) Render(w http.ResponseWriter, data interface{}) {
	t, err := template.ParseFiles("pages/layouts/root.go.html", p.PageFilePath)
	if err != nil {
		panic(err)
	}
	t.Execute(w, data)
}
