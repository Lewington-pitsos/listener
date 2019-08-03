package main

import (
	"time"

	"github.com/lewington/listener/globals"

	"github.com/lewington/listener/warn/hub"
)

func main() {
	hub.NewJSONHub([]globals.Component{
		// globals.ComponentRouter,
		globals.ComponentBettor,
		globals.ComponentScraper,
	})

	time.Sleep(time.Hour * 1000)
}
