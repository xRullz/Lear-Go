package main

import (
	"fmt"
	"net/http"
	"restapi/config"
	"restapi/routes"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

func main() {
	config.LoadConfig()
	config.ConnectDB()

	r := mux.NewRouter()
	routes.IndexRoute(r)

	log.Println("Server running on port 8080")
	http.ListenAndServe(fmt.Sprintf(":%v", "8080"), r)
}
