package ftcdata

import (
	"encoding/json"

	"github.com/rbrabson/ftc/ftc"
	"github.com/rbrabson/ftcrank/config"
)

const (
	ADVANCES_FROM_FILE_NAME = "advancesfrom.json"
	ADVANCES_TO_FILE_NAME   = "advancesto.json"
)

var (
	AdvancementsFrom = make([]FtcAdvancementsFrom, 0, len(Events))
	AdvancementsTo   = make([]FtcAdvancementsTo, 0, len(Events))
)

// AdvancementsFrom structure that contains the event to which the teams advanced
type FtcAdvancementsFrom struct {
	EventCode string `json:"eventCode"`
	ftc.AdvancementsFrom
}

// AdvancementsTo structure that contains the event from which the teams advanced
type FtcAdvancementsTo struct {
	EventCode string `json:"eventCode"`
	ftc.AdvancementsTo
}

// RetrieveAdvancements gets all the advancements to and from a given tournament
func RetrieveAdvancements() error {
	for _, event := range Events {
		switch event.TypeName {
		case EVENT_QUALIFIER:
			err := retrieveAdvancementsTo(event.Code)
			if err != nil {
				return err
			}
		case EVENT_CHAMPIONSHIP:
			err := retrieveAdvancementsTo(event.Code)
			if err != nil {
				return err
			}
			err = retrieveAdvancementsFrom(event.Code)
			if err != nil {
				return err
			}
		case EVENT_FIRST_CHAMPIONSHIP:
			err := retrieveAdvancementsFrom(event.Code)
			if err != nil {
				return err
			}
		}
	}

	// Save on disk
	return StoreAdvancements()
}

// retrieveAdvancementsFrom returns the list of teams in an event that advanced from another event.
func retrieveAdvancementsFrom(eventCode string) error {
	var err error

	// Get the events
	advancements, err := ftc.GetAdvancementsFrom(config.FTC_SEASON, eventCode)
	if err != nil {
		return err
	}
	for _, advancement := range advancements {
		ftcAdvancement := FtcAdvancementsFrom{
			EventCode:        eventCode,
			AdvancementsFrom: advancement,
		}
		AdvancementsFrom = append(AdvancementsFrom, ftcAdvancement)
	}

	return nil
}

// retrieveAdvancementsTo returns the list of teams that advance from an event to another event.
func retrieveAdvancementsTo(eventCode string) error {
	advancements, err := ftc.GetAdvancementsTo(config.FTC_SEASON, eventCode)
	if err != nil {
		return err
	}
	ftcAdvancement := FtcAdvancementsTo{
		EventCode:      eventCode,
		AdvancementsTo: *advancements,
	}
	AdvancementsTo = append(AdvancementsTo, ftcAdvancement)

	return nil
}

// StoreAdvancements writes the advancement data to the file system
func StoreAdvancements() error {
	err := storeAdvancementsFrom()
	if err != nil {
		return err
	}
	return storeAdvancementsTo()
}

// storeAdvancementsFrom writes the AdvancementsFrom data to the file system
func storeAdvancementsFrom() error {
	return writeFile(ADVANCES_FROM_FILE_NAME, AdvancementsFrom)
}

// storeAdvancementsTo writes the AdvancementsTo data to the file system
func storeAdvancementsTo() error {
	return writeFile(ADVANCES_TO_FILE_NAME, AdvancementsTo)
}

// LoadAdvancements loads the advancement data from the file system
func LoadAdvancements() error {
	err := loadAdvancementsFrom()
	if err != nil {
		return err
	}
	return loadAdvancementsTo()
}

// loadAdvancementsFrom loads the AdvancementsFrom data from the file system
func loadAdvancementsFrom() error {
	data, err := readFile(ADVANCES_FROM_FILE_NAME)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &AdvancementsFrom)
}

// loadAdvancementsTo loads the AdvancementsTo data from the file system
func loadAdvancementsTo() error {
	data, err := readFile(ADVANCES_FROM_FILE_NAME)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &AdvancementsTo)
}
