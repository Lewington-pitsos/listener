package hub

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/lewington/listener/globals"

	"github.com/lewington/listener/assist"
	"github.com/lewington/listener/warn/health"
)

// JSONHub serves the system's overall health as
// JSON.
type JSONHub struct {
	srv   *http.Server
	chart health.Chart
}

func (j *JSONHub) work() {
	go j.serve()
	go j.monitorHealth()
}

func (j *JSONHub) monitorHealth() {
	for {
		time.Sleep(globals.HealthCheckWait)
		j.chart.MarkOldSystemsAsRed()
	}
}

// Close stops the hub's internal server.
func (j *JSONHub) Close() {
	j.srv.Close()
	http.DefaultServeMux = new(http.ServeMux)
}

func (j *JSONHub) serve() {
	http.HandleFunc("/health", j.jsonHealth)
	http.HandleFunc("/update", j.updateHealth)

	fmt.Println("JSONHub serving and listening on port", globals.HubPort)

	if err := j.srv.ListenAndServe(); err != http.ErrServerClosed {
		panic(err)
	}

	fmt.Println("JSONHub closing")
}

func (j *JSONHub) jsonHealth(w http.ResponseWriter, req *http.Request) {
	bytes, err := json.Marshal(j.chart)
	assist.Check(err)

	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func (j *JSONHub) updateHealth(w http.ResponseWriter, req *http.Request) {
	bodyBytes, err := assist.SafeBytes(req.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("could not decode JSON body"))
		return
	}
	var sys health.System

	err = json.Unmarshal(bodyBytes, &sys)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("could not decode JSON body to system"))
		return
	}

	err = j.chart.Update(&sys)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("system is unknown"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Status OK"))
}

// NewJSONHub initializes a JSONHub with a component with
// black status for each passed in component name. The
// JSONHub starts serving on initialization.
func NewJSONHub(componentNames []globals.Component) *JSONHub {
	allComponents := map[globals.Component]*health.SystemStamp{}

	for _, name := range componentNames {
		allComponents[name] = health.NewSystemStamp(
			name,
			health.StatusBlack,
		)
	}

	j := &JSONHub{
		&http.Server{Addr: globals.HubPort},
		*health.NewChart(allComponents),
	}

	j.work()

	return j
}
