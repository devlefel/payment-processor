package main

import (
	authenticatorhttp "authenticator/user/handler/http"
	"authenticator/user/repository/mysql"
	"authenticator/user/service"
	"database/sql"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/tinrab/retry"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	driver         = "mysql"
	dataSourceName = "root:dev@root.authXiulop25_pI98@tcp(172.28.1.1)/authenticator?parseTime=true"
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

	fmt.Println("Starting the Services...")
	serv := service.NewService(processRepository)
	fmt.Println("Service Started!")

	fmt.Println("Starting the Router...")
	r := chi.NewMux()
	authenticatorhttp.NewHandler(serv, r)
	fmt.Println("Router Started!")
	port := ":8081"
	fmt.Println(fmt.Sprintf("API Started! Listening on Port: %s", port))
	log.Fatal(http.ListenAndServe(port, r))
}
