package database

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	migrateSql "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewConnMysql() *gorm.DB {
	fmt.Println("Starting database MySql ....")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPwd := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	if dbHost == "" || dbPort == "" || dbUser == "" || dbPwd == "" || dbName == "" {
		log.Fatalln("Invalid data mysql")
	}

	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPwd, dbHost, dbPort, dbName)

	conn, err := gorm.Open(mysql.Open(connString), &gorm.Config{})
	if err != nil {
		log.Fatalln("Failed connection database MySql")
	}

	fmt.Println("Successfully connect to DB :) ")

	dir, _ := os.Getwd()
	migrationsPath := dir + "/database/migrations"

	sqlDB, _ := conn.DB()
	driver, err := migrateSql.WithInstance(sqlDB, &migrateSql.Config{})
	if err != nil {
		log.Fatalf("Error creating MySQL driver instance: %v", err)
	}

	// Create a new migration instance
	m, err := migrate.NewWithDatabaseInstance("file://"+migrationsPath, "mysql", driver)
	if err != nil {
		log.Fatalf("Error creating migration instance: %v", err)
	}

	// Run the migration to the latest version
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Error applying migration: %v", err)
	}

	return conn
}
