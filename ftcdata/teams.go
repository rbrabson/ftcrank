package ftcdata

import (
	"encoding/json"
	"sort"

	"github.com/rbrabson/ftc/ftc"
	"github.com/rbrabson/ftcrank/config"
)

const (
	TEAMS_FILE_NAME = "teams.json"
)

var (
	Teams []ftc.Team = make([]ftc.Team, 0, 11301)
)

// Have to get the teams and move through the pages

// RetrieveTeams retrieves the set of teams for a given year using the FTC developer API
func RetrieveTeams() error {
	var err error
	// Get the first time. Then  move through each subsequent page to build up the Teams list
	Teams, err = ftc.GetTeams(config.FTC_SEASON)
	if err != nil {
		return err
	}

	// Sort based on the event start time and, if the same start time, the event code
	sort.Slice(Teams, func(i, j int) bool {
		return Teams[i].TeamNumber < Teams[j].TeamNumber
	})

	// Save on disk
	return StoreTeams()
}

// StoreEvents stores the events in the file system
func StoreTeams() error {
	return writeFile(TEAMS_FILE_NAME, Teams)
}

// LoadEvents loads the events from the file system
func LoadTeams() error {
	data, err := readFile(TEAMS_FILE_NAME)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &Teams)
}
