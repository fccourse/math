package main

import (
	"fmt"
	"testing"
)

func TestSoma(t *testing.T) {
  total := Soma(15, 15)
  if total != 30 {
    fmt.Errorf("Result must be 30 not %d", total)
  }
}
