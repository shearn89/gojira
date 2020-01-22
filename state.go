package gojira

import (
  "os"
  "encoding/json"
  "io/ioutil"

  log "github.com/sirupsen/logrus"
)

// State Methods
func (s *State) IsValidTransition(target string) bool {
  for _,t := range(s.Transitions) {
    if t == target {
      return true
    }
  }
  return false
}

func LoadStateModel(path string) (*StateModel,error) {
  log.WithFields(log.Fields{"path": path}).Debug("loading state model from file")
  // open the file
  jsonFile,err := os.Open(path)
  if err != nil {
    log.Error(err)
    return nil, err
  }
  defer jsonFile.Close()

  // read as byte array
  bv,err := ioutil.ReadAll(jsonFile)
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


