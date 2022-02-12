package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LookupDetail struct {
	gorm.Model
	LookupDetailUuid uuid.UUID `gorm:"type:uuid" json:"lookup_detail_uuid"`
	Name             string
	Description      string
	ActiveFlag       bool
	LookupID         uint
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
