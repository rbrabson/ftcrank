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
	Scores    []*ftc.MatchScores
}

// RetrieveScores retrieves the scores from the FTC API
func RetrieveScores() error {
	// TODO: combine all the scores for a single event into a single entry

	// Iterate over all the events
	for _, event := range Events {
		// If an event that contains scores
		if event.TypeName == EVENT_QUALIFIER || event.TypeName == EVENT_CHAMPIONSHIP || event.TypeName == EVENT_FIRST_CHAMPIONSHIP {
			// Get the qualification scores
			scores, err := ftc.GetEventScores(config.FTC_SEASON, event.Code, ftc.QUALIFIER)
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
				continue
			}
			score := FtcScores{
				EventCode: event.Code,
				Scores:    scores,
			}
			// Add the playoff scores
			scores, err = ftc.GetEventScores(config.FTC_SEASON, event.Code, ftc.PLAYOFF)
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
				continue
			}
			score.Scores = append(score.Scores, scores...)

			// Save the scores for the event
			Scores = append(Scores, score)
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

// UpdateScores adds or updates the scores for the given event
func UpdateScores(eventCode string) error {
	// Get the qualification scores
	scores, err := ftc.GetEventScores(config.FTC_SEASON, eventCode, ftc.QUALIFIER)
	if err != nil {
		return err
	}
	score := FtcScores{
		EventCode: eventCode,
		Scores:    scores,
	}
	// Add the playoff scores
	scores, err = ftc.GetEventScores(config.FTC_SEASON, eventCode, ftc.PLAYOFF)
	if err != nil {
		return err
	}
	score.Scores = append(score.Scores, scores...)

	// Add or update the scores
	updated := false
	for i := range Scores {
		// Update the existng scores for the event
		if Scores[i].EventCode == eventCode {
			Scores[i] = score
			updated = true
			break
		}
	}
	// Add the scores for the event
	if !updated {
		Scores = append(Scores, score)
	}

	return StoreScores()
}
