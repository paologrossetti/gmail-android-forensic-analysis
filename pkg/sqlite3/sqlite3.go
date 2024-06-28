package sqlite3

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func open(filename string) *sql.DB {

	conn, err := sql.Open("sqlite3", filename)
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}
	return conn
}

func createMap(rows *sql.Rows, columns []string) []map[string]interface{} {
	// for each database row / record, a map with the column names and row values is added to the allMaps slice
	var results []map[string]interface{}

	for rows.Next() {
		values := make([]interface{}, len(columns))
		pointers := make([]interface{}, len(columns))
		for i, _ := range values {
			pointers[i] = &values[i]
		}
		err := rows.Scan(pointers...)
		if err != nil {
			log.Fatal(err)
		}
		resultMap := make(map[string]interface{})
		for i, val := range values {
			resultMap[columns[i]] = val
		}
		results = append(results, resultMap)
	}
	return results
}

func SelectAll(filename string, table string) []map[string]interface{} {
	conn := open(filename)

	rows, err := conn.Query(fmt.Sprintf("SELECT * FROM %s", table))
	if err != nil {
		log.Fatal(err)
	}

	columns, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}

	var results []map[string]interface{} = createMap(rows, columns)

	defer rows.Close()
	defer conn.Close()

	return results
}

func SelectAllWhere(filename string, table string, columncondition string, value interface{}) []map[string]interface{} {
	conn := open(filename)

	rows, err := conn.Query(fmt.Sprintf("SELECT * FROM %s WHERE %s = ?", table, columncondition), value)
	if err != nil {
		log.Fatal(err)
	}

	columns, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}

	var results []map[string]interface{} = createMap(rows, columns)

	defer rows.Close()
	defer conn.Close()

	return results
}

func SelectSome(filename string, table string, columns []string) []map[string]interface{} {
	conn := open(filename)

	var selectString = strings.Join(columns, ", ")

	rows, err := conn.Query(fmt.Sprintf("SELECT %s FROM %s", selectString, table))
	if err != nil {
		log.Fatal(err)
	}

	var results []map[string]interface{} = createMap(rows, columns)

	defer rows.Close()
	defer conn.Close()

	return results
}

func SelectSomeWhere(filename string, table string, columns []string, columncondition string, value interface{}) []map[string]interface{} {
	conn := open(filename)

	var selectString = strings.Join(columns, ", ")

	rows, err := conn.Query(fmt.Sprintf("SELECT %s FROM %s WHERE %s = ?", selectString, table, columncondition), value)
	if err != nil {
		log.Fatal(err)
	}

	var results []map[string]interface{} = createMap(rows, columns)

	defer rows.Close()
	defer conn.Close()

	return results
}
