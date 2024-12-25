package emailproviders

//go:generate go run generate/main.go generate/domains.txt domains.go

func Exists(email string) bool {
	at := lastIndexByte(email, '@')
	if at <= 0 || len(email[at:]) < 3 {
		return false
	}

	_, exists := publicProviders[email[at+1:]]
	return exists
}

// lastIndexByte returns the index of the last instance of c in s, or -1 if c is not present in s.
func lastIndexByte(s string, c byte) int {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == c {
			return i
		}
	}
	return -1
}
