package models

type Token struct {
	ID           uint   `json:"id" gorm:"primaryKey"`
	UserID       uint   `json:"user"`
	RefreshToken string `json:"refresh_token"`
}

type TokenPair struct {
	AccessToken  string
	RefreshToken string
}

type TokenPayload struct {
	ID    uint
	Email string
}
