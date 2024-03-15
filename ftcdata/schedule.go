package ftcdata

import (
	"encoding/json"
	"fmt"

	"github.com/rbrabson/ftc/ftc"
	"github.com/rbrabson/ftcrank/config"
)

const (
	SCHEDULES_FILE_NAME = "schedules.json"
)

var (
	// The FTC events for a given year
	Schedules = make([]FtcSchedules, len(Events))
)

// ftcrank is an FTC ranking with the event code included
type FtcSchedules struct {
	EventCode string `json:"eventCode"`
	Schedules []*ftc.EventSchedule
}

// RetrieveSchedules retrieves the schedule data from the FTC API
func RetrieveSchedules() error {
	for _, event := range Events {
		// An event with matches that score
		if event.TypeName == EVENT_QUALIFIER || event.TypeName == EVENT_CHAMPIONSHIP || event.TypeName == EVENT_FIRST_CHAMPIONSHIP {
			// Get the qualification matches
			schedules, err := ftc.GetEventSchedule(config.FTC_SEASON, event.Code, ftc.QUALIFIER)
			if err != nil {
				fmt.Printf("Warning: Retrieving schedule for %s, error=%s\n", event.Code, err.Error())
				continue
			}
			schedule := FtcSchedules{
				EventCode: event.Code,
				Schedules: schedules,
			}
			// Add the playoff matches
			schedules, err = ftc.GetEventSchedule(config.FTC_SEASON, event.Code, ftc.PLAYOFF)
			if err != nil {
				fmt.Printf("Warning: Retrieving schedule for %s, error=%s\n", event.Code, err.Error())
				continue
			}
			schedule.Schedules = append(schedule.Schedules, schedules...)

			Schedules = append(Schedules, schedule)
		}
	}

	// Save on disk
	return StoreSchedules()
}

// StoreSchedules stores the schedules in the file system
func StoreSchedules() error {
	return writeFile(SCHEDULES_FILE_NAME, Schedules)
}

// LoadSchedules loads the schedules from the file system
func LoadSchedules() error {
	data, err := readFile(SCHEDULES_FILE_NAME)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &Schedules)
}

// UpdateSchedules adds or updates the schedule for the given event
func UpdateSchedules(eventCode string) error {
	// Get the qualification schedules
	schedules, err := ftc.GetEventSchedule(config.FTC_SEASON, eventCode, ftc.QUALIFIER)
	if err != nil {
		return err
	}
	schedule := FtcSchedules{
		EventCode: eventCode,
		Schedules: schedules,
	}

	// Add the playoff schedules
	schedules, err = ftc.GetEventSchedule(config.FTC_SEASON, eventCode, ftc.PLAYOFF)
	if err != nil {
		return err
	}
	schedule.Schedules = append(schedule.Schedules, schedules...)

	// Update or add the schedules
	updated := false
	for i := range Schedules {
		// Update the existing schedule
		if Schedules[i].EventCode == eventCode {
			Schedules[i] = schedule
			updated = true
			break
		}
	}
	// Add a new schedule
	if !updated {
		Schedules = append(Schedules, schedule)
	}

	return StoreSchedules()
}
