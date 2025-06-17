package in_memory_repositories_test

import (
	"testing"

	"github.com/AliakseiYafremau/GoToDo/repositories/in_memory"
)

func setup_repo() *in_memory_repositories.InMemoryTaskRepository {
	return in_memory_repositories.NewInMemoryTaskRepository(in_memory_repositories.NewInMemoryDB())
}

func TestCreateTask(t *testing.T) {
	repo := setup_repo()

	task_text := "Test task"
	task, err := repo.Create(task_text)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if task.ID != 1 {
		t.Errorf("unexpected id, expected 1, got %d", task.ID)
	}

	if task.Text != task_text {
		t.Errorf("unexpected text, expected %s, got %s", task_text, task.Text)
	}
}
