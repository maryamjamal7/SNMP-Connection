package domain

type SNMPPort interface {
	GetData(ip string) ([]SNMPData, error)
}

type StoragePort interface {
	Save(data SNMPData) error
	GetAll() ([]SNMPData, error)
}
