package models

type User struct {
	ID             uint   `json:"id" validate:"required,omitempty" gorm:"primaryKey"`
	Email          string `json:"email" gorm:"unique;not null"`
	Password       string `json:"password"`
	Role           string `json:"role"`
	IsActivated    bool   `json:"is_activated"`
	ActivationLink string `json:"activation_link"`
}
