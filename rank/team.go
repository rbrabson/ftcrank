package rank

import (
	"github.com/eullerpereira94/openskill"
	"github.com/rbrabson/ftc/ftc"
	"github.com/rbrabson/ftcrank/ftcdata"
)

var (
	ftcTeamMap = make(map[int]*ftc.Team)
	TeamMap    = make(map[int]*Team)
)

// A team used for rankings
type Team struct {
	Info    *ftc.Team
	Ratings []*MatchRating
}

// Rankings at a given match
type MatchRating struct {
	EventCode   string
	EventName   string
	StartRating openskill.Rating
	EndRating   openskill.Rating
}

// GetTeam returns the team with the given team number. If the team does not yet
// exist, it is created.
func GetTeam(teamNumber int) *Team {
	team := TeamMap[teamNumber]
	if team == nil {
		team = newTeam(teamNumber)
		TeamMap[team.Info.TeamNumber] = team
	}

	return team
}

// newTeam creates a new team
func newTeam(teamNumber int) *Team {
	// If the FTC teams aren't loaded in the map, load them now
	if len(ftcTeamMap) == 0 {
		loadFtcTeams()
	}

	// Use the fTC team to create a new ranking team
	ftcTeam := ftcTeamMap[teamNumber]
	team := Team{
		Info:    ftcTeam,
		Ratings: make([]*MatchRating, 0),
	}

	return &team
}

// loadFtcTeams loads the teams into the ftcTeamMap for easier access
func loadFtcTeams() {
	for _, team := range ftcdata.Teams {
		ftcTeamMap[team.TeamNumber] = &team
	}
}
