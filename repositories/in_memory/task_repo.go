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

func (repo *InMemoryRepository) GetById(task_id int) (models.Task, error) {
	task, ok := repo.DB.Tasks[task_id]

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

func (repo *InMemoryRepository) GetAll() ([]models.Task, error) {
	result := make([]models.Task, 0, len(repo.DB.Tasks))

	for _, task := range repo.DB.Tasks {
		result = append(result, task)
	}

	return result, nil
}

func (repo *InMemoryRepository) Update(task_id int, text string) (models.Task, error) {
	_, ok := repo.DB.Tasks[task_id]

	if !ok {
		return models.Task{}, errors.ErrTaskDoesNotExist
	}

	repo.DB.Tasks[task_id] = models.Task{ID: task_id, Text: text}

	return repo.DB.Tasks[task_id], nil
}

func (repo *InMemoryRepository) Delete(task_id int) error {
	_, ok := repo.DB.Tasks[task_id]

	if !ok {
		return errors.ErrTaskDoesNotExist
	}

	delete(repo.DB.Tasks, task_id)

	return nil
}
