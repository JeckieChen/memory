package util

import (
	"log"
	"os"
)

func GetPwd() (dir string, err error) {
	d, e := os.Getwd()
	if e != nil {
		log.Fatal(e)
		return "", nil
	}
	return d, e
}
