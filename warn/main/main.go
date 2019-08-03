package main

import (
	"time"

	"bitbucket.org/lewington/autoroller/globals"

	"bitbucket.org/lewington/autoroller/warn/hub"
)

func main() {
	hub.NewJSONHub([]globals.Component{
		// globals.ComponentRouter,
		globals.ComponentBettor,
		globals.ComponentScraper,
	})

	time.Sleep(time.Hour * 1000)
}
