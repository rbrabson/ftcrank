package ftcdata

import (
	"encoding/json"
	"fmt"

	"github.com/rbrabson/ftc/ftc"
	"github.com/rbrabson/ftcrank/config"
)

const (
	MATCHES_FILE_NAME = "matches.json"
)

var (
	Matches = make([]*FtcEventMatches, 0)
)

type FtcEventMatches struct {
	EventCode string
	Matches   []*ftc.Match
}

// RetrieveMatches gets the matches that have occurred during an FTC season.
func RetrieveMatches() error {
	// Iterate over all the events
	for _, event := range Events {
		if !(event.TypeName == EVENT_QUALIFIER || event.TypeName == EVENT_CHAMPIONSHIP || event.TypeName == EVENT_FIRST_CHAMPIONSHIP) {
			continue
		}
		var ftcMatch *FtcEventMatches
		matches, err := ftc.GetMatchResults(config.FTC_SEASON, event.Code, ftc.QUALIFIER)
		if err != nil {
			fmt.Printf("Warning: Retrieving awards for %s, tournamenLevel=%s, error=%s\n", event.Code, ftc.QUALIFIER, err.Error())
		} else {
			ftcMatch = &FtcEventMatches{
				EventCode: event.Code,
				Matches:   matches,
			}
		}
		matches, err = ftc.GetMatchResults(config.FTC_SEASON, event.Code, ftc.PLAYOFF)
		if err != nil {
			fmt.Printf("Warning: Retrieving awards for %s, tournamenLevel=%s, error=%s\n", event.Code, ftc.PLAYOFF, err.Error())
		} else {
			if ftcMatch == nil {
				ftcMatch = &FtcEventMatches{
					EventCode: event.Code,
					Matches:   matches,
				}
			} else {
				ftcMatch.Matches = append(ftcMatch.Matches, matches...)
			}
		}
		if ftcMatch != nil {
			Matches = append(Matches, ftcMatch)
		}
	}

	// Save on disk
	return StoreMatches()
}

// StoreMatches writes the Match data to the file system
func StoreMatches() error {
	return writeFile(MATCHES_FILE_NAME, Matches)
}

// LoadMatches loads the Match data from the file system
func LoadMatches() error {
	data, err := readFile(MATCHES_FILE_NAME)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &Matches)
}
