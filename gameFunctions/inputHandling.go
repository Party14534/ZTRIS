package gamefunctions

import (
	"github.com/gdamore/tcell/v2"
)

func GetInput(event *tcell.EventKey) *tcell.EventKey {
  
  if event.Key() == tcell.KeyDown {
    dropping = true
  } else {
    dropping = false
  }

  if event.Rune() == ' ' {
    t.dropTetronimo()
  } else if event.Key() == tcell.KeyLeft {
    t.moveTetronimo(false)
  } else if event.Key() == tcell.KeyRight {
    t.moveTetronimo(true)
  } else if event.Rune() == 'z' {
    t.rotateLeft()
  } else if event.Rune() == 'x' {
    t.rotateRight()
  }  
  return event
}
