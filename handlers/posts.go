package handlers

import (
	"effie/models"
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/codegangsta/martini-contrib/sessions"
	"github.com/coopernurse/gorp"
	"net/http"
	"strconv"
)

type PostForm struct {
	Title string `form:"title"`
	Body  string `form:"body"`
}

func PostsIndex(dbmap *gorp.DbMap, r render.Render) {
	postRepository := models.PostRepository{DbMap: dbmap}

	posts, _ := postRepository.All()

	// TODO: show flash message

	r.HTML(200, "posts_index", posts)
}

func NewPost(w http.ResponseWriter, req *http.Request, r render.Render) {
	r.HTML(200, "post_form", nil)
}

func CreatePost(w http.ResponseWriter, req *http.Request, dbmap *gorp.DbMap,
	session sessions.Session, postForm PostForm, r render.Render) {

	// TODO: validate
	postRepository := models.PostRepository{DbMap: dbmap}

	post := models.Post{Title: postForm.Title, Body: postForm.Body}
	postRepository.Create(&post)

	session.AddFlash("Your post has been saved")
	http.Redirect(w, req, "/admin/posts", 301)
}

func EditPost(w http.ResponseWriter, req *http.Request, dbmap *gorp.DbMap, params martini.Params, r render.Render) {
	postRepository := models.PostRepository{DbMap: dbmap}

	id, _ := strconv.Atoi(params["id"])

	post, _ := postRepository.FindByID(id)

	r.HTML(200, "post_form", post)
}
