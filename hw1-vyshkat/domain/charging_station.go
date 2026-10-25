package domain

type ChargingStation struct {
	InventoryNumber   string
	Name              string
	EnergyConsumption float64
}

func (cs ChargingStation) GetInventoryNumber() string {
	return cs.InventoryNumber
}

func (cs ChargingStation) GetName() string {
	return cs.Name
}

func (cs ChargingStation) GetEnergyConsumption() float64 {
	return cs.EnergyConsumption
}
