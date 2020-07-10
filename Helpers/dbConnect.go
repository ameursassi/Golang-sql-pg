package helpers

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
)

func GetVarFromEnv() string {
	host := os.Getenv("HOST")
	if host == "" {
		host = "localhost"
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "5432"
	}
	user := os.Getenv("USER")
	if user == "" {
		user = "Ameur"
	}
	password := os.Getenv("PASSWORD")
	if password == "" {
		password = "06205670"
	}
	dbname := os.Getenv("DATABASE")
	if dbname == "" {
		dbname = "DBYuso"
	}
	if os.Getenv("DATABASE_URL") != "" {
		ConnectionString := os.Getenv("DATABASE_URL")
		return ConnectionString
	}
	ConnectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	return ConnectionString
}
func DbConnect() *gorm.DB {
	connection := GetVarFromEnv()
	db, err := gorm.Open("postgres", connection)
	if err != nil {
		log.Fatal(err)
	}
	// defer db.Close()
	return db
}
