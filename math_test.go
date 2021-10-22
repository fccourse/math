package main

import "testing"

func TestSum(t *testing.T) {
  total := Sum(15, 15)
  if total != 30 {
    t.Errorf("Result must be 30 not %d", total)
  }
}

func TestSub(t *testing.T) {
  total := Sub(15, 15)
  if total != 0 {
    t.Errorf("Result must be 0 not %d", total)
  }
}


func TestMul(t *testing.T) {
  total := Mul(15, 15)
  if total != 225 {
    t.Errorf("Result must be 225 not %d", total)
  }
}

func TestDiv(t *testing.T) {
  total := Sum(15, 15)
  if total != 1 {
    t.Errorf("Result must be 1 not %d", total)
  }
}

