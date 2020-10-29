package startup

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/classAndrew/GoSync/pkg/routes"
	"github.com/classAndrew/GoSync/pkg/serverinfo"
	"github.com/classAndrew/GoSync/pkg/status"

	"github.com/classAndrew/GoSync/pkg/setup"
)

//ServeFiles called by main
func ServeFiles() {
	StartServer(&serverinfo.Server)
}

// StartServer starts server, loads configs
func StartServer(server *serverinfo.GoSyncServer) {
	if !setup.CheckSetup() {
		setup.FirstSetup()
	}
	contents, _ := ioutil.ReadFile("./config.json")
	json.Unmarshal(contents, &serverinfo.Server)
	log.Println(serverinfo.Server.RootDirectory)
	status.StartTime = time.Now()
	http.HandleFunc("/override", routes.OverrideRoute)
	http.HandleFunc("/status", routes.StatusRoute)
	http.Handle("/", http.FileServer(http.Dir(serverinfo.Server.RootDirectory)))

	log.Printf("Listening on port %d\n", server.Port)
	err := http.ListenAndServe(":"+strconv.Itoa(serverinfo.Server.Port), nil)

	if err != nil {
		log.Fatal(err)
	}
}
