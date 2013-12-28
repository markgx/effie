package repositories

import (
	"effie/models"
	r "github.com/dancannon/gorethink"
)

type UserRepository struct {
	*r.Session
}

func (ur *UserRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User

	rows, err := r.Table("users").Filter(r.Row.Field("username").Eq(username)).Run(ur.Session)

	if err != nil {
		return nil, err
	}

	if rows.Next() {
		rows.Scan(&user)
	}

	return &user, nil
}
