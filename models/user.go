package models

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Username string `gorm:"unique" json:"username" binding: "required"`
	Password string `json:"password" binding: "required"`
}
