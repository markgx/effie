package repositories

import (
	"effie/models"
	"github.com/coopernurse/gorp"
)

type UserRepository struct {
	*gorp.DbMap
}

func (r *UserRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User

	if err := r.DbMap.SelectOne(&user, "SELECT * FROM users WHERE username = ?", username); err != nil {
		return nil, err
	}

	if user.ID == 0 {
		return nil, nil
	}

	return &user, nil
}
