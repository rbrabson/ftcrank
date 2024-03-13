package config

import (
	"os"

	"github.com/eullerpereira94/openskill"
	"github.com/joho/godotenv"
)

const (
	DEFAULT_MU              float64 = 25.0
	DEFAULT_SIGNMA          float64 = 3
	SKILL_UNCERTANTY_DEGREE float64 = DEFAULT_MU / DEFAULT_SIGNMA
)

var (
	DEFAULT_RATING_PARMS = &openskill.NewRatingParams{
		AveragePlayerSkill:     DEFAULT_MU,
		SkillUncertaintyDegree: SKILL_UNCERTANTY_DEGREE,
	}
	DEFAULT_OPTIONS = openskill.Options{}
)

var (
	// Base directory under which data is stored
	STORAGE_DIRECTORY string
	FTC_SEASON        string
)

func init() {
	godotenv.Load()

	STORAGE_DIRECTORY = os.Getenv("STORAGE_DIRECTORY")
	FTC_SEASON = os.Getenv("FTC_SEASON")
}
