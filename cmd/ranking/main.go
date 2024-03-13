package main

import (
	"fmt"

	"github.com/rbrabson/ftcrank/ftcdata"
)

func main() {
	err := ftcdata.LoadEvents()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = ftcdata.RetrieveRankings()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = ftcdata.LoadRankings()
	if err != nil {
		fmt.Println(err)
		return
	}
}
