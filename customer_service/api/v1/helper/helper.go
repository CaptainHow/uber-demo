package helper

import (
	"encoding/json"
	"net/http"
)

func RespondWithError(w http.ResponseWriter, code int, errorMsg string) {
	RespondWithJSON(w, code, map[string]string{"message": errorMsg})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload any) {
	res, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(res)
}