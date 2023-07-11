package database

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
)

func InitialiseDB() {
	var err error
	fmt.Println("initialising DB")
	DBConn, err = gorm.Open(sqlite.Open("jobs.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("Error initialising :", err.Error())
		panic("failed to connect database")
	}
	fmt.Println("Database connection successfully Opened !")
	DBConn.AutoMigrate(Job{}, HostPort{})
	migrateHostPortData()
	fmt.Println("Migration completed successfully")
}

func migrateHostPortData() {
	ports := make([]HostPort, 10)
	port1 := HostPort{Port: "8000", Status: "Free"}
	port2 := HostPort{Port: "8001", Status: "Free"}
	port3 := HostPort{Port: "8002", Status: "Free"}
	port4 := HostPort{Port: "8003", Status: "Free"}
	port5 := HostPort{Port: "8004", Status: "Free"}
	port6 := HostPort{Port: "8005", Status: "Free"}
	port7 := HostPort{Port: "8006", Status: "Free"}
	port8 := HostPort{Port: "8007", Status: "Free"}
	port9 := HostPort{Port: "8008", Status: "Free"}
	port10 := HostPort{Port: "8009", Status: "Free"}
	ports = append(ports, port1)
	ports = append(ports, port2)
	ports = append(ports, port3)
	ports = append(ports, port4)
	ports = append(ports, port5)
	ports = append(ports, port6)
	ports = append(ports, port7)
	ports = append(ports, port8)
	ports = append(ports, port9)
	ports = append(ports, port10)

	err := DBConn.Create(ports).Error
	if err != nil {
		fmt.Println("Error adding ports data : " + err.Error())
	}
}
