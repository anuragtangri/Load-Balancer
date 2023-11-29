package main

// load balancer design struct
// user configuration from CLI (read)
// load distribution algorithm (round robin, weighted round robin , least connection)

// define api end point to get the request ?
// distribute the request to a bunch of backend servers ?

import (
	"fmt"
	"io"
	"net/http"
)

type LoadBalancer struct {
	name            string
	ip_address      string
	algorithm       string
	backend_servers []string
}

func forwardRequest(w http.ResponseWriter, r *http.Request, host string) {
	// Create a new URL by copying the original request URL
	newURL := *r.URL

	// Update the scheme and host of the new URL to point to the target server
	newURL.Scheme = "http"
	newURL.Host = host // Replace with the target server's address

	// Create a new request using the modified URL
	fmt.Println("the method is", r.Method)
	newRequest, err := http.NewRequest(r.Method, newURL.String(), r.Body)
	if err != nil {
		http.Error(w, "Error creating new request", http.StatusInternalServerError)
		return
	}

	// Copy headers from the original request to the new request
	newRequest.Header = make(http.Header)
	copyHeaders(newRequest.Header, r.Header)

	// Create an HTTP client
	client := http.Client{}

	// Send the new request to the target server
	response, err := client.Do(newRequest)
	if err != nil {
		http.Error(w, "Error forwarding request", http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()

	// Copy the response from the target server to the original response writer
	copyHeaders(w.Header(), response.Header)
	w.WriteHeader(response.StatusCode)
	io.Copy(w, response.Body)
}

func copyHeaders(dest http.Header, src http.Header) {
	for key, values := range src {
		for _, value := range values {
			dest.Add(key, value)
		}
	}
}

func distributeReq(w http.ResponseWriter, r *http.Request) {
	// you can add more server
	servers := [2]string{"localhost:5000", "localhost:5001"}
	// forwardRequest(w, r, servers[0])
	forwardRequest(w, r, servers[1])
	// idx := 0
	// for {
	// 	fmt.Println(servers[idx])
	// 	forwardRequest(w, r, servers[idx])
	// 	idx = (idx + 1) % 2
	// 	time.Sleep(5000)
	// }
}

func main() {

	// forwarding request is find
	// how to round robin ?
	http.HandleFunc("/", distributeReq)

	port := 8080
	fmt.Printf("Server is listening on port %d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
