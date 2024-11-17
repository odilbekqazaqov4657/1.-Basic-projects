package main

import (
	"app/api"
	"app/storage"
)

func main() {

	postRepo := storage.NewPostRepo()

	api.Api(postRepo)
}
