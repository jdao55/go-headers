package main

import (
	"fmt"
	"os"
	"net/http"
	"time"
	"bytes"
	"text/tabwriter"
	"sort"
	//"strings"
)

func checkRedirect(req *http.Request, via []*http.Request) error {
	return http.ErrUseLastResponse
}


func PrintHeader(resp *http.Header) {
	var (
		buffer bytes.Buffer
		tab_writer = tabwriter.NewWriter(&buffer, 0, 0, 2, ' ', 0)
	)
	header_list := make([]string, 0, len(*resp))
	for header_key := range *resp{
		header_list = append(header_list, header_key)
	}

	sort.Strings(header_list)

	
	for _, header_key := range header_list[:] {
		for _, headerValue := range (*resp)[header_key] {

			if len(headerValue)>60 {
				fmt.Fprintf(tab_writer, "%s\t%s\n", header_key, headerValue[:60])
				headerValue = headerValue[60:]
				for len(headerValue) > 60 {
					fmt.Fprintf(tab_writer, "\t%s\n", headerValue[:60])
					headerValue = headerValue[60:]
				}
				fmt.Fprintf(tab_writer, "\t%s\n",  headerValue)
			} else {
					fmt.Fprintf(tab_writer, "%s\t%s\n", header_key, headerValue)
			}
			
}
	}

	tab_writer.Flush()
	fmt.Println(buffer.String())
}



func main() {

	var client = http.Client{
		CheckRedirect: checkRedirect,
		Transport: &http.Transport{
			DisableCompression: true,
		},
	}
	
	url := os.Args[1]
	req, err := http.NewRequest("GET", url, nil)
	if err!=nil {
		os.Exit(1)
	}

	start_time := time.Now()
	resp, err := client.Do(req)
	req_duration := time.Since(start_time)

	if err!=nil {
		os.Exit(1)
	}
	fmt.Println("");
	fmt.Println("GET ", url)
	fmt.Println("Took: ", req_duration)
	fmt.Println(resp.Proto, resp.Status)
	PrintHeader(&(resp.Header))
	os.Exit(0)	
}
