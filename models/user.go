package models

type User struct {
	ID       int    `gorethink:"id,omitempty"`
	Username string `gorethink:"username"`
	Password string `gorethink:"password"`
}
