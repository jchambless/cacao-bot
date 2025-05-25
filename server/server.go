package server

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
)

type spaHandler struct {
	StaticPath string
	IndexPath  string
}

func CreateRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/health", healthCheck).Methods("GET")

	// Register the SPA handler for serving static files
	spa := &spaHandler{StaticPath: "web/build", IndexPath: "index.html"}
	r.PathPrefix("/").Handler(spa)

	return r
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}

func (h *spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join(h.StaticPath, r.URL.Path)

	fi, err := os.Stat(path)
	if os.IsNotExist(err) || fi.IsDir() {
		http.ServeFile(w, r, filepath.Join(h.StaticPath, h.IndexPath))
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.FileServer(http.Dir(h.StaticPath)).ServeHTTP(w, r)
}
