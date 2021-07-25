package utils

import (
	"fmt"
	"os"
	"time"
)

// OpenFile opens and closes file 10 times.
func OpenFile(path string) error {
	for i := 0; i < 10; i++ {
		file, err := os.Open(path)
		if err != nil {
			fmt.Printf("error during opening %v", path)
			return err
		}

		fmt.Printf("file has been opened %v", path)

		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				fmt.Printf("error during closing %v", err)
			}
		}(file)

		time.Sleep(1 * time.Second)
	}

	return nil
}
