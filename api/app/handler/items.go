package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"toko-ijah/api/app/model"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func GetAllItem(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	items := []model.Item{}
	db.Find(&items)
	respondJSON(w, http.StatusOK, items)
}

func CreateItem(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	item := model.Item{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&item); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	item.CreatedAt = time.Now().Unix()
	item.UpdatedAt = time.Now().Unix()

	if err := db.Save(&item).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, item)
}

func GetItem(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	sku := vars["sku"]
	item := getItemOr404(db, sku, w, r)
	if item == nil {
		return
	}
	respondJSON(w, http.StatusOK, item)
}

func UpdateItem(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	sku := vars["sku"]
	item := getItemOr404(db, sku, w, r)
	if item == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&item); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	item.UpdatedAt = time.Now().Unix()

	if err := db.Save(&item).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, item)
}

func DeleteItem(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	sku := vars["sku"]
	item := getItemOr404(db, sku, w, r)
	if item == nil {
		return
	}
	if err := db.Delete(&item).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

// getEmployeeOr404 gets a employee instance if exists, or respond the 404 error otherwise
func getItemOr404(db *gorm.DB, sku string, w http.ResponseWriter, r *http.Request) *model.Item {
	item := model.Item{}
	if err := db.First(&item, model.Item{SKU: sku}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &item
}
