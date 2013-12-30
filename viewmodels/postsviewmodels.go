package viewmodels

import "effie/models"

type PostsIndexViewModel struct {
	BaseViewModel
	Posts *[]models.Post
}
