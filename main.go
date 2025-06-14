// Пакет запуска приложения
package main

import (
	"fmt"

	"github.com/AliakseiYafremau/GoToDo/interfaces"
	"github.com/AliakseiYafremau/GoToDo/models"
	"github.com/AliakseiYafremau/GoToDo/repositories/in_memory"
)

func main() {
	task_1 := models.Task{ID: 1, Text: "task number 1"}
	task_2 := models.Task{ID: 2, Text: "task number 2"}
	task_3 := models.Task{ID: 3, Text: "task number 3"}
	task_4 := models.Task{ID: 4, Text: "task number 4"}
	task_5 := models.Task{ID: 5, Text: "task number 5"}

	var memory_repo interfaces.TaskRepository

	tasks := make(map[int]models.Task)
	tasks[task_1.ID] = task_1
	tasks[task_2.ID] = task_2
	tasks[task_3.ID] = task_3
	tasks[task_4.ID] = task_4
	tasks[task_5.ID] = task_5

	memory_repo = repositories.NewInMemoryRepository(tasks)

	fmt.Println(memory_repo.GetById(1))
	fmt.Println(memory_repo.GetById(2))
	fmt.Println(memory_repo.GetById(3))
	fmt.Println(memory_repo.GetById(4))
	fmt.Println(memory_repo.GetById(5))
}
