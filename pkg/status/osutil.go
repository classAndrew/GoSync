package status

import (
	"log"
	"os"
	"path/filepath"
	"time"
)

// StartTime starting time
var StartTime time.Time

// GetFilesSize return total size of dir
func GetFilesSize(dir string) int {
	size := 0
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		size += int(info.Size())
		return nil
	})
	if err != nil {
		log.Fatal(err.Error())
	}
	return size
}

// ListFilesDir returns a slice of all filenames in a dir and file number count
func ListFilesDir(dir string) (int, []string) {
	var name []string
	var num int
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			path += "/"
		} else {
			num++
		}
		name = append(name, path)
		return nil
	})
	return num, name
}

// GetUptime get how long the file server has been running
func GetUptime() time.Duration {
	return time.Since(StartTime)
}
