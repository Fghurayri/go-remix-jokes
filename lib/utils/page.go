package utils

import (
	"html/template"
	"net/http"
)

const ROOT_LAYOUT_FILE = "pages/layouts/root.go.html"
const NAV_LAYOUT_FILE = "pages/layouts/nav.go.html"

type Page struct {
	PageFilePath string
}

func NewPage(path string) *Page {
	return &Page{
		PageFilePath: path,
	}
}

func (p *Page) Render(w http.ResponseWriter, data interface{}) {
	t, err := template.ParseFiles(ROOT_LAYOUT_FILE, NAV_LAYOUT_FILE, p.PageFilePath)
	if err != nil {
		panic(err)
	}
	t.Execute(w, data)
}
