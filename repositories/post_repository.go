package repositories

import (
	"effie/models"
	r "github.com/dancannon/gorethink"
)

type PostRepository struct {
	*r.Session
}

func (pr *PostRepository) All() (*[]models.Post, error) {
	var posts []models.Post

	rows, err := r.Table("posts").GetAll().Run(pr.Session)

	if err != nil {
		return nil, err
	}

	rows.ScanAll(&posts)

	return &posts, nil
}

func (r *PostRepository) FindByID(id int) (*models.Post, error) {
	var post models.Post

	// if err := r.DbMap.SelectOne(&post, "SELECT * FROM posts WHERE ID=?", id); err != nil {
	// 	return nil, err
	// }

	return &post, nil
}

func (r *PostRepository) Create(post *models.Post) error {
	// return r.DbMap.Insert(post)
	return nil
}
