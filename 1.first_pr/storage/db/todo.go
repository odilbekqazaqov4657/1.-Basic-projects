package db

import (
	"app/database"
	"app/models"
)

type TodoRepo struct {
}

func NewTodoRepo() TodoRepo {
	return TodoRepo{}
}

func (t TodoRepo) CreateTodo(todo models.Todo) error {

	database.Todos = append(database.Todos, todo)

	return nil
}

func (t TodoRepo) GetTodo(userId string) ([]models.Todo, error) {

	var todos []models.Todo

	for _, e := range database.Todos {

		if e.UserId == userId {

			todos = append(todos, e)
		}
	}

	return todos, nil
}
