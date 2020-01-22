package gojira

import "testing"


func Test_TransitionTicket(t *testing.T) {
  sm,err := LoadStateModel("config/stateModel.json")
  if err != nil {
    t.Errorf("failed to load state model")
  }
  var cases = []struct {
    ticket Ticket
    target string
    expected bool
  }{
    { Ticket{Title:"test ticket", State:"raised"}, "open", false },
    { Ticket{Title:"test ticket", State:"raised"}, "foo", true },
  }

  for _,tt := range cases {
    err = tt.ticket.Transition(*sm, tt.target)
    if tt.expected && err == nil {
      t.Errorf("expected an error and got none")
    }
    if (!tt.expected && tt.ticket.State != tt.target) {
      t.Errorf("expected transition to update state, but it didn't")
    }
  }
}

