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

	err = ftcdata.RetrieveScores()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = ftcdata.LoadScores()
	if err != nil {
		fmt.Println(err)
		return
	}
}
