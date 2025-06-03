package models

type UserPreference struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	UserID   uint   `json:"user_id" gorm:"uniqueIndex"`
	Location string `json:"location,omitempty"`
	JobType  string `json:"job_type,omitempty"`
	Keywords string `json:"keywords,omitempty"`
}
