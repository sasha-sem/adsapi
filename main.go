package main

import (
	"sashasem/adsapi/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	v1 := r.Group("api/v1")
	{
		v1.POST("/ads", controllers.PostAd)
		v1.GET("/ads", controllers.GetAds)
		v1.GET("/ads/:id", controllers.GetAd)
	}

	r.Run(":8080")
}
