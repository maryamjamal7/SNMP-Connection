package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"snmp-connector/internal/adapters/api"
	"snmp-connector/internal/adapters/snmp"
	"snmp-connector/internal/adapters/storage"
	"snmp-connector/internal/domain"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	connStr := "host=localhost port=5432 user=mariam password=123 dbname=mydb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	snmpAdapter := snmp.NewSNMPAdapter("public")

	storageAdapter := storage.NewPostgresAdapter(db)
	service := domain.NewSNMPService(snmpAdapter, storageAdapter)
	handler := api.NewHandler(service)

	r := mux.NewRouter()
	r.HandleFunc("/poll", handler.FetchAndStore).Methods("GET")
	r.HandleFunc("/data", handler.ListData).Methods("GET")

	fmt.Println("🚀 SNMP Connector running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
