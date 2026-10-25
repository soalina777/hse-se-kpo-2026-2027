package service

import (
	"testing"
	"vyshkat/domain"
)

func TestParkServiceAddValidVehicle(t *testing.T) {
	var dsc domain.DefaultServiceCenter

	parkService := NewParkService(dsc)

	scooter := domain.Scooter{
		InventoryNumber:   "SC-001",
		Name:              "Suzuki",
		EnergyConsumption: 3.5,
		BeginnerRating:    7,
	}

	if !parkService.AddVehicle(scooter) {
		t.Error("Valid scooter must be accepted")
	}

	vehicles := parkService.GetVehicles()

	if len(vehicles) != 1 {
		t.Errorf("Expected 1 vehicle, got %d", len(vehicles))
	}
}

func TestParkServiceRejectsInvalidVehicle(t *testing.T) {
	var dsc domain.DefaultServiceCenter

	parkService := NewParkService(dsc)

	scooter := domain.Scooter{
		InventoryNumber:   "SC-001",
		Name:              "",
		EnergyConsumption: 3.5,
		BeginnerRating:    7,
	}

	if parkService.AddVehicle(scooter) {
		t.Error("Invalid scooter must be rejected")
	}

	vehicles := parkService.GetVehicles()

	if len(vehicles) != 0 {
		t.Errorf("Expected 0 vehicle, got %d", len(vehicles))
	}
}

type FakeServiceCenter struct {
	accept bool
}

func (m FakeServiceCenter) Inspect(vehicle domain.Vehicle) bool {
	return m.accept
}

func TestParkServiceRejectsVehicleWithFakeServiceCenter(t *testing.T) {
	fake := FakeServiceCenter{
		accept: false,
	}

	parkService := NewParkService(fake)

	scooter := domain.Scooter{
		InventoryNumber:   "SC-001",
		Name:              "Suzuki",
		EnergyConsumption: 3.5,
		BeginnerRating:    7,
	}

	if parkService.AddVehicle(scooter) {
		t.Error("Every vehicle must be rejected ")
	}

	vehicles := parkService.GetVehicles()

	if len(vehicles) != 0 {
		t.Errorf("Expected 0 vehicle, got %d", len(vehicles))
	}
}

func TestParkServiceAcceptsVehicleWithFakeServiceCenter(t *testing.T) {
	fake := FakeServiceCenter{
		accept: true,
	}

	parkService := NewParkService(fake)

	scooter := domain.Scooter{
		InventoryNumber:   "SC-001",
		Name:              "Suzuki",
		EnergyConsumption: 3.5,
		BeginnerRating:    100,
	}

	if !parkService.AddVehicle(scooter) {
		t.Error("Every Vehicle must be accepted")
	}

	vehicles := parkService.GetVehicles()

	if len(vehicles) != 1 {
		t.Errorf("Expected 1 vehicle, got %d", len(vehicles))
	}
}

func TestParkServiceRejectsDuplicateNumberForVehicle(t *testing.T) {
	fake := FakeServiceCenter{
		accept: true,
	}

	parkService := NewParkService(fake)

	scooter1 := domain.Scooter{
		InventoryNumber:   "SC-001",
		Name:              "Suzuki",
		EnergyConsumption: 3.5,
		BeginnerRating:    7,
	}

	if !parkService.AddVehicle(scooter1) {
		t.Error("Valid scooter must be accepted")
	}

	scooter2 := domain.Scooter{
		InventoryNumber:   "SC-001",
		Name:              "Arai",
		EnergyConsumption: 3.5,
		BeginnerRating:    6,
	}

	if parkService.AddVehicle(scooter2) {
		t.Error("Vehicle with duplicate inventory number must be rejected")
	}

	vehicles := parkService.GetVehicles()

	if len(vehicles) != 1 {
		t.Errorf("Expected 1 vehicle, got %d", len(vehicles))
	}
}

func TestParkServiceRejectsDuplicateNumberForItem(t *testing.T) {
	fake := FakeServiceCenter{
		accept: true,
	}

	parkService := NewParkService(fake)

	scooter1 := domain.Scooter{
		InventoryNumber:   "SC-001",
		Name:              "Suzuki",
		EnergyConsumption: 3.5,
		BeginnerRating:    7,
	}

	if !parkService.AddVehicle(scooter1) {
		t.Error("Valid scooter must be accepted")
	}

	helmet1 := domain.Helmet{
		InventoryNumber: "SC-001",
		Name:            "Shark",
	}

	if parkService.AddItem(helmet1) {
		t.Error("Item with duplicate inventory number must be rejected")
	}

	inventoryItems := parkService.GetInventoryItems()

	if len(inventoryItems) != 1 {
		t.Errorf("Expected 1 item, got %d", len(inventoryItems))
	}
}

