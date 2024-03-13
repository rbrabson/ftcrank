package skill

import (
	"github.com/eullerpereira94/openskill"
	"github.com/rbrabson/ftcrank/config"
)

// Rate rates a group of teams for classification.
func Rate(winningAlliance []*openskill.Rating, losingAlliance []*openskill.Rating) ([]*openskill.Rating, []*openskill.Rating) {
	teams := []openskill.Team{winningAlliance, losingAlliance}
	updatedTeams := openskill.Rate(teams, config.DEFAULT_OPTIONS)

	for i := range winningAlliance {
		winningAlliance[i] = updatedTeams[0][i]
	}
	for i := range losingAlliance {
		losingAlliance[i] = updatedTeams[1][i]
	}
	return winningAlliance, losingAlliance
}
