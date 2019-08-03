package health

import "github.com/lewington/listener/globals"

// System represents an isolated component
// of the overall betting software.
type System struct {
	Name   globals.Component
	Status Status
}

// Stamp creates a new SystemStamp from the current
// system.
func (s *System) Stamp() *SystemStamp {
	return NewSystemStamp(s.Name, s.Status)
}
