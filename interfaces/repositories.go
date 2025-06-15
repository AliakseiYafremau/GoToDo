package interfaces

import "github.com/AliakseiYafremau/GoToDo/models"

type TaskRepository interface {
	GetById(task_id int) (models.Task, error)
	Create(text string) (models.Task, error)
	GetAll() ([]models.Task, error)
	Update(task_id int, text string) (models.Task, error)
	Delete(task_id int) error
}
