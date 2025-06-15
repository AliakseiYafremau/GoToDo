package in_memory_repositories

import "github.com/AliakseiYafremau/GoToDo/models"

type InMemoryDB struct {
	Tasks       map[int]models.Task
	max_task_id int
}

func NewInMemoryDB() *InMemoryDB {
	return &InMemoryDB{
		Tasks:       make(map[int]models.Task),
		max_task_id: 0,
	}
}
