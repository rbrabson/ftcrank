package rank

import (
	"sort"
	"strings"

	"github.com/eullerpereira94/openskill"
	"github.com/rbrabson/ftc/ftc"
	"github.com/rbrabson/ftcrank/ftcdata"
	"github.com/rbrabson/ftcrank/skill"
)

var (
	RankedTeams []*Team
)

// RankTeams generates the rankings for all teams that participated in FTC
// for the given season.
func RankTeams() []*Team {
	// Make sure the matches are loaded
	if len(MatchMap) == 0 {
		loadMatches()
	}

	// For each event, rank the teams in said event. This must be called in the
	// order the events take place for the results to be meaningful.
	for _, event := range ftcdata.Events {
		matches := MatchMap[event.Code]
		if matches != nil && len(matches.Matches) > 0 {
			rankEventMatches(event, matches)
		}
	}

	// Sort the RankedTeams based on their average skill and uncertantity degree
	RankedTeams = getSortedTeams()
	for i, team := range RankedTeams {
		team.GlobalRating = i + 1
	}

	return RankedTeams
}

// rankEventMatches ranks all matches that occurred at an event
func rankEventMatches(event ftc.Event, matches *ftcdata.FtcEventMatches) {
	for _, match := range matches.Matches {
		if len(match.Teams) > 0 {
			rankEventMatch(event, &match)
		}
	}
}

// rankEventMatch ranks a single match that ocurred at an event
func rankEventMatch(event ftc.Event, match *ftc.Match) {
	redAlliance := make([]*Team, 0, 2)
	blueAlliance := make([]*Team, 0, 2)

	for _, t := range match.Teams {
		if t.OnField {
			team := GetTeam(t.TeamNumber)
			lastRating := len(team.Ratings) - 1
			if lastRating < 0 {
				rating := &MatchRating{
					EventCode:   event.Code,
					EventName:   event.Name,
					StartRating: *skill.NewRating(),
					EndRating:   *skill.NewRating(),
				}
				team.Ratings = append(team.Ratings, rating)
			} else if team.Ratings[lastRating].EventCode != event.Code {
				prevRating := team.Ratings[lastRating]
				rating := &MatchRating{
					EventCode:   event.Code,
					EventName:   event.Name,
					StartRating: prevRating.EndRating,
					EndRating:   prevRating.EndRating,
				}
				team.Ratings = append(team.Ratings, rating)
			}

			// For team, make sure there is a rating for the given match at the end
			if strings.HasPrefix(t.Station, "Red") {
				redAlliance = append(redAlliance, team)
			} else {
				blueAlliance = append(blueAlliance, team)
			}
		}
	}
	// Update the rankings from the match
	if redAllianceWon(match) {
		rateMatch(redAlliance, blueAlliance)
	} else {
		rateMatch(blueAlliance, redAlliance)
	}
}

// rateMatch gets the openskill ratings for a match between two alliances
func rateMatch(winningAlliance []*Team, losingAlliance []*Team) {
	winningRatings := getAllianceRatings(winningAlliance)
	losingRatings := getAllianceRatings(losingAlliance)
	winningRatings, losingRatings = skill.Rate(winningRatings, losingRatings)
	updateAllianceRatings(winningAlliance, winningRatings)
	updateAllianceRatings(losingAlliance, losingRatings)
}

// getAllianceRatings gets the end ratings for the last event for each team in the alliance
func getAllianceRatings(alliance []*Team) []*openskill.Rating {
	ratings := make([]*openskill.Rating, 0, len(alliance))
	for _, team := range alliance {
		lastRating := team.Ratings[len(team.Ratings)-1]
		ratings = append(ratings, &lastRating.EndRating)
	}
	return ratings
}

// updateAllianceRatings updates the alliance ratings based on the results of calling skill.Rate
func updateAllianceRatings(alliance []*Team, ratings []*openskill.Rating) {
	for i, team := range alliance {
		lastRating := team.Ratings[len(team.Ratings)-1]
		lastRating.EndRating = *ratings[i]
	}
}

// redAllianceWon returns an indication as to whether the red or blue alliance won
func redAllianceWon(match *ftc.Match) bool {
	if match.ScoreRedFinal > match.ScoreBlueFinal {
		return true
	}
	if match.ScoreBlueFinal > match.ScoreRedFinal {
		return false
	}
	if match.ScoreRedAuto > match.ScoreBlueAuto {
		return true
	}
	if match.ScoreBlueAuto > match.ScoreRedAuto {
		return false
	}
	if match.ScoreRedFoul < match.ScoreBlueFoul {
		return true
	}
	return false
}

// getSortedTEams sorts the teams based on the average player skill (mu) and skill uncertanty degree (sigma)
func getSortedTeams() []*Team {
	// Get the teams to be sorted into their ranking
	teams := make([]*Team, 0, len(TeamMap))
	for _, v := range TeamMap {
		if len(v.Ratings) > 0 {
			teams = append(teams, v)
		}
	}
	// Stort the teams based on the `mu` and, if equal, `sigma`` values. Only include teams with at least one rating.
	sort.Slice(teams, func(i, j int) bool {
		team1 := teams[i]
		team2 := teams[j]

		if len(team2.Ratings) == 0 {
			return true
		}
		if len(team1.Ratings) == 0 {
			return false
		}
		rating1 := team1.Ratings[len(team1.Ratings)-1].EndRating
		rating2 := team2.Ratings[len(team2.Ratings)-1].EndRating
		if rating1.AveragePlayerSkill != rating2.AveragePlayerSkill {
			return rating1.AveragePlayerSkill > rating2.AveragePlayerSkill
		} else {
			return rating1.SkillUncertaintyDegree < rating2.SkillUncertaintyDegree
		}
	})

	return teams
}
