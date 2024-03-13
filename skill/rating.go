package skill

import (
	"github.com/eullerpereira94/openskill"
	"github.com/rbrabson/ftcrank/config"
)

// NewRating returns a new openskill rating using the default values
func NewRating() *openskill.Rating {
	ratingParms := openskill.NewRatingParams{
		AveragePlayerSkill:     config.DEFAULT_MU,
		SkillUncertaintyDegree: config.DEFAULT_SIGNMA,
	}
	rating := openskill.NewRating(&ratingParms, nil)
	return rating
}
