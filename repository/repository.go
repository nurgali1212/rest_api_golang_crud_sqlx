package repository

import (
	"fmt"
	"rest_api_golang_crud_sqlx/database"
	"rest_api_golang_crud_sqlx/model"
)

type BookList interface {
	GetAllBook() ([]model.Book, error)
	CreateBook(book model.Book) (model.Book, error)
	GetByIdBook(bookId string) (model.Book, error)
	DeleteBook(bookId string) error
	UpdateBook(input model.Book) error
}

type CategoryList interface {
	GetAllCategoryRepo() ([]model.Category, error)
	CreateCategoryRepo(category model.Category) (model.Category, error)
	GetByIdCategoryRepo(categoryId string) (model.Category, error)
	DeleteCategoryRepo(categoryId string) error
	UpdateCategoryRepo(input model.Category) error
}

type CreateBookInput struct {
	Title      string `json:"title"`
	Author     string `json:"author" `
	CategoryID uint   `json:"category_id"`
}
type Repository struct {
	db *database.DB
}

func NewRepository(db *database.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// CATEGORY GET
func (r *Repository) GetAllCategoryRepo() (category []model.Category, err error) {
	return category, r.db.Conn.Select(&category, "SELECT * FROM category ORDER BY id")

}

//POST CATEGORY
func (r *Repository) CreateCategoryRepo(category model.Category) (model.Category, error) {
	fmt.Println(category)
	_, err := r.db.Conn.Exec("INSERT INTO category(genre) VALUES($1)", category.Genre)

	if err != nil {
		return category, err
	}
	return category, nil
}
//GET ID CATEGORY
func (r *Repository) GetByIdCategoryRepo(categoryId string) (category model.Category, err error) {
	return category, r.db.Conn.Get(&category, "SELECT * FROM category WHERE id = $1 ORDER BY id", categoryId)
}
//PUT CATEGORY
func (r *Repository) UpdateCategory(input model.Category) (err error) {
	_, err = r.db.Conn.NamedExec("UPDATE category SET genre=:genre WHERE id=:id", input)

	return
}

//DELETE CATEGORY
func (r *Repository) DeleteCategoryRepo(categoryId string) (err error) {
	_, err = r.db.Conn.Exec("DELETE FROM category WHERE id=$1;", categoryId)
	return
}


//GET BOOK
func (r *Repository) GetAllBook() (books []model.Book, err error) {
	return books, r.db.Conn.Select(&books, "SELECT * FROM book ORDER BY id")

}
//GET ID BOOK
func (r *Repository) GetByIdBook(bookId string) (book model.Book, err error) {
	return book, r.db.Conn.Get(&book, "SELECT * FROM book WHERE id = $1 ORDER BY id", bookId)
}

//POST BOOK
func (r *Repository) CreateBook(book model.Book) (model.Book, error) {
	fmt.Println(book)
	_, err := r.db.Conn.Exec("INSERT INTO book(title, author, category_id) VALUES($1, $2, $3)", book.Title, book.Author, book.CategoryId)

	if err != nil {
		return book, err
	}
	return book, nil
}
//DELETE BOOK
func (r *Repository) DeleteBook(bookId string) (err error) {
	_, err = r.db.Conn.Exec("DELETE FROM book WHERE id=$1;", bookId)
	return
}
//PUT BOOK

func (r *Repository) UpdateBook(input model.Book) (err error) {
	_, err = r.db.Conn.NamedExec("UPDATE book SET title=:title, author=:author WHERE id=:id", input)

	return
}
