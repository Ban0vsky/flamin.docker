package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Ban0vsky/flamin.docker/worker/database"
	"github.com/Ban0vsky/flamin.docker/worker/models"

	"github.com/gorilla/mux"
)

/**
* Returns an specific user
* Required arguments: string ID
 */
func ReadUser(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn
	vars := mux.Vars(r)
	ID := vars["ID"]
	var user models.User
	err := db.Where("id = ?", ID).Find(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

/**
* Updates an user
* Required arguments: string ID, string username
 */
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn
	ID := r.FormValue("ID")
	username := r.FormValue("username")

	var user models.User
	err := db.Where("id = ?", ID).Find(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user.Username = username

	err2 := db.Save(&user)
	if err2 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}
