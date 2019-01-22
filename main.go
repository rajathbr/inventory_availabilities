package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/magiconair/properties"

	_ "github.com/lib/pq"
)

type Availability struct {
	ItemID                string `json:"id,omitempty"`
	UnitOfMeasure         string `json:"uom,omitempty"`
	Onhand                int16  `json:"onhand,omitempty"`
	Demand                int16  `json:"demand,omitempty"`
	AvailableToPromiseQty int16  `json:"atpqty,omitonempty"`
}

var availabilites []Availability

func GetATPData(w http.ResponseWriter, r *http.Request) {
	p := properties.MustLoadFile("application.properties", properties.UTF8)
	host := p.GetString("host", "localhost")
	port := p.GetInt("port", 5432)
	user := p.GetString("user", "atpdbusr")
	password := p.GetString("password", "password")
	dbname := p.GetString("dbname", "atpdb")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("db connection got")

	sqlStatement := `SELECT * FROM atp where item_id=$1;`
	var avail Availability

	params := mux.Vars(r)
	row := db.QueryRow(sqlStatement, params["itemid"])

	fmt.Println("db row got")

	selecterr := row.Scan(&avail.ItemID, &avail.UnitOfMeasure, &avail.Onhand, &avail.Demand, &avail.AvailableToPromiseQty)

	switch selecterr {
	case sql.ErrNoRows:
		fmt.Println("Nothing to return")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - Item not found"))
		return
	case nil:
		fmt.Println(avail)
		json.NewEncoder(w).Encode(avail)
	default:
		panic(selecterr)
	}
}

//used this method for initial version without database
func GetAvailabilities(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range availabilites {
		if item.ItemID == params["itemid"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

}

func main() {
	/*availabilites = append(availabilites, Availability{ItemID: "342146", UnitOfMeasure: "EACH", Onhand: 40, Demand: 20, AvailableToPromiseQty: 20})
	availabilites = append(availabilites, Availability{ItemID: "234224", UnitOfMeasure: "EACH", Onhand: 60, Demand: 20, AvailableToPromiseQty: 40})*/

	router := mux.NewRouter()
	router.HandleFunc("/availabilities/{itemid}", GetATPData).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
