package restaurants

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

type Tables interface {
	ReservedTables(int) []int
}

func (r *Restaurant) ReservTables(name, phomeNumber, orderTime string, numberperson int) (*Restaurant, error) {

	if fmt.Sprintf(strings.TrimSpace(name)) == "" {
		fmt.Println("Для бронирования мест укажите имя")
		return r, fmt.Errorf("Не указано имя для бронирования мест")
	}
	if fmt.Sprintf(strings.TrimSpace(phomeNumber)) == "" {
		fmt.Println("Для бронирования мест укажите контактный номер телефона")
		return r, fmt.Errorf("Не указан контактный телефон для бронирования мест")
	}
	tc, err := time.Parse("3:04 PM", orderTime)
	if err != nil {
		fmt.Println("Не корректно введено время %w", err)
	}
	tb, _ := time.Parse("3:04 PM", r.LastBooking)
	to, _ := time.Parse("3:04 PM", r.OpenTime)

	if tb.Sub(tc) < 0 {
		fmt.Println("Ресторан скоро закрывается")
		return r, fmt.Errorf("Ресторан скоро закрывается")
	} else if tc.Sub(to) < 0 {
		fmt.Println("Ресторан еще закрыт")
		return r, fmt.Errorf("Ресторан еще закрыт")
	} else {
		fmt.Println("Можно забронировать")
	}
	if r.FreeTables == nil {
		fmt.Println("Свободных мест нет")
		return r, fmt.Errorf("Свободных мест нет")
	}
	if numberperson <= 0 {
		fmt.Println("Не верно указанно количество бронируемых мест")
		return r, fmt.Errorf("Не верно указанно количество бронируемых мест")
	}
	sum := 0
	for i := 0; i < len(r.FreeTables); i++ {
		sum += r.FreeTables[i]
	}
	fmt.Println("Всего мест в ресторане:", sum)

	sort.Ints(r.FreeTables)

	p := numberperson
	if numberperson <= sum {
		for q := 0; q < len(r.FreeTables); q++ {
			if p >= r.FreeTables[len(r.FreeTables)-1]-1 {
				p -= r.FreeTables[len(r.FreeTables)-1]
				r.ReservedTables = append(r.ReservedTables, r.FreeTables[len(r.FreeTables)-1])
				copy(r.FreeTables[len(r.FreeTables)-1:], r.FreeTables[len(r.FreeTables):])
				r.FreeTables[len(r.FreeTables)-1] = 0
				sort.Ints(r.FreeTables)
			}

			if p < r.FreeTables[len(r.FreeTables)-1]-1 && p > 0 && r.FreeTables[q] != 0 {
				p -= r.FreeTables[q]
				r.ReservedTables = append(r.ReservedTables, r.FreeTables[q])
				copy(r.FreeTables[q:], r.FreeTables[q+1:])
				r.FreeTables[len(r.FreeTables)-1] = 0
				sort.Ints(r.FreeTables)
			}
			sort.Ints(r.FreeTables)

		}
		fmt.Printf("\nРесторан: %s\nЗарезервированные столы: %v\nСвободные столы: %v\n", r.NameRestaurant, r.ReservedTables, r.FreeTables)
		return &Restaurant{ReservedTables: r.ReservedTables, FreeTables: r.FreeTables}, nil
	} else {
		fmt.Printf("\nThere are not enough seats in the restaurant\nВ ресторане недостаточно мест\n\n")
		return r, fmt.Errorf("\nThere are not enough seats in the restaurant\nВ ресторане недостаточно мест")
	}

}
