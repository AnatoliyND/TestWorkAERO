package web

import (
	"database/sql"
	"fmt"
	"main/restaurants"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Booking struct { //создание структуры, которая описывает нашу таблицу
	Id                                                uint16
	NameRestauran, ClientName, PhoneNumber, VisitTime string
	NumberOfPerson                                    int
}

var allRestaurants = []restaurants.Restaurant{restaurants.RestaurantJuvenility, restaurants.RestaurantKaravella, restaurants.RestaurantMeatAndSalad}

//создаем слайс(список) с типом данных Restaurant, в который будем сохранять новую бронь
var showInfo = Booking{} //внутрь этого объекта будем помещать ту информацию, которую нужно передать

func index(w http.ResponseWriter, r *http.Request) { //с помощью параметра w мы можем обращаться к определенной странице и что-либо показывать на этой странице
	//r параметр, который передается. Через этот параметр можем отследить запрос подключения к странице
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html") //внутри переменной t мы прописываем шаблон который хотим подключить.+
	// обращаемся к пакету template с помощью метода ParseFiles обращаемся к файлам, которые хотим обработать

	if err != nil {
		fmt.Fprintf(w, err.Error()) //обработка ошибки
	}
	person := r.FormValue("person")
	restaurantsForYou := restaurants.ActualReataurants(person)
	t.ExecuteTemplate(w, "index", restaurantsForYou) //используем ExecuteTemplate(), т.к. внутри шаблонов  динамическое подключение

}

func save_Booking(w http.ResponseWriter, r *http.Request) { //метод для обработки данных и переадресации пользователя на какую-либо страницу
	//создаем переменные для получения данных из заполняемой на сайте формы
	t, err := template.ParseFiles("templates/save_Booking.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "save_Booking", nil) //используем ExecuteTemplate(), т.к. внутри шаблонов будем создавать динамическое подключение

	nameRestaurant := r.FormValue("nameR") // в метод r.FormValue передаем название того поля из которого хотим получить значение
	name := r.FormValue("name")
	phone := r.FormValue("phone")
	person := r.FormValue("person")
	visit := r.FormValue("visit")

	if nameRestaurant == "" || name == "" || phone == "" || person == "" || visit == "" { //делаем проверку, что указанные поля заполнены
		fmt.Fprintf(w, "Не все данные введены!")
	} else {

		db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/restaurants") //подключение к базе данных sql. "mysql" - указывает+
		// к какой субд подключаемся. root:root - логин пароль. @tcp(127.0.0.1:3306) - сетевой адрес БД. golang - незвание БД к которой подключаемся
		if err != nil {
			panic(err)
		}
		defer db.Close()

		//установка данных
		insert, err := db.Query(fmt.Sprintf("INSERT INTO `clients`(`nameR`, `name`, `phone`, `person`, `visit`) VALUES('%s', '%s', '%s', '%s', '%s')", nameRestaurant, name, phone, person, visit)) // команда sql+
		//для добавления новой записи в таблицу `Bookings` в поля `nameR`, `name`, `phone`, `visit`. Добавляем значения VALUES, перечисляем добавляемые значения
		if err != nil {
			panic(err)
		}
		defer insert.Close() //закрываем insert после передачи данных

		http.Redirect(w, r, "/", http.StatusSeeOther) // метод http.Redirect переадресовывает нас на страницу. http.StatusSeeOther позволяет делать +
		// переадресацию с верным кодом ответа
	}
}

func HandleFunc() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/", index).Methods("GET")
	rtr.HandleFunc("/", index).Methods("POST")
	rtr.HandleFunc("/createBooking", createBooking).Methods("GET")
	rtr.HandleFunc("/save_Booking", save_Booking).Methods("GET")
	rtr.HandleFunc("/save_Booking", save_Booking).Methods("POST")
	//rtr.HandleFunc("/rest/{id:[0-9]+}", show_restaurant).Methods("GET") //Создаем шаблон для отслеживания URL адресов. /rest/{id:[0-9]+} - говорит о том, +
	// что будем обрабатывать все URL адреса, которые начинаются со слова rest

	http.Handle("/", rtr)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))) //обработка всех url адресов начинающихся со /static/+
	//StripPrefix - удаляет определенный префикс с указанной строки. FileServer ижем необходимый фаил в папке которая находтся внутри этого проекта с названием static

	http.ListenAndServe(":8080", nil)
}
