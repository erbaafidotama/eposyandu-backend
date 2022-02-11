package models

import (
	"time"

	"gorm.io/gorm"
)

type Warga struct {
	gorm.Model
	NamaWarga                    string
	KelaminLookupDetailID        int
	NoKk                         string
	Nik                          string
	Alamat                       string
	RtLookupDetailID             int
	RwLookupDetailID             int
	NoTelp                       string
	StatusKeluargaLookupDetailID int
	TanggalLahir                 time.Time `json:"tanggal_lahir"`
	TempatLahir                  string
	LulusPosyandu                bool
}
