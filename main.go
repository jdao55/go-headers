package main

import (
	"fmt"
	"os"
	"net/http"
	"time"
	"bytes"
	"text/tabwriter"
	"sort"
	"flag"
	"strings"
)

func checkRedirect(req *http.Request, via []*http.Request) error {
	return http.ErrUseLastResponse
}

func formatLine(str string, max_len int) string{
	str_arr := strings.Split(str, ";")
	var rstr string = ""
	for _, tstr :=range(str_arr) {
		tstr = strings.Trim(tstr," \t")
		if len(tstr)> max_len  {
			tstr = splitString(tstr, max_len)
		}
		rstr+=(tstr+"\n\t")
	}
	return rstr[:len(rstr)-2]
	}

//splits longs stirng with new lines
func splitString(str string, n int) string {
	str2:=""
	for i:=0; i< len(str)/n; i++{
		str2+=str[n*i:n*(i+1)]
		if i < (len(str)/n -1) {
			str2+="\n>>>\t"
		}

	}
	return str2

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
			fmt.Fprintf(tab_writer, "%s\t%s\n", header_key, formatLine(headerValue,60))
		}
	}

	tab_writer.Flush()
	fmt.Println(buffer.String())
}



func main() {

	gzip_ptr:=flag.Bool("gzip", false, "turn on gzip")
	flag.Parse();

	var client = http.Client{
		CheckRedirect: checkRedirect,
		Transport: &http.Transport{
			DisableCompression: true,
		},
	}



	url := flag.Args()[0]
	req, err := http.NewRequest("GET", url, nil)
	if err!=nil {
		os.Exit(1)
	}

	if (*gzip_ptr) {
		req.Header.Add("Accept-Encoding", "gzip, deflate")
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
