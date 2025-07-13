package storage

import (
	"database/sql"
	"os"
	"testing"

	"snmp-connector/internal/domain"

	_ "github.com/lib/pq"
)

func setupTestDB(t *testing.T) *sql.DB {
	connStr := os.Getenv("TEST_DB_CONN")
	if connStr == "" {
		connStr = "host=localhost port=5432 user=mariam password=123 dbname=mydb sslmode=disable"
	}
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		t.Fatalf("Failed to connect to test DB: %v", err)
	}
	return db
}
func TestSave_InvalidData(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	adapter := NewPostgresAdapter(db)

	invalidData := domain.SNMPData{
		DeviceIP:    "", // This is NOT NULL in DB — should fail
		OID:         ".1.3.6.1.2.1.1.1.0",
		Value:       "Some value",
		Type:        "OctetString",
		RetrievedAt: "2024-01-01T10:00:00Z",
	}

	err := adapter.Save(invalidData)
	if err == nil {
		t.Error("Expected error due to missing DeviceIP, but got none")
	}
}
