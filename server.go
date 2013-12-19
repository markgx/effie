package main

import (
	"database/sql"
	"effie/handlers"
	"effie/middleware"
	"effie/models"
	"github.com/BurntSushi/toml"
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/binding"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/codegangsta/martini-contrib/sessions"
	"github.com/coopernurse/gorp"
	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	Database
}

type Database struct {
	DSN string
}

func DB(dsn string) martini.Handler {
	return func(c martini.Context) {
		db, err := sql.Open("mysql", dsn)

		if err != nil {
			panic(err)
		}

		dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{}}
		dbmap.AddTableWithName(models.User{}, "users").SetKeys(true, "ID")

		c.Map(dbmap)
		defer db.Close()
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
	m.Post("/admin/posts", middleware.Authenticate, handlers.CreatePost)

	m.Get("/login", handlers.Login)
	m.Post("/login", binding.Form(handlers.LoginForm{}), handlers.LoginPost)
	m.Get("/logout", handlers.LogOut)
}
