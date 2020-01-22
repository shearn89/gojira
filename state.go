package gojira

import (
	"encoding/json"
	"io/ioutil"
	"os"

	log "github.com/sirupsen/logrus"
)

// IsValidTransition - checks if the target state is a valid transition from this one.
func (s *State) IsValidTransition(target string) bool {
	for _, t := range s.Transitions {
		if t == target {
			return true
		}
	}
	return false
}

// LoadStateModel - loads the state model from a json file.
func LoadStateModel(path string) (*StateModel, error) {
	log.WithFields(log.Fields{"path": path}).Debug("loading state model from file")
	// open the file
	jsonFile, err := os.Open(path)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer jsonFile.Close()

	// read as byte array
	bv, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	var sm StateModel

	err = json.Unmarshal(bv, &sm)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	log.WithFields(log.Fields{"path": path}).Debug("loaded state model from file")
	return &sm, nil
}
