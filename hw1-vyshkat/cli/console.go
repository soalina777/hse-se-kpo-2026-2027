package cli

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"vyshkat/domain"
	"vyshkat/service"
)

type Console struct {
	parkService *service.ParkService
	reader      *bufio.Reader
}

func NewConsole(parkService *service.ParkService) *Console {
	return &Console{
		parkService: parkService,
		reader:      bufio.NewReader(os.Stdin),
	}
}

func (c *Console) readString(text string) (string, error) {
	fmt.Print(text)
	line, err := c.reader.ReadString('\n')
	if err != nil && err != io.EOF {
		return "", err
	}
	if err == io.EOF && line == "" {
		return "", io.EOF
	}
	return strings.TrimSpace(line), nil
}

func (c *Console) readInt(text string) (int, error) {
	s, err := c.readString(text)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(s)
}

func (c *Console) readFloat(text string) (float64, error) {
	s, err := c.readString(text)
	if err != nil {
		return 0, err
	}
	return strconv.ParseFloat(s, 64)
}

func (c *Console) ShowOptions() {
	fmt.Println("\n===== ВышКат =====")
	fmt.Println("1. Добавить транспорт")
	fmt.Println("2. Добавить инвентарный объект")
	fmt.Println("3. Показать весь инвентарь")
	fmt.Println("4. Показать количество транспорта")
	fmt.Println("5. Показать энергопотребление")
	fmt.Println("6. Показать устройства для новичков")
	fmt.Println("0. Выход")
}

func (c *Console) Run() {
	for {
		c.ShowOptions()

		option, err := c.readInt("Ваш выбор: ")
		if err != nil {
			if err == io.EOF {
				return
			}
			fmt.Println("Введите целое число")
			continue
		}

		switch option {
		case 1:
			c.AddVehicle()
		case 2:
			c.AddInventoryItem()
		case 3:
			c.ShowInventory()
		case 4:
			c.ShowVehicleCount()
		case 5:
			c.ShowEnergyConsumption()
		case 6:
			c.ShowBeginnerDevices()
		case 0:
			fmt.Println("До свидания!")
			return
		default:
			fmt.Println("Неизвестная команда")
		}
	}
}

func (c *Console) AddVehicle() {
	fmt.Println("\nВыберите тип транспорта:")
	fmt.Println("1. Электросамокат")
	fmt.Println("0. Назад")

	choice, err := c.readInt("Ваш выбор: ")
	if err != nil {
		fmt.Println("Ошибка ввода")
		return
	}

	switch choice {
	case 1:
		c.AddScooter()
	case 0:
		return
	default:
		fmt.Println("Неизвестная команда")
	}
}

func (c *Console) AddScooter() {
	number, err := c.readString("Инвентарный номер: ")
	if err != nil {
		fmt.Println("Ошибка ввода")
		return
	}

	name, err := c.readString("Название: ")
	if err != nil {
		fmt.Println("Ошибка ввода")
		return
	}

	energy, err := c.readFloat("Энергопотребление за сутки: ")
	if err != nil {
		fmt.Println("Введите число")
		return
	}

	rating, err := c.readInt("Рейтинг для новичков (1-10): ")
	if err != nil {
		fmt.Println("Введите целое число")
		return
	}

	scooter := domain.Scooter{
		InventoryNumber:   number,
		Name:              name,
		EnergyConsumption: energy,
		BeginnerRating:    rating,
	}

	if c.parkService.AddVehicle(scooter) {
		fmt.Println("Транспорт успешно добавлен!")
	} else {
		fmt.Println("Транспорт не прошёл проверку или номер уже занят")
	}
}

func (c *Console) AddInventoryItem() {
	fmt.Println("\nВыберите тип объекта:")
	fmt.Println("1. Шлем")
	fmt.Println("2. Зарядная станция")
	fmt.Println("0. Назад")

	choice, err := c.readInt("Ваш выбор: ")
	if err != nil {
		fmt.Println("Ошибка ввода")
		return
	}

	switch choice {
	case 1:
		c.AddHelmet()
	case 2:
		c.AddChargingStation()
	case 0:
		return
	default:
		fmt.Println("Неизвестная команда")
	}
}

func (c *Console) AddHelmet() {
	number, err := c.readString("Инвентарный номер: ")
	if err != nil {
		fmt.Println("Ошибка ввода")
		return
	}

	name, err := c.readString("Название: ")
	if err != nil {
		fmt.Println("Ошибка ввода")
		return
	}

	helmet := domain.Helmet{
		InventoryNumber: number,
		Name:            name,
	}

	if c.parkService.AddItem(helmet) {
		fmt.Println("Шлем успешно добавлен!")
	} else {
		fmt.Println("Не удалось добавить шлем")
	}
}

func (c *Console) AddChargingStation() {
	number, err := c.readString("Инвентарный номер: ")
	if err != nil {
		fmt.Println("Ошибка ввода")
		return
	}

	name, err := c.readString("Название: ")
	if err != nil {
		fmt.Println("Ошибка ввода")
		return
	}

	energy, err := c.readFloat("Энергопотребление за сутки: ")
	if err != nil {
		fmt.Println("Введите число")
		return
	}

	station := domain.ChargingStation{
		InventoryNumber:   number,
		Name:              name,
		EnergyConsumption: energy,
	}

	if c.parkService.AddItem(station) {
		fmt.Println("Зарядная станция успешно добавлена!")
	} else {
		fmt.Println("Не удалось добавить зарядную станцию")
	}
}

func (c *Console) ShowInventory() {
	items := c.parkService.GetInventoryItems()

	fmt.Println("\n===== Весь инвентарь =====")
	if len(items) == 0 {
		fmt.Println("Инвентарь пуст")
		return
	}

	for _, item := range items {
		fmt.Printf("Номер: %s, Название: %s\n",
			item.GetInventoryNumber(), item.GetName())
	}
}

func (c *Console) ShowVehicleCount() {
	vehicles := c.parkService.GetVehicles()
	fmt.Printf("Количество транспорта: %d\n", len(vehicles))
}

func (c *Console) ShowEnergyConsumption() {
	total := c.parkService.GetTotalEnergyConsumption()
	fmt.Printf("Суммарное энергопотребление за сутки: %.2f\n", total)
}

func (c *Console) ShowBeginnerDevices() {
	devices := c.parkService.GetBeginnerDevices()

	fmt.Println("\n===== Устройства для новичков =====")
	if len(devices) == 0 {
		fmt.Println("Подходящих устройств нет")
		return
	}

	for _, device := range devices {
		fmt.Printf("Номер: %s | Название: %s\n",
			device.GetInventoryNumber(), device.GetName())
	}
}
