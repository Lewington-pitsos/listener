package hub

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/lewington/listener/globals"

	"github.com/lewington/listener/assist"
	"github.com/lewington/listener/warn/health"

	"github.com/lewington/listener/testhelp"
)

func TestJSONHub(t *testing.T) {
	globals.HubPort = ":7070"
	j := NewJSONHub([]globals.Component{globals.ComponentBettor, globals.ComponentScraper})

	testhelp.SlightDelay()

	resp, err := http.Get("http://localhost:7070/health")
	testhelp.ExpectNoError(t, err)

	var systemChart health.Chart

	err = json.Unmarshal(assist.StrictBytes(resp.Body), &systemChart)
	testhelp.ExpectNoError(t, err)
	for _, system := range systemChart.Systems {
		if system.Status != health.StatusBlack {
			t.Fatalf("expected %v for all initial systems, got %v", health.StatusBlack, system.Status)
		}
	}

	_, exists := systemChart.Systems["bettor"]

	testhelp.ExpectTrue(t, exists, "expected bettor system to exist")

	sys := health.System{
		"bettor",
		health.StatusRed,
	}

	sysBytes, err := json.Marshal(sys)
	assist.Check(err)

	resp, err = http.Post("http://localhost:7070/update", "application/json", bytes.NewBuffer(sysBytes))
	testhelp.ExpectNoError(t, err)

	if resp.Status != "200 OK" {
		t.Fatalf("expected successful status from post, got %v", resp)
	}

	resp, err = http.Get("http://localhost:7070/health")
	testhelp.ExpectNoError(t, err)
	var systemChart2 health.Chart

	err = json.Unmarshal(assist.StrictBytes(resp.Body), &systemChart2)
	testhelp.ExpectNoError(t, err)

	bettor, exists := systemChart2.Systems["bettor"]
	testhelp.ExpectTrue(t, exists, "expected bettor system to exist")
	if bettor.Status != health.StatusRed {
		t.Fatalf("expected %v for updated system, got %v", health.StatusRed, bettor.Status)
	}

	scraper, exists := systemChart2.Systems["scraper"]
	testhelp.ExpectTrue(t, exists, "expected bettor system to exist")
	if scraper.Status != health.StatusBlack {
		t.Fatalf("expected %v for non updated system, got %v", health.StatusBlack, scraper.Status)
	}

	resp, err = http.Post("http://localhost:7070/update", "application/json", nil)
	testhelp.ExpectNoError(t, err)
	if resp.Status != "400 Bad Request" {
		t.Fatalf("expected bad request status from post, got %v", resp)
	}

	sys = health.System{
		"non existasnt",
		health.StatusRed,
	}

	sysBytes, err = json.Marshal(sys)
	assist.Check(err)

	resp, err = http.Post("http://localhost:7070/update", "application/json", bytes.NewBuffer(sysBytes))
	testhelp.ExpectNoError(t, err)

	if resp.Status != "400 Bad Request" {
		t.Fatalf("expected bad request status from post, got %v", resp)
	}

	j.Close()
}

func TestJSONHubMarkingRedForLackOfUpdates(t *testing.T) {
	globals.HubPort = ":7070"
	globals.HealthCheckWait = time.Millisecond * 160
	j := NewJSONHub([]globals.Component{globals.ComponentBettor, globals.ComponentScraper})

	testhelp.SlightDelay()

	resp, err := http.Get("http://localhost:7070/health")
	testhelp.ExpectNoError(t, err)

	var systemChart health.Chart

	err = json.Unmarshal(assist.StrictBytes(resp.Body), &systemChart)
	testhelp.ExpectNoError(t, err)
	for _, system := range systemChart.Systems {
		if system.Status != health.StatusBlack {
			t.Fatalf("expected %v for all initial systems, got %v", health.StatusBlack, system.Status)
		}
	}

	assist.Wait(200)

	resp, err = http.Get("http://localhost:7070/health")
	testhelp.ExpectNoError(t, err)
	err = json.Unmarshal(assist.StrictBytes(resp.Body), &systemChart)
	testhelp.ExpectNoError(t, err)
	for _, system := range systemChart.Systems {
		if system.Status != health.StatusRed {
			t.Fatalf("expected %v for all systems not recently updated, got %v", health.StatusRed, system.Status)
		}
	}

	sys := health.System{
		"bettor",
		health.StatusGreen,
	}

	sysBytes, err := json.Marshal(sys)
	assist.Check(err)

	resp, err = http.Post("http://localhost:7070/update", "application/json", bytes.NewBuffer(sysBytes))
	testhelp.ExpectNoError(t, err)

	if resp.Status != "200 OK" {
		t.Fatalf("expected successful status from post, got %v", resp)
	}

	resp, err = http.Get("http://localhost:7070/health")
	testhelp.ExpectNoError(t, err)
	err = json.Unmarshal(assist.StrictBytes(resp.Body), &systemChart)
	testhelp.ExpectNoError(t, err)

	if systemChart.Systems["bettor"].Status != health.StatusGreen {
		t.Fatalf("expected %v for recently updated system, got %v", health.StatusGreen, systemChart.Systems["bettor"].Status)
	}
	if systemChart.Systems["scraper"].Status != health.StatusRed {
		t.Fatalf("expected %v for not recently updated system, got %v", health.StatusRed, systemChart.Systems["bettor"].Status)
	}

	assist.Wait(200)

	resp, err = http.Get("http://localhost:7070/health")
	testhelp.ExpectNoError(t, err)
	err = json.Unmarshal(assist.StrictBytes(resp.Body), &systemChart)
	testhelp.ExpectNoError(t, err)
	for _, system := range systemChart.Systems {
		if system.Status != health.StatusRed {
			t.Fatalf("expected %v for all systems not recently updated, got %v", health.StatusRed, system.Status)
		}
	}
	j.Close()
}
