package gojira

type State struct {
  Transitions []string `json:transitions,omitempty`
}

type StateModel map[string]State
