package controller

import (
	"fmt"
	"net/http"
)

// Index display landing page
func Index() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "ok")
	}
}
