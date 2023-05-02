package lib

import (
	"net/http"
)

func LocalHost() int {
	if http.ListenAndServe(":8000", nil) != nil {
		return 1
	}

	return 0
}
