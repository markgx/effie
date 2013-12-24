package repositories

import (
	"effie/models"
	"github.com/coopernurse/gorp"
)

type PostRepository struct {
	*gorp.DbMap
}

func (r *PostRepository) All() (*[]models.Post, error) {
	var posts []models.Post

	if _, err := r.DbMap.Select(&posts, "SELECT * FROM posts"); err != nil {
		return nil, err
	}

	return &posts, nil
}

func (r *PostRepository) FindByID(id int) (*models.Post, error) {
	var post models.Post

	if err := r.DbMap.SelectOne(&post, "SELECT * FROM posts WHERE ID=?", id); err != nil {
		return nil, err
	}

	return &post, nil
}

func (r *PostRepository) Create(post *models.Post) error {
	return r.DbMap.Insert(post)
}
