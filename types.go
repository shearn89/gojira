package gojira

// State - a simple representaion of a state and what it is allowed to transition to.
type State struct {
	Transitions    []string `json:"transitions,omitempty"`
	IsInitialState bool     `json:"initial,omitempty"`
}

// StateModel - convenience type wrapper for map.
type StateModel map[string]State

// Ticket - a simple ticket representation.
type Ticket struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	State string `json:"state"`
}
