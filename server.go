package main

import (
	"eposyandu/config"
	"eposyandu/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// db := config.GetDB()
	//when server.go start, it will be run function InitDB (connecting to database)
	// config.InitDB()
	config.GetDB()

	router := gin.Default()

	v1 := router.Group("api/v1") // /api/v1
	{
		v1.POST("/login", routes.Login) // /api/v1/login
		admin := v1.Group("/admin")     // /api/v1/sewaAset
		{
			admin.GET("/account", routes.GetAccount)   // /api/v1/admin/account
			admin.POST("/account", routes.PostAccount) // /api/v1/admin/account
		}
		// v1.POST("/warga", middleware.IsAuth(), routes.PostWarga)
		// v1.GET("/warga", middleware.IsAuth(), routes.GetWarga)
		v1.POST("/lookup", routes.PostLookup)
		v1.GET("/lookup", routes.GetLookup)
		v1.GET("/lookup/:lookup_uuid", routes.GetLookupByUuid)
		v1.PUT("/lookup/:lookup_uuid", routes.PutLookup)
		v1.POST("/lookupDetail", routes.PostLookupDetail)
	}

	router.Run()
}
