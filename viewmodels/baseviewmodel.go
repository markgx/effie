package viewmodels

import (
	"github.com/codegangsta/martini-contrib/sessions"
)

type BaseViewModel struct {
	session         sessions.Session
	Message         string
	ErrorMessage    string
	IsAuthenticated bool
	CurrentUsername string
}

func NewBaseViewModel(s sessions.Session) *BaseViewModel {
	bvm := BaseViewModel{session: s}

	if s != nil {
		username := s.Get("username")
		if username != nil {
			bvm.IsAuthenticated = true
			bvm.CurrentUsername = username.(string)
		}
	}

	return &bvm
}
