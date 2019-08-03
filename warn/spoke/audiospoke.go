package spoke

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"bitbucket.org/lewington/autoroller/globals"

	"bitbucket.org/lewington/autoroller/assist"
	"bitbucket.org/lewington/autoroller/audio"
	"bitbucket.org/lewington/autoroller/warn/health"
)

const port string = ":8083"

// AudioSpoke only periodically pings an endpoint
// to get updates on system, health. It plays sounds
// depending on that health.
type AudioSpoke struct {
	alerts map[globals.Component]string
	url    string
	closed bool
	// warnType -> alert sound file
}

// SpinUp starts a http server to listen for warnings
// and serve alerts.
func (p *AudioSpoke) SpinUp() {
	go p.keepPinging()
}

// AddSpoke adds a new warning-alert pair.
func (p *AudioSpoke) AddSpoke(comp globals.Component, alertFile string) {
	p.alerts[comp] = alertFile
}

// Close stops the server listening.
func (p *AudioSpoke) Close() {
	p.closed = true
}

func (p *AudioSpoke) keepPinging() {
	for {
		if p.closed {
			break
		} else {
			p.ping()
			time.Sleep(time.Second * 5)
		}

	}
}

func (p *AudioSpoke) ping() {
	resp, err := http.Get(p.url)

	if err != nil {
		fmt.Println(err)
		go audio.PlaySound("whoosh.mp3")
		time.Sleep(time.Minute * 5)
	}

	var systemChart health.Chart

	err = json.Unmarshal(assist.StrictBytes(resp.Body), &systemChart)
	assist.Check(err)

	for _, system := range systemChart.Systems {
		for name, alertFile := range p.alerts {
			if name == system.Name && system.Status == health.StatusRed {
				go audio.PlaySound(alertFile)
				return
			}
		}
	}
}

// NewAudio initializes a ParamHub using the
// port constant to determine the port it will
// listen to.
func NewAudio(address string) *AudioSpoke {
	return &AudioSpoke{
		map[globals.Component]string{},
		address,
		false,
	}
}
