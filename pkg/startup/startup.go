package startup

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/classAndrew/GoSync/pkg/setup"
)

// GoSyncServer server for gosync files
type GoSyncServer struct {
	Port          int
	RootDirectory string
}

// StartServer starts server, loads configs
func StartServer(server *GoSyncServer) {
	if !setup.CheckSetup() {
		setup.FirstSetup()
	}
	contents, _ := ioutil.ReadFile("./config.json")
	json.Unmarshal(contents, server)
	handler := func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("This is an overridden route\n"))
	}
	http.HandleFunc("/override", handler)
	http.Handle("/", http.FileServer(http.Dir(server.RootDirectory)))

	err := http.ListenAndServe(":"+strconv.Itoa(server.Port), nil)
	if err != nil {
		log.Fatal(err)
	}
}