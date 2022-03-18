package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func getFeedback(c *gin.Context) {
	db := connect()

	id := 10
	rows, err := db.Query("SELECT * FROM customerloyalty.feedback WHERE FeedbackId =?", id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil {
		fmt.Println("Failed to get columns", err)
		return
	}

	// Result is your slice string.
	rawResult := make([][]byte, len(cols))
	result := make([]string, len(cols))

	dest := make([]interface{}, len(cols)) // A temporary interface{} slice
	for i, _ := range rawResult {
		dest[i] = &rawResult[i] // Put pointers to each string in the interface slice
	}

	for rows.Next() {
		err = rows.Scan(dest...)
		if err != nil {
			fmt.Println("Failed to scan row", err)
			return
		}

		for i, raw := range rawResult {
			if raw == nil {
				result[i] = "\\N"
			} else {
				result[i] = string(raw)
			}
		}

		fmt.Printf("%#v\n", result)
	}

	c.IndentedJSON(http.StatusOK, gin.H{"System": "abcd1234!"})
}

func getFeelGood(c *gin.Context) {
	db := connect()

	rows, err := db.Query("SELECT * FROM customerloyalty.feel_good")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var data []feelgood

	for rows.Next() {
		var fg feelgood
		err = rows.Scan(&fg.Id, &fg.Name, &fg.Locale)
		if err != nil {
			fmt.Println("Failed to scan row", err)
			return
		}
		data = append(data, fg)
	}

	c.IndentedJSON(http.StatusOK, data)
}

func getCancelReason(c *gin.Context) {
	locale := c.Param("locale")
	db := connect()

	rows, err := db.Query("call ap_cancelreason_sel(?, ?)", locale, nil)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var data []cancelreason

	for rows.Next() {
		var cr cancelreason
		err = rows.Scan(&cr.CancelReasonId, &cr.Locale, &cr.System, &cr.CancelReasonName)
		if err != nil {
			fmt.Println("Failed to scan row", err)
			return
		}
		data = append(data, cr)
	}

	c.IndentedJSON(http.StatusOK, data)
}

func connect() *sql.DB {
	driver := "mysql"
	user := "root"
	password := "admin"
	endpoint := "localhost"
	port := "3306"
	dbName := "customerloyalty"
	charset := "charset=utf8mb4"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", user, password, endpoint, port, dbName, charset)
	fmt.Println(dsn)

	db, err := sql.Open(driver, dsn)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}
