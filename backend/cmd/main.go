package main

import (
	"fmt"
	"log"
	"net/http"
	"real-time-forum/backend/config"
	"real-time-forum/backend/database"
	"real-time-forum/backend/router"
)

func main() {
	db, err := database.InitDb()
	if err != nil {
		log.Fatal(err)
	}
	router.SetupRouter(db)
	fs := http.FileServer(
		http.Dir("./frontend"),
	)
	http.Handle("/", fs)
	addr := config.Server.Host + ":" + config.Server.Port
	fmt.Println("http://"+addr)
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}
