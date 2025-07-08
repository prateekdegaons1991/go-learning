package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

const (
	// URL is the URL to send the request to
	URL = "https://dummy-json.mock.beeceptor.com/todos"
)

func main() {

	fmt.Println("WebRequest in Go")

	response, err := http.Get(URL)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close() //caller's responsibility to close the response body
	fmt.Println("Response Status:", response.Status)
	fmt.Printf("Response is of Type : %T\n", response)

	data, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	responseData := string(data)
	fmt.Println("Response Data:", responseData)

	writeResponseToFile(responseData, "response.json")

}

// Function to write data to a file
func writeResponseToFile(responseData string, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.WriteString(responseData)
	if err != nil {
		panic(err)
	}

	fmt.Println("Response written to file:", filename)
}
