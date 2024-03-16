package models

import "gorm.io/gorm"

type GeoData struct {
	gorm.Model
	UserID   uint   `json:"user_id"`
	FileName string `json:"file_name"`
	FileType string `json:"file_type"`
	FilePath string `json:"file_path"`
}
