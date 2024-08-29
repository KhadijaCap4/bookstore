package handlers

import (
    "example/bookstore/models"
    "example/bookstore/database"
    "github.com/gin-gonic/gin"
    "net/http"
)

func GetBooks(c *gin.Context) {
    var books []models.Book
    if err := database.DB.Find(&books).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
    if err := database.DB.Create(&book).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, book)
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
