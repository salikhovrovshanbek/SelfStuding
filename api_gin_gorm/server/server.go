package server

import (
	"api_gin_gorm/DB"
	"api_gin_gorm/structures"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FindBooks(c *gin.Context) {
	var books []structures.Book
	DB.DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}

func CreateBook(c *gin.Context) {
	var input structures.CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book := structures.Book{
		Title:  input.Title,
		Author: input.Author,
	}
	DB.DB.Create(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func FindBook(c *gin.Context) {
	var book structures.Book

	if err := DB.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, book)
}

func UpdateBook(c *gin.Context) {
	var book structures.Book

	if err := DB.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
	}

	var input structures.UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	DB.DB.Model(&book).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func DeleteBook(c *gin.Context) {
	var book structures.Book

	if err := DB.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found"})
	}

	DB.DB.Delete(&book)

	c.JSON(http.StatusOK, gin.H{"data": true})
}


