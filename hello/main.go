package main

import (
	"fmt"
	"os"
)


func main() {
    fmt.Println("Hello, World!")
    var secretValue string
	var err error

    secretValue, err = fetchAzureSecret()

	if err != nil {
		fmt.Println("Error fetching secret:", err)
		os.Exit(1)
	}

	writeOutput(secretValue)

	if err != nil {
		fmt.Println("Error writing secret to file:", err)
		os.Exit(1)
	}
}

func writeOutput(output string) {
	f, err := os.Create("/tmp/secret")
	if err != nil {
		return
	}
	defer f.Close()

	f.WriteString(output)
	fmt.Println("Secret successfully fetched and saved to /tmp/secret")
}