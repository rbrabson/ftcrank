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

	err = ftcdata.RetrieveAwards()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = ftcdata.LoadAwards()
	if err != nil {
		fmt.Println(err)
		return
	}
}
