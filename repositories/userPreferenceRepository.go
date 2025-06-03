package repositories

import (
	"github.com/alysonsz/gopportunities.git/models"
	"gorm.io/gorm"
)

type UserPreferenceRepository interface {
	Save(preference *models.UserPreference) error
	FindByUserID(userID uint) (*models.UserPreference, error)
	Update(preference *models.UserPreference) error
}

type userPreferenceRepository struct {
	db *gorm.DB
}

func NewUserPreferenceRepository(db *gorm.DB) UserPreferenceRepository {
	return &userPreferenceRepository{db}
}

func (r *userPreferenceRepository) Save(preference *models.UserPreference) error {
	return r.db.Create(preference).Error
}

func (r *userPreferenceRepository) FindByUserID(userID uint) (*models.UserPreference, error) {
	var pref models.UserPreference
	if err := r.db.Where("user_id = ?", userID).First(&pref).Error; err != nil {
		return nil, err
	}
	return &pref, nil
}

func (r *userPreferenceRepository) Update(preference *models.UserPreference) error {
	return r.db.Save(preference).Error
}
