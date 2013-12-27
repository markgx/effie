package handlers

import (
	"effie/repositories"
	"fmt"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/codegangsta/martini-contrib/sessions"
	"github.com/coopernurse/gorp"
	"net/http"
)

type LoginForm struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func Login(r render.Render) {
	r.HTML(200, "login", nil)
}

func LoginPost(w http.ResponseWriter, req *http.Request, loginForm LoginForm, dbmap *gorp.DbMap, session sessions.Session) string {
	userRepository := repositories.UserRepository{DbMap: dbmap}

	// TODO: verify log in

	user, err := userRepository.FindByUsername(loginForm.Username)

	if err != nil {
		panic(err)
	}

	if user != nil {
		session.Set("user_id", user.ID)

		if returnUrl := session.Get("return_url"); returnUrl != nil {
			session.Set("return_url", nil)
			http.Redirect(w, req, returnUrl.(string), 301)
		}

		http.Redirect(w, req, "/admin", 301)
	}

	return fmt.Sprintf("U:%s P:%s v:%+v", loginForm.Username, loginForm.Password, user)
}

func LogOut(w http.ResponseWriter, req *http.Request, session sessions.Session) {
	session.Delete("user_id")
	http.Redirect(w, req, "/login", 301)
}
