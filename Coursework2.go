///////////////////////////////////////////////////////

// Theory of Computation Assignment 1 Section 1

// NAME:

// STUDENT NUMBER:

///////////////////////////////////////////////////////


package main

import "fmt"

// Booleans

type V struct { }

type BoolPair struct { t chan V; f chan V }

type Bool = chan BoolPair

func True(x Bool) {
  p := <- x ; p.t <- V{} 
}

func False(x Bool) {
  p := <- x ; p.f <- V{} 
}

func Not(x Bool, y Bool) {
  p := <- x
  t, f := make(chan V), make(chan V)
  y <- BoolPair { t, f }
  select {
    case <- t: p.f <- V{}
    case <- f: p.t <- V{}
  } 
}

func decodeBool(x Bool) bool {
  t := make(chan V)
  f := make(chan V)
  x <- BoolPair{ t, f }
  select {
    case <- t: return true
    case <- f: return false
  }
}

// Natural numbers

type NatPair struct { z chan V; s chan chan NatPair }

type Nat = chan NatPair

func Z(x Nat) {
  p := <- x; p.z <- V{}
}

func S(n func(Nat)) func(Nat) {
  return func (x Nat) {
    p := <- x
    y := make(Nat)
    p.s <- y
    go n(y)
  }
}

func Inc(x Nat, y Nat) {
  /////////////////////////////////////////
  // Question 1: complete this definition
  /////////////////////////////////////////
}

func decodeNat(x Nat) int {
  z := make(chan V)
  s := make(chan Nat)
  x <- NatPair{ z, s }
  select {
    case <- z: return 0
    case y := <- s: return decodeNat(y)+1
  }
}

// Lists of natural numbers

type ListPair struct { n chan V; c chan NatListPair }

type NatListPair struct { v Nat; t chan ListPair }

type List = chan ListPair

func Nil(x List) {
  p := <- x; p.n <- V{}
}

func Cons(h func(Nat), t func(List)) func(List) {
  return func(x List) {
    p := <- x
    y := make(List)
    z := make(Nat)
    nlp := NatListPair{ z, y }
    p.c <- nlp
    go h(z)
    go t(y)
  }
}

func decodeList(x List) []int {
  n := make(chan V)
  c := make(chan NatListPair)
  x <- ListPair{ n, c }
  select {
    case <- n: return []int{}
    case nlp := <- c: return append(decodeList(nlp.t),decodeNat(nlp.v))
  }
}

// Parameterise Even by the trigger channel of the process

type EvenTrigger struct { b Bool; n Nat }

func Even(trigger chan EvenTrigger) {
  p := <- trigger
  z := make(chan V)
  s := make(chan Nat)
  p.n <- NatPair{ z, s }
  select {
    case <- z:
      go True(p.b)
    case m := <- s:
      c := make(Bool)
      go Not(p.b,c)
      go Even(trigger)
      trigger <- EvenTrigger{ c, m }
  }
}

// Question 4: define the function Length

// Main

func main() {

  ////////////////////////////////////////////////////////////////////////////////////////////////////
  // Question 2: write code to check that 1 + 1 = 2
  //
  // Question 3: write code to construct a list containing 0, 1 and 2, and use decodeList to print it.
  //
  // Question 5: write code to check that the length of this list is 3.
  ////////////////////////////////////////////////////////////////////////////////////////////////////

}
