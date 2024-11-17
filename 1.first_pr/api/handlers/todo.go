package handlers

import (
	"app/models"
	"app/storage/db"
	"fmt"
	"time"
)

var todoRepo = db.NewTodoRepo()

func (h Handler) CreateTodo() {

	if UserToken != nil {
		if UserToken.Subscribes.Todo {

			var newTodo models.Todo

			fmt.Println("Task: ")
			fmt.Scanln(&newTodo.Task)

			newTodo.CreatedAt = time.Now()

			newTodo.UserId = UserToken.UserId

			todoRepo.CreateTodo(newTodo)
		}
	}
}

func (h Handler) GetTodos() {
	if UserToken != nil {
		if UserToken.Subscribes.Todo {

			todos, err := todoRepo.GetTodo(UserToken.UserId)

			if err != nil {
				fmt.Println("Error on getting data from database: ", err)
				return
			}
			for _, e := range todos {
				fmt.Println(e)
			}
		}
	}
}
