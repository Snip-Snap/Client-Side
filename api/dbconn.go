package api

import (
	"bufio"
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func parseCreds(fn string) string {
	infile, err := os.Open(fn)
	CheckError(err)

	filescanner := bufio.NewScanner(infile)

	var creds string
	for filescanner.Scan() {
		creds = filescanner.Text()
	}
	infile.Close()
	return creds
}

func ConnectPSQL() {

	creds := parseCreds("/run/secrets")
	//code when you run manually without docker
	//creds := parseCreds("../../dbcreds.config")

	var err error

	DB, err = sql.Open("postgres", creds)
	CheckError(err)
}

func ClosePSQL() {
	DB.Close()
}
