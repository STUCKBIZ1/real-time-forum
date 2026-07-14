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
	addr := config.Server.Host + ":" + config.Server.Port
	fmt.Println(addr)
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}
