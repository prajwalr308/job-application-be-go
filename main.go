package main

import (
	"job-app/main/db"
	"job-app/main/routes"

	"github.com/gin-gonic/gin"
)

/*  dlv dap -l 127.0.0.1:38697 --log --log-output="dap"   */
func main() {
	db.InitDB()
	//create a new gin router
	router := gin.Default()
	//register a routes
	routes.RegisterRoutes(router)

	//run
	router.Run(":8080")

}
