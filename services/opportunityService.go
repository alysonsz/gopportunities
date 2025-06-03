package services

import (
	"errors"
	"fmt"
	"strings"

	"github.com/alysonsz/gopportunities.git/models"
	"github.com/alysonsz/gopportunities.git/repositories"
)

type OpportunityService interface {
	CreateOpportunity(opportunity *models.Opportunity) error
	GetOpportunity(id uint) (*models.Opportunity, error)
	GetAllOpportunities(status, location string) ([]models.Opportunity, error)
	UpdateOpportunity(id uint, opportunity *models.Opportunity) error
	DeleteOpportunity(id uint) error
}

type opportunityService struct {
	repo                repositories.OpportunityRepository
	notificationService *NotificationService
}

func NewOpportunityService(repo repositories.OpportunityRepository, ns *NotificationService) OpportunityService {
	return &opportunityService{repo: repo, notificationService: ns}
}

func (s *opportunityService) CreateOpportunity(opportunity *models.Opportunity) error {
	opportunity.Status = strings.ToLower(opportunity.Status)
	if opportunity.Status != "open" && opportunity.Status != "closed" {
		return errors.New("status inválido, deve ser 'open' ou 'closed'")
	}

	err := s.repo.Create(opportunity)
	if err == nil {
		s.notificationService.NotifyAll(fmt.Sprintf("Nova oportunidade: %s", opportunity.Title))
	}
	return err
}

func (s *opportunityService) GetOpportunity(id uint) (*models.Opportunity, error) {
	return s.repo.GetByID(id)
}

func (s *opportunityService) GetAllOpportunities(status, location string) ([]models.Opportunity, error) {
	return s.repo.GetAll(status, location)
}

func (s *opportunityService) UpdateOpportunity(id uint, updatedOpportunity *models.Opportunity) error {
	existingOpportunity, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	existingOpportunity.Title = updatedOpportunity.Title
	existingOpportunity.Description = updatedOpportunity.Description
	existingOpportunity.Location = updatedOpportunity.Location
	existingOpportunity.Status = strings.ToLower(updatedOpportunity.Status)

	return s.repo.Update(existingOpportunity)
}

func (s *opportunityService) DeleteOpportunity(id uint) error {
	return s.repo.Delete(id)
}
