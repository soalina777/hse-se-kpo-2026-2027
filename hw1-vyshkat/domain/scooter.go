package domain

type Scooter struct {
	InventoryNumber   string
	Name              string
	EnergyConsumption float64
	BeginnerRating    int
}

func (s Scooter) GetInventoryNumber() string {
	return s.InventoryNumber
}

func (s Scooter) GetName() string {
	return s.Name
}

func (s Scooter) GetEnergyConsumption() float64 {
	return s.EnergyConsumption
}

func (s Scooter) GetBeginnerRating() int {
	return s.BeginnerRating
}

func (s Scooter) IsVehicle() {}
