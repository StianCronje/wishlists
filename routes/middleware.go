package routes

import (
	"context"
	"net/http"
	"os"
	"path/filepath"
	"wishlists/helpers"
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

func addUserContext(hf http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("jwt_token")
		if err != nil {
			rw.WriteHeader(http.StatusUnauthorized)
			rw.Write([]byte{})
			return
		}

		user, err := helpers.GetUserFromToken(cookie.Value)
		if err != nil {
			rw.WriteHeader(http.StatusUnauthorized)
			rw.Write([]byte{})
			return
		}
		ctx := context.WithValue(r.Context(), "user", user)
		hf.ServeHTTP(rw, r.WithContext(ctx))
	}
}