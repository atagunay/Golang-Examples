package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Shipment struct {
	Id      int
	Address string
}

func main() {

	shipments := []Shipment{
		{
			Id:      1,
			Address: "istanbul",
		},
		{
			Id:      2,
			Address: "Ankara",
		},
	}

	http.HandleFunc("/shipments", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodGet && request.Method != http.MethodPost {
			writer.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		if request.Method == http.MethodPost {
			var shipment Shipment
			readByte, _ := io.ReadAll(request.Body)
			err := json.Unmarshal(readByte, &shipment)
			if err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
				return
			}

			fmt.Println(shipment)
			writer.WriteHeader(http.StatusCreated)
			return
		}

		shipmentsByte, err := json.Marshal(shipments)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		writer.Header().Set("content-type", "application/json")
		writer.Write(shipmentsByte)
	})

	http.ListenAndServe(":80", nil)
}
