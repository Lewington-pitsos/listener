package main

import (
	"fmt"
	"time"

	"github.com/lewington/listener/globals"
	"github.com/lewington/listener/warn/spoke"
)

func main() {
	s := spoke.NewAudio(fmt.Sprintf("http://%v:8083/health", globals.HubIP))
	s.AddSpoke("bettor", "screaming.mp3")
	s.AddSpoke("scraper", "screaming-20.mp3")
	s.SpinUp()

	time.Sleep(time.Hour * 1000)
}
