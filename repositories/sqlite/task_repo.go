package sqlite_repositories

import (
	"database/sql"

	"github.com/AliakseiYafremau/GoToDo/errors"
	"github.com/AliakseiYafremau/GoToDo/models"
)

type SqliteTaskRepository struct {
	DB *sql.DB
}

func NewSqliteTaskRepository(db *sql.DB) *SqliteTaskRepository {
	return &SqliteTaskRepository{
		DB: db,
	}
}

func (repo *SqliteTaskRepository) GetById(task_id int) (models.Task, error) {
	rows, err := repo.DB.Query("SELECT * FROM tasks WHERE id = ?", task_id)

	if err != nil {
		return models.Task{}, errors.ErrExecuteFailed
	}

	defer rows.Close()

	var received_id int

	var received_text string
	err = rows.Scan(&received_id, &received_text)

	if err != nil {
		return models.Task{}, errors.ErrExecuteFailed
	}

	err = rows.Err()

	if err != nil {
		return models.Task{}, errors.ErrExecuteFailed
	}

	return models.Task{
		ID:   received_id,
		Text: received_text,
	}, nil
}

func (repo *SqliteTaskRepository) Create(text string) (models.Task, error) {
	result, err := repo.DB.Exec("INSERT INTO tasks(text) VALUES (?)", text)

	if err != nil {
		return models.Task{}, errors.ErrExecuteFailed
	}

	task_id, err := result.LastInsertId()

	if err != nil {
		return models.Task{}, errors.ErrExecuteFailed
	}

	return models.Task{
		ID:   int(task_id),
		Text: text,
	}, nil
}

func (repo *SqliteTaskRepository) GetAll() ([]models.Task, error) {
	rows, err := repo.DB.Query("SELECT * FROM tasks")

	if err != nil {
		return make([]models.Task, 0), err
	}

	defer rows.Close()

	var result []models.Task

	for rows.Next() {
		var task_id int

		var text string

		err := rows.Scan(&task_id, &text)

		if err != nil {
			return make([]models.Task, 0), err
		}

		result = append(result, models.Task{ID: task_id, Text: text})
	}

	err = rows.Err()

	if err != nil {
		return result, nil
	}

	return result, nil
}

func (repo *SqliteTaskRepository) Update(task_id int, text string) (models.Task, error) {
	_, err := repo.DB.Exec("UPDATE tasks SET text = ? WHERE id = ?", text, task_id)

	if err != nil {
		return models.Task{}, errors.ErrExecuteFailed
	}

	return repo.GetById(task_id)
}

func (repo *SqliteTaskRepository) Delete(task_id int) error {
	_, err := repo.DB.Exec("DELETE FROM tasks WHERE id = ?", task_id)

	if err != nil {
		return errors.ErrExecuteFailed
	}

	return nil
}
