package routes

import (
	"eposyandu/config"
	"eposyandu/models"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type anakRequest struct {
	// LookupUuid     uuid.UUID `json:"lookup_uuid"`
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

func PostAnak(c *gin.Context) {
	db := config.GetDB()
	var anakReq anakRequest

	if err := c.BindJSON(&anakReq); err != nil {
		fmt.Println("ERROR BINDJSON", err)
	}

	anak := models.AllAnak{
		AnakUuid:       uuid.New(),
		NamaAnak:       anakReq.NamaAnak,
		NoKk:           anakReq.NoKk,
		Nik:            anakReq.Nik,
		TempatLahir:    anakReq.TempatLahir,
		TanggalLahir:   anakReq.TanggalLahir,
		JenisKelaminId: anakReq.JenisKelaminId,
		NamaOrangTua:   anakReq.NamaOrangTua,
		NikOrangTua:    anakReq.NikOrangTua,
		Alamat:         anakReq.Alamat,
		Rt:             anakReq.Rt,
		Rw:             anakReq.Rw,
		AnakKe:         anakReq.AnakKe,
		LulusPosyandu:  anakReq.LulusPosyandu,
	}

	if err := db.Create(&anak).Error; err != nil {
		c.JSON(500, gin.H{
			"status": "gagal create",
		})
	} else {
		c.JSON(200, gin.H{
			"status": "berhasil create",
			"data":   anak,
		})
	}
}

// func PutLookup(c *gin.Context) {
// 	db := config.GetDB()
// 	var lookupReq lookupRequest

// 	if err := c.BindJSON(&lookupReq); err != nil {
// 		fmt.Println("ERROR BINDJSON", err)
// 	}

// 	var flag bool
// 	var activeFlag bool = lookupReq.ActiveFlag
// 	if !activeFlag {
// 		flag = true
// 	} else {
// 		flag = lookupReq.ActiveFlag
// 	}

// 	lookup := models.Lookup{
// 		LookupCode:  lookupReq.LookupCode,
// 		Description: lookupReq.Description,
// 		ActiveFlag:  flag,
// 	}

// 	if err := db.Model(&lookup).Updates(lookup).Error; err != nil {
// 		c.JSON(500, gin.H{
// 			"status": "gagal update",
// 		})
// 	} else {
// 		c.JSON(200, gin.H{
// 			"status": "berhasil update",
// 			"data":   lookup,
// 		})
// 	}
// }

func GetAnak(c *gin.Context) {
	db := config.GetDB()
	var anak []models.AllAnak

	if err := db.Find(&anak).Error; err != nil {
		c.JSON(500, gin.H{
			"status": "gagal get",
		})
	} else {
		c.JSON(200, gin.H{
			"status": "berhasil get",
			"data":   anak,
		})
	}
}

func GetAnakByUuid(c *gin.Context) {
	db := config.GetDB()
	var anak models.AllAnak

	uuid := c.Param("anak_uuid")
	if err := db.Where("anak_uuid = ?", uuid).First(&anak).Error; err != nil {
		c.JSON(500, gin.H{
			"status": "gagal get",
		})
	} else {
		c.JSON(200, gin.H{
			"status": "berhasil get",
			"data":   anak,
		})
	}
}
