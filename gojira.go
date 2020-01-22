package gojira

import (
  log "github.com/sirupsen/logrus"
)

func init() {
  log.SetLevel(log.DebugLevel)
}

func main() {
  log.Info("starting app")
}
