package main

import (
	//"subsdb"
	//"encoding/json"
	//"fmt"
	"strconv"

	"os"
	//"io/ioutil"
	"net/http"
	//"encoding/json"
	//"database/sql"
	//"database/sql"
	"log"
	//"strconv"
	"time"
	//"yaml"

	//"github.com/goccy/go-yaml"
	//"github.com/kelseyhightower/envconfig"
	"github.com/joho/godotenv"
	//_ "github.com/lib/pq"
	//"strings"
	//"github.com/google/uuid"
	//"strings"
	//"db/subscriptionsdb/Cre/AddSubscriptions"
	//"./db/subsdb"
)




type Subscription struct {
SubscriptionId int `json:"SubscriptionId" `	
ServiceName string `json: "ServiceName"`
Price int `json:"Price"`
UserId string `json:"UserId"`
StartDate time.Time `json: "StartDate"`
CreateDate time.Time `json: "CreateDate"`
}

func main() {
	/**
	* Загрузка кофигурации из .env файла в переменные окружения.
	*/
	err := godotenv.Load()
	if err != nil {
		log.Println("env error:", err)
	}

	//check env
	dbuser := os.Getenv("DB_USER")

	log.Println("env user:", dbuser)


	log.Println("Start")
	router := http.NewServeMux()
	// Migration if true - up, false - down.
	need_init, err := strconv.ParseBool(os.Getenv("SERVE_INIT_DB"))
	if need_init {
		log.Println("init db")
		InitDb(true)
	}
	// "ручки".
	router.HandleFunc("/api/create", HandlerCreate) 	//"C"
	router.HandleFunc("/api/read/{id}", HandlerRead) 	// "R" 
	router.HandleFunc("/api/edit", HandlerEdit) 		// "U"
	router.HandleFunc("/api/delete/{id}", HandlerDelete) // "D"
	router.HandleFunc("/api/list", HandlerList) 		 // "L"
	router.HandleFunc("/api/sum", HandlerSumByCondition) // "sums"
        
	port := os.Getenv("SERVE_PORT")
	log.Println("serve on " + port)
	http.ListenAndServe(port, router)
}