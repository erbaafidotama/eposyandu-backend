package routes

import (
	"admin-rt/config"
	"admin-rt/models"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type wargaRequest struct {
	NamaWarga    string `json:"nama_warga"`
	NoKk         string `json:"no_kk"`
	Nik          string `json:"nik"`
	Alamat       string `json:"alamat"`
	TypeKeluarga int    `json:"type_keluarga"`
	TanggalLahir string `json:"tanggal_lahir"`
	TempatLahir  string `json:"tempat_lahir"`
}

func PostWarga(c *gin.Context) {
	db := config.GetDB()
	var wargaReq wargaRequest

	if err := c.BindJSON(&wargaReq); err != nil {
		fmt.Println("ERROR BINDJSON", err)
	}
	// convert string date to date db
	dateStr := wargaReq.TanggalLahir
	format := "2006-01-02"
	date, _ := time.Parse(format, dateStr)

	warga := models.Warga{
		NamaWarga:    wargaReq.NamaWarga,
		NoKk:         wargaReq.NoKk,
		Nik:          wargaReq.Nik,
		Alamat:       wargaReq.Alamat,
		TanggalLahir: date,
		TypeKeluarga: wargaReq.TypeKeluarga,
		TempatLahir:  wargaReq.TempatLahir,
	}

	db.Create(&warga)
	c.JSON(200, gin.H{
		"status": "berhasil post",
		"data":   wargaReq,
	})
}

// GetWarga Route
func GetWarga(c *gin.Context) {
	db := config.GetDB()
	wargas := []models.Warga{}

	// select * from User
	if err := db.Find(&wargas).Error; err != nil {
		// return error
		c.JSON(404, gin.H{
			"status":  "error",
			"message": "record not found",
		})
		c.Abort()
		return
	}

	// return complete
	c.JSON(200, gin.H{
		"message": "GET data Account",
		"data":    wargas,
	})
}
