package graphqltest

import (
	"bufio"
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

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
//	creds := parseCreds("../../dbcreds.config")

	var err error

	db, err = sql.Open("postgres", creds)
	CheckError(err)
}

func ClosePSQL() {
	db.Close()
}
