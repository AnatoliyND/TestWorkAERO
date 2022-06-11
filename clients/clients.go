package clients

import (
	"fmt"
	"main/restaurants"
	"sort"
)

type Client struct {
	NameClient     string
	PhoneNumber    string
	NumberOfPerson int
	OrderTime      string
}

type People []Client

type ByWaitingForCooking []restaurants.Restaurant

func (a ByWaitingForCooking) Len() int { return len(a) }
func (a ByWaitingForCooking) Less(i, j int) bool {
	return float64(a[i].WaitingForCooking) < float64(a[j].WaitingForCooking)
}
func (a ByWaitingForCooking) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

type ByAverageCheck []restaurants.Restaurant

func (a ByAverageCheck) Len() int           { return len(a) }
func (a ByAverageCheck) Less(i, j int) bool { return a[i].AverageChek < a[j].AverageChek }
func (a ByAverageCheck) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func NewClient(name, pnumber, time string, person int) *Client {
	return &Client{NameClient: name, PhoneNumber: pnumber, OrderTime: time, NumberOfPerson: person}
}

func (c *Client) ActualReataurants() []restaurants.Restaurant {
	var Restaurants01 = []restaurants.Restaurant{restaurants.RestaurantJuvenility, restaurants.RestaurantKaravella, restaurants.RestaurantMeatAndSalad}

	sort.Sort(ByAverageCheck(Restaurants01))
	sort.Sort(ByWaitingForCooking(Restaurants01))

	var Restaurants02 = []restaurants.Restaurant{}
	for i := 0; i < len(Restaurants01); i++ {
		if c.NumberOfPerson <= 0 {
			fmt.Println("Не верно указанно количество человек")
			return nil
		}
		sum := 0
		for q := 0; q < len(Restaurants01[i].FreeTables); q++ {
			sum += Restaurants01[i].FreeTables[q]
		}
		if sum >= c.NumberOfPerson {
			fmt.Printf("\nРесторан: %s\nСвободно мест: %d\nВремя ожидания приготовления: %d\nСредний чек: %02.f\n", Restaurants01[i].NameRestaurant, sum, Restaurants01[i].WaitingForCooking, Restaurants01[i].AverageChek)
			Restaurants02 = append(Restaurants02, Restaurants01[i])
		}
	}
	if len(Restaurants02) <= 0 {
		fmt.Println("Нет подходящих ресторанов для заявленного количества людей")
	}
	return Restaurants02
}
