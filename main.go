package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Order struct {
	Order_Id   int     `json:"order_id"`
	Code       string  `json:"city"`
	Qtd        int     `json:"quantity"`
	Unit_price float64 `json:"unity_price"`
}

func main() {

	//define routes
	http.HandleFunc("/orders", getOrders)

	//starting servers
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func getOrders(w http.ResponseWriter, r *http.Request) {
	customers := []Order{
		{Order_Id: 1, Code: "ITSA4", Qtd: 500, Unit_price: 2.5},
		{Order_Id: 2, Code: "WEGE3", Qtd: 100, Unit_price: 12.34},
		{Order_Id: 3, Code: "WEGE3", Qtd: 200, Unit_price: 17.5},
	}

	w.Header().Add("Content-Type", "application/json")

	json.NewEncoder(w).Encode(customers)
}
