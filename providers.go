package emailproviders

import (
	"log"
	"strings"
)

func Exists(email string) bool {
	log.Printf("exists? %s", email)
	at := strings.LastIndexByte(email, '@')
	if at <= 0 || len(email[at:]) < 3 {
		log.Printf("false")
		return false
	}
	log.Printf("exists? %s", email[at+1:])

	_, exists := publicProviders[email[at+1:]]

	log.Print(exists)

	return exists
}

var publicProviders = map[string]struct{}{
	//go:generate go run generate/main.go
	"gmail.com": struct{}{},
}
