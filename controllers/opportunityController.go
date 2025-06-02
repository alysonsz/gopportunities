package controllers

import (
	"net/http"
	"strconv"

	"github.com/alysonsz/gopportunities.git/models"
	"github.com/alysonsz/gopportunities.git/services"
	"github.com/gin-gonic/gin"
)

type OpportunityController interface {
	Create(ctx *gin.Context)
	Get(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type opportunityController struct {
	service services.OpportunityService
}

func NewOpportunityController(service services.OpportunityService) OpportunityController {
	return &opportunityController{service: service}
}

// @Summary Create a new opportunity
// @Description Create a new opportunity with provided data
// @Tags opportunities
// @Accept json
// @Produce json
// @Param opportunity body models.Opportunity true "Opportunity data"
// @Success 201 {object} models.Opportunity
// @Failure 400 {object} map[string]string
// @Security ApiKeyAuth
// @Router /api/v1/opportunities [post]
func (ctrl *opportunityController) Create(ctx *gin.Context) {
	var opportunity models.Opportunity
	if err := ctx.ShouldBindJSON(&opportunity); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.service.CreateOpportunity(&opportunity); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, opportunity)
}

// @Summary Get an opportunity by ID
// @Description Retrieve opportunity details by its ID
// @Tags opportunities
// @Produce json
// @Param id path int true "Opportunity ID"
// @Success 200 {object} models.Opportunity
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Security ApiKeyAuth
// @Router /api/v1/opportunities/{id} [get]
func (ctrl *opportunityController) Get(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	opportunity, err := ctrl.service.GetOpportunity(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Opportunity not found"})
		return
	}

	ctx.JSON(http.StatusOK, opportunity)
}

// @Summary List all opportunities
// @Description Get a list of opportunities with optional filters
// @Tags opportunities
// @Produce json
// @Param status query string false "Filter by status"
// @Param location query string false "Filter by location"
// @Success 200 {array} models.Opportunity
// @Failure 500 {object} map[string]string
// @Security ApiKeyAuth
// @Router /api/v1/opportunities [get]
func (ctrl *opportunityController) GetAll(ctx *gin.Context) {
	status := ctx.Query("status")
	location := ctx.Query("location")

	opportunities, err := ctrl.service.GetAllOpportunities(status, location)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, opportunities)
}

// @Summary Update an opportunity by ID
// @Description Update an existing opportunity
// @Tags opportunities
// @Accept json
// @Produce json
// @Param id path int true "Opportunity ID"
// @Param opportunity body models.Opportunity true "Updated data"
// @Success 200 {object} models.Opportunity
// @Failure 400 {object} map[string]string
// @Security ApiKeyAuth
// @Router /api/v1/opportunities/{id} [put]
func (ctrl *opportunityController) Update(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var opportunity models.Opportunity
	if err := ctx.ShouldBindJSON(&opportunity); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.service.UpdateOpportunity(uint(id), &opportunity); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, opportunity)
}

// @Summary Delete an opportunity by ID
// @Description Delete an opportunity using its ID
// @Tags opportunities
// @Produce json
// @Param id path int true "Opportunity ID"
// @Success 204 {object} nil
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security ApiKeyAuth
// @Router /api/v1/opportunities/{id} [delete]
func (ctrl *opportunityController) Delete(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := ctrl.service.DeleteOpportunity(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
