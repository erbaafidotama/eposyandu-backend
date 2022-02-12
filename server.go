package main

import (
	"eposyandu/config"
	"eposyandu/middleware"
	"eposyandu/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// db := config.GetDB()
	//when server.go start, it will be run function InitDB (connecting to database)
	// config.InitDB()
	config.GetDB()

	// router := gin.Default()
	router := gin.New()
	router.Use(middleware.CORSMiddleware())

	router.POST("/login", routes.Login)
	admin := router.Group("/admin") //admin is the prefix
	{
		admin.GET("/account", routes.GetAccount)   // /admin/account
		admin.POST("/account", routes.PostAccount) // /admin/account
	}
	lookup := router.Group("/lookup") //lookup is the prefix
	{
		lookup.GET("/", routes.GetLookup)                   // /lookup
		lookup.GET("/:lookup_uuid", routes.GetLookupByUuid) // /lookup/:lookup_uuid
		lookup.POST("/", routes.PostLookup)                 // /lookup
		lookup.PUT("/:lookup_uuid", routes.PutLookup)       // /lookup/:lookup_uuid
	}
	lookupDetail := router.Group("/lookup_detail") //lookup_detail is the prefix
	{
		lookupDetail.GET("/", routes.GetLookupDetail)                          // /lookup_detail
		lookupDetail.GET("/:lookup_detail_uuid", routes.GetLookupDetailByUuid) // /lookup_detail/:lookup_detail_uuid
		lookupDetail.POST("/", routes.PostLookupDetail)                        // /lookup_detail
		lookupDetail.PUT("/:lookup_detail_uuid", routes.PutLookupDetail)       // /lookup_detail/:lookup_detail_uuid
	}
	warga := router.Group("/warga") //warga is the prefix
	{
		warga.GET("/", routes.GetWarga) // /warga
		// warga.GET("/:warga_uuid", routes.GetWargaByUuid) // /warga/:warga_uuid
		warga.POST("/", routes.PostWarga) // /warga
		// warga.PUT("/:warga_uuid", routes.PutWarga)       // /warga/:warga_uuid
	}

	// router.Run(":8080") // if you want to run on port 8080
	router.Run()
}
