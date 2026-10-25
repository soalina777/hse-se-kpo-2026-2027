package domain

type ServiceCenter interface {
	Inspect(vehicle Vehicle) bool
}

type DefaultServiceCenter struct {
}

func (dsc DefaultServiceCenter) Inspect(vehicle Vehicle) bool {
	if vehicle.GetInventoryNumber() == "" {
		return false
	}

	if vehicle.GetName() == "" {
		return false
	}

	if vehicle.GetBeginnerRating() < 1 || vehicle.GetBeginnerRating() > 10 {
		return false
	}

	return true
}
