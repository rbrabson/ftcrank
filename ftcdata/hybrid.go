package ftcdata

import (
	"encoding/json"
	"fmt"

	"github.com/rbrabson/ftc/ftc"
	"github.com/rbrabson/ftcrank/config"
)

const (
	HYBRID_SCHEDULES_FILE_NAME = "hybrid.json"
)

var (
	// The FTC events for a given year
	HybridSchedules = make([]FtcHybridSchedules, len(Events))
)

// ftcrank is an FTC ranking with the event code included
type FtcHybridSchedules struct {
	EventCode string `json:"eventCode"`
	Schedules []*ftc.HybridSchedule
}

// RetrieveHybridSchedules retrieves the schedule data from the FTC API
func RetrieveHybridSchedules() error {
	for _, event := range Events {
		// If an event that has scoring
		if event.TypeName == EVENT_QUALIFIER || event.TypeName == EVENT_CHAMPIONSHIP || event.TypeName == EVENT_FIRST_CHAMPIONSHIP {
			// Get the qualification matches
			schedules, err := ftc.GetHybridSchedule(config.FTC_SEASON, event.Code, ftc.QUALIFIER)
			if err != nil {
				fmt.Printf("Warning: Retrieving rankings for %s, error=%s\n", event.Code, err.Error())
				continue
			}
			schedule := FtcHybridSchedules{
				EventCode: event.Code,
				Schedules: schedules,
			}

			// Get the playoff matches and append them to the list of event schedules
			schedules, err = ftc.GetHybridSchedule(config.FTC_SEASON, event.Code, ftc.PLAYOFF)
			if err != nil {
				fmt.Printf("Warning: Retrieving rankings for %s, error=%s\n", event.Code, err.Error())
				continue
			}
			schedule.Schedules = append(schedule.Schedules, schedules...)

			// Append the even schedule to the list of schedules
			HybridSchedules = append(HybridSchedules, schedule)
		}
	}

	// Save on disk
	return StoreHybridSchedules()
}

// StoreHybridSchedules stores the schedules in the file system
func StoreHybridSchedules() error {
	return writeFile(HYBRID_SCHEDULES_FILE_NAME, HybridSchedules)
}

// LoadHybridSchedules loads the schedules from the file system
func LoadHybridSchedules() error {
	data, err := readFile(HYBRID_SCHEDULES_FILE_NAME)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &HybridSchedules)
}

// UpdateHybridSchedules updates the hybrid scores for the given event
func UpdateHybridSchedules(eventCode string) error {
	schedules, err := ftc.GetHybridSchedule(config.FTC_SEASON, eventCode, ftc.QUALIFIER)
	if err != nil {
		return err
	}
	schedule := FtcHybridSchedules{
		EventCode: eventCode,
		Schedules: schedules,
	}

	// Get the playoff matches and append them to the list of event schedules
	schedules, err = ftc.GetHybridSchedule(config.FTC_SEASON, eventCode, ftc.PLAYOFF)
	if err != nil {
		return err
	}
	schedule.Schedules = append(schedule.Schedules, schedules...)

	// Update or add the schedules
	updated := false
	for i := range HybridSchedules {
		// Update the existing schedule
		if HybridSchedules[i].EventCode == eventCode {
			HybridSchedules[i] = schedule
			updated = true
			break
		}
	}
	// Add a new schedule
	if !updated {
		HybridSchedules = append(HybridSchedules, schedule)
	}

	// Save on disk
	return StoreHybridSchedules()
}
