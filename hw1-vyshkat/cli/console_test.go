package cli

import (
	"bufio"
	"strings"
	"testing"

	"vyshkat/domain"
	"vyshkat/service"
)

func TestReadString(t *testing.T) {
	console := &Console{
		reader: bufio.NewReader(strings.NewReader("  Scooter  \n")),
	}

	result, err := console.readString("Введите название: ")

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if result != "Scooter" {
		t.Errorf("Expected Scooter, got %s", result)
	}
}

func TestReadInt(t *testing.T) {
	console := &Console{
		reader: bufio.NewReader(strings.NewReader(" 42 \n")),
	}

	result, err := console.readInt("Введите число: ")

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if result != 42 {
		t.Errorf("Expected 42, got %d", result)
	}
}

func TestReadFloat(t *testing.T) {
	console := &Console{
		reader: bufio.NewReader(strings.NewReader(" 10.5 \n")),
	}

	result, err := console.readFloat("Введите энергопотребление: ")

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if result != 10.5 {
		t.Errorf("Expected 10.5, got %f", result)
	}
}

func TestReadIntInvalidInput(t *testing.T) {
	console := &Console{
		reader: bufio.NewReader(strings.NewReader("abc\n")),
	}

	_, err := console.readInt("Введите число: ")

	if err == nil {
		t.Error("Expected error for invalid input")
	}
}

func TestReadFloatInvalidInput(t *testing.T) {
	console := &Console{
		reader: bufio.NewReader(strings.NewReader("abc\n")),
	}

	_, err := console.readFloat("Введите число: ")

	if err == nil {
		t.Error("Expected error for invalid input")
	}
}

func TestConsoleAddVehicle(t *testing.T) {
	parkService := service.NewParkService(domain.DefaultServiceCenter{})

	input := "1\nSC-001\nNinebot\n3.5\n8\n"

	console := &Console{
		parkService: parkService,
		reader:      bufio.NewReader(strings.NewReader(input)),
	}

	console.AddVehicle()

	vehicles := parkService.GetVehicles()

	if len(vehicles) != 1 {
		t.Fatalf("Expected 1 vehicle, got %d", len(vehicles))
	}
}

func TestConsoleAddVehicleReturnBack(t *testing.T) {
	parkService := service.NewParkService(domain.DefaultServiceCenter{})

	input := " 0\n"

	console := &Console{
		parkService: parkService,
		reader:      bufio.NewReader(strings.NewReader(input)),
	}

	console.AddVehicle()

	vehicles := parkService.GetVehicles()

	if len(vehicles) != 0 {
		t.Errorf("Expected 0 vehicles, got %d", len(vehicles))
	}
}

func TestConsoleAddScooter(t *testing.T) {
	parkService := service.NewParkService(domain.DefaultServiceCenter{})

	input := "SC-001\nNinebot\n3.5\n8\n"

	console := &Console{
		parkService: parkService,
		reader:      bufio.NewReader(strings.NewReader(input)),
	}

	console.AddScooter()

	vehicles := parkService.GetVehicles()

	if len(vehicles) != 1 {
		t.Fatalf("Expected 1 vehicle, got %d", len(vehicles))
	}

	scooter, ok := vehicles[0].(domain.Scooter)
	if !ok {
		t.Fatal("Expected Scooter")
	}

	if scooter.InventoryNumber != "SC-001" {
		t.Errorf("Expected SC-001, got %s", scooter.InventoryNumber)
	}

	if scooter.Name != "Ninebot" {
		t.Errorf("Expected Ninebot, got %s", scooter.Name)
	}

	if scooter.EnergyConsumption != 3.5 {
		t.Errorf("Expected energy 3.5, got %f", scooter.EnergyConsumption)
	}

	if scooter.BeginnerRating != 8 {
		t.Errorf("Expected rating 8, got %d", scooter.BeginnerRating)
	}
}

func TestConsoleAddInventoryItemHelmet(t *testing.T) {
	parkService := service.NewParkService(domain.DefaultServiceCenter{})

	input := "1\nH-001\nHelmet Pro\n"

	console := &Console{
		parkService: parkService,
		reader:      bufio.NewReader(strings.NewReader(input)),
	}

	console.AddInventoryItem()

	items := parkService.GetInventoryItems()

	if len(items) != 1 {
		t.Fatalf("Expected 1 item, got %d", len(items))
	}

	helmet, ok := items[0].(domain.Helmet)
	if !ok {
		t.Fatal("Expected Helmet")
	}

	if helmet.InventoryNumber != "H-001" {
		t.Errorf("Expected H-001, got %s", helmet.InventoryNumber)
	}

	if helmet.Name != "Helmet Pro" {
		t.Errorf("Expected Helmet Pro, got %s", helmet.Name)
	}
}

func TestConsoleAddInventoryItemChargingStation(t *testing.T) {
	parkService := service.NewParkService(domain.DefaultServiceCenter{})

	input := "2\n CS-001\n Charging Station\n 5.5\n"

	console := &Console{
		parkService: parkService,
		reader:      bufio.NewReader(strings.NewReader(input)),
	}

	console.AddInventoryItem()

	items := parkService.GetInventoryItems()

	if len(items) != 1 {
		t.Fatalf("Expected 1 item, got %d", len(items))
	}

	chargingStation, ok := items[0].(domain.ChargingStation)

	if !ok {
		t.Fatal("Expected Charging Station")
	}

	if chargingStation.InventoryNumber != "CS-001" {
		t.Errorf("Expected CS-001, got %s", chargingStation.InventoryNumber)
	}

	if chargingStation.Name != "Charging Station" {
		t.Errorf("Expected Charging Station, got %s", chargingStation.Name)
	}

	if chargingStation.EnergyConsumption != 5.5 {
		t.Errorf("Expected 5.5, got %f", chargingStation.EnergyConsumption)
	}
}

func TestConsoleAddInventoryItemReturnBack(t *testing.T) {
	parkService := service.NewParkService(domain.DefaultServiceCenter{})

	input := " 0\n"

	console := &Console{
		parkService: parkService,
		reader:      bufio.NewReader(strings.NewReader(input)),
	}

	console.AddInventoryItem()

	items := parkService.GetInventoryItems()

	if len(items) != 0 {
		t.Fatalf("Expected 0 items, got %d", len(items))
	}
}

func TestRunAndMainOperations(t *testing.T) {
	parkService := service.NewParkService(domain.DefaultServiceCenter{})

	input := "1\n1\nSC-001\nXiaomi\n3.5\n8\n3\n4\n5\n6\n0\n"

	console := &Console{
		parkService: parkService,
		reader:      bufio.NewReader(strings.NewReader(input)),
	}

	console.Run()

	vehicles := parkService.GetVehicles()

	if len(vehicles) != 1 {
		t.Fatalf("Expected 1 vehicle, got %d", len(vehicles))
	}

	if parkService.GetTotalEnergyConsumption() != 3.5 {
		t.Errorf("Expected energy consumption 3.5")
	}

	beginnerDevices := parkService.GetBeginnerDevices()

	if len(beginnerDevices) != 1 {
		t.Errorf("Expected 1 beginner device, got %d", len(beginnerDevices))
	}
}
