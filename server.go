package main

import (
	"golang/config"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.SetupDatabaseConnection()
	// authController controller.AuthController = controller.NewAuthController(authService, jwtService)
	// userRepository repository.UserRepository = repository.NewUserRepository(db)
	// jwtService     service.JWTService        = service.NewJWTService()
	// userService    service.UserService       = service.NewUserService(userRepository)
	// authService    service.AuthService       = service.NewAuthService(userRepository)
	// userController controller.UserController = controller.NewUserController(userService, jwtService)
	// bookController controller.BookController = controller.NewBookController(bookService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	// authRoutes := r.Group("api/auth")
	// {
	// 	authRoutes.POST("/login", authController.Login)
	// 	authRoutes.POST("/register", authController.Register)
	// }

	// userRoutes := r.Group("api/user", middleware.AuthorizeJWT(jwtService))
	// {
	// 	userRoutes.GET("/profile", userController.Profile)
	// 	userRoutes.PUT("/profile", userController.Update)
	// }

	r.Run()
}
