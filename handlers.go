package main

import (
	//"subsdb"
	"encoding/json"
	"fmt"

	"net/http"
	"log"
	"strconv"
	"time"
	//"yaml"
)

//@Summary creates subscription
//@Accept json {"ServiceName":'Yandex Plus',"Price":400,"UserId":'60601fee-2bf1-4721-ae6f-7636e79a0cba',"StartDate":'2025-01-07T23:20:50.52Z'}
//@Produces status 200 - success
func HandlerCreate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Creating subscription\n")


var subscription Subscription 
log.Println("unmarshaling", subscription)
decoder := json.NewDecoder(r.Body)
decoder.DisallowUnknownFields()
um_err := decoder.Decode(&subscription)

if um_err != nil {
	log.Panic(um_err)
}
}

//@Summary create subscription
//@Accept id in path type int
//@Produces json
func HandlerRead(w http.ResponseWriter, r *http.Request) {

	log.Println("read by id")
	id := r.PathValue("id")
	log.Println("read by id:", id)
	int_id, err := strconv.Atoi(id)
	if err != nil {
		log.Println("Input error ", err)
	}
	subscription := ReadById(int_id)
	jdata, err := json.Marshal(subscription)
	if err != nil {
		log.Println("json error: ", err)
	}
	w.Write(jdata)
}

//@Summary updates subscription
//@Accept json {"ServiceName":'Yandex Plus',"Price":400,"UserId":'60601fee-2bf1-4721-ae6f-7636e79a0cba',"StartDate":'2025-01-07T23:20:50.52Z'}
//@Produces status 200 - success

func HandlerEdit(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Editing")
	var subscription Subscription 
log.Println("unmarshaling")
//um_err := json.Unmarshal(bytes, &subscription)
//bytes, read_err := io.ReadAll(r.Body)
decoder := json.NewDecoder(r.Body)
decoder.DisallowUnknownFields()
//str_map := map[string]string{}
um_err := decoder.Decode(&subscription)
if um_err != nil {
	log.Println(w, "error", um_err)
}
log.Println("To update: ", subscription)
Update(subscription)
}

//@Summary deletes a subscription
//@Accept id in path type int
//@Produces status: 200 - success
func HandlerDelete(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Deleting")
	id := r.PathValue("id")
	log.Println("delete id ", id)
	int_id, err := strconv.Atoi(id)
	if err != nil {
		log.Println("atoi error:", err)
	}
	res := Delete(int_id)

	log.Println("delete res:", res)

	if res != nil {
	}
	
}

//@Summary returns array of subscriptions
//@Accept no parameters
//@Produces json
func HandlerList(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "preparing list")

	subscriptions := List()
	if (subscriptions != nil && len(subscriptions) > 0) {
		jdata, err := json.Marshal(subscriptions)
		if err != nil {
			log.Println("list marshar error", err)
			w.WriteHeader(http.StatusConflict)
		}

		w.Write(jdata)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

type SumConditions struct {
	FromDate time.Time `json:"FromDate"`
	ToDate time.Time `json:"ToDate"`
	UserId string `json: "UserId"`
	ServiceName string `json: "ServiceName"`
}

//@Summary updates subscription
//@Accept json {"FromDate":"2025-01-07T23:20:50.52Z", "ToDate": "2028-01-07T23:20:50.52Z", "UserId": "60601fee-2bf1-4721-ae6f-7636e79a0cba" }
// cant sum by user_id, ServiceName
//@Produces int - sum.
func HandlerSumByCondition(w http.ResponseWriter, r *http.Request) {

	var cond SumConditions
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	//str_map := map[string]string{}
	err := decoder.Decode(&cond)
	if err != nil {
		log.Println("decode conditions error:", err)
	}
	sum := SumByConditions(cond)

	fmt.Fprintf(w, "%d", sum)

}