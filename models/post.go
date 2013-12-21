package models

import (
	"github.com/coopernurse/gorp"
)

type Post struct {
	ID    int
	Title string
	Body  string
}

type PostRepository struct {
	*gorp.DbMap
}

func (r *PostRepository) All() (*[]Post, error) {
	var posts []Post

	if _, err := r.DbMap.Select(&posts, "SELECT * FROM posts"); err != nil {
		return nil, err
	}

	return &posts, nil
}

func (r *PostRepository) FindByID(id int) (*Post, error) {
	var post Post

	if err := r.DbMap.SelectOne(&post, "SELECT * FROM posts WHERE ID=?", id); err != nil {
		return nil, err
	}

	return &post, nil
}
