package predict

import (
	"fmt"
	"os"
	"strings"

	"github.com/eullerpereira94/openskill"
	"github.com/rbrabson/ftcrank/ftcdata"
	"github.com/rbrabson/ftcrank/rank"
	"github.com/rbrabson/ftcrank/skill"
)

type MatchPrediction struct {
	MatchNumber        int
	RedAlliance        []*rank.Team
	BlueAlliance       []*rank.Team
	RedWinProbability  float64
	BlueWinProbability float64
}

func PredictMatches(eventCode string, teamNumber ...int) []*MatchPrediction {
	// Get the schedule for the event
	var schedules *ftcdata.FtcSchedules
	for _, s := range ftcdata.Schedules {
		if s.EventCode == eventCode {
			schedules = &s
			break
		}
	}
	if schedules == nil {
		fmt.Println(os.ErrInvalid)
		os.Exit(1)
	}

	matches := make([]*MatchPrediction, 0, 200)
	for i, schedule := range schedules.Schedules {
		redAlliance := make([]*rank.Team, 0, 2)
		blueAlliance := make([]*rank.Team, 0, 2)
		for _, team := range schedule.Teams {
			if strings.HasPrefix(team.Station, "Red") {
				redAlliance = append(redAlliance, rank.TeamMap[team.TeamNumber])
			} else {
				blueAlliance = append(blueAlliance, rank.TeamMap[team.TeamNumber])
			}
		}
		redAllianceSkills := make([]*openskill.Rating, 0, len(redAlliance))
		blueAllianceSkills := make([]*openskill.Rating, 0, len(blueAlliance))
		for _, team := range redAlliance {
			for _, rating := range team.Ratings {
				if rating.EventCode == eventCode {
					redAllianceSkills = append(redAllianceSkills, &rating.StartRating)
				}
			}
		}
		for _, team := range blueAlliance {
			for _, rating := range team.Ratings {
				if rating.EventCode == eventCode {
					blueAllianceSkills = append(blueAllianceSkills, &rating.StartRating)
				}
			}
		}

		if len(teamNumber) > 0 {
			if redAlliance[0].Info.TeamNumber != teamNumber[0] && redAlliance[1].Info.TeamNumber != teamNumber[0] && blueAlliance[0].Info.TeamNumber != teamNumber[0] && blueAlliance[1].Info.TeamNumber != teamNumber[0] {
				continue
			}
		}

		results := skill.PredictWin(redAllianceSkills, blueAllianceSkills)
		match := &MatchPrediction{
			MatchNumber:        i + 1,
			RedAlliance:        redAlliance,
			BlueAlliance:       blueAlliance,
			RedWinProbability:  results[0],
			BlueWinProbability: results[1],
		}
		matches = append(matches, match)
	}

	return matches
}
