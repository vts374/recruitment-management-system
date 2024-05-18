package models

import "time"

type User struct {
	ID              uint `gorm:"primaryKey"`
	Name            string
	Email           string `gorm:"unique"`
	Address         string
	UserType        string // "Applicant" or "Admin"
	PasswordHash    string
	ProfileHeadline string
	Profile         Profile
}

type Profile struct {
	ID                uint `gorm:"primaryKey"`
	UserID            uint `gorm:"unique"`
	ResumeFileAddress string
	Skills            string
	Education         string
	Experience        string
	Name              string
	Email             string
	Phone             string
}

type Job struct {
	ID                uint `gorm:"primaryKey"`
	Title             string
	Description       string
	PostedOn          time.Time
	TotalApplications int
	CompanyName       string
	PostedBy          uint // UserID of the admin who posted the job
}
