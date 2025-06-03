package services

import (
	"github.com/alysonsz/gopportunities.git/models"
	"github.com/alysonsz/gopportunities.git/repositories"
)

type UserPreferenceService interface {
	SetPreference(userID uint, pref *models.UserPreference) error
	GetPreference(userID uint) (*models.UserPreference, error)
}

type userPreferenceService struct {
	repo repositories.UserPreferenceRepository
}

func NewUserPreferenceService(repo repositories.UserPreferenceRepository) UserPreferenceService {
	return &userPreferenceService{repo}
}

func (s *userPreferenceService) SetPreference(userID uint, pref *models.UserPreference) error {
	pref.UserID = userID
	existing, err := s.repo.FindByUserID(userID)
	if err == nil && existing != nil {
		pref.ID = existing.ID
		return s.repo.Update(pref)
	}
	return s.repo.Save(pref)
}

func (s *userPreferenceService) GetPreference(userID uint) (*models.UserPreference, error) {
	return s.repo.FindByUserID(userID)
}
