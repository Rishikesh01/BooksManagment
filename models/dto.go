package models

type DeleteRequest struct {
	ID int
}

type GetBook struct {
	ID int `json:"id"`
}

type GetAuthorsBooks struct {
	Author string `json:"author"`
}

type GetBookWithTitle struct {
	Title string `json:"title"`
}

type DisplayBook struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

type UpdateBook struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

type SaveBook struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}
