package routes

import (
	"github.com/gin-gonic/gin"
	"job-app/main/controller"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/jobs", controller.GetJobs)
	router.POST("/job", controller.CreateJob)
}
