package handlers

import (
	"log"
	"net/http"
)

func (h *Handler) HelloHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method)
	log.Println(r.Header)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello World!"))
}
