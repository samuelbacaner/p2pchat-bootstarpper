package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/patrickmn/go-cache"
)

type AddressRequest struct {
	Address string `json:"address"`
}

type AddressResponse struct {
	Addresses []string `json:"addresses"`
}

type ErrorMsg struct {
	Error string `json:"error"`
}

func main() {

	c := cache.New(2*time.Minute, 1*time.Minute)

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != http.MethodPost {
			http.Error(w, "only POST allowed", http.StatusMethodNotAllowed)
			return
		}

		var req AddressRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, fmt.Sprintf("unable to parse request body: %s ", err.Error()), http.StatusBadRequest)
			return
		}

		c.Set(req.Address, "nil", 2*time.Minute)

		addresses := make([]string, 0, len(c.Items()))
		for k := range c.Items() {
			addresses = append(addresses, k)
		}

		resp := AddressResponse{
			Addresses: addresses,
		}
		json.NewEncoder(w).Encode(&resp)
	})

	http.ListenAndServe(":8080", mux)
}
