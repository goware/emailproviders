package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatal("expecting <input.txt> <output.go> arguments")
	}

	input, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	output, err := os.OpenFile(os.Args[2], os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer output.Close()

	fmt.Fprintln(output, "package emailproviders")
	fmt.Fprintln(output)
	fmt.Fprintln(output, "var publicProviders = map[string]struct{}{")

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		fmt.Fprintf(output, "	\"%s\": struct{}{}, \n", scanner.Text())
	}

	fmt.Fprintln(output, "}")

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
