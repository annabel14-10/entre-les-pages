package delivery

import (
	"book-service/domain"
	"book-service/middleware"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	BookUsecase domain.BookUsecase
}

func NewBookHandler(r *gin.Engine, usecase domain.BookUsecase) {
	handler := &BookHandler{
		BookUsecase: usecase,
	}

	api := r.Group("/api")
	{
		api.GET("/books", handler.FetchBooks)
		api.GET("/books/:id", handler.GetBook)

		protected := api.Group("/")
		protected.Use(middleware.RequireAuth())
		{
			protected.POST("/books", handler.AddBook)
			protected.PUT("/books/:id", handler.UpdateBook)
			protected.DELETE("/books/:id", handler.DeleteBook)

			protected.POST("/books/:id/chapters", handler.AddChapter)
			protected.PUT("/chapters/:id", handler.UpdateChapter)
			protected.DELETE("/chapters/:id", handler.DeleteChapter)

			api.Static("/uploads", "./uploads")
		}
	}
}

func (h *BookHandler) FetchBooks(c *gin.Context) {
	level := c.Query("level")
	books, err := h.BookUsecase.FetchAllBooks(level)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, books)
}

func (h *BookHandler) AddBook(c *gin.Context) {
	title := c.PostForm("title")
	author := c.PostForm("author")
	level := c.PostForm("level")

	if title == "" || author == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title and Author are required"})
		return
	}

	var imageUrl string

	file, err := c.FormFile("image")
	if err == nil {
		os.MkdirAll("uploads", os.ModePerm)

		filename := fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename)
		savePath := filepath.Join("uploads", filename)

		if err := c.SaveUploadedFile(file, savePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image"})
			return
		}
		imageUrl = "/api/uploads/" + filename
	}

	book := domain.Book{
		Title:    title,
		Author:   author,
		Level:    level,
		ImageURL: imageUrl,
	}

	err = h.BookUsecase.AddBook(&book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create book"})
		return
	}

	c.JSON(http.StatusCreated, book)
}

func (h *BookHandler) GetBook(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID format"})
		return
	}

	book, err := h.BookUsecase.FetchBookByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, book)
}

func (h *BookHandler) UpdateBook(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	title := c.PostForm("title")
	author := c.PostForm("author")
	level := c.PostForm("level")

	existingBook, err := h.BookUsecase.FetchBookByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	imageUrl := existingBook.ImageURL

	file, err := c.FormFile("image")
	if err == nil {
		os.MkdirAll("uploads", os.ModePerm)
		filename := fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename)
		savePath := filepath.Join("uploads", filename)

		if err := c.SaveUploadedFile(file, savePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save new image"})
			return
		}
		imageUrl = "/api/uploads/" + filename
	}

	existingBook.Title = title
	existingBook.Author = author
	existingBook.Level = level
	existingBook.ImageURL = imageUrl

	err = h.BookUsecase.ModifyBook(existingBook)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update book"})
		return
	}

	c.JSON(http.StatusOK, existingBook)
}

func (h *BookHandler) DeleteBook(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID format"})
		return
	}

	if err := h.BookUsecase.RemoveBook(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete book"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully!"})
}

func (h *BookHandler) AddChapter(c *gin.Context) {
	bookIDParam := c.Param("id")
	bookID, err := strconv.Atoi(bookIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	var chapter domain.Chapter
	if err := c.ShouldBindJSON(&chapter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	chapter.BookID = uint(bookID)

	if err := h.BookUsecase.AddChapter(&chapter); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Chapter added!", "chapter": chapter})
}

func (h *BookHandler) UpdateChapter(c *gin.Context) {
	chapterIDParam := c.Param("id")
	chapterID, err := strconv.Atoi(chapterIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid chapter ID"})
		return
	}

	var chapter domain.Chapter
	if err := c.ShouldBindJSON(&chapter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	chapter.ID = uint(chapterID)

	if err := h.BookUsecase.ModifyChapter(&chapter); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Chapter updated!", "chapter": chapter})
}

func (h *BookHandler) DeleteChapter(c *gin.Context) {
	chapterIDParam := c.Param("id")
	chapterID, err := strconv.Atoi(chapterIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid chapter ID"})
		return
	}

	if err := h.BookUsecase.RemoveChapter(uint(chapterID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete chapter"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Chapter deleted!"})
}