package domain

import (
	"testing"
)

func TestDefaultServiceCenterInspect(t *testing.T) {
	var dsc DefaultServiceCenter

	scooter := Scooter{
		InventoryNumber:   "SC-001",
		Name:              "Suzuki",
		EnergyConsumption: 3.5,
		BeginnerRating:    6,
	}

	if !dsc.Inspect(scooter) {
		t.Error("Scooter must be accepted")
	}
}

func TestDefaultServiceCenterRejectsEmptyNumber(t *testing.T) {
	var dsc DefaultServiceCenter

	scooter := Scooter{
		InventoryNumber:   "",
		Name:              "Suzuki",
		EnergyConsumption: 3.5,
		BeginnerRating:    6,
	}

	if dsc.Inspect(scooter) {
		t.Error("Scooter must be rejected")
	}
}

func TestDefaultServiceCenterRejectsEmptyName(t *testing.T) {
	var dsc DefaultServiceCenter

	scooter := Scooter{
		InventoryNumber:   "SC-001",
		Name:              "",
		EnergyConsumption: 3.5,
		BeginnerRating:    6,
	}

	if dsc.Inspect(scooter) {
		t.Error("Scooter must be rejected")
	}
}

func TestDefaultServiceCenterRejectsLowRating(t *testing.T) {
	var dsc DefaultServiceCenter

	scooter := Scooter{
		InventoryNumber:   "SC-001",
		Name:              "Suzuki",
		EnergyConsumption: 3.5,
		BeginnerRating:    0,
	}

	if dsc.Inspect(scooter) {
		t.Error("Scooter with rating 0 must be rejected")
	}
}

func TestDefaultServiceCenterRejectsHighRating(t *testing.T) {
	var dsc DefaultServiceCenter

	scooter := Scooter{
		InventoryNumber:   "SC-001",
		Name:              "Suzuki",
		EnergyConsumption: 3.5,
		BeginnerRating:    11,
	}

	if dsc.Inspect(scooter) {
		t.Error("Scooter with rating 11 must be rejected")
	}
}

func TestDefaultServiceCenterAcceptsBoundaryRatings(t *testing.T) {
	var dsc DefaultServiceCenter

	scooter1 := Scooter{
		InventoryNumber:   "SC-001",
		Name:              "Suzuki",
		EnergyConsumption: 3.5,
		BeginnerRating:    10,
	}

	if !dsc.Inspect(scooter1) {
		t.Error("Scooter must be accepted")
	}

	scooter2 := Scooter{
		InventoryNumber:   "SC-002",
		Name:              "Suzuki",
		EnergyConsumption: 3.5,
		BeginnerRating:    1,
	}

	if !dsc.Inspect(scooter2) {
		t.Error("Scooter must be accepted")
	}
}
