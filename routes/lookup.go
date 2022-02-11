package routes

import (
	"eposyandu/config"
	"eposyandu/models"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type lookupRequest struct {
	// LookupUuid     uuid.UUID `json:"lookup_uuid"`
	LookupCode     string `json:"lookup_code"`
	Description    string `json:"description"`
	ActiveFlag     bool   `json:"active_flag"`
	LookupDetailID uint   `json:"lookup_detail_id"`
}

func PostLookup(c *gin.Context) {
	db := config.GetDB()
	var lookupReq lookupRequest

	if err := c.BindJSON(&lookupReq); err != nil {
		fmt.Println("ERROR BINDJSON", err)
	}

	var flag bool
	var activeFlag bool = lookupReq.ActiveFlag
	if !activeFlag {
		flag = true
	} else {
		flag = lookupReq.ActiveFlag
	}

	lookup := models.Lookup{
		LookupCode:  lookupReq.LookupCode,
		Description: lookupReq.Description,
		ActiveFlag:  flag,
		LookupUuid:  uuid.New(),
	}

	db.Create(&lookup)
	c.JSON(200, gin.H{
		"status": "berhasil post",
		"data":   lookupReq,
	})
}

func PutLookup(c *gin.Context) {
	db := config.GetDB()
	var lookupReq lookupRequest

	if err := c.BindJSON(&lookupReq); err != nil {
		fmt.Println("ERROR BINDJSON", err)
	}

	var flag bool
	var activeFlag bool = lookupReq.ActiveFlag
	if !activeFlag {
		flag = true
	} else {
		flag = lookupReq.ActiveFlag
	}

	lookup := models.Lookup{
		LookupCode:  lookupReq.LookupCode,
		Description: lookupReq.Description,
		ActiveFlag:  flag,
	}

	if err := db.Model(&lookup).Updates(lookup).Error; err != nil {
		c.JSON(500, gin.H{
			"status": "gagal update",
		})
	} else {
		c.JSON(200, gin.H{
			"status": "berhasil update",
			"data":   lookup,
		})
	}
	// c.JSON(200, gin.H{
	// 	"status": "berhasil put",
	// 	"data":   lookupReq,
	// })
}

func GetLookup(c *gin.Context) {
	db := config.GetDB()
	var lookup []models.Lookup

	db.Find(&lookup)

	c.JSON(200, gin.H{
		"status": "berhasil get",
		"data":   lookup,
	})
}

func GetLookupByUuid(c *gin.Context) {
	db := config.GetDB()
	var lookup models.Lookup

	uuid := c.Param("lookup_uuid")
	db.Where("lookup_uuid = ?", uuid).First(&lookup)

	c.JSON(200, gin.H{
		"status": "berhasil get",
		"data":   lookup,
	})
}
