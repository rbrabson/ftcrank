package ftcdata

import (
	"encoding/json"
	"sort"
	"time"

	"github.com/rbrabson/ftc/ftc"
	"github.com/rbrabson/ftcrank/config"
)

const (
	EVENTS_FILE_NAME = "events.json"
)

// Event Types
const (
	EVENT_QUALIFIER          = "Qualifier"
	EVENT_CHAMPIONSHIP       = "Championship"
	EVENT_FIRST_CHAMPIONSHIP = "FIRST Championship"
)

var (
	// The FTC events for a given year
	Events = make([]*ftc.Event, 0)
)

// RetrieveEvents retrieves the set of events for a given year using the FTC developer API
func RetrieveEvents() error {
	var err error

	// Get the events
	Events, err = ftc.GetEvents(config.FTC_SEASON)
	if err != nil {
		return err
	}

	// Sort based on the event start time and, if the same start time, the event code
	sort.Slice(Events, func(i, j int) bool {
		d1 := time.Time(Events[i].DateStart)
		d2 := time.Time(Events[j].DateStart)
		if d1.Equal(d2) {
			return Events[i].Code < Events[j].Code
		}
		return d1.Before(d2)
	})

	// Save on disk
	return StoreEvents()
}

// StoreEvents stores the events in the file system
func StoreEvents() error {
	return writeFile(EVENTS_FILE_NAME, Events)
}

// LoadEvents loads the events from the file system
func LoadEvents() error {
	data, err := readFile(EVENTS_FILE_NAME)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &Events)
}
