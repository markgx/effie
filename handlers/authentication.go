package handlers

import (
	"crypto/sha256"
	"effie/repositories"
	"encoding/base64"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/codegangsta/martini-contrib/sessions"
	r "github.com/dancannon/gorethink"
	"net/http"
)

type LoginForm struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func Login(r render.Render) {
	r.HTML(200, "login", struct{ ErrorMessage string }{})
}

func LoginPost(w http.ResponseWriter, req *http.Request, loginForm LoginForm, rdbSession *r.Session,
	session sessions.Session, r render.Render) {
	userRepository := repositories.UserRepository{Session: rdbSession}

	user, err := userRepository.FindByUsername(loginForm.Username)

	if err != nil {
		panic(err)
	}

	if user != nil && hashString(loginForm.Password) == user.PasswordHash {
		session.Set("user_id", user.ID)
		session.Set("username", user.Username)

		if returnUrl := session.Get("return_url"); returnUrl != nil {
			session.Set("return_url", nil)
			http.Redirect(w, req, returnUrl.(string), 301)
		}

		http.Redirect(w, req, "/admin", 301)
	}

	// authentication failed, show error
	r.HTML(200, "login", struct{ ErrorMessage string }{
		ErrorMessage: "Your username and/or password was incorrect",
	})
}

func LogOut(w http.ResponseWriter, req *http.Request, session sessions.Session) {
	session.Delete("user_id")
	session.Delete("username")
	http.Redirect(w, req, "/login", 301)
}

func hashString(s string) string {
	hb := sha256.Sum256([]byte(s))
	h := base64.URLEncoding.EncodeToString(hb[:])
	return h
}
