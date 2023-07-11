package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Job struct {
	gorm.Model

	Name    string `json:"name"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

func GetJobs() ([]Job, error) {
	var jobs []Job
	err := DBConn.Find(&jobs).Error
	return jobs, err
}

func GetJob(name string) (Job, error) {
	var job Job
	err := DBConn.Where("name = ?", name).First(&job).Error
	return job, err
}

func CreateJob(job *Job) error {
	return DBConn.Create(&job).Error
}

func UpdateJob(job *Job) error {
	return DBConn.Save(&job).Error
}
