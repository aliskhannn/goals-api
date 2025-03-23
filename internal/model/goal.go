package model

type Goal struct {
	ID          int    `json:"id"`
	Title       string `json:"title" validate:"required,max=255"`
	Description string `json:"description" validate:"required"`
	Completed   bool   `json:"completed"`
}
