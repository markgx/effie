package main

import (
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/codegangsta/martini-contrib/sessions"
)

func main() {
	m := martini.Classic()

	store := sessions.NewCookieStore([]byte("changethis"))
	m.Use(sessions.Sessions("effie_session", store))

	m.Use(render.Renderer(render.Options{
		Directory: "templates",
		Layout:    "layout",
	}))

	m.Get("/admin", func(r render.Render) {
		r.HTML(200, "admin_index", nil)
	})

	m.Run()
}
