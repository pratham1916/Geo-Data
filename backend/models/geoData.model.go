package models

import "gorm.io/gorm"

type GeoData struct {
	gorm.Model
	UserID    uint    `json:"userId"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
