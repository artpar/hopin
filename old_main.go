//package main
//
//import (
//	"github.com/gorilla/mux"
//	"net/http"
//	"fmt"
//	"database/sql"
//	_ "github.com/go-sql-driver/mysql"
//	logging "github.com/op/go-logging"
//	"os"
//	"strconv"
//	"github.com/artpar/hopin/helper"
//)
//
//var log = logging.MustGetLogger("example")
//var format = "%{time:15:04:05.000000} ▶ %{level:.4s} %{id:03x} %{message}"
//
//type Person struct {
//	id    int
//	email string
//	regId string
//}
//
//var con *sql.DB
//var err error
//
//func main() {
//
//	args := helper.ProcessArguments(os.Args)
//	fmt.Println(args)
//	profile := args["profile"]
//	configuration := helper.GetConfiguration(profile)
//
//	logBackend := logging.NewLogBackend(os.Stderr, "", 0)
//	logFile, err := os.OpenFile(configuration.LogFileLocation, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
//	logBackendFile := logging.NewLogBackend(logFile, "", 0)
//	syslogBackend, err := logging.NewSyslogBackend("")
//	if err != nil {
//		log.Fatal(err)
//	}
//	logging.SetLevel(logging.DEBUG, "hopinserver")
//	logging.SetBackend(logBackend, syslogBackend, logBackendFile)
//	logging.SetFormatter(logging.MustStringFormatter(format))
//
//	r := mux.NewRouter()
//
//	store := configuration.Database;
//
//	log.Info("Host: %s\n", store.Host)
//	con, err = sql.Open("mysql", store.User+":"+store.Password+"@/"+store.Database)
//	if err != nil {
//		panic(err)
//	}
//
//
//	defer con.Close()
//
//	r.HandleFunc("/updateregid", AddPersonHandler)
//	http.Handle("/", r)
//	log.Info("serving at %s:%s", configuration.Server.Host, strconv.Itoa(configuration.Server.Port))
//	http.ListenAndServe(configuration.Server.Host+":"+strconv.Itoa(configuration.Server.Port), nil)
//}
//
//func getPersonByEmail(email string) *Person {
//	row := con.QueryRow("select id, email, regid from user where email = ?", email)
//	if err != nil {
//		panic(err)
//	}
//	x := new(Person)
//	row.Scan(&x.id, &x.email, &x.regId)
//	return x
//}
//
//func addNewPerson(x *Person) {
//	con.Exec("insert into user (email) value (?)", x.email)
//	updatePersonRegId(x, x.regId)
//}
//
//func updatePersonRegId(x *Person, r string) {
//	con.Exec("update user set regid=? where email=?", r, x.email)
//}
//
//
//
//func AddPersonHandler(response http.ResponseWriter , request *http.Request) {
//	//	vars := mux.Vars(request)
//	email := request.FormValue("email")
//	regId := request.FormValue("regid")
//	log.Info("new person - %s - %s\n", email, regId)
//	x := getPersonByEmail(email)
//	if (x.id == 0) {
//		x = &Person{email:email, regId: regId}
//		addNewPerson(x)
//	} else if x.regId != regId {
//		updatePersonRegId(x, regId)
//	}
//	x = getPersonByEmail(email)
//
//	if x.email != email {
//		response.Write([]byte("Not Ok"))
//		return
//	}
//	response.Write([]byte("Ok"))
//}
//
