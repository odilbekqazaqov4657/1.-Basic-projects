package storage

import modelsanddb "app/models_and_db"

type PostRepo struct {
}

func NewPostRepo() PostRepo {
	return PostRepo{}
}

func (p PostRepo) Create(post modelsanddb.Post) error {

	modelsanddb.Posts = append(modelsanddb.Posts, post)

	return nil
}

func (p PostRepo) GetPostsByUserId(userId string) (posts []modelsanddb.Post, err error) {

	for _, post := range modelsanddb.Posts {
		if post.UserId == userId {
			posts = append(posts, post)
		}
	}
	return
}

func (p PostRepo) GetPosts() (posts []modelsanddb.Post, err error) {

	posts = modelsanddb.Posts

	return
}
