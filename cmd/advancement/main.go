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

	err = ftcdata.RetrieveAdvancements()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = ftcdata.LoadAdvancements()
	if err != nil {
		fmt.Println(err)
		return
	}
}
