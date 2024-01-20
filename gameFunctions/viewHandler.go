package gamefunctions

import "github.com/rivo/tview"

type gameState struct {
  board [15][10]string
  gameOver bool
}

func (g *gameState) toString() string {
  var s []string

  // Draw to string row by row
  for i := 0; i < 15; i++ {
    line := ""
    for j := 0; j < 10; j++ {
      if g.board[i][j] != "" {
        line += "|X"
      } else {
        line += "| "
      }
    }
    line += "|\n"
    s = append(s, line)
  }

  for _, block := range t.footprint {
    var posY int = int(t.pos.y + block.y)
    var posX int = int(t.pos.x + block.x)
    runeLine := []rune(s[posY])
    runeLine[posX*2] = '|'
    runeLine[posX*2 + 1] = 'O'
    s[posY] = string(runeLine)
  }
  
  var tString string
  for _, line := range s {
    tString += line
  }

  return tString
}

func updateView(textView *tview.TextView) {
  textView.SetText(state.toString())
}
