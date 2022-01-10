package routes

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
)

func Setup() *mux.Router {
	spa := spaHandler{staticPath: "ui/build", indexPath: "index.html"}

	r := mux.NewRouter()
	r.HandleFunc("/api/register/", Register).Methods(http.MethodPost)
	r.HandleFunc("/api/login/", Login).Methods(http.MethodPost)
	r.HandleFunc("/api/logout/", Logout).Methods(http.MethodPost)
	r.HandleFunc("/api/user/", GetUser).Methods(http.MethodGet)

	r.HandleFunc("/api/items/", CreateItem).Methods(http.MethodPost)
	r.HandleFunc("/api/items/", GetItems).Methods(http.MethodGet)
	r.PathPrefix("/").Handler(spa)

	return r
}



type spaHandler struct {
	staticPath string
	indexPath  string
}

func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	path = filepath.Join(h.staticPath, path)

	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}