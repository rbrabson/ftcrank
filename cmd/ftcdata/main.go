package main

import (
	"fmt"

	"github.com/rbrabson/ftcrank/ftcdata"
)

func main() {
	var err error

	err = ftcdata.RetrieveEvents()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = ftcdata.RetrieveTeams()
	if err != nil {
		fmt.Println(err)
		return
	}
}
