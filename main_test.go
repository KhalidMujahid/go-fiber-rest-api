package main

import "testing"

func sum(a int,b int) int {
  return a + b
}

func Test(t *testing.T) {
  sum := sum(2,3)

  if sum != 5 {
    t.Error("Expected ", 6 ," got ",sum) 
  }
}
