package config

import (
	"log"
	"os"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DeviceInfo struct {
	gorm.Model
	Hostname  string `gorm:"column:hostname" json:"hostname"`
	SessionID string `gorm:"column:session_id" json:"sessionID"`
	Key       string `gorm:"column:key" json:"key"`
}

type AgentConfig struct {
	gorm.Model
	EnrollmentUrl string `gorm:"column:enrollment_url" json:"enrollmentUrl"`
	EnrolledUrl   string `gorm:"column:enrolled_url" json:"enrolledUrl"`
	MQTT          string `gorm:"column:mqtt" json:"mqtt"`
	Topic         string `gorm:"column:topic" json:"topic"`
}

func GetDeviceInfo() *DeviceInfo {
	hostname := GetHostName()
	if hostname == "" {
		log.Println("Blank Hostname")
		return nil
	}

	return &DeviceInfo{
		Hostname:  hostname,
		SessionID: uuid.New().String(),
		Key:       os.Getenv("KEY"),
	}
}

func GetHostName() string {
	hostname, err := os.Hostname()
	if err != nil {
		log.Println("Error:", err)
		return ""
	}

	return hostname
}
