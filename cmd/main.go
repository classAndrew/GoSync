package main

import (
	"github.com/classAndrew/GoSync/pkg/startup"
)

func main() {
	var server startup.GoSyncServer
	startup.StartServer(&server)
}
