package handlers

import (
	"effie/models"
	"effie/repositories"
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/codegangsta/martini-contrib/sessions"
	r "github.com/dancannon/gorethink"
	"net/http"
	"strconv"
)

type PostForm struct {
	Title string `form:"title"`
	Body  string `form:"body"`
}

func PostsIndex(rdbSession *r.Session, r render.Render) {
	postRepository := repositories.PostRepository{Session: rdbSession}

	posts, _ := postRepository.All()

	// TODO: show flash message

	r.HTML(200, "posts_index", posts)
}

func NewPost(w http.ResponseWriter, req *http.Request, r render.Render) {
	r.HTML(200, "post_form", nil)
}

func CreatePost(w http.ResponseWriter, req *http.Request, rdbSession *r.Session,
	session sessions.Session, postForm PostForm, r render.Render) {

	// TODO: validate
	postRepository := repositories.PostRepository{Session: rdbSession}

	post := models.Post{Title: postForm.Title, Body: postForm.Body}
	postRepository.Create(&post)

	session.AddFlash("Your post has been saved")
	http.Redirect(w, req, "/admin/posts", 301)
}

func EditPost(w http.ResponseWriter, req *http.Request, rdbSession *r.Session, params martini.Params, r render.Render) {
	postRepository := repositories.PostRepository{Session: rdbSession}

	id, _ := strconv.Atoi(params["id"])

	post, _ := postRepository.FindByID(id)

	r.HTML(200, "post_form", post)
}
