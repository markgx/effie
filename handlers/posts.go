package handlers

import (
	"effie/models"
	"effie/repositories"
	"effie/viewmodels"
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/codegangsta/martini-contrib/sessions"
	r "github.com/dancannon/gorethink"
	"net/http"
)

type PostForm struct {
	Title string `form:"title"`
	Body  string `form:"body"`
}

func PostsIndex(rdbSession *r.Session, session sessions.Session, r render.Render) {
	postRepository := repositories.PostRepository{Session: rdbSession}

	posts, err := postRepository.All()
	if err != nil {
		panic(err)
	}

	// TODO: show flash message

	viewModel := viewmodels.PostsIndexViewModel{
		BaseViewModel: *viewmodels.NewBaseViewModel(session),
		Posts:         posts,
	}

	r.HTML(200, "posts/index", viewModel)
}

func NewPost(w http.ResponseWriter, req *http.Request, r render.Render) {
	r.HTML(200, "posts/form", nil)
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

	id, _ := params["id"]

	post, _ := postRepository.FindByID(id)

	r.HTML(200, "posts/form", post)
}
