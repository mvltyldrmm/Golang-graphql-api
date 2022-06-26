package models

import (
	"time"

	"github.com/google/uuid"
)

type Partnership struct {
	CompanyId string    `json:"company_id"`
	UserId    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	DetailId  string    `json:"detail_id"`
	Id        uuid.UUID `json:"id"`
}
