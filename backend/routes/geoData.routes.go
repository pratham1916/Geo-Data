package routes

import (
	"encoding/json"
	"geo-data/models"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	"strings"
)

func CreateGeoData(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            http.Error(w, "Unsupported HTTP method", http.StatusMethodNotAllowed)
            return
        }
        if err := r.ParseMultipartForm(10 << 20); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        userID := r.FormValue("user_id")
        file, _, err := r.FormFile("file_path")
        if err != nil {
            http.Error(w, "File upload error: "+err.Error(), http.StatusInternalServerError)
            return
        }
        defer file.Close()

        fileBytes, err := ioutil.ReadAll(file)
        if err != nil {
            http.Error(w, "Error reading file: "+err.Error(), http.StatusInternalServerError)
            return
        }

        var geoJSON map[string]interface{}
        if err := json.Unmarshal(fileBytes, &geoJSON); err != nil {
            http.Error(w, "Error parsing JSON: "+err.Error(), http.StatusInternalServerError)
            return
        }

        geometry, err := json.Marshal(geoJSON["features"].([]interface{})[0].(map[string]interface{})["geometry"])
        if err != nil {
            http.Error(w, "Error processing geometry: "+err.Error(), http.StatusInternalServerError)
            return
        }

        geoData := models.GeoData{
            UserID:   userID,
            Geometry: string(geometry),
        }

        if err := db.Create(&geoData).Error; err != nil {
            http.Error(w, "Error saving geometry to database: "+err.Error(), http.StatusInternalServerError)
            return
        }

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]interface{}{
            "success": true,
            "message": "Geometry saved successfully",
            "user_id": userID,
        })
    }
}

func ListGeoData(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Unsupported HTTP method", http.StatusMethodNotAllowed)
			return
		}

		var geodata []models.GeoData 
		db.Find(&geodata)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(geodata)
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
