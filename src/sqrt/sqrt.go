package main

import (
	"fmt"
	"math"
	"net/http"
)

func sqrt() string {
	x := 0.0001
	for i := 0; i < 10; i++ {
		x += math.Sqrt(x)
	}
	return "Code.education Rocks!"
}

func HttpServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, sqrt())
}

func main() {
	http.HandleFunc("/", HttpServer)
	http.ListenAndServe(":80", nil)
}
