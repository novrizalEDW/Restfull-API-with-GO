package delete_data

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

func DeleteData(w http.ResponseWriter, r *http.Request) {

	var result []Pasien
	var pasien = Pasien{}

	params := mux.Vars(r)
	for _, item := range result {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	db, err := connect_db.Connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	cabe := params["id"]

	_, err = db.Exec("delete from zainab_author where id = ? ", cabe)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("delete success!")

	result = append(result, pasien)
	json.NewEncoder(w).Encode(result)
}
