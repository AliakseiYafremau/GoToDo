package repositories

import (
	"github.com/AliakseiYafremau/GoToDo/errors"
	"github.com/AliakseiYafremau/GoToDo/models"
)

type InMemoryRepository struct {
	DB *InMemoryDB
}

func NewInMemoryRepository(db *InMemoryDB) *InMemoryRepository {
	return &InMemoryRepository{
		DB: db,
	}
}

func (repo InMemoryRepository) GetById(id int) (models.Task, error) {
	task, ok := repo.DB.Tasks[id]

	if !ok {
		return models.Task{}, errors.ErrTaskDoesNotExist
	}

	return task, nil
}

func (repo *InMemoryRepository) Create(text string) (models.Task, error) {
	repo.DB.max_task_id++
	task := models.Task{ID: repo.DB.max_task_id, Text: text}

	repo.DB.Tasks[repo.DB.max_task_id] = task

	return task, nil
}

func (repo InMemoryRepository) GetAll() ([]models.Task, error) {
	var result []models.Task

	for _, task := range repo.DB.Tasks {
		result = append(result, task)
	}

	return result, nil
}
