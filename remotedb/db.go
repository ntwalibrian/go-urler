package remotedb

import (
	"context"
	"log"
	
	"github.com/jackc/pgx/v5"
	
)

//initialize supabase connection 
func Init() *pgx.Conn {
	// remove this put it in an env file
	connStr := "postgresql://postgres.wgqofcnauzlwfkwyqchr:nzN61UpIw0Q8Innk@aws-0-eu-central-1.pooler.supabase.com:5432/postgres"

	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	// Example query to test connection
	var version string
	if err := conn.QueryRow(context.Background(), "SELECT version()").Scan(&version); err != nil {
		log.Fatalf("Query failed: %v", err)
	}
	log.Println("Connected to:", version)

	return conn
}