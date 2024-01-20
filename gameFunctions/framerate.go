package gamefunctions

import (
  "time"
)

var beginning = time.Now()
// var conversion = 1000000000

func framePassed(framerate int16) (bool) {
  targetDuration := time.Second / 60
  duration := time.Now().Sub(beginning)

  if targetDuration <= duration {
    beginning = time.Now()
    return true
  }

  return false
}
