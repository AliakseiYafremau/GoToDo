// Пакет запуска приложения
package main

import (
	"fmt"

	"github.com/AliakseiYafremau/GoToDo/interfaces"
	"github.com/AliakseiYafremau/GoToDo/repositories/in_memory"
)

func main() {
	var task_memory_repo interfaces.TaskRepository

	in_memory_db := repositories.NewInMemoryDB()
	task_memory_repo = repositories.NewInMemoryRepository(in_memory_db)

	task_memory_repo.Create("Task number 1")
	task_memory_repo.Create("Task number 2")
	task_memory_repo.Create("Task number 3")
	task_memory_repo.Create("Task number 4")
	task_memory_repo.Create("Task number 5")

	fmt.Println(task_memory_repo.GetAll())

	fmt.Println(task_memory_repo.GetById(1))
	fmt.Println(task_memory_repo.GetById(2))
	fmt.Println(task_memory_repo.GetById(3))
	fmt.Println(task_memory_repo.GetById(4))
	fmt.Println(task_memory_repo.GetById(5))
}
