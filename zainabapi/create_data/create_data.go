package create_data

import (
	"encoding/json"
	"fmt"
	"net/http"
	"zainabapi/connect_db"
)

type Pasien struct {
	ID       int    `json:"id"`
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

func CreateData(w http.ResponseWriter, r *http.Request) {

	var status []ResponData
	var pasien = Pasien{}
	var respon ResponData

	db, err := connect_db.Connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewDecoder(r.Body).Decode(&pasien)

	stmt, err := db.Prepare("insert into zainab_author values (?, ?, ?, ?, ?, ?)")

	stmt.Exec(nil, pasien.Nama, pasien.Alamat, pasien.Username, pasien.Password, pasien.Email)

	if err != nil {
		fmt.Println(err.Error())

		respon.Status = "Error"
		respon.Kode = "code error"

		status = append(status, respon)

		json.NewEncoder(w).Encode(status)

		return
	} else {

		respon.Status = "succes"
		respon.Kode = "200"

		status = append(status, respon)
		fmt.Println("insert success!")

		json.NewEncoder(w).Encode(status)

	}

}
