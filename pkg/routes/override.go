package routes

import (
	"net/http"
	"strconv"

	"github.com/classAndrew/GoSync/pkg/status"

	"github.com/classAndrew/GoSync/pkg/serverinfo"
)

// OverrideRoute handler
func OverrideRoute(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Overridden route\n"))
}

// StatusRoute gets status of the server and returns as response
func StatusRoute(w http.ResponseWriter, req *http.Request) {
	stat := status.GetFilesSize(serverinfo.Server.RootDirectory)
	w.Write([]byte(serverinfo.Server.RootDirectory + "\n" + strconv.Itoa(stat)))
}
