package route

import (
	"goback/src/controller"
	"goback/src/repository"
	"goback/src/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRouter(db *gorm.DB, r *gin.Engine) {
	routerGroup := r.Group("users")

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	routerGroup.GET("/:id", userController.GetUserByID)
	routerGroup.GET("/", userController.GetAll)
	routerGroup.POST("/create", userController.AddUser)
}
