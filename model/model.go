package model

type Book struct {
	Id         uint   `json:"id" db:"id"`
	Title      string `json:"title" db:"title"`
	Author     string `json:"author" db:"author"`
	CategoryId int    `json:"category_id" db:"category_id"`
}

type Category struct {
	Id    uint   `json:"id"`
	Genre string `json:"genre"`
}

type UpdateBookInput struct {
	Title  *string `json:"title"`
	Author *string `json:"author"`
}
type UpdateCategoryinput struct {
	Genre	*string	`json:"genre"`
}
