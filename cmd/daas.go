package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// type BaseResult struct {
// 	Result  bool
// 	Message string
// }

func handleRequests() {
	myrouter := mux.NewRouter().StrictSlash(false)

	for _, srvcs := range settings.Services {
		fmt.Printf("registering %s with url %s\n", srvcs.Name, srvcs.Url)
		myrouter.HandleFunc(srvcs.Url, srvcs.dbServiceGet).Methods("GET")
	}

	httpport := fmt.Sprintf(":%d", settings.Db.ServicePort)
	log.Fatal(http.ListenAndServe(httpport, myrouter))
}

func main() {
	loadConfig("config.yaml")
	ConnectMySql()
	handleRequests()
}
