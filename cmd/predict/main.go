package main

import (
	"github.com/rbrabson/ftcrank/ftcdata"
	"github.com/rbrabson/ftcrank/predict"
	"github.com/rbrabson/ftcrank/rank"
)

func main() {
	ftcdata.LoadAll()
	rank.RankTeams()

	predict.PredictMatches("USNCCMP", 7083)
}
