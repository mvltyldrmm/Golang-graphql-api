package models

import (
	"github.com/google/uuid"
	"time"
)

type Companies struct {
	Id            uuid.UUID `json:"id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	LanguageId    uuid.UUID `json:"language_id"`
	CountryId     uuid.UUID `json:"country_id"`
	CompanyName   string    `json:"company_name" validate:"required"`
	CompanyTitle  string    `json:"company_title" validate:"required"`
	CompanyDesc   string    `json:"company_desc" validate:"required"`
	CompanyEmail  string    `json:"company_email" validate:"required"`
	IsDeleted     bool      `json:"is_deleted"`
	DeletedAt     time.Time `json:"deleted_at"`
	IsActived     bool      `json:"is_active"`
	CompanyNumber string    `json:"company_number" validate:"required"`
	AdressLine1   string    `json:"address_line1" validate:"required"`
	AdressLine2   string    `json:"address_line2" validate:"required"`
}
