package config

import (
	"log"
	"os"

	"gorm.io/gorm"
)

type DeviceInfo struct {
	gorm.Model
	Hostname  string `gorm:"column:hostname" json:"hostname"`
	SessionID string `gorm:"column:session_id" json:"sessionID"`
	Key       string `gorm:"column:key" json:"key"`
}

// Request body
type Login struct {
	Username  string `gorm:"username;index:idx_username,unique" json:"username,omitempty"`
	DeviceKey string `gorm:"column:device_key" json:"device_key,omitempty"`
	Hostname  string `gorm:"column:hostname" json:"hostname,omitempty"`
}

type LoginResponse struct {
	ClientID  string `gorm:"client_id" json:"client_id,omitempty"`
	SessionID string `gorm:"column:session_id" json:"sessionID,omitempty"`
	Topic     string `gorm:"column:topic" json:"topic"`
}

func GetDeviceInfo() *Login {
	hostname := GetHostName()
	if hostname == "" {
		log.Println("Blank Hostname")
		return nil
	}

	return &Login{
		Hostname:  hostname,
		Username:  os.Getenv("USERNAME"),
		DeviceKey: os.Getenv("KEY"),
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
