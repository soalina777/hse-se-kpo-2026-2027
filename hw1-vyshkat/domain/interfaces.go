package domain

type InventoryItem interface {
	GetInventoryNumber() string
	GetName() string
}

type EnergyConsumer interface {
	GetEnergyConsumption() float64
}

type BeginnerFriendly interface {
	GetBeginnerRating() int
}

type Vehicle interface {
	InventoryItem
	BeginnerFriendly
	IsVehicle()
}
