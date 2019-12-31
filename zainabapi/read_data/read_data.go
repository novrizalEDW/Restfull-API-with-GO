package read_data

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

// function Untuk Read Data
func ReadData(w http.ResponseWriter, r *http.Request) {
	var result []Pasien

	db, err := connect_db.Connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	rows, err := db.Query("select * from zainab_author")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	for rows.Next() {
		var bungkus1 = Pasien{}
		var err = rows.Scan(&bungkus1.ID, &bungkus1.Nama, &bungkus1.Alamat, &bungkus1.Username, &bungkus1.Password, &bungkus1.Email)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		result = append(result, bungkus1)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