func TestParkServiceTotalEnergyConsumption(t *testing.T) {
	fake := FakeServiceCenter{
		accept: true,
	}

	parkService := NewParkService(fake)

	scooter1 := domain.Scooter{
		InventoryNumber:   "SC-001",
		EnergyConsumption: 3.5,
	}
	if !parkService.AddVehicle(scooter1) {
		t.Fatal("Failed to add scooter1")
	}

	scooter2 := domain.Scooter{
		InventoryNumber:   "SC-002",
		EnergyConsumption: 2.5,
	}
	if !parkService.AddVehicle(scooter2) {
		t.Fatal("Failed to add scooter2")
	}

	chargingStation := domain.ChargingStation{
		InventoryNumber:   "CS-001",
		EnergyConsumption: 5.0,
	}
	if !parkService.AddItem(chargingStation) {
		t.Fatal("Failed to add charging station")
	}

	helmet := domain.Helmet{
		InventoryNumber: "HL-001",
	}
	if !parkService.AddItem(helmet) {
		t.Fatal("Failed to add helmet")
	}

	total := parkService.GetTotalEnergyConsumption()

	if total != 11.0 {
		t.Errorf("Expected total enerdy consumption : 11.0, but got %f",
			total)
	}
}

func TestParkServiceGetBeginnerDevices(t *testing.T) {
	fake := FakeServiceCenter{
		accept: true,
	}

	parkService := NewParkService(fake)

	scooter1 := domain.Scooter{
		InventoryNumber: "SC-001",
		Name:            "Suzuki",
		BeginnerRating:  7,
	}
	if !parkService.AddVehicle(scooter1) {
		t.Fatal("Scooter must be accepted")
	}

	scooter2 := domain.Scooter{
		InventoryNumber: "SC-002",
		Name:            "Xiaomi",
		BeginnerRating:  3,
	}
	if !parkService.AddVehicle(scooter2) {
		t.Fatal("Scooter must be accepted")
	}

	scooter3 := domain.Scooter{
		InventoryNumber: "SC-003",
		Name:            "Honda",
		BeginnerRating:  10,
	}
	if !parkService.AddVehicle(scooter3) {
		t.Fatal("Scooter must be accepted")
	}

	beginnerDevices := parkService.GetBeginnerDevices()

	if len(beginnerDevices) != 2 {
		t.Fatalf("Expected 2 devices. Got %d devices", len(beginnerDevices))
	}

	if beginnerDevices[0].GetInventoryNumber() != "SC-001" {
		t.Error("Expected Suzuki as first beginner device")
	}

	if beginnerDevices[1].GetInventoryNumber() != "SC-003" {
		t.Error("Expected Honda as second beginner device")
	}
}

func TestParkServiceRejectsVehicleAsItem(t *testing.T) {
	fake := FakeServiceCenter{
		accept: true,
	}

	parkService := NewParkService(fake)

	scooter := domain.Scooter{
		InventoryNumber:   "SC-001",
		Name:              "Suzuki",
		EnergyConsumption: 3.5,
		BeginnerRating:    6,
	}

	if parkService.AddItem(scooter) {
		t.Error("Vehicle must not be accepted with AddItem")
	}

	inventoryItems := parkService.GetInventoryItems()
	if len(inventoryItems) != 0 {
		t.Errorf("Expected 0 items. Got %d items", len(inventoryItems))
	}

	vehicles := parkService.GetVehicles()
	if len(vehicles) != 0 {
		t.Errorf("Expected 0 vehicles. Got %d vehicles", len(vehicles))
	}
}

func TestAddVehicleNegativeEnergy(t *testing.T) {
	fake := FakeServiceCenter{
		accept: true,
	}

	parkService := NewParkService(fake)

	scooter := domain.Scooter{
		InventoryNumber:   "SC-001",
		EnergyConsumption: -3,
	}

	if parkService.AddVehicle(scooter) {
		t.Error("Vehicle with negative energy must not be accepted")
	}

	vehicles := parkService.GetVehicles()
	if len(vehicles) != 0 {
		t.Errorf("Expected 0 vehicles. Got %d vehicles", len(vehicles))
	}
}

func TestAddItemNegativeEnergy(t *testing.T) {
	fake := FakeServiceCenter{
		accept: true,
	}

	parkService := NewParkService(fake)

	chargingStation := domain.ChargingStation{
		InventoryNumber:   "CS-001",
		EnergyConsumption: -5,
	}

	if parkService.AddItem(chargingStation) {
		t.Error("Item with negative energy must not be accepted")
	}

	inventoryItems := parkService.GetInventoryItems()
	if len(inventoryItems) != 0 {
		t.Errorf("Expected 0 items. Got %d items", len(inventoryItems))
	}
}

func TestAddVehicleEmptyInventoryNumber(t *testing.T) {
	park := NewParkService(FakeServiceCenter{accept: true})

	scooter := domain.Scooter{
		InventoryNumber:   "",
		Name:              "Test Scooter",
		EnergyConsumption: 3.5,
		BeginnerRating:    7,
	}

	if park.AddVehicle(scooter) {
		t.Error("Vehicle with empty inventory number must be rejected")
	}
}

func TestAddItemEmptyInventoryNumber(t *testing.T) {
	park := NewParkService(domain.DefaultServiceCenter{})

	helmet := domain.Helmet{
		InventoryNumber: "",
		Name:            "Test Helmet",
	}

	if park.AddItem(helmet) {
		t.Error("Item with empty inventory number must be rejected")
	}
}
