package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/ErwinSalas/go-eda/services/order-service/pkg/api"
	"github.com/ErwinSalas/go-eda/services/order-service/pkg/app"
	"github.com/ErwinSalas/go-eda/services/order-service/pkg/datastore"
)

func initDb() (*sql.DB, error) {
	dbName := os.Getenv("MYSQL_DATABASE")
	dbUser := os.Getenv("MYSQL_USER")
	dbPassword := os.Getenv("MYSQL_PASSWORD")

	if dbName == "" || dbUser == "" || dbPassword == "" {
		log.Fatal("missing required environment variables")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(mysql:3306)/%s", dbUser, dbPassword, dbName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("error opening connection %v", err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("db ping error : %v", err)
	}

	return db, nil
}

func main() {
	db, err := initDb()

	if err != nil {
		log.Fatal("Error initializing database")
		return
	}
	defer db.Close()

	mysqlDatastore, err := datastore.NewDataStore(db)
	if err != nil {
		fmt.Println("Error creating datastore:", err)
		return
	}

	// client, err := InitSQS()

	// if err != nil {
	// 	log.Fatal("Error initializing SQS")
	// 	return
	// }

	// // queueURL := os.Getenv("QUEUE_URL")
	// fmt.Println("queueURL:", queueURL)

	// queueService := queue.NewSQSService(queueURL, client)
	ctx := context.Background()

	mysqlDatastore.Migrate(ctx)
	app := app.NewApp(mysqlDatastore)

	api.NewRouter(3001, *app)

}
