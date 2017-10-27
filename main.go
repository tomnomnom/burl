package main

import (
	"bufio"
	"crypto/tls"
	"flag"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"time"
)

func main() {
	flag.Parse()

	var input io.Reader
	input = os.Stdin

	if flag.NArg() > 0 {
		file, err := os.Open(flag.Arg(0))
		if err != nil {
			fmt.Printf("failed to open file: %s\n", err)
			os.Exit(1)
		}
		input = file
	}

	sc := bufio.NewScanner(input)

	for sc.Scan() {
		url, err := url.ParseRequestURI(sc.Text())
		if err != nil {
			fmt.Printf("invalid url: %s\n", sc.Text())
			continue
		}

		if !resolves(url) {
			fmt.Printf("does not resolve: %s\n", url)
			continue
		}

		resp, err := fetchURL(url)
		if err != nil {
			fmt.Printf("failed to fetch: %s (%s)\n", url, err)
			continue
		}

		if resp.StatusCode != http.StatusOK {
			fmt.Printf("non-200 response code: %s (%s)\n", url, resp.Status)
		}
	}

	if sc.Err() != nil {
		fmt.Printf("error: %s\n", sc.Err())
	}
}

func resolves(u *url.URL) bool {
	addrs, _ := net.LookupHost(u.Hostname())
	return len(addrs) != 0
}

func fetchURL(u *url.URL) (*http.Response, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := http.Client{
		Transport: tr,
		Timeout:   20 * time.Second,
	}

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "burl/0.1")

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	return resp, err
}
