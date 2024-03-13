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

	err = ftcdata.RetrieveHybridSchedules()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = ftcdata.LoadHybridSchedules()
	if err != nil {
		fmt.Println(err)
		return
	}
}
