package main

// load balancer design struct
// user configuration from CLI (read)
// load distribution algorithm
import (
	"fmt"
	"net/http"
)

type LoadBalancer struct {
	name            string
	ip_address      string
	algorithm       string
	backend_servers []string
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, World!")
	})

	port := 8080
	fmt.Printf("Server is listening on port %d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
