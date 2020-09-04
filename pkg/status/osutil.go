package status

import (
	"log"
	"os"
	"path/filepath"
)

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
