package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"

	"database/sql"

	"github.com/gorilla/mux"

	"github.com/stclaird/hostedby/pkg/model"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	log.Println("Entering health check end point")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "API OK")
}

func Find(w http.ResponseWriter, r *http.Request) {
	log.Println("Entering Find IP in CIDR")

	Data := model.CidrObject{
		Net:           "Null",
		Start_ip:      0,
		End_ip:        0,
		Url:           "Null",
		Cloudplatform: "Null",
		Iptype:        "Null",
		Error:         "Null",
	}

	v := r.URL.Query()

	ipstr := v.Get("ip")
	ip := net.ParseIP(ipstr)
	ipdec, err := ipv4toDecimal(ip)

	if err != nil {
		Data.Error = fmt.Sprint("", err)

	} else {
		model.DB, _ = sql.Open("sqlite3", "cloudIP.sqlite3.db")
		Data = model.IpinCidr(model.DB, ipdec)
		Data.Error = "None"
	}

	jData, _ := json.Marshal(Data)

	if Data.Net == "" {
		w.WriteHeader(http.StatusNotFound)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jData)
}

func main() {
	log.Println("Starting API server")
	//Create a new router
	router := mux.NewRouter()
	log.Println("creating routes")
	//Create Endpoints
	router.HandleFunc("/health", HealthCheck).Methods("GET")
	router.HandleFunc("/findip", Find).Methods("GET")

	http.Handle("/", router)

	//Listen to requests
	http.ListenAndServe(":8080", router)

}
