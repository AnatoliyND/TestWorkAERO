package restaurants

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
