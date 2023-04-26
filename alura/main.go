package main

import (
	"net/http"
)

func main() {
	LocalHost()
}

func LocalHost() int {
	if http.ListenAndServe(":8000", nil) != nil {
		return 1
	}

	return 0
}
