package story

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

// NewHandler returns a handler for this adventure
func NewHandler(a Adventure) http.Handler {
	h := handler{a}
	return h
}

type handler struct {
	adventure Adventure
}

var chapterTemplate = template.Must(template.ParseFiles("story/chapter.html"))

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimLeft(r.URL.Path, "/")
	if path == "" || path == "/" {
		path = "intro"
	}
	if chapter, ok := h.adventure[path]; ok {
		err := chapterTemplate.Execute(w, chapter)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	fmt.Fprintf(w, "chapter %s not found", path)
}
