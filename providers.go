package emailproviders

import (
	"log"
	"strings"
)

//go:generate go run generate/main.go generate/domains.txt domains.go

func Exists(email string) bool {
	at := strings.LastIndexByte(email, '@')
	if at <= 0 || len(email[at:]) < 3 {
		log.Printf("false")
		return false
	}

	_, exists := publicProviders[email[at+1:]]
	return exists
}
