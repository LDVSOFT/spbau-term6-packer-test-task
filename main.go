package main

import (
	"net/http"
	"fmt"
	"os"
	"bufio"
	"io/ioutil"
	"bytes"
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
	body, err := ioutil.ReadAll(bufio.NewReader(response.Body))
	response.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading response: %s", err.Error())
		os.Exit(1)
	}

	body_size := len(body)
	body_words := 0
	scanner := bufio.NewScanner(bytes.NewReader(body))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		body_words++
	}

	fmt.Printf("%d %d\n", body_size, body_words)
}
