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

//splits longs stirng with new lines
func splitString(str string, n int) string {
	for i:=1; i< len(str)/n+1; i++{
		str=str[:n*i+i-1]+"\n\t"+str[n*i+i:];
	}
	return str
	
}

//prints header info
func PrintHeader(resp *http.Header) {
	var (
		buffer bytes.Buffer
		tab_writer = tabwriter.NewWriter(&buffer, 0, 0, 2, ' ', 0)
	)
	//sort headers
	header_list := make([]string, 0, len(*resp))
	for header_key := range *resp{
		header_list = append(header_list, header_key)
	}
	sort.Strings(header_list)
	
	
	for _, header_key := range header_list[:] {
		for _, headerValue := range (*resp)[header_key] {
			fmt.Fprintf(tab_writer, "%s\t%s\n", header_key, splitString(headerValue,60))
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
