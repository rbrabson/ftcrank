package rank

import (
	"github.com/rbrabson/ftcrank/ftcdata"
)

var (
	MatchMap = make(map[string]*ftcdata.FtcEventMatches)
)

// loadMatches loads the FTC matches into a map for easy access
func loadMatches() {
	for _, match := range ftcdata.Matches {
		MatchMap[match.EventCode] = match
	}
}
