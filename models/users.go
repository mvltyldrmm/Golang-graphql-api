package models

import (
	"time"

	"github.com/google/uuid"
)

type Users struct {
	Id         uuid.UUID `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	LanguageId uuid.UUID `json:"language_id"`
	CountryId  uuid.UUID `json:"country_id"`
	Name       string    `json:"name"`
	Surname    string    `json:"surname"`
	Email      string    `json:"email"`
	IsDeleted  bool      `json:"is_deleted"`
	DeletedAt  time.Time `json:"deleted_at"`
	IsActıve   bool      `json:"is_active"`
	Number     string    `json:"number"`
}
