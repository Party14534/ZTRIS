package gamefunctions

import (
	"github.com/rivo/tview"
)

var t *Tetronimo
var state *gameState

func RunGame(textView *tview.TextView, app *tview.Application) {
  state = new(gameState)
  t = CreateTetronimo(RandTetronimo())

  for { if framePassed(60) {
    // Update the active tetronimo
    UpdateTetronimo(app)
    
    // Update the view
    updateView(textView)  

    // Draw to screen
    app.SetRoot(textView, true).SetFocus(textView).Draw()
  
  } } 
}
