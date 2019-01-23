package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"gopkg.in/resty.v1"
)

func callbackRequest(r *http.Request) string {
	cbAddress := os.Getenv("CALLBACK_URL")
	res := ""
	if len(cbAddress) == 0 {
		return res
	}

	req_id := r.Header.Get("x-request-id")
	trace_id := r.Header.Get("x-b3-traceid")
	span_id := r.Header.Get("x-b3-spanid")
	parent_span_id := r.Header.Get("x-b3-parentspanid")
	sampled := r.Header.Get("x-b3-sampled")
	flags := r.Header.Get("x-b3-flags")
	span_context := r.Header.Get("x-ot-span-context")

	resp, err := resty.R().SetHeader("x-request-id", req_id).SetHeader("x-b3-traceid", trace_id).SetHeader("x-b3-spanid", span_id).SetHeader("x-b3-parentspanid", parent_span_id).SetHeader("x-b3-sampled", sampled).SetHeader("x-b3-flags", flags).SetHeader("x-ot-span-context", span_context).Get(cbAddress)
	
	if err != nil {
		fmt.Printf("[Callback] Error on GET %s - %s", cbAddress, err)
	} else {
		res = fmt.Sprintf("Code: %d - Body: %s", resp.StatusCode(), resp.String())
		fmt.Printf("[Callback Response] %s", res)
	}
	return res
}

func responseCallback(w http.ResponseWriter, r *http.Request) {
	content := os.Getenv("CONTENT")
	if len(content) == 0 {
		content = "Hello from tiny mock server"
	}

	code := os.Getenv("CODE")
	if len(code) == 0 {
		code = "200"
	}

	callbackRequest(r)
	i, _ := strconv.Atoi(code)
	w.WriteHeader(i)
	fmt.Fprintf(w, content) // send data to client side
}

func pingExternalService() {
	pingHost := os.Getenv("PING_HOST")
	if len(pingHost) == 0 {
		return
	}
	pingInterval := os.Getenv("PING_INTERVAL")
	interval := 1
	if len(pingInterval) > 0 {
		interval, _ = strconv.Atoi(pingInterval)
	}

	for true {
		response, err := http.Get(pingHost)
		if err != nil {
			fmt.Printf("[Polling] Error on GET %s - %s", pingHost, err)
		} else {
			defer response.Body.Close()
			contents, err := ioutil.ReadAll(response.Body)
			if err != nil {
				fmt.Printf("[Polling] Err reading response %s", err)
			}
			fmt.Printf("[Polling] [%d] %s\n", response.StatusCode, string(contents))
		}
		time.Sleep(time.Duration(interval) * time.Second)
	}
}

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "9090"
	}
	go pingExternalService()
	fmt.Printf("Started on port: %s\n", port)
	http.HandleFunc("/", responseCallback)                    // set router
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
