package repository

import (
	"fmt"
	"strings"
	"test_project/model"

	"github.com/jmoiron/sqlx"
)

type TodoList interface {
	GetAll() ([]model.Todo, error)
	GetListById(id string) (model.Todo, error)
	UpdateList(listId int, input model.UpdateListInput) error
	DeleteList(id string) error
	CreateList(input model.Todo) (model.Todo, error)
}

type Repository struct {
	TodoList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		TodoList: NewTodoListPostgres(db),
	}
}

type TodoListPostgres struct {
	db *sqlx.DB
}

func NewTodoListPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{db: db}
}
func (r *TodoListPostgres) GetAll() ([]model.Todo, error) {
	var lists []model.Todo
	query := fmt.Sprintln("SELECT * FROM todos ORDER BY id")
	err := r.db.Select(&lists, query)
	return lists, err
}
func (r *TodoListPostgres) GetListById(id string) (model.Todo, error) {
	var list model.Todo
	query := fmt.Sprintln("SELECT * FROM todos where id=$1")
	err := r.db.Get(&list, query, id)
	return list, err
}
func (r *TodoListPostgres) UpdateList(listId int, input model.UpdateListInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE todos SET %s,completed_at=NOW()  WHERE id=$%d", setQuery, argId)
	args = append(args, listId)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *TodoListPostgres) DeleteList(id string) error {

	_, err := r.db.Exec("Delete from todos  where id= $1", id)
	return err
}

func (r *TodoListPostgres) CreateList(input model.Todo) (model.Todo, error) {
	var id string
	var outErr model.Todo
	query := "insert into todos (description,created_at) values ($1,Now()) returning id"

	row := r.db.QueryRow(query, input.Description)
	if err := row.Scan(&id); err != nil {
		fmt.Println("Error in row.Scan")
		return outErr, err
	}
	output, err := r.GetListById(string(id))
	if err != nil {
		fmt.Println("Error in the GetListById")
		return outErr, err
	}
	return output, err
}
