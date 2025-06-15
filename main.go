// Пакет запуска приложения
package main

import (
	"fmt"

	"github.com/AliakseiYafremau/GoToDo/interfaces"
	"github.com/AliakseiYafremau/GoToDo/repositories/in_memory"
	"github.com/AliakseiYafremau/GoToDo/repositories/sqlite"
)

func main() {
	var task_memory_repo interfaces.TaskRepository

	in_memory_db := in_memory_repositories.NewInMemoryDB()
	sqlite_repositories.NewSqliteDB()
	task_memory_repo = in_memory_repositories.NewInMemoryTaskRepository(in_memory_db)

	task_memory_repo.Create("Task number 1")
	task_memory_repo.Create("Task number 2")
	task_memory_repo.Create("Task number 3")
	task_memory_repo.Create("Task number 4")
	task_memory_repo.Create("Task number 5")

	fmt.Println(task_memory_repo.GetAll())

	task_memory_repo.Delete(1)

	fmt.Println(task_memory_repo.GetAll())

	task_memory_repo.Update(2, "Task number 2. CHANGED")

	fmt.Println(task_memory_repo.GetAll())
}
