package main

import (
	"encoding/json"
	"go-cpe/config"
	"go-cpe/routeros"
	"log"
	"net/http"
)

func main() {
	username, password := config.LoadEnv()
	manager := routeros.NewConnectionManager(username, password)

	http.HandleFunc("/connect", func(w http.ResponseWriter, r *http.Request) {

		ipAddress := r.URL.Query().Get("ip")
		if ipAddress == "" {
			respondWithJSON(w, http.StatusBadRequest, routeros.Response{
				Status:  "error",
				Message: "'ip' required",
			})
			return
		}

		address := ipAddress + ":8728"

		client, err := manager.GetConnection(address)
		if err != nil {
			respondWithJSON(w, http.StatusInternalServerError, routeros.Response{
				Status:  "error",
				Message: "Failed to connect to RouterOS: " + err.Error(),
			})
			return
		}

		data, err := routeros.FetchRouterData(client)
		if err != nil {
			respondWithJSON(w, http.StatusInternalServerError, routeros.Response{
				Status:  "error",
				Message: "Error retrieving data: " + err.Error(),
			})
			return
		}

		respondWithJSON(w, http.StatusOK, routeros.Response{
			Status:  "success",
			Message: "Data received successfully",
			Data:    data,
		})
	})

	log.Println("The server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func respondWithJSON(w http.ResponseWriter, status int, response routeros.Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error sending JSON response: %v", err)
	}
}
