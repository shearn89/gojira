package gojira

import "testing"

func Test_IsValidTransition(t *testing.T) {
  var cases = []struct {
    start State
    target string
    expected bool
  }{
    { State{[]string{"out"}}, "out", true },
    { State{[]string{"out"}}, "foo", false },
  }

  for _, tt := range cases {
    result := tt.start.IsValidTransition(tt.target)
    if result != tt.expected {
      t.Errorf("got %v, wanted %v", result, tt.expected)
    }
  }
}

func Test_LoadStateModel(t *testing.T) {
  var cases = []struct {
    path string
    expected bool
  }{
    { "config/stateModel.json", false },
    { "test/notHere.json", true },
    { "test/invalid.json", true },
  }

  for _, tt := range cases {
    result,err := LoadStateModel(tt.path)
    if tt.expected && err == nil {
      t.Errorf("expected an error and got none")
    }
    if result == nil && !tt.expected {
      t.Errorf("got no StateModel when we expected one")
    }
  }
}
