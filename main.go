package main

import (
	"net/http"
	"fmt"
	"os"
	"bufio"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "Only one argument expected but %d given.\n", len(args))
		os.Exit(2)
	}
	url := args[0]

	response, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get url: %s", err.Error())
		os.Exit(1)
	}
	defer response.Body.Close()

	scanner := bufio.NewScanner(bufio.NewReader(response.Body))
	scanner.Split(bufio.ScanBytes)
	var body_size int = 0
	for scanner.Scan() {
		body_size += len(scanner.Bytes())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read response: %s", err.Error())
		os.Exit(1)
	}
	fmt.Println(body_size)
}
