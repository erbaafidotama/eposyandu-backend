package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Account model
type AllAnak struct {
	gorm.Model
	AnakUuid       uuid.UUID `gorm:"type:uuid" json:"anak_uuid"`
	NamaAnak       string    `json:"nama_anak"`
	NoKk           string    `json:"no_kk"`
	Nik            string    `json:"nik"`
	TempatLahir    string    `json:"tempat_lahir"`
	TanggalLahir   time.Time `json:"tanggal_lahir"`
	JenisKelaminId string    `json:"jenis_kelamin_id"`
	NamaOrangTua   string    `json:"nama_orang_tua"`
	NikOrangTua    string    `json:"nik_orang_tua"`
	Alamat         string    `json:"alamat"`
	Rt             string    `json:"rt"`
	Rw             string    `json:"rw"`
	AnakKe         string    `json:"anak_ke"`
	LulusPosyandu  bool      `json:"lulus_posyandu"`
}
