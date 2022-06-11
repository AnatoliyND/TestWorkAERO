package web

import (
	"database/sql"
	"fmt"
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

var booking = []Booking{} //создаем слайс(список) с типом данных Booking, в который будем сохранять новую бронь
var showInfo = Booking{}  //внутрь этого объекта будем помещать ту информацию, которую нужно передать

func index(w http.ResponseWriter, r *http.Request) { //с помощью параметра w мы можем обращаться к определенной странице и что-либо показывать на этой странице
	//r параметр, который передается. Через этот параметр можем отследить запрос подключения к странице
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html") //внутри переменной t мы прописываем шаблон который хотим подключить.+
	// обращаемся к пакету template с помощью метода ParseFiles обращаемся к файлам, которые хотим обработать

	if err != nil {
		fmt.Fprintf(w, err.Error()) //обработка ошибки
	}

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/restaurants") //"mysql" параметр указывает к какой СУБД мы поключаемся, во втором параметре указываются данные для подключения - логин, пароль, хост, доп порт, название БД к которой подключаемся
	if err != nil {
		panic(err) //выполняем проверку на наличие ошибки, при ее присутствии вызываем панику с выводом этой ошибки
	}
	defer db.Close() // закрываем БД

	res, err := db.Query("SELECT * FROM `clients`") //делаем выборку данных. SELECT * FROM `clients` - команда позволяет вытянуть все данные из +
	//таблички `clients`, либо вместо * можно указать конкретные поля

	if err != nil {
		panic(err)
	}

	booking = []Booking{} //перед циклом обращаемся к данным по бронированию и указываем, что это пустой список(необходимо, чтобы брони не дублировались +
	//при обновлении страницы) каждый раз когда будем попадать на главную страницу список будут назначаться пустым(ранее добавленые данные сохраняются)
	for res.Next() { //перебираем все res. Метод Next возвращает нам либо true - если есть следующая строка, которую можно обработать,+
		// либо false - если нет строк, которые можно обработать
		var post Booking                                                                                                          //создаем объект на основе структуры Booking
		err = res.Scan(&post.Id, &post.NameRestauran, &post.ClientName, &post.PhoneNumber, &post.VisitTime, &post.NumberOfPerson) //убеждаемся существуют ли какие-либо данные в ряде, который рассматриваем и вытягиваем их
		if err != nil {
			panic(err)
		}
		booking = append(booking, post)

	}

	t.ExecuteTemplate(w, "index", booking) //используем ExecuteTemplate(), т.к. внутри шаблонов  динамическое подключение
}

func save_Booking(w http.ResponseWriter, r *http.Request) { //метод для обработки данных и переадресации пользователя на какую-либо страницу
	//создаем переменные для получения данных из заполняемой на сайте формы
	nameRestaurant := r.FormValue("nameR")
	name := r.FormValue("name") // в метод r.FormValue передаем название того поля из которого хотим получить значение
	phone := r.FormValue("phone")
	person := r.FormValue("person")
	visit := r.FormValue("visit")

	if nameRestaurant == "" || name == "" || phone == "" || person == "" || visit == "" { //делаем проверку, что указанные поля заполнены
		fmt.Fprintf(w, "Не все данные введены!")
	} else {

		db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/golang") //подключение к базе данных sql. "mysql" - указывает+
		// к какой субд подключаемся. root:root - логин пароль. @tcp(127.0.0.1:3306) - сетевой адрес БД. golang - незвание БД к которой подключаемся
		if err != nil {
			panic(err)
		}
		defer db.Close()

		//установка данных
		insert, err := db.Query(fmt.Sprintf("INSERT INTO `clients`(`nameR`, `name`, `phone`, `person`, `visit`) VALUES('%s', '%s', '%s', '%d', '%s')", nameRestaurant, name, phone, person, visit)) // команда sql+
		//для добавления новой записи в таблицу `Bookings` в поля `name`, `phone`, `visit`. Добавляем значения VALUES, перечисляем добавляемые значения
		if err != nil {
			panic(err)
		}
		defer insert.Close() //закрываем insert после передачи данных

		http.Redirect(w, r, "/", http.StatusSeeOther) // метод http.Redirect переадресовывает нас на страницу. http.StatusSeeOther позволяет делать +
		// переадресацию с верным кодом ответа
	}
}

func createBooking(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/createBooking.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "createBooking", nil) //используем ExecuteTemplate(), т.к. внутри шаблонов будем создавать динамическое подключение
}

func show_post(w http.ResponseWriter, r *http.Request) { //функция обрабатывает страничку для отображения полной информации про какую-либо статью
	vars := mux.Vars(r) //создаем объект vars на основе библиотеки mux и применяем метод Vars, в который передаем параметр r

	t, err := template.ParseFiles("templates/show.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/golang")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	res, err := db.Query(fmt.Sprintf("SELECT * FROM `Bookings` WHERE `id` = '%s'", vars["id"]))

	if err != nil {
		panic(err)
	}

	showInfo = Booking{}
	for res.Next() {
		var post Booking
		err = res.Scan(&post.Id, &post.NameRestauran, &post.ClientName, &post.PhoneNumber, &post.VisitTime, &post.NumberOfPerson)
		if err != nil {
			panic(err)
		}
		showInfo = post

	}

	t.ExecuteTemplate(w, "show", showInfo)

}

func HandleFunc() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/", index).Methods("GET")
	rtr.HandleFunc("/createBooking", createBooking).Methods("GET")
	rtr.HandleFunc("/save_Booking", save_Booking).Methods("POST")
	rtr.HandleFunc("/post/{id:[0-9]+}", show_post).Methods("GET") //Создаем шаблон для отслеживания URL адресов. /post/{id:[0-9]+} - говорит о том, +
	// что будем обрабатывать все URL адреса, которые начинаются со слова post

	http.Handle("/", rtr)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/")))) //обработка всех url адресов начинающихся со /static/+
	//StripPrefix - удаляет определенный префикс с указанной строки. FileServer ижем необходимый фаил в папке которая находтся внутри этого проекта с названием static

	http.ListenAndServe(":8080", nil)
}
