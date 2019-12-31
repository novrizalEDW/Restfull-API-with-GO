package search_data

import (
	"encoding/json"
	"fmt"
	"net/http"
	"zainabapi/connect_db"

	"github.com/gorilla/mux"
)

type Pasien struct {
	ID       string `json:"id"`
	Nama     string `json:"nama"`
	Alamat   string `json:"alamat"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func SearchData(w http.ResponseWriter, r *http.Request) {

	var pasien = Pasien{}

	var result []Pasien
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r) // Get params
	for _, item := range result {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	var db, err = connect_db.Connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()
	cabe := params["id"]
	err = db.
		QueryRow("select id, nama, alamat, username, password, email from zainab_author where id ="+cabe).
		Scan(&pasien.ID, &pasien.Nama, &pasien.Alamat, &pasien.Username, &pasien.Password, &pasien.Email)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	result = append(result, pasien)
	json.NewEncoder(w).Encode(result)
}
