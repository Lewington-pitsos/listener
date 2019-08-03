package main

import (
	"fmt"
	"time"

	"bitbucket.org/lewington/autoroller/globals"
	"bitbucket.org/lewington/autoroller/warn/spoke"
)

func main() {
	s := spoke.NewAudio(fmt.Sprintf("http://%v:8083/health", globals.HubIP))
	s.AddSpoke("bettor", "screaming.mp3")
	s.AddSpoke("scraper", "screaming-20.mp3")
	s.SpinUp()

	time.Sleep(time.Hour * 1000)
}
