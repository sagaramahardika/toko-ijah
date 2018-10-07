package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"toko-ijah/api/app/model"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func GetAllIncomingItem(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	incomingItems := []model.IncomingItem{}
	db.Find(&incomingItems)
	respondJSON(w, http.StatusOK, incomingItems)
}

func CreateIncomingItem(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	incomingItem := model.IncomingItem{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&incomingItem); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&incomingItem).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, incomingItem)
}

func GetIncomingItem(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err == nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	item := getIncomingItemOr404(db, id, w, r)
	if item == nil {
		return
	}

	respondJSON(w, http.StatusOK, item)
}

func UpdateIncomingItem(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err == nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	incomingItem := getIncomingItemOr404(db, id, w, r)
	if incomingItem == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&incomingItem); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&incomingItem).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, incomingItem)
}

func DeleteIncomingItem(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err == nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	incomingItem := getIncomingItemOr404(db, id, w, r)
	if incomingItem == nil {
		return
	}

	if err := db.Delete(&incomingItem).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusNoContent, nil)
}

// getIncomingItemOr404 gets a incoming item instance if exists, or respond the 404 error otherwise
func getIncomingItemOr404(db *gorm.DB, id int64, w http.ResponseWriter, r *http.Request) *model.IncomingItem {
	incomingItem := model.IncomingItem{}

	if err := db.First(&incomingItem, model.IncomingItem{ID: id}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}

	return &incomingItem
}
