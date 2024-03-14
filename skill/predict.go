package skill

import (
	"github.com/eullerpereira94/openskill"
)

// PredictWin returns the probability of each team has to win ordered by the
// order of the teams.
func PredictWin(alliance1, alliance2 []*openskill.Rating) []float64 {
	teams := []openskill.Team{alliance1, alliance2}
	winLikelihood := openskill.PredictWin(teams, nil)
	return winLikelihood
}

// PredictRank calculates and predicts the ranks of teams based on pairwise
// probabilities.
//
// Returns:
//
//	predictions: A 2D slice of float64 containing the predicted ranks and corresponding probabilities.
//	             Each inner slice has two elements: rank and probability.
//	             The outer slice represents the predictions for each team.
//	             The length of predictions slice is equal to the number of teams.
func PredictRank(teams []*openskill.Rating) [][]float64 {
	teamList := make([]openskill.Team, 0, len(teams))
	for _, t := range teams {
		tr := openskill.Team{t}
		teamList = append(teamList, tr)
	}
	ranks := openskill.PredictRank(teamList, nil)
	return ranks
}
