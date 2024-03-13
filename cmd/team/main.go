package main

import (
	"fmt"

	"github.com/rbrabson/ftcrank/ftcdata"
)

func main() {
	err := ftcdata.RetrieveTeams()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = ftcdata.LoadTeams()
	if err != nil {
		fmt.Println(err)
		return
	}
}
