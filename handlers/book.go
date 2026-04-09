package handlers

import (
    "example/bookstore/database"
    "example/bookstore/models"
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
    books := []models.Book{}
    if err := database.DB.Find(&books).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch books"})
        return
    }
    c.JSON(http.StatusOK, books)
}

func PostBooks(c *gin.Context) {
    var book models.Book
    if err := c.ShouldBindJSON(&book); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    book.Title = strings.TrimSpace(book.Title)
    book.Author = strings.TrimSpace(book.Author)

    if book.Title == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Title is required and cannot be empty"})
        return
    }
    if book.Author == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Author is required and cannot be empty"})
        return
    }

    if err := database.DB.Create(&book).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create book"})
        return
    }
    c.JSON(http.StatusCreated, book)
}

func GetBookByID(c *gin.Context) {
    id := c.Param("id")
    var book models.Book
    if err := database.DB.First(&book, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
        return
    }
    c.JSON(http.StatusOK, book)
}
