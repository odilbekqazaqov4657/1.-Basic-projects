package api

import (
	modelsanddb "app/models_and_db"
	"app/storage"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
)

type Token struct {
	UserId string
}

var UserToken *Token

func Api(postRepo storage.PostRepo) {

	var endPoint int

	h := NewHandler(postRepo)

	for endPoint != -1 {

		fmt.Println(`
			0. Log-out
			1. Log-in
			2. Create post
			3. Get posts
			4. Get posts by user id 
			-1. Exit
	`)

		fmt.Println("Endpointlardan birini tanlang: ")

		fmt.Scanln(&endPoint)

		switch endPoint {
		case 0:
			if UserToken != nil {
				h.LogOut()
			} else {
				fmt.Println("You are not logged in")
			}
		case 1:
			h.LogIn()
		case 2:
			if UserToken != nil {
				h.CreatePost()
			} else {
				fmt.Println("You are not logged in")
			}
		case 3:
			if UserToken != nil {
				h.GetPosts()
			} else {
				fmt.Println("You are not logged in")
			}
		case 4:
			if UserToken != nil {
				h.GetPostsByUserId()
			} else {
				fmt.Println("You are not logged in")
			}
		}
	}
}

type Handler struct {
	postRepo storage.PostRepo
}

func NewHandler(postRepo storage.PostRepo) Handler {
	return Handler{postRepo: postRepo}
}

func (h Handler) LogIn() {

	var login modelsanddb.Login

	fmt.Print("Enter username: ")
	fmt.Scanln(&login.Username)

	fmt.Print("Enter password: ")
	fmt.Scanln(&login.Password)

	userRepo := storage.NewUserRepo()

	user, err := userRepo.Login(login)

	if err != nil {
		log.Println("error on login", err)
		return
	}

	UserToken = &Token{}
	UserToken.UserId = user.UserId

	fmt.Println("You are logged in !!!")
}

func (h Handler) LogOut() {
	UserToken = nil
	fmt.Println("You are logged out !")
}

func (h Handler) CreatePost() {

	var post modelsanddb.Post

	fmt.Println("Enter post title ")
	fmt.Scanln(&post.Title)

	uid := uuid.New().String()
	post.PostId = uid

	post.CreatedAt = time.Now()

	post.UserId = UserToken.UserId

	h.postRepo.Create(post)
}

func (h Handler) GetPosts() {

	posts, err := h.postRepo.GetPosts()

	if err != nil {
		return
	}
	fmt.Println(posts)
}

func (h Handler) GetPostsByUserId() {

	posts, err := h.postRepo.GetPostsByUserId(UserToken.UserId)

	if err != nil {
		return
	}
	fmt.Println(posts)
}
