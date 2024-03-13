package ftcdata

import (
	"encoding/json"

	"github.com/rbrabson/ftc/ftc"
	"github.com/rbrabson/ftcrank/config"
)

// Save the awards given at a given match. It already includes
// the event code, so just iterate through each event and add
// the results
const (
	AWARDS_FILE_NAME = "awards.json"
)

var (
	Awards []ftc.TeamAward = make([]ftc.TeamAward, len(Events))
)

func RetrieveAwards() error {
	// Iterate over all the events
	for _, event := range Events {
		awards, err := ftc.GetEventAwards(config.FTC_SEASON, event.Code)
		if err != nil {
			return err
		}
		Awards = append(Awards, awards...)
	}

	// Save on disk
	return StoreAwards()
}

// StoreAwards writes the Awards data to the file system
func StoreAwards() error {
	return writeFile(AWARDS_FILE_NAME, Awards)
}

// LoadAwards loads the Awards data from the file system
func LoadAwards() error {
	data, err := readFile(AWARDS_FILE_NAME)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &Awards)
}
