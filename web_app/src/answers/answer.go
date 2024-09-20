package answers

import (
	"encoding/json"
	"log"
	"net/http"
)

type Error struct {
	Erro string `json:"erro"`
}

func JSON (w http.ResponseWriter, statusCode int, dados interface{})  {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	
	if dados != nil{
		if err := json.NewEncoder(w).Encode(dados); err != nil{
			log.Fatal(err)
		}
	}
}

func Erro(w http.ResponseWriter, r http.Response)  {
	var erro Error
	json.NewDecoder(r.Body).Decode(&erro)
	JSON(w, r.StatusCode, erro)
}