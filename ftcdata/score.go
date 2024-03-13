package ftcdata

import (
	"encoding/json"
	"fmt"

	"github.com/rbrabson/ftc/ftc"
	"github.com/rbrabson/ftcrank/config"
)

const (
	SCORES_FILE_NAME = "scores.json"
)

var (
	// The FTC events for a given year
	Scores = make([]FtcScores, len(Events))
)

// FtcScores is an FTC scores with the event code included
type FtcScores struct {
	EventCode string `json:"eventCode"`
	Scores    []ftc.MatchScores
}

// RetrieveScores retrieves the scores from the FTC API
func RetrieveScores() error {
	// Iterate over all the events
	for _, event := range Events {
		if event.TypeName == EVENT_QUALIFIER || event.TypeName == EVENT_CHAMPIONSHIP || event.TypeName == EVENT_FIRST_CHAMPIONSHIP {
			scores, err := ftc.GetEventScores(config.FTC_SEASON, event.Code, ftc.QUALIFIER)
			if err != nil {
				fmt.Printf("Warning: Retrieving scores for %s, tournamenLevel=%s, error=%s\n", event.Code, ftc.QUALIFIER, err.Error())
			} else {
				score := FtcScores{
					EventCode: event.Code,
					Scores:    scores,
				}
				Scores = append(Scores, score)
			}
			scores, err = ftc.GetEventScores(config.FTC_SEASON, event.Code, ftc.PLAYOFF)
			if err != nil {
				fmt.Printf("Warning: Retrieving scores for %s, tournamenLevel=%s, error=%s\n", event.Code, ftc.PLAYOFF, err.Error())
			} else {
				score := FtcScores{
					EventCode: event.Code,
					Scores:    scores,
				}
				Scores = append(Scores, score)
			}

		}
	}

	// Save on disk
	return StoreScores()
}

// StoreScores stores the scores in the file system
func StoreScores() error {
	return writeFile(SCORES_FILE_NAME, Scores)
}

// LoadScores loads the scores from the file system
func LoadScores() error {
	data, err := readFile(SCORES_FILE_NAME)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &Scores)
}
