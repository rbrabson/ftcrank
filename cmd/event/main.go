package main

import (
	"fmt"

	"github.com/rbrabson/ftcrank/ftcdata"
)

func main() {
	err := ftcdata.RetrieveEvents()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = ftcdata.LoadEvents()
	if err != nil {
		fmt.Println(err)
		return
	}
}
