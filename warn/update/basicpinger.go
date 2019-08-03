package update

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lewington/listener/assist"
	"github.com/lewington/listener/globals"
	"github.com/lewington/listener/warn/health"
)

// BasicPinger makes post requests to a JSONHub.
type BasicPinger struct {
	endpoint string
}

// Update makes an update Post to the hub and
// returns an error if anything goes wrong.
func (b *BasicPinger) Update(component globals.Component, status health.Status) error {
	var err error
	fmt.Println("attempting to ping")
	sysBytes, err := json.Marshal(health.System{
		component,
		status,
	})
	assist.Check(err)

	success := false
	for i := 0; i < 4; i++ {
		resp, err := http.Post(b.endpoint, "application/json", bytes.NewBuffer(sysBytes))

		if err != nil {
			fmt.Println("failure to post update error: %v, response: %v", err, resp)
		} else {
			resp.Body.Close()
			if resp.Status == "200 OK" {
				success = true
				break
			} else {
				fmt.Println("anomalous response detected when updating to hub")
			}
		}
	}
	if !success {
		assist.Panicf("failed to update health to hub")
	}

	return nil
}

// NewBasicPinger initializes a BasicPinger.
func NewBasicPinger(address string) Pinger {
	return &BasicPinger{address}
}

// JSONHubPinger initializes a BasicPinger with
// a specific address.
func JSONHubPinger() Pinger {
	return NewBasicPinger(fmt.Sprintf("http://%v:8083/update", globals.HubIP))
}
