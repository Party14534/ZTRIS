package main

import (
	"fmt"
	"ztris/gameFunctions"
	"github.com/rivo/tview"
)

func main() {
  app := tview.NewApplication()

  textView := tview.NewTextView().
    SetText("|     |").
    SetTextAlign(tview.AlignCenter).
    SetDynamicColors(true)

  textView.SetBorder(true)

  
  app = app.SetRoot(textView, true).SetInputCapture(gamefunctions.GetInput); 

  go gamefunctions.RunGame(textView, app)

  if err := app.Run(); err != nil {
    app.Stop()
    panic(err)
  }

  fmt.Println("Your Score: ", gamefunctions.Score);
}
