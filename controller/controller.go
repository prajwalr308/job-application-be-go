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
	var jobs []model.Job
	sqlStatement := `SELECT * FROM jobs`
	if db.DB == nil {
		log.Fatal("Database connection is not initialized")
	}
	rows, err := db.DB.Query(sqlStatement)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()
	for rows.Next() {
		var job model.Job
		err := rows.Scan(&job.ID, &job.Title, &job.Description, &job.Status, &job.Created_at)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		jobs = append(jobs, job)
	}
	c.JSON(http.StatusOK, jobs)

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

func UpdateJob(c *gin.Context) {
	var updatedJob model.Job
	if err := c.ShouldBindJSON(&updatedJob); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := c.Param("id")
	sqlStatement := `UPDATE jobs SET title=$2, description=$3, status=$4 WHERE id=$1`

	result, err := db.DB.Exec(sqlStatement, id, updatedJob.Title, updatedJob.Description, updatedJob.Status)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No job found with the given ID"})
		return
	}
	c.JSON(http.StatusOK, updatedJob)
}
