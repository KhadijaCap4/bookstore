package handlers

import (
    "encoding/json"
    "example/bookstore/database"
    "example/bookstore/models"
    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
    "net/http"
    "net/http/httptest"
    "strconv"
    "strings"
    "testing"
)

func setupRouter() *gin.Engine {
    r := gin.Default()
    r.GET("/books", GetBooks)
    r.POST("/books", PostBooks)
    r.GET("/books/:id", GetBookByID)
    return r
}

func TestGetBooks(t *testing.T) {
	database.Connect()
	r := gin.Default()
	r.GET("/books", GetBooks)

	req, _ := http.NewRequest("GET", "/books", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestPostAndGetBookByID(t *testing.T) {
    database.Connect()
    r := setupRouter()

    // Ajouter un livre
    book := `{"title":"Les voisins", "author":"Danielle Steel", "price":19.99}`
    req, _ := http.NewRequest("POST", "/books", strings.NewReader(book))
    req.Header.Set("Content-Type", "application/json")
    resp := httptest.NewRecorder()
    r.ServeHTTP(resp, req)

    assert.Equal(t, http.StatusOK, resp.Code)

    // Récupérer l'ID du livre inséré
    var createdBook models.Book
    err := json.Unmarshal(resp.Body.Bytes(), &createdBook)
    if err != nil {
        t.Fatalf("Erreur lors de l'unmarshal de la réponse : %v", err)
    }
    bookID := createdBook.ID

    // Vérifier que l'ID est bien récupéré
    if bookID == 0 {
        t.Fatalf("L'ID du livre n'a pas été récupéré")
    }

    // Récupérer le livre par son ID
    req, _ = http.NewRequest("GET", "/books/"+strconv.Itoa(int(bookID)), nil)
    resp = httptest.NewRecorder()
    r.ServeHTTP(resp, req)

    assert.Equal(t, http.StatusOK, resp.Code)

    // Vérifier les détails du livre récupéré
    var retrievedBook models.Book
    err = json.Unmarshal(resp.Body.Bytes(), &retrievedBook)
    if err != nil {
        t.Fatalf("Erreur lors de l'unmarshal de la réponse : %v", err)
    }

    assert.Equal(t, createdBook.Title, retrievedBook.Title)
    assert.Equal(t, createdBook.Author, retrievedBook.Author)
    assert.Equal(t, createdBook.Price, retrievedBook.Price)
}