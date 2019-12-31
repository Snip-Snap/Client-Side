package graphqltest

import (
	"bufio"
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

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

func ConnectPSQL() {

	creds := parseCreds("../../dbcreds.config")
	var err error

	db, err = sql.Open("postgres", creds)
	if err != nil {
		log.Fatal(err)
	}
}

func ClosePSQL() {
	db.Close()
}
