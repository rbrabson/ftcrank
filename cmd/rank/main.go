package main

import (
	"fmt"

	"github.com/rbrabson/ftcrank/ftcdata"
	"github.com/rbrabson/ftcrank/rank"
)

func main() {
	ftcdata.LoadAll()
	rank.RankTeams()

	var count int

	/*
		// Print out the "Top 10" in North Carolina
		count = 1
		for i, team := range rank.RankedTeams {
			if *team.Info.HomeRegion == "USNC" {
				rating := team.Ratings[len(team.Ratings)-1].EndRating
				fmt.Printf("Rank: %2d, GlobalRank: %3d, Team: %5d %s, mu: %.2f, sigma: %.2f\n", count, i+1, team.Info.TeamNumber, team.Info.NameShort, rating.AveragePlayerSkill, rating.SkillUncertaintyDegree)
				count++
				if count > 10 {
					break
				}
			}
		}
	*/

	// Print out the all teams in North Carolina
	count = 1
	for i, team := range rank.RankedTeams {
		if team.Info.HomeRegion != nil && *team.Info.HomeRegion == "USNC" {
			rating := team.Ratings[len(team.Ratings)-1].EndRating
			fmt.Printf("Rank: %3d, GlobalRank: %4d, Team: %5d %s, mu: %.2f, sigma: %.2f\n", count, i+1, team.Info.TeamNumber, team.Info.NameShort, rating.AveragePlayerSkill, rating.SkillUncertaintyDegree)
			count++
		}
	}

	/*
		// Print out the teams in the NC State Championship
		count = 1
		for i, team := range rank.RankedTeams {
			lastRating := team.Ratings[len(team.Ratings)-1]
			if lastRating.EventCode == "USNCCMP" {
				rating := team.Ratings[len(team.Ratings)-1].EndRating
				fmt.Printf("Rank: %2d, GlobalRank: %4d, Team: %5d %s, mu: %.2f, sigma: %.2f\n", count, i+1, team.Info.TeamNumber, team.Info.NameShort, rating.AveragePlayerSkill, rating.SkillUncertaintyDegree)
				count++
			}
		}
	*/

	/*
		// Get starting rankings among teams at the NC State Championship
		ncTeams := make([]*rank.Team, 0, 110)
		for _, team := range rank.RankedTeams {
			lastRating := team.Ratings[len(team.Ratings)-1]
			if lastRating.EventCode == "USNCCMP" {
				ncTeams = append(ncTeams, team)
			}
		}
		sort.Slice(ncTeams, func(i, j int) bool {
			team1 := ncTeams[i]
			team2 := ncTeams[j]

			if len(team2.Ratings) == 0 {
				return true
			}
			if len(team1.Ratings) == 0 {
				return false
			}
			rating1 := team1.Ratings[len(team1.Ratings)-1].StartRating
			rating2 := team2.Ratings[len(team2.Ratings)-1].StartRating
			if rating1.AveragePlayerSkill != rating2.AveragePlayerSkill {
				return rating1.AveragePlayerSkill > rating2.AveragePlayerSkill
			} else {
				return rating1.SkillUncertaintyDegree < rating2.SkillUncertaintyDegree
			}
		})
		for i, team := range ncTeams {
			rating := team.Ratings[len(team.Ratings)-1].StartRating
			fmt.Printf("Rank: %2d, Team: %5d %s, mu: %.2f, sigma: %.2f\n", i+1, team.Info.TeamNumber, team.Info.NameShort, rating.AveragePlayerSkill, rating.SkillUncertaintyDegree)
		}
	*/
}
