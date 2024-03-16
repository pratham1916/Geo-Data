package routes

import (
	"encoding/json"
	"geo-data/models"
	"gorm.io/gorm"
	"net/http"
	"strings"
)


func CreateOrListGeoData(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			var geoData models.GeoData
			if err := json.NewDecoder(r.Body).Decode(&geoData); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			db.Create(&geoData)
			json.NewEncoder(w).Encode(geoData)
		} else if r.Method == http.MethodGet {
			var geodata []models.GeoData
			db.Find(&geodata)
			json.NewEncoder(w).Encode(geodata)
		} else {
			http.Error(w, "Unsupported HTTP method", http.StatusMethodNotAllowed)
		}
	}
}

func RetrieveUpdateOrDeleteGeoData(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/geodata/")
		switch r.Method {
		case http.MethodGet:
			var geoData models.GeoData
			if result := db.First(&geoData, id); result.Error != nil {
				http.Error(w, "GeoData not found", http.StatusNotFound)
				return
			}
			json.NewEncoder(w).Encode(geoData)
		case http.MethodPut, http.MethodDelete:
			modifyGeoData(w, r, db, id)
		default:
			http.Error(w, "Unsupported HTTP method", http.StatusMethodNotAllowed)
		}
	}
}

func modifyGeoData(w http.ResponseWriter, r *http.Request, db *gorm.DB, id string) {
	switch r.Method {
	case http.MethodPut:
		var updatedGeoData models.GeoData
		if err := json.NewDecoder(r.Body).Decode(&updatedGeoData); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		db.Model(&models.GeoData{}).Where("id = ?", id).Updates(updatedGeoData)
		json.NewEncoder(w).Encode(updatedGeoData)
	case http.MethodDelete:
		db.Delete(&models.GeoData{}, id)
		w.WriteHeader(http.StatusNoContent)
	}
}
