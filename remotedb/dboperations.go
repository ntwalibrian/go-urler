package remotedb

import (
	"context"
	"fmt"
	"log"
)

func DBwrite(shortKey string, url string) {
	db := Init()
	defer db.Close(context.Background())

	query := `INSERT INTO urls (key,url) VALUES ($1,$2)`

	_,err := db.Exec(context.Background(), query, shortKey, url)
	if err != nil {
		log.Fatalf("Error inserting into database: %v" , err)
	}

	fmt.Println("URL saved successfully: " + url + "--->" + shortKey )

}

func DBread() map[string]string{
	db := Init()
	defer db.Close(context.Background())

	query := `SELECT key,url FROM urls`
	results := make(map[string]string)

	rows,err := db.Query(context.Background(), query)
	if err != nil {
		log.Fatalf("Error executing in database: %v" , err)
	}
	defer rows.Close()

	for rows.Next() {
		var shortKey string
		var url string

		if err := rows.Scan(&shortKey, &url); err != nil {
			log.Fatalf("Error scanning: %v" , err)
		}

		results[shortKey] = url
	}

	// Check for any errors that occurred during iteration
	if err := rows.Err(); err != nil {
		log.Fatalf("Error iterating over rows: %v", err)
	}

	return results
}