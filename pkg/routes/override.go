package routes

import (
	"net/http"
)

// OverrideRoute handler
func OverrideRoute(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("overridden route\n"))
}
