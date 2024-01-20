package gamefunctions

import (
	"testing"
	"ztris/gameFunctions"
)

func TestBagWorks(t *testing.T) {
  t.Run("Generates unique values after N times", func(t *testing.T) {
    N := 1000

    var tTypes [][7]int
    for i := 0; i < N; i++ {
      var row [7]int
      
      for j := 0; j < 7; j++ {
        row[j] = gamefunctions.RandTetronimo()
      }
      tTypes = append(tTypes, row)
    }
    
    var allUnique bool = true
    for i := 0; i < N; i++ {
      if !uniqueValues(t, tTypes[i]) {
        allUnique = false
        break
      }
    }

    if !allUnique {
      t.Error("Not all values were unique\n")
    }
  
  }) 
}

func BenchmarkBag(b *testing.B) {
  for i := 0; i < b.N; i++ {
    gamefunctions.RandTetronimo()
  }
}

func uniqueValues(t testing.TB, slice [7]int) bool {
  t.Helper()
  m := make(map[int]bool)

  for _, val := range slice {
    _, ok := m[val]
    if ok { return false }
    m[val] = true
  }

  return true
}
