package models

type Post struct {
	ID    int    `gorethink:"id,omitempty"`
	Title string `gorethink:"title"`
	Body  string `gorethink:"body"`
}
