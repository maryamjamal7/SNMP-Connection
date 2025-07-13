package storage

import (
	"database/sql"
	"snmp-connector/internal/domain"
)

type PostgresAdapter struct {
	DB *sql.DB
}

func NewPostgresAdapter(db *sql.DB) *PostgresAdapter {
	return &PostgresAdapter{DB: db}
}

func (p *PostgresAdapter) Save(data domain.SNMPData) error {
	query := `INSERT INTO snmp_data1 (device_ip, oid, value, type, retrieved_at)
              VALUES ($1, $2, $3, $4, $5)`
	_, err := p.DB.Exec(query, data.DeviceIP, data.OID, data.Value, data.Type, data.RetrievedAt)
	return err
}

func (p *PostgresAdapter) GetAll() ([]domain.SNMPData, error) {
	rows, err := p.DB.Query(`SELECT id, device_ip, oid, value, type, retrieved_at FROM snmp_data1 ORDER BY retrieved_at DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []domain.SNMPData
	for rows.Next() {
		var d domain.SNMPData
		err := rows.Scan(&d.ID, &d.DeviceIP, &d.OID, &d.Value, &d.Type, &d.RetrievedAt)
		if err != nil {
			return nil, err
		}
		list = append(list, d)
	}
	return list, nil
}
