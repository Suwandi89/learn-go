package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Book struct {
	BookID      int    `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"desc"`
}

var BookDatas = []Book{}

func GetAllBooks(ctx *gin.Context) {
	if len(BookDatas) == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprint("No books found"),
		})
		return
	}

	ctx.JSON(http.StatusOK, BookDatas)
}

func GetBook(ctx *gin.Context) {
	bookID := ctx.Param("bookID")
	condition := false

	var bookData Book

	for i, book := range BookDatas {
		bID := strconv.Itoa(book.BookID)
		if bookID == bID {
			condition = true
			bookData = BookDatas[i]
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("book with id %v not found", bookID),
		})
		return
	}

	ctx.JSON(http.StatusOK, bookData)
}

func CreateBook(ctx *gin.Context) {
	var newBook Book

	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newBook.BookID = len(BookDatas) + 1
	BookDatas = append(BookDatas, newBook)

	ctx.JSON(http.StatusCreated, fmt.Sprint("Created"))
}

func UpdateBook(ctx *gin.Context) {
	bookID := ctx.Param("bookID")
	condition := false

	var updatedBook Book

	if err := ctx.ShouldBindJSON(&updatedBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for i, book := range BookDatas {
		bID := strconv.Itoa(book.BookID)
		if bookID == bID {
			condition = true
			BookDatas[i] = updatedBook
			BookDatas[i].BookID, _ = strconv.Atoi(bookID)
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("book with id %v not found", bookID),
		})
		return
	}

	ctx.JSON(http.StatusOK, fmt.Sprint("Updated"))
}

func DeleteBook(ctx *gin.Context) {
	bookID := ctx.Param("bookID")
	condition := false

	var bookIndex int

	for i, book := range BookDatas {
		bID := strconv.Itoa(book.BookID)
		if bookID == bID {
			condition = true
			bookIndex = i
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("book with id %v not found", bookID),
		})
		return
	}

	copy(BookDatas[bookIndex:], BookDatas[bookIndex+1:])
	BookDatas[len(BookDatas)-1] = Book{}
	BookDatas = BookDatas[:len(BookDatas)-1]

	ctx.JSON(http.StatusOK, fmt.Sprint("Deleted"))
}
