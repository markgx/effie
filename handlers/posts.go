package handlers

import (
	"github.com/codegangsta/martini-contrib/render"
	"net/http"
)

func NewPost(w http.ResponseWriter, req *http.Request, r render.Render) {
	r.HTML(200, "post_form", nil)
}

func CreatePost(w http.ResponseWriter, req *http.Request, r render.Render) {
	// TODO: validate and save

	http.Redirect(w, req, "/admin", 301)
}
