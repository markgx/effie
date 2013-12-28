package models

type User struct {
	ID       string `gorethink:"id,omitempty"`
	Username string `gorethink:"username"`
	Password string `gorethink:"password"`
}
