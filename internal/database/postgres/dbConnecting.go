package postgres

import (
	"database/sql"
	"fmt"
	"log"
)

// TODO: Вынести в pkg вместе с конфигурационным файлом для переиспользования
func ConnectDB() *sql.DB {
	c := NewConfig()

	psqlInfo := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", c.User, c.Password, c.Host, c.Port, c.DBname)

	fmt.Println(psqlInfo) ///{postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable}

	db, err := sql.Open("postgres", psqlInfo)
	if err!= nil {
		log.Fatal("Failed to connect database")
	}

	err = db.Ping()
	if err!= nil {
		log.Fatal("Failed to ping database")
	}

	log.Println("Connected to database")
	return db
}

