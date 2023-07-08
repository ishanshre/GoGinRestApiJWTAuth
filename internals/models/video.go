package models

type Person struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
}
type Video struct {
	Title       string `json:"title" binding:"required,min=2,max=20"`
	Description string `json:"description" binding:"required,max=100"`
	URL         string `json:"url" binding:"required,url"`
	Author      Person `json:"author" binding:"required"`
}
