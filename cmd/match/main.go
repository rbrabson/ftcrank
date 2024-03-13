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

	err = ftcdata.RetrieveMatches()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = ftcdata.LoadMatches()
	if err != nil {
		fmt.Println(err)
		return
	}
}
