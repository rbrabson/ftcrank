package ftcdata

import (
	"encoding/json"
	"fmt"

	"github.com/rbrabson/ftc/ftc"
	"github.com/rbrabson/ftcrank/config"
)

const (
	RANKINGS_FILE_NAME = "rankings.json"
)

var (
	// The FTC events for a given year
	Rankings = make([]ftcrank, len(Events))
)

// ftcrank is an FTC ranking with the event code included
type ftcrank struct {
	EventCode string `json:"eventCode"`
	Rankings  []ftc.Ranking
}

// RetrieveRankings gets the rankings from the FTC API server
func RetrieveRankings() error {
	for _, event := range Events {
		if event.TypeName == EVENT_QUALIFIER || event.TypeName == EVENT_CHAMPIONSHIP || event.TypeName == EVENT_FIRST_CHAMPIONSHIP {
			rankings, err := ftc.GetRankings(config.FTC_SEASON, event.Code)
			if err != nil {
				fmt.Printf("Warning: Retrieving rankings for %s, error=%s\n", event.Code, err.Error())
			} else {
				ranking := ftcrank{
					EventCode: event.Code,
					Rankings:  rankings,
				}
				Rankings = append(Rankings, ranking)
			}
		}
	}

	// Save on disk
	return StoreRankings()
}

// StoreRankings stores the rankings in the file system
func StoreRankings() error {
	return writeFile(RANKINGS_FILE_NAME, Rankings)
}

// LoadRankings loads the rankings from the file system
func LoadRankings() error {
	data, err := readFile(RANKINGS_FILE_NAME)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &Rankings)
}
