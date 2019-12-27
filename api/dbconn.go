package main

import (
	"bufio"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func parseCreds(fn string) string {
	infile, err := os.Open(fn)
	if err != nil {
		log.Fatal(err)
	}

	filescanner := bufio.NewScanner(infile)

	var creds string
	for filescanner.Scan() {
		creds = filescanner.Text()
	}
	infile.Close()
	return creds
}

func connectPSQL() *sql.DB {

	creds := parseCreds("../../dbcreds.config")

	db, err := sql.Open("postgres", creds)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
