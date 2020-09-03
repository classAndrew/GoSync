package serverinfo

// GoSyncServer server for gosync files
type GoSyncServer struct {
	Port          int
	RootDirectory string
}

// Server server struct
var Server GoSyncServer
