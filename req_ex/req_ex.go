package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
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

func post() {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	request := map[string]interface{}{
		"id":   1,
		"type": "create",
		"data": "test data",
	}

	if err := enc.Encode(request); err != nil {
		fmt.Printf("error: can't encode - %s\n", err)
		return
	}

	resp, err := http.Post("http://httpbin.org/post", "application/json", &buf)
	if err != nil {
		fmt.Printf("error post: %s\n", err)
		return
	}

	io.Copy(os.Stdout, resp.Body)
}

func auth() {
	user, passwd := "bugs", "duck season"
	url := fmt.Sprintf("http://httpbin.org/basic-auth/%s/%s", user, passwd)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("error auth: %s\n", err)
		return
	}

	req.SetBasicAuth(user, passwd)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("error auth: %s\n", err)
		return
	}

	io.Copy(os.Stdout, resp.Body)
}

func main() {
	// simpleGet()
	// timeoutGet()
	// post()
	auth()
}
