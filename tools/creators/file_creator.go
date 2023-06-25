package creators

import (
	"fmt"
	"os"
)

func CreateFile(fileName string) error {
	_, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("creating file: %w", err)
	}

	return nil
}
