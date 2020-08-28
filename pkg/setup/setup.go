package setup

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
)

// CheckSetup to check if the server has been set up already
func CheckSetup() bool {
	if _, err := os.Stat("./config.json"); err != nil {
		log.Println("Fileserver not setup yet!")
		return false
	}
	return true
}

// FirstSetup to setup server
func FirstSetup() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to the first time server setup.")
	fmt.Print("Root directory of file server: ")
	rootdir, _ := reader.ReadString('\n')
	if _, err := os.Stat(rootdir[:len(rootdir)-1]); err != nil {
		log.Fatal("Entered directory does not exist!")
	}
	log.Print("Entered: ", rootdir)
	fmt.Print("Enter the port: ")
	port, _ := reader.ReadString('\n')
	portint, _ := strconv.Atoi(port[:len(port)-1])
	f, _ := os.Create("config.json")
	defer f.Close()
	fmt.Println(portint)
	marshalled, _ := json.Marshal(struct {
		RootDir string
		Port    int
	}{rootdir[:len(rootdir)-1], portint})
	f.Write(marshalled)

}
