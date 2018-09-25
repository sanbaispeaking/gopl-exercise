// Package intset provides a set of integers based on a bit vector.
package main

import (
	"bytes"
	"fmt"
)

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// AddAll adds a list of non-negative values to the set.
func (s *IntSet) AddAll(seq ...int) {
	for _, x := range seq {
		s.Add(x)
	}
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// IntersectWith sets s to the intersection of s and t
func (s *IntSet) IntersectWith(t *IntSet) {
	for i := range s.words {
		if i < len(t.words) {
			s.words[i] &= t.words[i]
		}
	}
}

// DifferenceWith sets s to the diff of s and t
func (s *IntSet) DifferenceWith(t *IntSet) {
	sDup := s.Copy()
	sDup.IntersectWith(t)
	for i := range s.words {
		s.words[i] &= ^sDup.words[i]
	}

}

// SymmetricDifference returns a new set of elements present at either s or t
func (s *IntSet) SymmetricDifference(t *IntSet) *IntSet {
	longer, shorter := s, t
	if len(longer.words) < len(shorter.words) {
		longer, shorter = shorter, longer
	}
	longer = longer.Copy()
	for i := range longer.words {
		if i < len(shorter.words) {
			longer.words[i] ^= shorter.words[i]
		}
	}

	return longer
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// Len reports number of elements
func (s *IntSet) Len() int {
	var count int
	for _, word := range s.words {
		count += popcount(word)
	}
	return count
}

// Remove removes the non-negative value x from the set.
func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	l := len(s.words)
	if word <= l {
		s.words[word] &= ^(1 << bit)
		if s.words[word] == 0 && word == l {
			s.words = s.words[:l-1]
		}
	}
}

// Clear drop all elemtents
func (s *IntSet) Clear() {
	s.words = s.words[:0]
}

// Copy makes a deep copy of the set
func (s *IntSet) Copy() *IntSet {
	var ns = IntSet{make([]uint64, len(s.words))}
	copy(ns.words, s.words)
	return &ns
}

// Copyied from Section 2.6.2
//!+
// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// popcount returns the population count (number of set bits) of x.
func popcount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

//!-

func main() {
	var s, t IntSet
	s.AddAll(1, 2, 3, 4)
	t.AddAll(2, 3, 5, 6)
	fmt.Println("The symmetric diff of s and t: ", s.SymmetricDifference(&t))

	sDup := s.Copy()
	sDup.IntersectWith(&t)
	fmt.Println("The intersection of s and t: ", sDup)

	s.UnionWith(&t)
	s.DifferenceWith(&t)
	fmt.Println("(s union t) - t: ", &s)
}
