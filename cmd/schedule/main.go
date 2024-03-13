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

	err = ftcdata.RetrievSchedules()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = ftcdata.LoadSchedules()
	if err != nil {
		fmt.Println(err)
		return
	}
}
