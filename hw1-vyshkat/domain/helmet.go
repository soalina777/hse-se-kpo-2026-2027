package domain

type Helmet struct {
	InventoryNumber string
	Name            string
}

func (h Helmet) GetInventoryNumber() string {
	return h.InventoryNumber
}

func (h Helmet) GetName() string {
	return h.Name
}
