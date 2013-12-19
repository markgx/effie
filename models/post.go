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
