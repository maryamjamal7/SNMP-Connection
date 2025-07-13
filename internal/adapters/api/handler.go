package api

import (
	"encoding/json"
	"net/http"
	"snmp-connector/internal/domain"
)

type Handler struct {
	Service *domain.SNMPService
}

func NewHandler(service *domain.SNMPService) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) FetchAndStore(w http.ResponseWriter, r *http.Request) {
	ip := r.URL.Query().Get("ip")
	if ip == "" {
		http.Error(w, "ip parameter is required", http.StatusBadRequest)
		return
	}

	err := h.Service.FetchAndStore(ip)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("SNMP data fetched and saved."))
}

func (h *Handler) ListData(w http.ResponseWriter, r *http.Request) {
	data, err := h.Service.ListData()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
