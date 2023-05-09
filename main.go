package main

import (
	"database/sql"
	"log"
	"sync"
	"time"

	"github.com/go-sql-driver/mysql"
)

func main() {
	db, err := connectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	const numOfConns = 10
	db.SetMaxIdleConns(numOfConns)
	db.SetMaxOpenConns(numOfConns)
	// notice:
	//  Comment out `SetConnMaxLifetime` once
	// db.SetConnMaxLifetime(time.Second * 1)

	var wg sync.WaitGroup
	for i := 0; i < numOfConns; i++ {
		wg.Add(1)
		go func() {
			rows, err := db.Query("SELECT * from users")
			if err != nil {
				log.Fatal(err)
			}
			if err := rows.Close(); err != nil {
				log.Fatal(err)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	log.Printf("idle conns: %d", db.Stats().Idle)

	for {
		rows, err := db.Query("SELECT * from users")
		if err != nil {
			log.Fatal(err)
		}
		if err := rows.Close(); err != nil {
			log.Fatal(err)
		}
		log.Print("OK")
		// notice:
		//  if you want to make an error, change the value to `wait_timeout` or higher.
		//  ex: time.Sleep(3 * time.Second)
		time.Sleep(1 * time.Second)
	}
}

func connectDB() (*sql.DB, error) {
	c := mysql.Config{
		DBName:    "db",
		User:      "user",
		Passwd:    "password",
		Addr:      "localhost:23306",
		Net:       "tcp",
		ParseTime: true,
		// https://github.com/go-sql-driver/mysql/blob/master/CHANGELOG.md#version-14-2018-06-03
		AllowNativePasswords: true,
	}

	db, err := sql.Open("mysql", c.FormatDSN())
	if err != nil {
		return nil, err
	}

	return db, nil
}
