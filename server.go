package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/binding"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/codegangsta/martini-contrib/sessions"
	"net/http"
)

type Config struct {
	Database
}

type Database struct {
	DSN string
}

type LoginForm struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func main() {
	var config Config

	if _, err := toml.DecodeFile("config.toml", &config); err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return
	}

	m := martini.Classic()

	store := sessions.NewCookieStore([]byte("changethis"))
	m.Use(sessions.Sessions("effie_session", store))

	m.Use(render.Renderer(render.Options{
		Directory: "templates",
		Layout:    "layout",
	}))

	m.Get("/admin", func(w http.ResponseWriter, req *http.Request, session sessions.Session, r render.Render) {
		u := session.Get("user_id")

		if u == nil {
			http.Redirect(w, req, "/login", 301)
		}

		r.HTML(200, "admin_index", nil)
	})

	m.Get("/login", func(r render.Render) {
		r.HTML(200, "login", nil)
	})

	m.Post("/login", binding.Form(LoginForm{}), func(loginForm LoginForm) string {
		// TODO: verify log in

		// TODO: if success, set session

		return fmt.Sprintf("U:%s P:%s", loginForm.Username, loginForm.Password)
	})

	m.Run()
}
