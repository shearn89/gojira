package gojira

import (
  "errors"

  log "github.com/sirupsen/logrus"
)

var (
  TransitionError = errors.New("invalid transition requested for state")
)

// Ticket methods
func (t *Ticket) Transition(sm StateModel, target string) error {
  log.Debug("transitioning ticket")
  log.WithFields(log.Fields{
    "ticket": t.Title,
    "current": t.State,
    "target": target,
  }).Info("transitioning ticket")

  current := sm[t.State]
  if current.IsValidTransition(target) {
    log.WithFields(log.Fields{"transition": target}).Debug("transition is valid")
    t.State = target
  } else {
    log.Error(TransitionError)
    return TransitionError
  }

  return nil
}
