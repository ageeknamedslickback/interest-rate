package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/ageeknamedslickback/interest-rate/pkg/account"
	"github.com/gorilla/mux"
)

// CalculateInterestHandler is a REST handler to calculate the interest given a balance
func CalculateInterestHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	log.Print(vars)
	balance, err := strconv.ParseFloat(vars["balance"], 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	interestAmount, err := account.CalculateInterestApplicable(balance)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"interest": interestAmount})
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/interest/{balance}", CalculateInterestHandler)

	// Small scale `mux` based server
	log.Fatal(http.ListenAndServe(":8000", r))
}
