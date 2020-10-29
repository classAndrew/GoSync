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
	uptime := status.GetUptime().Hours()
	w.Write([]byte(strconv.FormatFloat(uptime, 'f', 8, 64) + " Hours up\n"))
	w.Write([]byte(serverinfo.Server.RootDirectory + " Root with " + strconv.Itoa(stat) + " Bytes Total\n"))
	count, names := status.ListFilesDir(serverinfo.Server.RootDirectory)
	w.Write([]byte(strconv.Itoa(count) + " Total Files\n\n"))
	for i := 1; i < len(names); i++ {
		w.Write([]byte(names[i] + "\n"))
	}
}
