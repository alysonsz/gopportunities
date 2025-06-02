package router

import (
	"github.com/alysonsz/gopportunities.git/controllers"
	"github.com/alysonsz/gopportunities.git/models"
	"github.com/alysonsz/gopportunities.git/repositories"
	"github.com/alysonsz/gopportunities.git/services"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initializeRoutes(router *gin.Engine) {
	db, err := gorm.Open(sqlite.Open("gopportunities.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migração automática da tabela Opportunity
	db.AutoMigrate(&models.Opportunity{})

	opportunityRepo := repositories.NewOpportunityRepository(db)
	opportunityService := services.NewOpportunityService(opportunityRepo)
	opportunityController := controllers.NewOpportunityController(opportunityService)

	v1 := router.Group("/api/v1")
	{
		v1.POST("/opportunities", opportunityController.Create)
		v1.GET("/opportunities/:id", opportunityController.Get)
		v1.GET("/opportunities", opportunityController.GetAll)
		v1.PUT("/opportunities/:id", opportunityController.Update)
		v1.DELETE("/opportunities/:id", opportunityController.Delete)
	}
}
