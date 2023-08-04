package main

import (
	"database/sql"
	_ "dm"
	"fmt"
)

func main() {
	var db *sql.DB
	var err error

	driverName := "dm"
	dataSourceName := "dm://SYSDBA:SYSDBA@localhost:8888"

	if db, err = sql.Open(driverName, dataSourceName); err != nil {
		fmt.Println(err)
	}

	if err = db.Ping(); err != nil {
		fmt.Println(err)
	}

	var sql = "SELECT * FROM TEST"
	rows, err := db.Query(sql)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	var id int
	for rows.Next() {
		if err = rows.Scan(&id); err != nil {
			fmt.Println(err)
		}
		fmt.Printf("id: %v\n", id)
	}

	if err = db.Close(); err != nil {
		fmt.Printf("db close failed: %s.\n", err)
	}
}
