package update_data

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

type ResponData struct {
	Status string `json:"status"`
	Kode   string `json:"code"`
}

func UpdateData(w http.ResponseWriter, r *http.Request) {
	var status []ResponData
	var pasien = Pasien{}
	var respon ResponData

	var result []Pasien

	params := mux.Vars(r) // Get params
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

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewDecoder(r.Body).Decode(&pasien)

	cabe := params["id"]
	_, err = db.Exec("update zainab_author set nama = ?, alamat = ?, username = ?, password = ?, email = ? where id = ?", pasien.Nama, pasien.Alamat, pasien.Username, pasien.Password, pasien.Email, cabe)
	if err != nil {
		fmt.Println(err.Error())

		respon.Status = "update Error"
		respon.Kode = "code error"

		status = append(status, respon)

		json.NewEncoder(w).Encode(status)

		return
	} else {

		respon.Status = "update succes"
		respon.Kode = "200"

		status = append(status, respon)
		fmt.Println("update success!")

		json.NewEncoder(w).Encode(status)

	}

}
