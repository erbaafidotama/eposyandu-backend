package models

import (
	"time"

	"gorm.io/gorm"
)

type Warga struct {
	gorm.Model
	NamaWarga    string
	NoKk         string
	Nik          string
	Alamat       string
	TypeKeluarga int
	TanggalLahir time.Time `json:"tanggal_lahir"`
	TempatLahir  string
}
