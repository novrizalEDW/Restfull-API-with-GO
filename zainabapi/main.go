package main

import (
	"log"
	"net/http"
	"zainabapi/create_data"
	"zainabapi/delete_data"
	"zainabapi/homepage"
	"zainabapi/read_data"
	"zainabapi/search_data"
	"zainabapi/update_data"

	"github.com/gorilla/mux"
)

func main() {
	HandleRoute()
}

func HandleRoute() {
	r := mux.NewRouter()

	r.HandleFunc("/", homepage.HomePage).Methods("GET")
	r.HandleFunc("/data-nama", read_data.ReadData).Methods("GET")
	r.HandleFunc("/data-nama/{id}", search_data.SearchData).Methods("GET")
	r.HandleFunc("/data-nama", create_data.CreateData).Methods("POST")
	r.HandleFunc("/data-nama/{id}", delete_data.DeleteData).Methods("DELETE")
	r.HandleFunc("/data-nama/{id}", update_data.UpdateData).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8000", r))
}
