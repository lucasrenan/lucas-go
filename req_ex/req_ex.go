package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Response struct {
	Origin  string `json:"origin"`
	URL     string `json:"url"`
	Headers struct {
		UserAgent string `json:"User-Agent"`
	} `json:"headers"`
}

func simpleGet() {
	resp, err := http.Get("http://httpbin.org/get")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	fmt.Printf("status: %d\n", resp.StatusCode)
	// io.Copy(os.Stdout, resp.Body)

	var rep Response
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// if err := json.Unmarshal([]byte(body), &rep); err != nil {
	// 	log.Fatal(err)
	// }

	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&rep); err != nil {
		log.Fatalf("error: can't decode - %s", err)
	}

	fmt.Println(rep)
}

func timeoutGet() {
	req, err := http.NewRequest("GET", "http://httpbin.org/get", nil)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	req = req.WithContext(ctx)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("with timeout status: %d\n", resp.StatusCode)
}

func main() {
	simpleGet()
	timeoutGet()
}
