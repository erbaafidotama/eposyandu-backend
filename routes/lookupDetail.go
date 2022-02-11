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

	db.Create(&lookupDetail)
	c.JSON(200, gin.H{
		"status": "berhasil post",
		"data":   lookupDetailReq,
	})
}
