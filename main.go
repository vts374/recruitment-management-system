package main

import (
	"github.com/vts374/recruitment-management-system/controllers"
	"github.com/vts374/recruitment-management-system/models"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	r := gin.Default()

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.User{}, &models.Profile{}, &models.Job{})

	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.POST("/uploadResume", controllers.UploadResume)
	r.POST("/admin/job", controllers.CreateJob)
	r.GET("/admin/job/:job_id", controllers.GetJob)
	r.GET("/admin/applicants", controllers.GetApplicants)
	r.GET("/admin/applicant/:applicant_id", controllers.GetApplicant)
	r.GET("/jobs", controllers.GetJobs)
	r.GET("/jobs/apply", controllers.ApplyJob)

	r.Run()
}
