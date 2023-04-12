package routers

import (
	"mygram/controllers"
	"mygram/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)

		userRouter.POST("/login", controllers.UserLogin)
	}

	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middlewares.Authentication())
		photoRouter.POST("/create", controllers.CreatePhoto)

		photoRouter.GET("/getall", controllers.GetAllPhotos)

		photoRouter.GET("/getone/:photoId", controllers.GetPhoto)

		photoRouter.PUT("/update/:photoId", middlewares.PhotoAuthorization(), controllers.UpdatePhoto)

		photoRouter.DELETE("/delete/:photoId", middlewares.PhotoAuthorization(), controllers.DeletePhoto)
	}

	commentRouter := r.Group("/comments")
	{
		commentRouter.Use(middlewares.Authentication())
		commentRouter.POST("/create/:photoId", controllers.CreateComment)

		commentRouter.GET("/getall", controllers.GetAllComments)

		commentRouter.GET("/getall/:photoId", controllers.GetAllCommentsForPhoto)

		commentRouter.GET("/get/:commentId", controllers.GetComment)

		commentRouter.PUT("/update/:commentId", middlewares.CommentAuthorization(), controllers.UpdateComment)

		commentRouter.DELETE("/delete/:commentId", middlewares.CommentAuthorization(), controllers.DeleteComment)

	}

	return r
}
