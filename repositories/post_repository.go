package repositories

import (
	"effie/models"
	r "github.com/dancannon/gorethink"
)

type PostRepository struct {
	*r.Session
}

func (pr *PostRepository) All() (*[]models.Post, error) {
	rows, err := r.Table("posts").Run(pr.Session)

	if err != nil {
		return nil, err
	}

	posts := []models.Post{}
	if err = rows.ScanAll(&posts); err != nil {
		return nil, err
	}

	return &posts, nil
}

func (pr *PostRepository) FindByID(id string) (*models.Post, error) {
	row, err := r.Table("posts").Get(id).RunRow(pr.Session)

	if err != nil {
		return nil, err
	}

	if row.IsNil() {
		return nil, nil
	}

	var post models.Post
	if err = row.Scan(&post); err != nil {
		return nil, err
	}

	return &post, nil
}

func (pr *PostRepository) Create(post *models.Post) error {
	_, err := r.Table("posts").Insert(post).RunWrite(pr.Session)
	return err
}
