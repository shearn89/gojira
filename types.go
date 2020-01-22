package gojira

type State struct {
  Transitions []string `json:transitions,omitempty`
  IsInitialState bool `json:initial,omitempty`
}

type StateModel map[string]State

type Ticket struct {
  Title string
  State string
}
