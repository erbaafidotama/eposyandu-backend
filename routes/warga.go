package routes

import (
	"admin-rt/config"
	"admin-rt/models"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// PostWarga this
func PostWarga(c *gin.Context) {
	db := config.GetDB()
	typeKeluarga, err := strconv.Atoi(c.PostForm("type_keluarga"))
	if err != nil {
		fmt.Println("ERROR Covert Type")
	}
	// convert string date to date db
	dateStr := c.PostForm("tanggal_lahir")
	format := "2006-01-02"
	date, _ := time.Parse(format, dateStr)

	// make object from form body
	warga := models.Warga{
		NamaWarga:    c.PostForm("nama_warga"),
		NoKk:         c.PostForm("no_kk"),
		Nik:          c.PostForm("nik"),
		Alamat:       c.PostForm("alamat"),
		TanggalLahir: date,
		TypeKeluarga: typeKeluarga,
		TempatLahir:  c.PostForm("tempat_lahir"),
	}

	// crete data to db
	// config.DB.Create(&account)
	db.Create(&warga)

	c.JSON(200, gin.H{
		"status": "berhasil post",
		"data":   warga,
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
