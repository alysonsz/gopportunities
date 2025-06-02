package repositories

import (
	"github.com/alysonsz/gopportunities.git/models"
	"gorm.io/gorm"
)

// OpportunityRepository define operações para gerenciar oportunidades
type OpportunityRepository interface {
	Create(opportunity *models.Opportunity) error
	GetByID(id uint) (*models.Opportunity, error)
	GetAll(status string, location string) ([]models.Opportunity, error)
	Update(opportunity *models.Opportunity) error
	Delete(id uint) error
}

type opportunityRepository struct {
	db *gorm.DB
}

// NewOpportunityRepository cria uma instância do repositório
func NewOpportunityRepository(db *gorm.DB) OpportunityRepository {
	return &opportunityRepository{db: db}
}

// Create cria uma nova oportunidade
func (r *opportunityRepository) Create(opportunity *models.Opportunity) error {
	return r.db.Create(opportunity).Error
}

// GetByID retorna uma oportunidade específica pelo ID
func (r *opportunityRepository) GetByID(id uint) (*models.Opportunity, error) {
	var opportunity models.Opportunity
	err := r.db.First(&opportunity, id).Error
	return &opportunity, err
}

// GetAll retorna todas as oportunidades, com filtros opcionais
func (r *opportunityRepository) GetAll(status string, location string) ([]models.Opportunity, error) {
	var opportunities []models.Opportunity
	query := r.db.Model(&models.Opportunity{})

	if status != "" {
		query = query.Where("status = ?", status)
	}

	if location != "" {
		query = query.Where("location = ?", location)
	}

	err := query.Find(&opportunities).Error
	return opportunities, err
}

// Update atualiza uma oportunidade existente
func (r *opportunityRepository) Update(opportunity *models.Opportunity) error {
	return r.db.Save(opportunity).Error
}

// Delete remove uma oportunidade pelo ID
func (r *opportunityRepository) Delete(id uint) error {
	return r.db.Delete(&models.Opportunity{}, id).Error
}
