package model

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

type NetExists struct {
	Type     string
	Name     string
	Tbl_name string
	Rootpage string
	Sql      string
}

func NetTableExists(db *sql.DB) bool {
	r := NetExists{}
	sqlStatement := `SELECT * FROM sqlite_master WHERE type = 'table' AND tbl_name = 'net';`

	row := db.QueryRow(sqlStatement)

	switch err := row.Scan(&r.Type, &r.Name, &r.Tbl_name, &r.Rootpage, &r.Sql); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return false
	case nil:
		fmt.Println(r)
		return true
	default:
		panic(err)
	}
}

type CidrObject struct {
	Net           string
	Start_ip      int
	End_ip        int
	Url           string
	Cloudplatform string
	Iptype        string
	Error         string
}

func IpinCidr(db *sql.DB, ipdecimal int) CidrObject {
	//Cidr objects contain a start and end ip in decimal. To find if an ip address is stored

	r := CidrObject{}
	sqlStatement := `SELECT net, start_ip, end_ip, url, cloudplatform, iptype FROM net WHERE start_ip <= $1 AND end_ip >= $1;`

	row := db.QueryRow(sqlStatement, ipdecimal)

	switch err := row.Scan(&r.Net, &r.Start_ip, &r.End_ip, &r.Url, &r.Cloudplatform, &r.Iptype); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(r)
	default:
		panic(err)
	}

	return r
}
