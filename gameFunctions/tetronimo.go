package gamefunctions

import (
	"math"
	"math/rand"

	"github.com/rivo/tview"
)

var dropTime uint8 = 60
var dropping bool = false
var lockTime uint8 = 0
var playerBag = []int{0,1,2,3,4,5,6}
var defaultBag = []int{0,1,2,3,4,5,6}


type Pos struct {
  x int8
  y int8
}

type Tetronimo struct {
  pos Pos
  footprint [4]Pos
  frameSinceDrop uint8
}

func (t *Tetronimo) hasCollided() bool {
  
  for _, block := range t.footprint {
    if nextToBlock(block.x + t.pos.x, block.y + t.pos.y, state) {
      return true
    }
  }

  return false
}

func nextToBlock(x int8, y int8, state *gameState) bool {
  
  if y + 1 < 15 {
    if state.board[y+1][x] != "" {
      return true
    } 
  } else {
    return true
  }

  return false
}

func SetTetronimo(app *tview.Application) {
  for _, block := range t.footprint {
    posX := block.x + t.pos.x
    posY := block.y + t.pos.y
    if state.board[posY][posX] != "" {
      state.gameOver = true
      app.Stop()
      return
    }
    state.board[block.y + t.pos.y][block.x + t.pos.x] = "X"
  }
  clearLine()
}

func clearLine() {
  board := &state.board
  for i := 14; i >= 0; i-- {
    for j := 0; j < 10; j++ {
      if board[i][j] == "" { break }
      
      // Move rows down by one
      if j == 9 {
        for k := i; k > 0; k-- {
          for l := 0; l < 10; l++ {
            board[k][l] = board[k-1][l]             
          }
        }
        for l := 0; l < 10; l++ {
          board[0][l] = ""
        }

        i++
      }
    }
  } 
}

func (t *Tetronimo) canMoveDir(dir bool, state *gameState) bool {
  var direction int8 = -1
  if dir { direction = 1 }
  
  for _, block := range t.footprint {
    xPos := block.x + t.pos.x
    yPos := block.y + t.pos.y
    if xPos + direction < 0 || xPos + direction >= 10 {
      return false
    } else if state.board[yPos][xPos + direction] != "" {
      return false
    }
  }

  return true
}

func (t *Tetronimo) moveTetronimo(dir bool) {
  if t.canMoveDir(dir, state) {
    var direction int8 = -1
    if dir { direction = 1 }
    t.pos.x += direction
    lockTime = 0
  }
}

func (t *Tetronimo) lowerTetronimo() {
  t.pos.y++
}

func (t *Tetronimo) dropTetronimo() {
  for {
    if t.hasCollided() { return }
    t.lowerTetronimo()
  }
}

func (t *Tetronimo) getFootprintWidth() int8 {
  var min int8 = 100
  var max int8 = -100

  for _, block := range t.footprint {
    if block.x > max {
      max = block.x
    }
    if block.x < min {
      min = block.x
    }
  }

  return max - min + 1
}

func (t *Tetronimo) getFootprintHeight() int8 {
  var min int8 = 100
  var max int8 = -100

  for _, block := range t.footprint {
    if block.y > max {
      max = block.y
    }
    if block.y < min {
      min = block.y
    }
  }

  return max - min + 1
}

func (t *Tetronimo) rotateLeft() {

  boxSize := max( t.getFootprintWidth(), t.getFootprintHeight() )
  if boxSize == 2 { return }
  center := boxSize/2
  
  buffers := t.footprint

  for i := range t.footprint {
    buffer := t.footprint[i]
    newX := -(center - buffer.y)
    newY := (center - buffer.x)
    buffer.x = newX + center
    buffer.y = newY + center

    posY := buffer.y + t.pos.y
    posX := buffer.x + t.pos.x
    if posY >= 0 && posY < 15 {
      if posX >= 0 && posX < 10 {
        if state.board[posY][posX] != "" {
          return
        }
      }
    }

    buffers[i] = buffer
  }

  t.footprint = buffers
  t.correctOutOfBounds()
}

func (t *Tetronimo) correctOutOfBounds() {
  for i := range t.footprint {
    block := &t.footprint[i]
    posX := t.pos.x + block.x
    posY := t.pos.y + block.y

    if posY < 0 {
      t.pos.y += int8(math.Abs(float64(posY)))
    }
    if posY >= 15 {
      t.pos.y -= int8(math.Abs(float64(posY - 14)))
    }

    if posX < 0 {
      t.pos.x += int8(math.Abs(float64(posX)))
    }
    if posX >= 10 {
      t.pos.x -= int8(math.Abs(float64(posX - 9)))
    }
  }
}

func CreateTetronimo(blockType int) (*Tetronimo) {
  t := new(Tetronimo)
  t.pos.x = 4
  t.pos.y = 0
  t.frameSinceDrop = 0

  if blockType == 0 {
    t.footprint[0] = Pos{x: 0, y: 1}
    t.footprint[1] = Pos{x: 1, y: 1}
    t.footprint[2] = Pos{x: 2, y: 1}
    t.footprint[3] = Pos{x: 3, y: 1}
  } else if blockType == 1 {
    t.footprint[0] = Pos{x: 0, y: 0}
    t.footprint[1] = Pos{x: 0, y: 1}
    t.footprint[2] = Pos{x: 1, y: 1}
    t.footprint[3] = Pos{x: 2, y: 1}
  } else if blockType == 2 {
    t.footprint[0] = Pos{x: 0, y: 1}
    t.footprint[1] = Pos{x: 1, y: 1}
    t.footprint[2] = Pos{x: 2, y: 1}
    t.footprint[3] = Pos{x: 2, y: 0}
  } else if blockType == 3 {
    t.footprint[0] = Pos{x: 0, y: 0}
    t.footprint[1] = Pos{x: 1, y: 0}
    t.footprint[2] = Pos{x: 0, y: 1}
    t.footprint[3] = Pos{x: 1, y: 1}
  } else if blockType == 4 {
    t.footprint[0] = Pos{x: 0, y: 1}
    t.footprint[1] = Pos{x: 1, y: 1}
    t.footprint[2] = Pos{x: 1, y: 0}
    t.footprint[3] = Pos{x: 2, y: 0}
  } else if blockType == 5 {
    t.footprint[0] = Pos{x: 0, y: 1}
    t.footprint[1] = Pos{x: 1, y: 1}
    t.footprint[2] = Pos{x: 1, y: 0}
    t.footprint[3] = Pos{x: 2, y: 1}
  } else if blockType == 6 {
    t.footprint[0] = Pos{x: 0, y: 0}
    t.footprint[1] = Pos{x: 1, y: 0}
    t.footprint[2] = Pos{x: 1, y: 1}
    t.footprint[3] = Pos{x: 2, y: 1}
  }

  
  return t
}

func UpdateTetronimo(app *tview.Application) {
  if t.hasCollided() {
    lockTime++
    if lockTime >= 20 {
      SetTetronimo(app)
      *t = *CreateTetronimo(RandTetronimo())
    }
    return
  }

  lockTime = 0

  realDropTime := dropTime
  if dropping {
    realDropTime /= 3
  }

  if t.frameSinceDrop >= realDropTime {
    t.frameSinceDrop = 0
    t.lowerTetronimo()
  } else {
    t.frameSinceDrop++
  }
  
}


func RandTetronimo() int {
  
  if(len(playerBag) == 0) {
    playerBag = append(playerBag, defaultBag...)
  }

  index := rand.Intn(len(playerBag))
  tType := playerBag[index]
  
  playerBag = append(playerBag[:index], playerBag[index+1:]...)  

  return tType 
}

