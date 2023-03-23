package routers

import (
	"belajar-gin/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.GET("/books", controllers.GetAllBooks)

	router.GET("/books/:bookID", controllers.GetBook)

	router.POST("/books", controllers.CreateBook)

	router.PUT("/books/:bookID", controllers.UpdateBook)

	router.DELETE("/books/:bookID", controllers.DeleteBook)

	return router
}
