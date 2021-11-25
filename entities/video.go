package entities

type Video struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Title       string `json:"title" binding:"min=2,max=10"`
	Description string `json:"description" binding:"max=20"`
	URL         string `json:"url" binding:"required,url"`
	// Author      Person `json:"author" binding:"required"`
}
