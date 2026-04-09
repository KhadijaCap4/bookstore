package models

type Book struct {
    ID     uint    `json:"id" gorm:"primaryKey"`
    Title  string  `json:"title" gorm:"index"`
    Author string  `json:"author" gorm:"index"`
    Price  float64 `json:"price" gorm:"index"`
}
