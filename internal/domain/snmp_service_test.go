package domain

import (
	"errors"
	"testing"
)

// ---- Mock SNMP Adapter that returns error ----
type MockSNMPAdapter struct{}

func (m *MockSNMPAdapter) GetData(ip string) ([]SNMPData, error) {
	return nil, errors.New("failed to connect to device")
}

// ---- Mock Storage (we won't call it in this test, but required) ----
type MockStorage struct{}

func (m *MockStorage) Save(data SNMPData) error {
	return nil
}

func (m *MockStorage) GetAll() ([]SNMPData, error) {
	return nil, nil
}
func TestFetchAndStoreWithInvalidIP(t *testing.T) {
	snmp := &MockSNMPAdapter{}
	storage := &MockStorage{}
	service := NewSNMPService(snmp, storage)

	err := service.FetchAndStore("192.0.2.1") // Invalid IP

	if err == nil {
		t.Error("Expected error for invalid IP, got nil")
	}
}
