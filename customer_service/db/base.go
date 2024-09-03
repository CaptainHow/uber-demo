package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbClient *gorm.DB

func Init() error {
	var err error
	host := "localhost"
	port := "5432"
	user := "postgres"
	password := "postgres"
	dbname := "postgres"
	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", host, port, user, password, dbname)
	
	dbClient, err = gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		return err
	}

	sqlDb, _ := dbClient.DB()
	err = sqlDb.Ping()

	if err != nil {
		return err
	}

	log.Println("Db connection established")
	return nil
}

func GetDb() *gorm.DB {
	return dbClient
}

func CloseDb() {
	conn, _ := dbClient.DB()
	conn.Close()
}