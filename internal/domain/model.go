package domain

type SNMPData struct {
	ID          int
	DeviceIP    string
	OID         string
	Value       string
	Type        string
	RetrievedAt string
}
