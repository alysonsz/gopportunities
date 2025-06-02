package router

import (
	"github.com/alysonsz/gopportunities.git/authentication"
	"github.com/alysonsz/gopportunities.git/controllers"
	"github.com/alysonsz/gopportunities.git/models"
	"github.com/alysonsz/gopportunities.git/repositories"
	"github.com/alysonsz/gopportunities.git/services"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initializeRoutes(router *gin.Engine) {
	db, err := gorm.Open(sqlite.Open("./gopportunities.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database" + err.Error())
	}

	db.AutoMigrate(&models.Opportunity{}, &models.User{})

	opportunityRepo := repositories.NewOpportunityRepository(db)
	userRepo := repositories.NewUserRepository(db)
	opportunityService := services.NewOpportunityService(opportunityRepo)
	authService := services.NewAuthService(userRepo)
	opportunityController := controllers.NewOpportunityController(opportunityService)
	authController := controllers.NewAuthController(authService)

	v1 := router.Group("/api/v1")
	{
		v1.POST("/register", authController.Register)
		v1.POST("/login", authController.Login)

		authRoutes := v1.Group("/")
		authRoutes.Use(authentication.AuthJwt())
		{
			authRoutes.POST("/opportunities", opportunityController.Create)
			authRoutes.GET("/opportunities/:id", opportunityController.Get)
			authRoutes.GET("/opportunities", opportunityController.GetAll)
			authRoutes.PUT("/opportunities/:id", opportunityController.Update)
			authRoutes.DELETE("/opportunities/:id", opportunityController.Delete)
		}
	}
}
