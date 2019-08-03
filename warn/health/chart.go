package health

import (
	"fmt"
	"sync"

	"github.com/lewington/listener/assist"
	"github.com/lewington/listener/globals"
)

// Chart represents the health og the overall
// system.
type Chart struct {
	mutex   *sync.Mutex
	Systems map[globals.Component]*SystemStamp
}

// Update overwrites the existing system of
// the same name with the given system. If no
// such system exists, an error is returned.
func (c *Chart) Update(sys *System) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	for name := range c.Systems {
		if name == sys.Name {
			c.Systems[name] = sys.Stamp()
			return nil
		}
	}

	return fmt.Errorf("no existing systems matching %v", sys)
}

// MarkOldSystemsAsRed maeks all systems whose last update
// was more than a certain time ago as red.
func (c *Chart) MarkOldSystemsAsRed() {
	c.mutex.Lock()
	for _, system := range c.Systems {
		if system.Timestamp.Before(assist.Timestamp().Add(-globals.HealthCheckWait)) {
			system.Status = StatusRed
		}
	}
	c.mutex.Unlock()
}

// NewChart initializes a Chart.
func NewChart(systems map[globals.Component]*SystemStamp) *Chart {
	return &Chart{
		&sync.Mutex{},
		systems,
	}
}
