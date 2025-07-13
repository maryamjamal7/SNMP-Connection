package domain

type SNMPService struct {
	snmpClient  SNMPPort
	storageRepo StoragePort
}

func NewSNMPService(snmp SNMPPort, storage StoragePort) *SNMPService {
	return &SNMPService{snmpClient: snmp, storageRepo: storage}
}

func (s *SNMPService) FetchAndStore(ip string) error {
	dataList, err := s.snmpClient.GetData(ip)
	if err != nil {
		return err
	}

	for _, data := range dataList {
		err := s.storageRepo.Save(data)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *SNMPService) ListData() ([]SNMPData, error) {
	return s.storageRepo.GetAll()
}
