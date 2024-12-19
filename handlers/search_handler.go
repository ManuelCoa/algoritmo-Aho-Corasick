package handlers

import (
	"aho-corasick-service/algoritmo"
	"encoding/json"
	"net/http"
)

type SearchRequest struct {
	Text     string   `json:"text"`     // Texto donde buscar
	Patterns []string `json:"patterns"` // Patrones a buscar
}

type SearchResponse struct {
	Matches map[string][]int `json:"matches"` // Mapa con índices de ocurrencias para cada patrón
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	var req SearchRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	matches := algoritmo.AhoCorasickSearch(req.Text, req.Patterns)
	response := SearchResponse{Matches: matches}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
