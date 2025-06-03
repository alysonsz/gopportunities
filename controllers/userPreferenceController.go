package controllers

import (
	"net/http"
	"strconv"

	"github.com/alysonsz/gopportunities.git/models"
	"github.com/alysonsz/gopportunities.git/services"
	"github.com/gin-gonic/gin"
)

type UserPreferenceController interface {
	SetPreference(ctx *gin.Context)
	GetPreference(ctx *gin.Context)
}

type userPreferenceController struct {
	service services.UserPreferenceService
}

func NewUserPreferenceController(service services.UserPreferenceService) UserPreferenceController {
	return &userPreferenceController{service}
}

// @Summary Set user notification preferences
// @Description Allows a user to set preferences like location, job type, and keywords for notifications.
// @Tags preferences
// @Accept json
// @Produce json
// @Param preference body models.UserPreference true "User Preferences"
// @Success 200 {object} models.UserPreference
// @Failure 400 {object} map[string]string
// @Router /api/v1/preferences [post]
func (ctrl *userPreferenceController) SetPreference(ctx *gin.Context) {
	userID, _ := strconv.Atoi(ctx.GetString("userID"))
	var pref models.UserPreference
	if err := ctx.ShouldBindJSON(&pref); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if err := ctrl.service.SetPreference(uint(userID), &pref); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save preference"})
		return
	}

	ctx.JSON(http.StatusOK, pref)
}

// @Summary Get user notification preferences
// @Description Retrieve a user's current notification preferences.
// @Tags preferences
// @Produce json
// @Success 200 {object} models.UserPreference
// @Failure 404 {object} map[string]string
// @Router /api/v1/preferences [get]
func (ctrl *userPreferenceController) GetPreference(ctx *gin.Context) {
	userID, _ := strconv.Atoi(ctx.GetString("userID"))
	pref, err := ctrl.service.GetPreference(uint(userID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Preference not found"})
		return
	}

	ctx.JSON(http.StatusOK, pref)
}
