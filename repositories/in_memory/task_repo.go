package repositories

import (
	"github.com/AliakseiYafremau/GoToDo/errors"
	"github.com/AliakseiYafremau/GoToDo/models"
)

type InMemoryRepository struct {
	Tasks map[int]models.Task
}

func (db InMemoryRepository) GetById(id int) (models.Task, error) {
	task, ok := db.Tasks[id]

	if !ok {
		return models.Task{}, errors.ErrTaskDoesNotExist
	}

	return task, nil
}
