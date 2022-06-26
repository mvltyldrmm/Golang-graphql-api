package models

import (
	"time"

	"github.com/google/uuid"
)

type PartnershipDetails struct {
	Id            uuid.UUID `json:"id"`
	DeactiveAt    time.Time `json:"deactive_at"`
	PartnershipId uuid.UUID `json:"partnership_id"`
	IsDeactive    bool      `json:"is_deactive"`
	IsActive      bool      `json:"is_active"`
	AccountType   int       `json:"account_type"`
}
