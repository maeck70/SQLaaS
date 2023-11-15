package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type ActiveToken struct {
	Token        string    `db:"token"`
	Type         int       `db:"type"`
	LastSeenDttm time.Time `db:"last_seen_dttm"`
}

func (ds DaasService) dbServiceGet(w http.ResponseWriter, r *http.Request) {
	var parms []interface{}

	// Build the parameter list from the URL
	for _, p := range ds.Arguments {
		parms = append(parms, mux.Vars(r)[p.Name])
	}

	// Execute SQL
	qr, err := db.Queryx(ds.Sql, parms...)
	FailOnError(err, fmt.Sprintf("Unable to execute SQL: %s", ds.Sql))

	// Build the result
	var result []map[string]interface{}
	for qr.Next() {
		rowMap := make(map[string]interface{})
		err := qr.MapScan(rowMap)
		FailOnError(err, "Unable to scan row")

		// Convert string byte field values to strings
		for key, value := range rowMap {
			switch ds.Fields[key].Type {
			case "string":
				if b, ok := value.([]byte); ok {
					rowMap[key] = string(b)
				}
			case "int":
				if b, ok := value.([]byte); ok {
					num, err := strconv.Atoi(string(b))
					if err != nil {
						num = 0
					}
					rowMap[key] = num
				}
			}
		}

		log.Printf("Found row %+v\n", rowMap)
		result = append(result, rowMap)
	}
	qr.Close()

	log.Printf("Found %d rows\n", len(result))

	// Return Json for the result
	JsonResponse(w, result)
}
