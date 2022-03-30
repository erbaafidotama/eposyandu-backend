package models

import (
	"time"

	"gorm.io/gorm"
)

type Warga struct {
	gorm.Model
	NamaWarga                    string    `json:"nama_warga"`
	KelaminLookupDetailID        int       `json:"kelamin_lookup_detail_id"`
	NoKk                         string    `json:"no_kk"`
	Nik                          string    `json:"nik"`
	Alamat                       string    `json:"alamat"`
	RtLookupDetailID             int       `json:"rt_lookup_detail_id"`
	RwLookupDetailID             int       `json:"rw_lookup_detail_id"`
	NoTelp                       string    `json:"no_telp"`
	StatusKeluargaLookupDetailID int       `json:"status_keluarga_lookup_detail_id"`
	TanggalLahir                 time.Time `json:"tanggal_lahir"`
	TempatLahir                  string    `json:"tempat_lahir"`
	LulusPosyandu                bool      `json:"lulus_posyandu"`
}
