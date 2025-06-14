package repositories

import (
	"github.com/AliakseiYafremau/GoToDo/errors"
	"github.com/AliakseiYafremau/GoToDo/models"
)

type InMemoryRepository struct {
	Tasks  map[int]models.Task
	max_id int
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		Tasks:  make(map[int]models.Task),
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

func (db *InMemoryRepository) Create(text string) (models.Task, error) {
	db.max_id++
	task := models.Task{ID: db.max_id, Text: text}

	db.Tasks[db.max_id] = task
	return task, nil
}

func (db InMemoryRepository) GetAll() ([]models.Task, error) {
	var result []models.Task

	for _, task := range db.Tasks {
		result = append(result, task)
	}

	return result, nil
}
