package models

type Book struct {
    ID     uint    `json:"id" gorm:"primaryKey"`
    Title  string  `json:"title" binding:"required"`
    Author string  `json:"author" binding:"required"`
    Price  float64 `json:"price" binding:"required,gt=0"`
}
