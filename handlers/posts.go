package handlers

import (
	"effie/models"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/coopernurse/gorp"
	"net/http"
)

func PostsIndex(dbmap *gorp.DbMap, r render.Render) {
	postRepository := models.PostRepository{DbMap: dbmap}

	posts, _ := postRepository.All()

	r.HTML(200, "posts_index", posts)
}

func NewPost(w http.ResponseWriter, req *http.Request, r render.Render) {
	r.HTML(200, "post_form", nil)
}

func CreatePost(w http.ResponseWriter, req *http.Request, r render.Render) {
	// TODO: validate and save

	http.Redirect(w, req, "/admin", 301)
}
