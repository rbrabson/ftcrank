package ftcdata

import (
	"encoding/json"
	"fmt"

	"github.com/rbrabson/ftcrank/config"
	"github.com/rbrabson/ftcrank/store"
)

// writeFile writes the data to the requested file. The data is unmarshalled into JSON
// prior to being written into the provided filename. `data` should not be a pointer.
func writeFile(filename string, data any) error {
	// Write the events into the proper directory
	path := fmt.Sprintf("%s/%s/raw", config.STORAGE_DIRECTORY, config.FTC_SEASON)
	file := fmt.Sprintf("%s/%s", path, filename)
	bytes, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return err
	}
	return store.WriteFile(path, file, bytes)
}

// readFile read the data from the requested file, with the bytes read returned.
func readFile(filename string) ([]byte, error) {
	file := fmt.Sprintf("%s/%s/raw/%s", config.STORAGE_DIRECTORY, config.FTC_SEASON, filename)
	return store.ReadFile(file)
}
