package routes

import (
	"eposyandu/config"
	"eposyandu/models"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type lookupDetailRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ActiveFlag  bool   `json:"active_flag"`
	LookupID    uint   `json:"lookup_id"`
}

func PostLookupDetail(c *gin.Context) {
	db := config.GetDB()
	var lookupDetailReq lookupDetailRequest

	if err := c.BindJSON(&lookupDetailReq); err != nil {
		fmt.Println("ERROR BINDJSON", err)
	}

	var flag bool
	activeFlag := lookupDetailReq.ActiveFlag
	if !activeFlag {
		flag = true
	} else {
		flag = lookupDetailReq.ActiveFlag
	}

	lookupDetail := models.LookupDetail{
		Name:             lookupDetailReq.Name,
		Description:      lookupDetailReq.Description,
		ActiveFlag:       flag,
		LookupID:         lookupDetailReq.LookupID,
		LookupDetailUuid: uuid.New(),
	}

	if err := db.Create(&lookupDetail).Error; err != nil {
		c.JSON(500, gin.H{
			"status": "gagal create",
		})
	} else {
		c.JSON(200, gin.H{
			"status": "berhasil create",
			"data":   lookupDetail,
		})
	}
}

func PutLookupDetail(c *gin.Context) {
	db := config.GetDB()
	var lookupDetailReq lookupDetailRequest

	if err := c.BindJSON(&lookupDetailReq); err != nil {
		fmt.Println("ERROR BINDJSON", err)
	}

	var flag bool
	var activeFlag bool = lookupDetailReq.ActiveFlag
	if !activeFlag {
		flag = true
	} else {
		flag = lookupDetailReq.ActiveFlag
	}

	lookupDetail := models.LookupDetail{
		Name:        lookupDetailReq.Name,
		Description: lookupDetailReq.Description,
		ActiveFlag:  flag,
		LookupID:    lookupDetailReq.LookupID,
	}

	if err := db.Model(&lookupDetail).Updates(lookupDetail).Error; err != nil {
		c.JSON(500, gin.H{
			"status": "gagal update",
		})
	} else {
		c.JSON(200, gin.H{
			"status": "berhasil update",
			"data":   lookupDetail,
		})
	}
}

func GetLookupDetail(c *gin.Context) {
	db := config.GetDB()
	var lookupDetail []models.LookupDetail

	if err := db.Find(&lookupDetail).Error; err != nil {
		c.JSON(500, gin.H{
			"status": "gagal get",
		})
	} else {
		c.JSON(200, gin.H{
			"status": "berhasil get",
			"data":   lookupDetail,
		})
	}
}

func GetLookupDetailByUuid(c *gin.Context) {
	db := config.GetDB()
	var lookupDetail models.LookupDetail

	uuid := c.Param("lookup_detail_uuid")
	if err := db.Where("lookup_detail_uuid = ?", uuid).First(&lookupDetail).Error; err != nil {
		c.JSON(500, gin.H{
			"status": "gagal get",
		})
	} else {
		c.JSON(200, gin.H{
			"status": "berhasil get",
			"data":   lookupDetail,
		})
	}
}
