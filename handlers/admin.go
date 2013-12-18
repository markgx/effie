package handlers

import (
	"github.com/codegangsta/martini-contrib/render"
	"net/http"
)

func AdminHome(w http.ResponseWriter, r render.Render) {
	r.HTML(200, "admin_index", nil)
}
