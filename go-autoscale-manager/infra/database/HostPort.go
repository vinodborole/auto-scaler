package database

import "github.com/jinzhu/gorm"

type HostPort struct {
	gorm.Model

	Port   string `json:"port"`
	Status string `json:"status"`
}

func GetFreePort() (HostPort, error) {
	var hostPort HostPort
	err := DBConn.Where("Status=Free").First(&hostPort).Error
	return hostPort, err
}

func UpdatePortStatus(hostPort *HostPort) error {
	return DBConn.Save(&hostPort).Error
}
