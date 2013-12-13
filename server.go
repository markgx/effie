package main

import (
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	"net/http"
)

func main() {
	m := martini.Classic()

	m.Use(render.Renderer(render.Options{
		Directory: "templates",
		Layout:    "layout",
	}))

	m.Get("/", rootHandler)
	m.Run()
}

func rootHandler(r render.Render) {
	r.HTML(200, "root", nil)
}
