package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	for sc.Scan() {
		url, err := url.ParseRequestURI(sc.Text())
		if err != nil {
			fmt.Fprintf(os.Stderr, "invalid url: %s\n", sc.Text())
			continue
		}

		resp, err := fetchURL(url)
		statusCode := 0
		if resp != nil {
			statusCode = resp.StatusCode
		}

		if err != nil || statusCode != 200 {
			fmt.Printf("%s [%d]: %T %s\n", url, statusCode, err, err)
		}
	}

	if sc.Err() != nil {
		fmt.Printf("error: %s\n", sc.Err())
	}
}

func fetchURL(url *url.URL) (*http.Response, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := http.Client{
		Transport: tr,
		Timeout:   20 * time.Second,
	}

	req, err := http.NewRequest("GET", url.String(), nil)
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
