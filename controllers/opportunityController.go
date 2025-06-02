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
