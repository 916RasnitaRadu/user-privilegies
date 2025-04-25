package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func ViewExperience(w http.ResponseWriter, r *http.Request) {
	SendJSON(w, "Aici aveti xp")
	w.WriteHeader(200)
}

func HandleManageAvailability(w http.ResponseWriter, r *http.Request) {
	SendJSON(w, "Manajing Availability")
	w.WriteHeader(200)
}

func HandleIalamuie(w http.ResponseWriter, r *http.Request) {
	SendJSON(w, "Acuma iei la muie")
	w.WriteHeader(200)
}

func HandleDalamuie(w http.ResponseWriter, r *http.Request) {
	SendJSON(w, "Acuma dai la muie")
	w.WriteHeader(200)
}

func HandleEstiGata(w http.ResponseWriter, r *http.Request) {
	SendJSON(w, "Acum iei la muie da si dai")
	w.WriteHeader(200)
}

func SendJSON(w http.ResponseWriter, data any) {
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Error encoding response: %v", err)
	}
}
