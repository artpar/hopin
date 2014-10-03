package main

import "github.com/gorilla/mux"
import (
	"net/http"
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Handlers struct {
	db *sql.DB
	Handlers map[string]func(http.ResponseWriter, *http.Request)
}

type ConnectionParam struct {
	user     string
	password string
	port     int
	host     string
	database string
}

type Person struct {
	id    int
	email string
	regId string
}

var con *sql.DB
var err error

func main() {
	r := mux.NewRouter()

	store := &ConnectionParam {
		"root", "", 3306, "localhost", "hopin"};
	fmt.Printf("Host: %s\n", store.host)
	con, err = sql.Open("mysql", store.user+":"+store.password+"@/"+store.database)
	if err != nil {
		panic(err)
	}


	defer con.Close()

	r.HandleFunc("/updateregid", AddPersonHandler)
	http.Handle("/", r)
	http.ListenAndServe(":30824", nil)
}

func getPersonByEmail(email string) *Person {
	row := con.QueryRow("select id, email, regid from user where email = ?", email)
	if err != nil {
		panic(err)
	}
	x := new(Person)
	row.Scan(&x.id, &x.email, &x.regId)
	return x
}

func addNewPerson(x *Person) {
	con.Exec("insert into user (email) value (?)", x.email)
	updatePersonRegId(x, x.regId)
}

func updatePersonRegId(x *Person, r string) {
	con.Exec("update user set regid=? where email=?", r, x.email)
}



func AddPersonHandler(response http.ResponseWriter , request *http.Request) {
	//	vars := mux.Vars(request)
	email := request.FormValue("email")
	regId := request.FormValue("regid")
	fmt.Printf("new person - %s - %s\n", email, regId)
	x := getPersonByEmail(email)
	if (x.id == 0) {
		x = &Person{email:email, regId: regId}
		addNewPerson(x)
	} else if x.regId != regId {
		updatePersonRegId(x, regId)
	}
	x = getPersonByEmail(email)

	if x.email != email {
		response.Write([]byte("Not Ok"))
		return
	}
	response.Write([]byte("Ok"))
}

func NewRideHandler(response http.ResponseWriter, request *http.Request) {

}
