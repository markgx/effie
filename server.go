package main

import (
	"./models"
	"database/sql"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/binding"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/codegangsta/martini-contrib/sessions"
	"github.com/coopernurse/gorp"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
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

	m.Use(DB(config.Database.DSN))

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

	m.Post("/login", binding.Form(LoginForm{}), func(w http.ResponseWriter, req *http.Request, loginForm LoginForm, dbmap *gorp.DbMap, session sessions.Session) string {
		userRepository := models.UserRepository{DbMap: dbmap}

		// TODO: verify log in

		user, _ := userRepository.FindByUsername(loginForm.Username)

		if user != nil {
			session.Set("user_id", user.ID)
			http.Redirect(w, req, "/admin", 301)
		}

		return fmt.Sprintf("U:%s P:%s v:%+v", loginForm.Username, loginForm.Password, user)
	})

	m.Run()
}
