package repositories

import (
	"github.com/AliakseiYafremau/GoToDo/errors"
	"github.com/AliakseiYafremau/GoToDo/models"
)

type InMemoryRepository struct {
	Tasks  map[int]models.Task
	max_id int
}

func NewInMemoryRepository(tasks map[int]models.Task) InMemoryRepository {
	return InMemoryRepository{
		Tasks:  tasks,
		max_id: 0,
	}
}

func (db InMemoryRepository) GetById(id int) (models.Task, error) {
	task, ok := db.Tasks[id]

	if !ok {
		return models.Task{}, errors.ErrTaskDoesNotExist
	}

	return task, nil
}
