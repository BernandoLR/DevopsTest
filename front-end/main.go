package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func APIGetMessage(res http.ResponseWriter, req *http.Request) {

	msg,err := GetMessage()
	if err != nil {
		log.Println("Error cause", err)
		fmt.Fprintln(res, err)
	}

	fmt.Fprintln(res, msg)


}

func GetMessage() (string,error) {
	path := "http://backend-service.default.svc.cluster.local:8081/back-end"
	log.Println("Get message from", path)
	res, err := http.Get(path)
	if err != nil {
		log.Println("Error cause", err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("Error when read data, cause", err)
	}

	return string(body),nil

}

func main() {
	srv := mux.NewRouter()
	srv.HandleFunc("/", APIGetMessage).Methods("GET")
	fmt.Println("Starting front-end server at port 8080")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", srv))
}
