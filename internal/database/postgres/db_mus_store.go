package postgres

import (
	"database/sql"
	"fmt"
	"log"
)

func ConnectDB() *sql.DB {
	c := NewConfig()

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", c.Host, c.Port, c.User, c.Password, c.DBname)

	db, err := sql.Open("postgres", psqlInfo)
	if err!= nil {
		log.Fatal("Failed to connect database")
	}
	//db.Close()

	err = db.Ping()
	if err!= nil {
		log.Fatal("Failed to ping database")
	}

	log.Println("Connected to database")
	return db
}