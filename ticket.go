package gojira

import (
	"errors"

	log "github.com/sirupsen/logrus"
)

var (
	errTransition = errors.New("invalid transition requested for state")
)

// Transition - update a ticket to a new state.
func (t *Ticket) Transition(sm StateModel, target string) error {
	log.Debug("transitioning ticket")
	log.WithFields(log.Fields{
		"ticket":  t.ID,
		"current": t.State,
		"target":  target,
	}).Info("transitioning ticket")

	current := sm[t.State]
	if current.IsValidTransition(target) {
		log.WithFields(log.Fields{"transition": target}).Debug("transition is valid")
		t.State = target
	} else {
		log.Error(errTransition)
		return errTransition
	}

	return nil
}
