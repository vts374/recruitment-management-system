package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	// "strings"

	"github.com/gin-gonic/gin"
	"github.com/vts374/recruitment-management-system/models"
)

// Function to upload resume and parse it
func UploadResume(c *gin.Context) {
	// Ensure the user is authenticated and is an applicant
	// Upload and parse resume
	file, err := c.FormFile("resume")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to upload file"})
		return
	}

	// Save the file
	filePath := "./uploads/" + file.Filename
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	// Parse the resume
	parsedData, err := parseResume(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse resume"})
		return
	}

	c.JSON(http.StatusOK, parsedData)
}

// Function to create job openings
func CreateJob(c *gin.Context) {
	// Ensure the user is authenticated and is an admin
	// Create job opening
}

// Function to fetch job details
func GetJob(c *gin.Context) {
	// Ensure the user is authenticated
	// Retrieve job details
}

// Function to fetch all users
func GetApplicants(c *gin.Context) {
	// Ensure the user is authenticated and is an admin
	// Retrieve all users
	    // Ensure the user is authenticated and is an admin
    var applicants []models.User
    // Retrieve all applicants from the database

    c.JSON(http.StatusOK, applicants)
}

// Function to fetch a specific applicant's data
func GetApplicant(c *gin.Context) {
	// Ensure the user is authenticated and is an admin
	    // applicantID := c.Param("applicant_id")
    var applicant models.User
    // Retrieve the applicant's data from the database using applicantID

    c.JSON(http.StatusOK, applicant)

}

// Function to fetch all job openings
func GetJobs(c *gin.Context) {
	// Ensure the user is authenticated
	// Retrieve all job openings
	    var jobs []models.Job
    // Retrieve all job openings from the database

    c.JSON(http.StatusOK, jobs)
}

// Function to apply for a job
func ApplyJob(c *gin.Context) {
	// Ensure the user is authenticated and is an applicant
	// Apply for job
	 // Ensure the user is authenticated and is an applicant
    var application struct {
        JobID uint `json:"job_id"`
    }
    if err := c.ShouldBindJSON(&application); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
        return
}
func parseResume(filePath string, map[string]interface{}, error) {
    url := "https://api.apilayer.com/resume_parser/upload"
    method := "POST"

    payload, err := os.Open(filePath)
    if err != nil {
        return nil, err
    }
	defer payload.Close()

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/octet-stream")
	req.Header.Add("apikey", "gNiXyflsFu3WNYCz1ZCxdWDb7oQg1Nl1")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// func ScoreProfile(job Job, profile Profile) float64 {
// 	// Simple heuristic: Match skills and experience
// 	score := 0.0
// 	if strings.Contains(profile.Skills, "matching_skill") {
// 		score += 1.0
// 	}
// 	if strings.Contains(profile.Experience, "relevant_experience") {
// 		score += 1.0
// 	}
// 	return score
// }
