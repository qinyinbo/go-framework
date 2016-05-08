package storage

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

var sqlDb *sql.DB

func Connect(config string) (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:pass@tcp(127.0.0.1:3306)/user_db")
	return db, err
}

func GetRow(db *sql.DB, sqlQuery string, args ...interface{}) (map[string]string, error) {

	rows, err := db.Query(sqlQuery)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}
	// Make a slice for the values
	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	data := make(map[string]string) //map must be maked ?
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}
		var value string
		for i, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			data[strings.Title(columns[i])] = value
		}
	}
	if err = rows.Err(); err != nil {
		panic(err.Error())
	}
	fmt.Println(data)
	return data, nil

}
func GetAll(db *sql.DB, sqlQuery string, args ...interface{}) ([]map[string]string, error) {
	rows, err := db.Query(sqlQuery)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}
	// Make a slice for the values
	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	allData := make([]map[string]string, 0)
	data := make(map[string]string) //map must be maked ?
	ii := 0
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}
		var value string
		for i, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			data[strings.Title(columns[i])] = value
		}
		allData = append(allData, data)
		ii++
	}
	if err = rows.Err(); err != nil {
		panic(err.Error())
	}
	fmt.Println(allData)
	return allData, nil

}
func Insert(db *sql.DB, sqlQuery string, args ...interface{}) int64 {

	res, err := db.Exec(sqlQuery, args...)
	if err != nil {
		panic(err.Error())
	}

	id, err := res.LastInsertId()
	if err != nil {
		panic(err.Error())
	}
	return id

}
func Exec(db *sql.DB, sqlQuery string, args ...interface{}) (int64, error) {

	res, err := db.Exec(sqlQuery, args...)
	if err != nil {
		panic(err.Error())
	}
	count, err := res.RowsAffected()
	if err != nil {
		panic(err.Error())
	}
	return count, err
}
