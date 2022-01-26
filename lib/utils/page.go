package utils

import (
	"html/template"
	"net/http"
)

const ROOT_LAYOUT_FILE = "html/layouts/root.go.html"
const NAV_LAYOUT_FILE = "html/layouts/nav.go.html"

type Page struct {
	PageFilePath string
}

func NewPage(path string) *Page {
	return &Page{
		PageFilePath: path,
	}
}

func (p *Page) Render(w http.ResponseWriter, r *http.Request, d map[string]interface{}) {
	if d == nil {
		d = make(map[string]interface{})
	}

	d["IsSignedIn"] = IsSignedIn(r)

	t, err := template.ParseFiles(ROOT_LAYOUT_FILE, NAV_LAYOUT_FILE, p.PageFilePath)
	if err != nil {
		panic(err)
	}
	t.Execute(w, d)
}
