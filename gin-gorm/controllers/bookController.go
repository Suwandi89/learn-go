package controllers

import (
	"errors"
	"fmt"
	"gin-gorm/database"
	"gin-gorm/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllBooks(ctx *gin.Context) {
	db := database.GetDB()

	allBook := []models.Book{}
	db.Find(&allBook)

	if len(allBook) == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "No books found",
			"error_message": fmt.Sprint("There are no books found in the database."),
		})
		return
	}

	ctx.JSON(http.StatusOK, allBook)
}

func GetBook(ctx *gin.Context) {
	db := database.GetDB()

	bookID := ctx.Param("bookID")

	book := models.Book{}

	err := db.First(&book, "id = ?", bookID).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error_status":  "Failed to get book",
				"error_message": fmt.Sprintf("Book with id %s doesn't exist", bookID),
			})
			return
		}
		fmt.Println("Error finding book:", err)
	}

	ctx.JSON(http.StatusOK, book)
}

func CreateBook(ctx *gin.Context) {
	db := database.GetDB()

	newBook := models.Book{}

	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err := db.Create(&newBook).Error

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Create book failed",
			"error_message": fmt.Sprint("Error cerating book data:", err),
		})
		return
	}

	ctx.JSON(http.StatusCreated, newBook)
}

func UpdateBook(ctx *gin.Context) {
	db := database.GetDB()

	bookID := ctx.Param("bookID")

	updatedBook := models.Book{}

	if err := ctx.ShouldBindJSON(&updatedBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	bID, _ := strconv.Atoi(bookID)

	err := db.Model(&updatedBook).Where("id = ?", bID).Updates(models.Book{NameBook: updatedBook.NameBook, Author: updatedBook.Author}).Error

	if err != nil {
		fmt.Println("error updating book data:", err)
		return
	}

	err = db.First(&updatedBook, "id = ?", bID).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("Book data not found")
			return
		}
		fmt.Println("Error finding book:", err)
	}

	ctx.JSON(http.StatusOK, updatedBook)
}

func DeleteBook(ctx *gin.Context) {
	db := database.GetDB()

	bookID := ctx.Param("bookID")

	book := models.Book{}

	operation := db.Where("id = ?", bookID).Delete(&book)

	if operation.Error != nil {
		fmt.Println("Error deleting book:", operation.Error)
		return
	} else if operation.RowsAffected < 1 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Delete failed",
			"error_message": fmt.Sprintf("Book with id %s doesn't exist", bookID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Book deleted succesfully",
	})
}
