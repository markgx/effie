package middleware

import (
	"github.com/codegangsta/martini-contrib/sessions"
	"net/http"
)

func Authenticate(w http.ResponseWriter, req *http.Request, session sessions.Session) {
	u := session.Get("user_id")

	if u == nil {
		session.Set("return_url", req.URL.String())
		http.Redirect(w, req, "/login", 301)
	}
}
