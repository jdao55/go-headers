package main

import (
	"fmt"
	"os"
	"net/http"
)

func headerToString(header http.Header) string {
	headerString := ""
	return headerString
}


func main() {

	var client = http.Client{
		Transport: &http.Transport{
			DisableCompression: true,
		},
	}
	
	url := os.Args[1]
	req, err := http.NewRequest("GET", url, nil)
	if err!=nil {
		os.Exit(1)
	}
	resp, err := client.Do(req)
	fmt.Println(resp.Header)
	if err!=nil {
		os.Exit(1)
	}
	os.Exit(0)	
}
