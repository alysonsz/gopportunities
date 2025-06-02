package services

import (
	"errors"
	"strings"

	"github.com/alysonsz/gopportunities.git/models"
	"github.com/alysonsz/gopportunities.git/repositories"
)

// OpportunityService define métodos para regras de negócio
type OpportunityService interface {
	CreateOpportunity(opportunity *models.Opportunity) error
	GetOpportunity(id uint) (*models.Opportunity, error)
	GetAllOpportunities(status, location string) ([]models.Opportunity, error)
	UpdateOpportunity(id uint, opportunity *models.Opportunity) error
	DeleteOpportunity(id uint) error
}

type opportunityService struct {
	repo repositories.OpportunityRepository
}

// NewOpportunityService instancia um novo serviço
func NewOpportunityService(repo repositories.OpportunityRepository) OpportunityService {
	return &opportunityService{repo: repo}
}

func (s *opportunityService) CreateOpportunity(opportunity *models.Opportunity) error {
	opportunity.Status = strings.ToLower(opportunity.Status)
	if opportunity.Status != "open" && opportunity.Status != "closed" {
		return errors.New("status inválido, deve ser 'open' ou 'closed'")
	}
	return s.repo.Create(opportunity)
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
