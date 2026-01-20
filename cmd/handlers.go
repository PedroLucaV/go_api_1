package main

import (
	"encoding/json"
	"net/http"
)

func handleClientProfile(w http.ResponseWriter, r *http.Request) {
	GetClientProfile(w, r)
}

func GetClientProfile(w http.ResponseWriter, r *http.Request) {
	clientId := r.PathValue("clientId")
	if clientId == "" {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	clientProfile, ok := database[clientId]
	if !ok {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	response := ClientProfile{
		Email: clientProfile.Email,
		Name:  clientProfile.Name,
		Id:    clientProfile.Id,
	}

	json.NewEncoder(w).Encode(response)
}
