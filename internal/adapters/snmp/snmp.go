package snmp

import (
	"fmt"
	"time"

	"snmp-connector/internal/domain"

	"github.com/gosnmp/gosnmp"
)

type SNMPAdapter struct {
	Community string
}

func NewSNMPAdapter(community string) *SNMPAdapter {
	return &SNMPAdapter{Community: community}
}

func (a *SNMPAdapter) GetData(ip string) ([]domain.SNMPData, error) {
	g := &gosnmp.GoSNMP{
		Target:    ip,
		Port:      161,
		Community: a.Community,
		Version:   gosnmp.Version2c,
		Timeout:   2 * time.Second,
		Retries:   1,
	}

	err := g.Connect()
	if err != nil {
		return nil, err
	}
	defer g.Conn.Close()

	oids := []string{".1.3.6.1.2.1.1.5.0", ".1.3.6.1.2.1.1.3.0"} // sysName, sysUpTime
	fmt.Printf("Requesting OIDs from %s: %v\n", ip, oids)

	result, err := g.Get(oids)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Received %d variables\n", len(result.Variables))

	var results []domain.SNMPData
	for _, variable := range result.Variables {
		data := domain.SNMPData{
			DeviceIP:    ip,
			OID:         variable.Name,
			Value:       fmt.Sprintf("%v", variable.Value),
			Type:        variable.Type.String(),
			RetrievedAt: time.Now().Format(time.RFC3339),
		}
		results = append(results, data)
	}

	return results, nil
}
