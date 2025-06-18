package sqlite_repositories_test

import (
	"testing"

	"github.com/AliakseiYafremau/GoToDo/repositories/sqlite"
)

func setup_repo() (*sqlite_repositories.SqliteTaskRepository, error) {
	database, err := sqlite_repositories.NewSqliteDB(":memory:")

	if err != nil {
		return nil, err
	}

	return sqlite_repositories.NewSqliteTaskRepository(database), nil
}

func TestCreateTask(t *testing.T) {
	repo, err := setup_repo()

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

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
