package gojira

import (
  "errors"
  "os"
  "encoding/json"
  "io/ioutil"
)

var (
  TransitionError = errors.New("invalid transition requested for state")
)

func (s *State) IsValidTransition(target string) bool {
  for _,t := range(s.Transitions) {
    if t == target {
      return true
    }
  }
  return false
}

func LoadStateModel(path string) (*StateModel,error) {
  // open the file
  jsonFile,err := os.Open(path)
  if err != nil {
    return nil, err
  }
  defer jsonFile.Close()

  // read as byte array
  bv,err := ioutil.ReadAll(jsonFile)
  if err != nil {
    return nil, err
  }

  var sm StateModel

  err = json.Unmarshal(bv, &sm)
  if err != nil {
    return nil, err
  }

  return &sm, nil
}
