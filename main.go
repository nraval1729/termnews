package main

import (
	"github.com/nraval1729/termnews/news"
	"github.com/nraval1729/termnews/ui"
)

func main() {
	err := news.FetchPeriodically()
	if err != nil {
		panic(err)
	}

	err = ui.Run()
	if err != nil {
		panic(err)
	}
}
