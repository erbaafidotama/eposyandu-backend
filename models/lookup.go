package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Lookup struct {
	gorm.Model
	LookupUuid    uuid.UUID      `gorm:"type:uuid" json:"lookup_uuid"`
	LookupCode    string         `json:"lookup_code"`
	Description   string         `json:"description"`
	ActiveFlag    bool           `json:"active_flag"`
	LookupDetails []LookupDetail `json:"lookup_detail_id"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
