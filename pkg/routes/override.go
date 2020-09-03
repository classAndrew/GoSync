package routes

import (
	"net/http"
	"os"
	"strconv"

	"github.com/classAndrew/GoSync/pkg/serverinfo"
)

// OverrideRoute handler
func OverrideRoute(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Overridden route\n"))
}

// StatusRoute gets status of the server and returns as response
func StatusRoute(w http.ResponseWriter, req *http.Request) {
	status, err := os.Stat(serverinfo.Server.RootDirectory)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	w.Write([]byte(status.Name() + "\n" + strconv.Itoa(int(status.Size()))))
}
