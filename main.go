package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {
	print("hellow")
	log.Println("Starting the application...")

	router := mux.NewRouter()
	//dbClient := getDbClient()
	//RepositoryDb := domain.NewRepositoryDb(dbClient)
	//h := app.Handlers{Service.NewService(RepositoryDb)}

	router.HandleFunc("/api/Verify", VerifiLogin).Methods(http.MethodGet)
	router.HandleFunc("/api/Connection", VerifiLogin1).Methods(http.MethodGet)
	router.HandleFunc("/api/Testing", VerifiLogin2).Methods(http.MethodGet)

	//router.HandleFunc("/api/loadGrid/{status}", h.GetallStudent).Methods(http.MethodGet)
	//router.HandleFunc("/api/loadform/{student_id:[0-9]+}", h.GetStudentinfo).Methods(http.MethodGet)
	//router.HandleFunc("/api/delete/{student_id:[0-9]+}", h.DeleteStudentinfo).Methods(http.MethodPost)

	listenAddr := ":8282"

	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		listenAddr = ":" + val
	}
	fmt.Println("hi there")

	log.Printf("About to listen on %s. Go to https://127.0.0.1%s", listenAddr, listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, router))

}

type Success struct {
	Success bool
	ErrorID int
	Message string
}

func VerifiLogin(w http.ResponseWriter, r *http.Request) {

	response := Success{
		Success: true,
		ErrorID: 2,
		Message: "Backend connected Successfully",
	}

	writeResponse(w, http.StatusOK, response)

}

func VerifiLogin1(w http.ResponseWriter, r *http.Request) {

	response := Success{
		Success: true,
		ErrorID: 2,
		Message: "Backend connected Successfully version2",
	}

	writeResponse(w, http.StatusOK, response)

}

func VerifiLogin2(w http.ResponseWriter, r *http.Request) {

	response := Success{
		Success: true,
		ErrorID: 2,
		Message: "Backend connected Successfully version3",
	}

	writeResponse(w, http.StatusOK, response)

}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}

//func getDbClient() *sqlx.DB {
//
//	//dbUser := "srvadmin"
//	//dbAddr := "ccm-test-srv.database.windows.net"
//	//dbPort := "1433"
//	//dbPass := "Avasoft@123"
//	//dbName := "ava-eus-ccm-db"
//	//log.Println("DB DETAILS...")
//
//	//fmt.Println(dbUser, dbPass, dbAddr, dbPort, dbName, "line 78")
//	//dataSource := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", dbUser, dbPass, dbAddr, dbPort, dbName)
//
//	//dataSourse := fmt.Sprintf("sqlserver://sa:Admin123@:1433?database=Goapi")
//
//	//client, err := sqlx.Connect("mysql", "root:Admin@123@/Goapi")
//	client, err := sqlx.Connect("mysql", "root:Admin@123@/Goapi")
//	//client, err := sql.Open("mysql", "root:Admin@123@/Goapi")
//
//	if err != nil {
//		panic(err)
//	} else {
//		log.Println("Db connected successfully......")
//
//	}
//	// See "Important settings" section.
//	client.SetConnMaxLifetime(time.Minute * 3)
//	client.SetMaxOpenConns(10)
//	client.SetMaxIdleConns(10)
//
//	return client
//
//}
