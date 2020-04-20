package api

import (
	"log"
)

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func dbError(err error) bool{
	if err != nil{
		return true
	}
	return false
}