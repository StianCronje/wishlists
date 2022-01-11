package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Setup() *mux.Router {

	r := mux.NewRouter()
	r.StrictSlash(true)
	r.HandleFunc("/api/register", Register).Methods(http.MethodPost)
	r.HandleFunc("/api/login", Login).Methods(http.MethodPost)
	r.HandleFunc("/api/logout", Logout).Methods(http.MethodPost)
	r.HandleFunc("/api/user", addUserContext(GetUser)).Methods(http.MethodGet)

	r.HandleFunc("/api/items/", addUserContext(CreateItem)).Methods(http.MethodPost)
	r.HandleFunc("/api/items/", addUserContext(GetItems)).Methods(http.MethodGet)

	spa := spaHandler{staticPath: "ui/build", indexPath: "index.html"}
	r.PathPrefix("/").Handler(spa)

	return r
}