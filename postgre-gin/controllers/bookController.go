package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"postgre-gin/models"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var (
	db  *sql.DB
	err error
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "db-go-sql"
)

func ConnectDB() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err = sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}

func GetAllBooks(ctx *gin.Context) {
	db = ConnectDB()
	defer db.Close()

	var count int

	err = db.QueryRow(`SELECT COUNT(*) FROM book`).Scan(&count)
	if err != nil {
		panic(err)
	}

	if count == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprint("No books found"),
		})
		return
	}

	var BookDatas = []models.Book{}

	sqlStatement := `SELECT * FROM book`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var book = models.Book{}

		err = rows.Scan(&book.BookID, &book.Title, &book.Author, &book.Description)

		if err != nil {
			panic(err)
		}

		BookDatas = append(BookDatas, book)
	}

	ctx.JSON(http.StatusOK, BookDatas)
}

func GetBook(ctx *gin.Context) {
	db = ConnectDB()
	defer db.Close()

	bookID := ctx.Param("bookID")
	condition := false

	var bookData models.Book

	sqlStatement := `SELECT * FROM book`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var book = models.Book{}

		err = rows.Scan(&book.BookID, &book.Title, &book.Author, &book.Description)

		if err != nil {
			panic(err)
		}

		bID, _ := strconv.Atoi(bookID)

		if book.BookID == bID {
			condition = true
			bookData = book
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
	db = ConnectDB()
	defer db.Close()

	var newBook models.Book
	var count int

	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = db.QueryRow(`SELECT COUNT(*) FROM book`).Scan(&count)
	if err != nil {
		panic(err)
	}

	if count != 0 {
		err = db.QueryRow(`SELECT id FROM book ORDER BY id DESC LIMIT 1`).Scan(&count)
		if err != nil {
			panic(err)
		}
	}

	sqlStatement := `
	INSERT INTO book (id, title, author, description)
	VALUES ($1, $2, $3, $4)
	`

	db.QueryRow(sqlStatement, count+1, newBook.Title, newBook.Author, newBook.Description)

	newBook.BookID = count + 1

	ctx.JSON(http.StatusCreated, fmt.Sprintf("New book data created : %+v", newBook))
}

func UpdateBook(ctx *gin.Context) {
	db = ConnectDB()
	defer db.Close()

	bookID := ctx.Param("bookID")

	var updatedBook models.Book

	if err := ctx.ShouldBindJSON(&updatedBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	sqlStatement := `
	UPDATE book 
	SET title = $2, author = $3, description = $4
	WHERE id = $1;`

	bID, _ := strconv.Atoi(bookID)
	res, err := db.Exec(sqlStatement, bID, updatedBook.Title, updatedBook.Author, updatedBook.Description)

	if err != nil {
		panic(err)
	}

	count, err := res.RowsAffected()

	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, fmt.Sprintf("Updated data amount: %d", count))
}

func DeleteBook(ctx *gin.Context) {
	db = ConnectDB()
	defer db.Close()

	bookID := ctx.Param("bookID")

	sqlStatement := `
	DELETE FROM book 
	WHERE id = $1;`

	bID, _ := strconv.Atoi(bookID)
	res, err := db.Exec(sqlStatement, bID)

	if err != nil {
		panic(err)
	}

	count, err := res.RowsAffected()

	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, fmt.Sprintf("Deleted data amount: %d", count))

}
