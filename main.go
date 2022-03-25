package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	fmt.Printf("Starting server...")

	http.HandleFunc("/", helloworld)
	http.HandleFunc("/health_check", check)

	http.ListenAndServe(":3000", nil)
}

// helloworld just displays a banner message for testing
func helloworld(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("FIRSTNAME")
	status := http.StatusOK
	w.WriteHeader(status)
	w.Write([]byte(fmt.Sprintf(`
	<!DOCTYPE html>
	<html lang="en">
		<head><title>Hello %s</title></head>
		<body><h1>Hello %s!</h1></body>
	</html>
	`, name, name)))
}

func check(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`<h1>Health check</h1>`))
}
