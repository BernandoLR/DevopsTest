package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
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
	date:=time.Now().Day()
	url := "http://numbersapi.com/"
	path := fmt.Sprint(url, date)
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
	srv.HandleFunc("/back-end", APIGetMessage).Methods("GET")
	fmt.Println("Starting back-end server at port 8081")
	log.Fatal(http.ListenAndServe("0.0.0.0:8081", srv))

}
