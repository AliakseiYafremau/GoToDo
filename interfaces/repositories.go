package interfaces

import "github.com/AliakseiYafremau/GoToDo/models"

type TaskRepository interface {
	GetById(id int) (models.Task, error)
	// Create(text string) (models.Task, error)
	// Update(id int, text string) (models.Task, error)
	// Delete(id int) (error)
}
