package fixtures

import (
	"os"
)

// fixtures is a slice of fixture functions to be executed
var fixtures = []func(string) error{
	LoadTodos,
}

// Checks if the SQLite database file exists and loads fixtures if it doesn't.
func CheckAndLoadFixtures(filePath string) error {
	// Check if the SQLite file exists
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		// If the file doesn't exist, execute the fixture functions
		for _, fixture := range fixtures {
			err := fixture(filePath)
			if err != nil {
				return err
			}
		}
	} else if err != nil {
		return err
	}

	return nil
}
