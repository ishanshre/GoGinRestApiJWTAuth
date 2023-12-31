package models

type Author struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}
type Video struct {
	ID          int    `json:"id"`
	Title       string `json:"title" binding:"required,min=2,max=20"`
	Description string `json:"description" binding:"required,max=100"`
	URL         string `json:"url" binding:"required,url"`
	Author      Author `json:"author" binding:"required"`
}
