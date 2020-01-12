package main

import (
	"database/sql"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/tinrab/retry"
	"log"
	"net/http"
	"os"
	"processor/core/requester"
	"processor/payment/gateway"
	"processor/payment/repository/mysql"
	processorhttp "processor/payment/server/http"
	"processor/payment/service"
	"time"
)

const (
	driver         = "mysql"
	dataSourceName = "root:Xiulop25_pI98prd@root.process@tcp(172.30.1.1)/processor?parseTime=true"
)

func main() {
	fmt.Println("Starting API...")
	fmt.Println("Connecting to Database...")
	var db *sql.DB

	retry.ForeverSleep(2*time.Second, func(_ int) (err error) {
		db, err = sql.Open(driver, dataSourceName)
		if err != nil {
			fmt.Println("Error Connecting to Database: ", err)
			os.Exit(1)
		}
		return
	})

	processRepository, err := mysql.NewRepository(db)

	if err != nil {
		fmt.Println("Could not connect with the database")
		panic(err)
	}

	fmt.Println("Database Connected!")

	fmt.Println("Opening the Gateway...")
	req := requester.NewRequester()
	gate := gateway.NewGateway(processRepository, req)
	fmt.Println("Gateway Opened!")

	fmt.Println("Starting the Services...")
	serv := service.NewService(processRepository, gate)
	fmt.Println("Service Started!")

	fmt.Println("Starting the Router...")
	r := chi.NewMux()
	processorhttp.NewHandler(serv, r)
	fmt.Println("Router Started!")
	port := ":8082"
	fmt.Println(fmt.Sprintf("API Started! Listening on Port: %s", port))
	log.Fatal(http.ListenAndServe(port, r))
}
