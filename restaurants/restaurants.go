package restaurants

import (
	"fmt"
	"sort"
	"strconv"
)

type Restaurant struct {
	NameRestaurant    string
	allTables         []int
	FreeTables        []int
	ReservedTables    []int
	WaitingForCooking int
	OpenTime          string
	CloseTime         string
	LastBooking       string
	AverageChek       float64
}

var RestaurantKaravella = Restaurant{
	NameRestaurant:    "Каравелла",
	allTables:         []int{4, 4, 4, 4, 4, 4, 3, 3, 2, 2},
	FreeTables:        []int{4, 4, 4, 4, 4, 4, 3, 3, 2, 2},
	ReservedTables:    []int{},
	WaitingForCooking: 30,
	OpenTime:          "9:00 AM",
	CloseTime:         "11:00 PM",
	LastBooking:       "9:00 PM",
	AverageChek:       2000,
}

var RestaurantJuvenility = Restaurant{
	NameRestaurant:    "Молодость",
	allTables:         []int{3, 3, 3},
	FreeTables:        []int{3, 3, 3},
	ReservedTables:    []int{},
	WaitingForCooking: 15,
	OpenTime:          "9:00 AM",
	CloseTime:         "11:00 PM",
	LastBooking:       "9:00 PM",
	AverageChek:       1000,
}

var RestaurantMeatAndSalad = Restaurant{
	NameRestaurant:    "Мясо и салат",
	allTables:         []int{8, 8, 3, 3, 3, 3},
	FreeTables:        []int{8, 8, 3, 3, 3, 3},
	ReservedTables:    []int{},
	WaitingForCooking: 60,
	OpenTime:          "9:00 AM",
	CloseTime:         "11:00 PM",
	LastBooking:       "9:00 PM",
	AverageChek:       1500,
}

var Restaurants = []Restaurant{RestaurantKaravella, RestaurantJuvenility, RestaurantMeatAndSalad}

type ByWaitingForCooking []Restaurant

func (a ByWaitingForCooking) Len() int { return len(a) }
func (a ByWaitingForCooking) Less(i, j int) bool {
	return float64(a[i].WaitingForCooking) < float64(a[j].WaitingForCooking)
}
func (a ByWaitingForCooking) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

type ByAverageCheck []Restaurant

func (a ByAverageCheck) Len() int           { return len(a) }
func (a ByAverageCheck) Less(i, j int) bool { return a[i].AverageChek < a[j].AverageChek }
func (a ByAverageCheck) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func ActualReataurants(s string) []Restaurant {
	p, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("Не корректный ввод")
	}
	var Restaurants01 = []Restaurant{RestaurantJuvenility, RestaurantKaravella, RestaurantMeatAndSalad}

	sort.Sort(ByAverageCheck(Restaurants01))
	sort.Sort(ByWaitingForCooking(Restaurants01))

	var Restaurants02 = []Restaurant{}
	for i := 0; i < len(Restaurants01); i++ {
		if p <= 0 {
			fmt.Println("Не верно указанно количество человек")
			return nil
		}
		sum := 0
		for q := 0; q < len(Restaurants01[i].FreeTables); q++ {
			sum += Restaurants01[i].FreeTables[q]
		}
		if sum >= p {
			fmt.Printf("\nРесторан: %s\nСвободно мест: %d\nВремя ожидания приготовления: %d\nСредний чек: %02.f\n", Restaurants01[i].NameRestaurant, sum, Restaurants01[i].WaitingForCooking, Restaurants01[i].AverageChek)
			Restaurants02 = append(Restaurants02, Restaurants01[i])
		}
	}
	if len(Restaurants02) <= 0 {
		fmt.Println("Нет подходящих ресторанов для заявленного количества людей")
	}
	return Restaurants02
}
