package models

type Post struct {
	ID    string `gorethink:"id,omitempty"`
	Title string `gorethink:"title"`
	Body  string `gorethink:"body"`
}
