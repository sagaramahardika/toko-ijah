package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"toko-ijah/api/app/model"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func GetAllOutgoingItem(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	outgoingItem := []model.OutgoingItem{}
	db.Find(&outgoingItem)

	respondJSON(w, http.StatusOK, outgoingItem)
}

func CreateOutgoingItem(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	outgoingItem := model.OutgoingItem{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&outgoingItem); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&outgoingItem).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusCreated, outgoingItem)
}

func GetOutgoingItem(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	item := getOutgoingItemOr404(db, id, w, r)
	if item == nil {
		return
	}

	respondJSON(w, http.StatusOK, item)
}

func UpdateOutgoingItem(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	outgoingItem := getOutgoingItemOr404(db, id, w, r)
	if outgoingItem == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&outgoingItem); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&outgoingItem).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, outgoingItem)
}

func DeleteOutgoingItem(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	outgoingItem := getOutgoingItemOr404(db, id, w, r)
	if outgoingItem == nil {
		return
	}

	if err := db.Delete(&outgoingItem).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusNoContent, nil)
}

// getOutgoingItemOr404 gets a outgoing item instance if exists, or respond the 404 error otherwise
func getOutgoingItemOr404(db *gorm.DB, id int64, w http.ResponseWriter, r *http.Request) *model.OutgoingItem {
	outgoingItem := model.OutgoingItem{}

	if err := db.First(&outgoingItem, model.OutgoingItem{ID: id}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}

	return &outgoingItem
}
