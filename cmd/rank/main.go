package main

import (
	"errors"
	"fmt"
	"os"
	"sort"

	"github.com/eullerpereira94/openskill"
	"github.com/fatih/color"
	"github.com/rbrabson/ftcrank/ftcdata"
	"github.com/rbrabson/ftcrank/rank"
	"github.com/rodaine/table"
	"github.com/urfave/cli/v2"
)

var (
	appName  = "rank"
	Version  = "0.1.0"
	Revision = "v1beta1"
	usage    = "FTC team ranking utility"
)

var (
	// flags are the set of flags supported by the prediction application
	flags = []cli.Flag{
		&cli.StringFlag{
			Name:     "user",
			Aliases:  []string{"u"},
			EnvVars:  []string{"FTC_USERNAME"},
			Usage:    "Username used to log into the FTC API website",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "password",
			Aliases:  []string{"p"},
			EnvVars:  []string{"FTC_AUTHORIZATION_KEY"},
			Usage:    "The authorization token used as the password when logging into the FTC API website",
			Required: true,
		},
		&cli.StringFlag{
			Name:        "server",
			Aliases:     []string{"s"},
			EnvVars:     []string{"FTC_SERVER"},
			Usage:       "API server URL",
			DefaultText: "https://ftc-api.firstinspires.org/v2.0",
		},
		&cli.StringFlag{
			Name:        "dir",
			Aliases:     []string{"d"},
			EnvVars:     []string{"STORAGE_DIRECTORY"},
			Usage:       "Base directory under which data is stored",
			DefaultText: "",
			Required:    true,
		},
		&cli.StringFlag{
			Name:        "season",
			Aliases:     []string{"y"},
			EnvVars:     []string{"FTC_SEASON"},
			Usage:       "Season for the FTC data",
			DefaultText: "2023",
		},
		&cli.StringFlag{
			Name:    "event",
			Aliases: []string{"e"},
			Usage:   "FTC event from which the teams should be ranked",
		},
		&cli.BoolFlag{
			Name:    "start",
			Aliases: []string{"x"},
			Usage:   "Whether the teams at the event should be ranked at start or end of the event",
		},
		&cli.StringFlag{
			Name:    "region",
			Aliases: []string{"r"},
			Usage:   "FTC region to be ranked",
		},
		&cli.IntFlag{
			Name:    "limit",
			Aliases: []string{"l"},
			Usage:   "Maximum number of teams to be ranked",
		},
		&cli.StringFlag{
			Name:  "log",
			Usage: "Log level to be used when logging messages",
			Value: "Warn",
		},
	}
)

type TeamRating struct {
	Team   *rank.Team
	Rating *openskill.Rating
}

func printRanking(teams []*TeamRating) {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New("Rank", "Name", "mu", "sigma")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	for i, team := range teams {
		tbl.AddRow(i+1,
			team.Team.Info.DisplayTeamNumber+" "+team.Team.Info.NameShort,
			fmt.Sprintf("%.2f", team.Rating.AveragePlayerSkill),
			fmt.Sprintf("%.2f", team.Rating.SkillUncertaintyDegree),
		)
	}

	tbl.Print()
}

func runApp(cli *cli.Context) error {
	ftcdata.LoadAll()
	rank.RankTeams()

	// List the current ratings for teams within the requested region
	if cli.IsSet("region") {
		teams := make([]*TeamRating, 0, 110)
		count := 1
		for _, team := range rank.RankedTeams {
			if team.Info.HomeRegion != nil && *team.Info.HomeRegion == cli.String("region") {
				rating := &team.Ratings[len(team.Ratings)-1].EndRating
				teamRating := &TeamRating{
					Team:   team,
					Rating: rating,
				}
				teams = append(teams, teamRating)
				count++
			}
			if cli.IsSet("limit") && cli.Int("limit") < count {
				break
			}
		}
		printRanking(teams)

		return nil
	}

	// List the teams ratings at the end of the event
	if cli.IsSet("event") && !cli.Bool("start") {
		teams := make([]*TeamRating, 0, 200)
		count := 1
		for _, team := range rank.RankedTeams {
			lastRating := team.Ratings[len(team.Ratings)-1]
			if lastRating.EventCode == cli.String("event") {
				rating := &team.Ratings[len(team.Ratings)-1].EndRating
				teamRating := &TeamRating{
					Team:   team,
					Rating: rating,
				}
				teams = append(teams, teamRating)
				count++
			}
			if cli.IsSet("limit") && cli.Int("limit") < count {
				break
			}
		}
		printRanking(teams)

		return nil
	}

	// List the team ratings at the start of the event
	if cli.IsSet("event") {
		eventTeams := make([]*rank.Team, 0, 110)
		for _, team := range rank.RankedTeams {
			lastRating := team.Ratings[len(team.Ratings)-1]
			if lastRating.EventCode == cli.String("event") {
				eventTeams = append(eventTeams, team)
			}
		}
		fmt.Println(len(eventTeams))
		sort.Slice(eventTeams, func(i, j int) bool {
			team1 := eventTeams[i]
			team2 := eventTeams[j]

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
		teams := make([]*TeamRating, 0, 200)
		count := 1
		for _, team := range eventTeams {
			rating := &team.Ratings[len(team.Ratings)-1].StartRating
			teamRating := &TeamRating{
				Team:   team,
				Rating: rating,
			}
			teams = append(teams, teamRating)
			count++
			if cli.IsSet("limit") && cli.Int("limit") < count {
				break
			}
		}
		printRanking(teams)

		return nil
	}

	return errors.New("neither region nor event were set")
}

func main() {
	app := &cli.App{
		Name:    appName,
		Version: Version + "+" + Revision,
		Flags:   flags,
		Usage:   usage,
		Action:  runApp,
	}

	app.Run(os.Args)
}
