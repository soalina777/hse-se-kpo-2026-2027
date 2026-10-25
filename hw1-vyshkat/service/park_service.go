package service

import (
	"vyshkat/domain"
)

type ParkService struct {
	vehicles       []domain.Vehicle
	inventoryItems []domain.InventoryItem
	serviceCenter  domain.ServiceCenter
}

func NewParkService(serviceCenter domain.ServiceCenter) *ParkService {
	return &ParkService{serviceCenter: serviceCenter}
}

func (p *ParkService) HasInventoryNumber(number string) bool {
	for _, v := range p.inventoryItems {
		if v.GetInventoryNumber() == number {
			return true
		}
	}

	return false
}

func (p *ParkService) AddVehicle(vehicle domain.Vehicle) bool {
	if vehicle.GetInventoryNumber() == "" ||
		p.HasInventoryNumber(vehicle.GetInventoryNumber()) {
		return false
	}

	consumer, ok := vehicle.(domain.EnergyConsumer)
	if ok {
		if consumer.GetEnergyConsumption() < 0 {
			return false
		}
	}

	if p.serviceCenter.Inspect(vehicle) {
		p.vehicles = append(p.vehicles, vehicle)
		p.inventoryItems = append(p.inventoryItems, vehicle)
		return true
	}

	return false
}

func (p *ParkService) AddItem(item domain.InventoryItem) bool {
	_, ok := item.(domain.Vehicle)

	if ok {
		return false
	}

	if item.GetInventoryNumber() == "" ||
		p.HasInventoryNumber(item.GetInventoryNumber()) {
		return false
	}

	consumer, ok := item.(domain.EnergyConsumer)
	if ok {
		if consumer.GetEnergyConsumption() < 0 {
			return false
		}
	}

	p.inventoryItems = append(p.inventoryItems, item)
	return true
}

func (p *ParkService) GetVehicles() []domain.Vehicle {
	return append([]domain.Vehicle(nil), p.vehicles...)
}

func (p *ParkService) GetInventoryItems() []domain.InventoryItem {
	return append([]domain.InventoryItem(nil), p.inventoryItems...)
}

func (p *ParkService) GetTotalEnergyConsumption() float64 {
	var sum float64

	for _, v := range p.inventoryItems {
		// проверим реализует ли конкретный объект интерфейс EnergyConsumer
		consumer, ok := v.(domain.EnergyConsumer)

		if ok {
			sum += consumer.GetEnergyConsumption()
		}
	}

	return sum
}

func (p *ParkService) GetBeginnerDevices() []domain.InventoryItem {
	var beginnerDevices []domain.InventoryItem

	for _, v := range p.inventoryItems {
		beginnerDevice, ok := v.(domain.BeginnerFriendly)

		if ok && beginnerDevice.GetBeginnerRating() >= 6 {
			beginnerDevices = append(beginnerDevices, v)
		}
	}

	return beginnerDevices
}
