package main

import (
	"database/sql"
	"os"
	"strconv"

	_ "github.com/lib/pq"

	"fmt"
	"log"
)



func OpenDb(admin bool) *sql.DB {
		host := os.Getenv("DB_HOST")
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	if err != nil {
		log.Panic(err)
	}


	connection_string := ""
	if !admin {
	user:= os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	connection_string = fmt.Sprintf("host=%s port=%d" +
	" user=%s password=%s dbname=%s sslmode=disable",
host, port, user, password, dbname)
	} else {
		user := os.Getenv("DB_ADMIN_NAME")
		password = os.Getenv("DB_ADMIN_PASSWORD")
			connection_string = fmt.Sprintf("host=%s port=%d" +
	" user=%s password=%s dbname=%s sslmode=disable",
host, port, user, password, dbname)

	}
	log.Println("connect string: ", connection_string)
	db, err := sql.Open("postgres", connection_string)
	if err != nil {
		log.Println("open db error", err)
		panic(err)
	}
	return db
}

const (
	up_script = "./migrations/up_subscriptions.sql"
	down_script = "./migrations/down_subscriptions.sql"
)

func InitDb(up bool) sql.Result {
	
	var file string = ""
	if up {
		file = up_script
	} else {
		file = down_script
	}

	script, err := os.ReadFile(file)
	if err != nil {
		log.Println("script read error:", err)
	}
	db := OpenDb(true)
	str_script := string(script[:])
	res, err := db.Exec(str_script)
	if err != nil {
		log.Println("script error:", err)
	}
	return res
}

//todo изучить ORM для golang
func Create(s Subscription) {
	insert_sql := "insert into public.subscriptions(service_name,price," +
	"user_id, start_date)" +
	"values($1, $2, $3, $4)"
	fmt.Println("do create: ",s)
	db := OpenDb(false)
	_ , err := db.Exec(insert_sql, s.ServiceName, s.Price, s.UserId, s.StartDate)
	if err != nil {
		fmt.Println("error insertion in db ", err)
	} else {
		fmt.Println("succefully created")
	}
}

func ReadById(id int) *Subscription {
	select_sql := "select service_name, price, user_id, start_date from public.subscriptions " +
	"where subscription_id = $1" 
	db := OpenDb(false)
	subscription := Subscription{}
	err := db.QueryRow(select_sql, id).Scan(&subscription.ServiceName, &subscription.Price,
	 &subscription.UserId, &subscription.StartDate)
	if err != nil {
		log.Println("erro select db ", err)
	}

	fmt.Println("sub: ", subscription)

return &subscription
}

func Update(s Subscription) sql.Result {
	update_sql := "update public.subscriptions set service_name = $1, price = $2, user_id = $3, start_date = $4  where subscription_id = $5"
	db := OpenDb(false)
	res, err := db.Exec(update_sql, s.ServiceName, s.Price, s.UserId, s.StartDate, s.SubscriptionId)
	if err != nil {
		log.Println("update error ", err)
	}
	log.Println("Update res: ", res)
	return res
}

func Delete(id int) sql.Result {
	delete_sql := "delete from public.subscriptions where subscription_id = $1"
	db := OpenDb(false)
	res, err := db.Exec(delete_sql, id)
	if err != nil {
		log.Println("delete error:", err)

	}
	return res
}

func List() []Subscription {
	all_sql := "select * from public.subscriptions;"
	db := OpenDb(false)
	rows, err := db.Query(all_sql)
	if err != nil {
		fmt.Println("select error:", err)
	}
	var subscriptions []Subscription
	for rows.Next() {
		var s Subscription
		err := rows.Scan(&s.SubscriptionId, &s.ServiceName, &s.Price, &s.UserId, &s.StartDate,
		&s.CreateDate)
		if err != nil {
			log.Println("scan error:", err)
			return subscriptions
		}
		subscriptions = append(subscriptions, s)
	}
	return subscriptions
}

func SumByConditions(c SumConditions) (int) {
	base_sql := "select sum(price) from public.subscriptions where start_date > $1" +
	" and start_date < $2"
	user_sql := " and user_id = $3"
	service_sql := " and service_name = $3"
	var sum int
	var err any
	db := OpenDb(false)
	if len(c.UserId) > 0 {
		sql := base_sql + user_sql
		err = db.QueryRow(sql, c.FromDate, c.ToDate, c.UserId).Scan(&sum)

	} else {
		sql := base_sql + service_sql
		err = db.QueryRow(sql, c.FromDate, c.ToDate, c.ServiceName).Scan(&sum)

	}
	if err != nil {
		log.Println("query sum error:", err)
	}

	return sum

}
