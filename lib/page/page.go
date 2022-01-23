package page

import (
	"html/template"
	"net/http"
)

const ROOT_DIR = "pages/layouts/root.go.html"

type Page struct {
	PageFilePath string
}

func NewPage(path string) *Page {
	return &Page{
		PageFilePath: path,
	}
}

func (p *Page) Render(w http.ResponseWriter, data interface{}) {
	t, err := template.ParseFiles(ROOT_DIR, p.PageFilePath)
	if err != nil {
		panic(err)
	}
	t.Execute(w, data)
}
