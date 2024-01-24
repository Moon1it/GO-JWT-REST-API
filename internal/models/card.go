package models

type Card struct {
	ID         uint   `json:"id" gorm:"primaryKey"`
	UserID     uint   `json:"user"`
	Item       string `json:"item"`
	Definition string `json:"definition"`
}
