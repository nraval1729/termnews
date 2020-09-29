package main

import (
	"./news"
	"./ui"
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
