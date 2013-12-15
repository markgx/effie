package models

import (
	"github.com/coopernurse/gorp"
)

type User struct {
	ID       int
	Username string
	Password string
}

type UserRepository struct {
	*gorp.DbMap
}

func (r *UserRepository) FindByUsername(username string) (*User, error) {
	var user User

	if err := r.DbMap.SelectOne(&user, "SELECT * FROM users WHERE username = ?", username); err != nil {
		return nil, err
	}

	if user.ID == 0 {
		return nil, nil
	}

	return &user, nil
}
