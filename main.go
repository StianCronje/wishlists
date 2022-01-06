package main

import (
	"net/http"
	"os"
	"path/filepath"
	"wishlists/routes"

	"github.com/gorilla/mux"
)

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

func main() {
	spa := spaHandler{staticPath: "ui/build", indexPath: "index.html"}

	r := mux.NewRouter()
	r.HandleFunc("/api/items/", routes.GetItems).Methods(http.MethodGet)
	r.PathPrefix("/").Handler(spa)

	http.ListenAndServe(":5000", r)
}