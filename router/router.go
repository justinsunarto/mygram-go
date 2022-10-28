package router

import (
	"assignment_4/controllers"
	"assignment_4/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)

		userRouter.POST("/login", controllers.UserLogin)

		userRouter.PUT("/:userId", middlewares.Authentication(), controllers.UpdateUser)

		userRouter.DELETE("/:userId", middlewares.Authentication(), controllers.DeleteUser)
	}

	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middlewares.Authentication())
		photoRouter.POST("/", middlewares.Authentication(), controllers.CreatePhoto)
		photoRouter.GET("/", middlewares.Authentication(), controllers.GetPhoto)
		photoRouter.PUT("/:photoId", controllers.UpdatePhoto)
		photoRouter.DELETE("/:photoId", controllers.DeletePhoto)
	}

	commentRouter := r.Group("/comments")
	{
		commentRouter.Use(middlewares.Authentication())
		commentRouter.POST("/", controllers.CreateComments)
		commentRouter.GET("/", controllers.GetComment)
		commentRouter.PUT("/:commentId", controllers.UpdateComment)
		commentRouter.DELETE("/:commentId", controllers.DeleteComment)
	}

	sosmedRouter := r.Group("/socialmedias")
	{
		sosmedRouter.Use(middlewares.Authentication())
		sosmedRouter.POST("/", controllers.CreateSosmed)
		sosmedRouter.GET("/", controllers.GetSosmed)
		sosmedRouter.PUT("/:commentId", controllers.UpdateSosmed)
		sosmedRouter.DELETE("/:commentId", controllers.Deletesosmed)
	}

	return r
}
