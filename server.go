package main

import (
	"effie/handlers"
	"effie/middleware"
	"github.com/BurntSushi/toml"
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/binding"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/codegangsta/martini-contrib/sessions"
	r "github.com/dancannon/gorethink"
)

type Config struct {
	Database `toml:"database"`
}

type Database struct {
	DSN string `toml:"dsn"`
}

func DB(dsn string) martini.Handler {
	return func(c martini.Context) {
		session, err := r.Connect(map[string]interface{}{
			"address":  "localhost:28015",
			"database": "effie",
		})

		if err != nil {
			panic(err)
		}

		c.Map(session)
		defer session.Close()
		c.Next()
	}
}

func main() {
	var config Config
	loadConfig(&config)

	m := martini.Classic()

	store := sessions.NewCookieStore([]byte("changethis"))
	m.Use(sessions.Sessions("effie_session", store))

	m.Use(render.Renderer(render.Options{
		Directory: "templates",
		Layout:    "layout",
	}))

	m.Use(DB(config.Database.DSN))

	loadRoutes(m)
	m.Run()
}

func loadConfig(config *Config) {
	if _, err := toml.DecodeFile("config.toml", config); err != nil {
		panic(err)
	}
}

func loadRoutes(m *martini.ClassicMartini) {
	m.Get("/admin", middleware.Authenticate, handlers.AdminHome)

	m.Get("/admin/posts", middleware.Authenticate, handlers.PostsIndex)
	m.Get("/admin/posts/new", middleware.Authenticate, handlers.NewPost)
	m.Post("/admin/posts", middleware.Authenticate, binding.Form(handlers.PostForm{}), handlers.CreatePost)
	m.Get("/admin/posts/edit/:id", middleware.Authenticate, handlers.EditPost)

	m.Get("/login", handlers.Login)
	m.Post("/login", binding.Form(handlers.LoginForm{}), handlers.LoginPost)
	m.Get("/logout", handlers.LogOut)
}
