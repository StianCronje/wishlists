package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"wishlists/database"
	"wishlists/models"
)

func CreateItem(rw http.ResponseWriter, r *http.Request) {
	var data models.WishItem

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, err.Error())
		return
	}

	data.User = r.Context().Value("user").(models.User)

	database.DB.Create(&data)

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(data)
}

func GetItems(rw http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(models.User)

	var items []models.WishItem

	database.DB.Where("user_id = ?", user.ID).Find(&items)

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(items)
}