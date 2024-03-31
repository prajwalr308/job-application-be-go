package controller

import (
	"fmt"
	"job-app/main/db"
	"job-app/main/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetJobs(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": "GetJobs",
	})
}

func CreateJob(c *gin.Context) {
	var newJob model.Job

	if err := c.ShouldBindJSON(&newJob); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(c)
	fmt.Println(newJob)
	sqlStatement := `INSERT INTO jobs (title, description, status) VALUES ($1, $2, $3) RETURNING id`
	id := 0
	if db.DB == nil {
		log.Fatal("Database connection is not initialized")
	}
	err := db.DB.QueryRow(sqlStatement, newJob.Title, newJob.Description, newJob.Status).Scan(&id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newJob.ID = id
	c.JSON(http.StatusCreated, newJob)
}
