package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", rootPage)
	return r
}

func rootPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Cacao Bot")
}
