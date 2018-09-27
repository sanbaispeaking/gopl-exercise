package intset

import (
	"testing"
)

func TestIntersectWith(t *testing.T) {
	var scenarios = []struct {
		l, r []int
		want string
	}{
		{[]int{}, []int{}, "{}"},
		{[]int{1, 2, 3}, []int{}, "{}"},
		{[]int{}, []int{2, 3, 4}, "{}"},
		{[]int{1, 2, 3}, []int{5, 6, 7}, "{}"},
		{[]int{1, 2, 3}, []int{2, 3, 4, 5}, "{2 3}"},
		{[]int{1, 2, 3, 4}, []int{2, 3, 5}, "{2 3}"},
	}

	for _, scen := range scenarios {
		var l, r IntSet
		l.AddAll(scen.l...)
		r.AddAll(scen.r...)
		l.IntersectWith(&r)
		if l.String() != scen.want {
			t.Errorf("%v & %v = %s, want %s", scen.l, scen.r, &l, scen.want)
		}
	}
}

func TestDifferenceWith(t *testing.T) {
	var scenarios = []struct {
		l, r []int
		want string
	}{
		{[]int{}, []int{}, "{}"},
		{[]int{1, 2, 3}, []int{}, "{1 2 3}"},
		{[]int{}, []int{2, 3, 4}, "{}"},
		{[]int{1, 2, 3}, []int{5, 6, 7}, "{1 2 3}"},
		{[]int{1, 2, 3}, []int{2, 3, 4, 5}, "{1}"},
		{[]int{1, 2, 3, 4}, []int{2, 3, 5}, "{1 4}"},
	}

	for _, scen := range scenarios {
		var l, r IntSet
		l.AddAll(scen.l...)
		r.AddAll(scen.r...)
		l.DifferenceWith(&r)
		if l.String() != scen.want {
			t.Errorf("%v & %v = %s, want %s", scen.l, scen.r, &l, scen.want)
		}
	}
}

func TestSymmetricDifferenceWith(t *testing.T) {
	var scenarios = []struct {
		l, r []int
		want string
	}{
		{[]int{}, []int{}, "{}"},
		{[]int{1, 2, 3}, []int{}, "{1 2 3}"},
		{[]int{}, []int{2, 3, 4}, "{2 3 4}"},
		{[]int{1, 2, 3}, []int{5, 6, 7}, "{1 2 3 5 6 7}"},
		{[]int{1, 2, 3}, []int{2, 3, 4, 5}, "{1 4 5}"},
		{[]int{1, 2, 3, 4}, []int{2, 3, 5}, "{1 4 5}"},
	}

	for _, scen := range scenarios {
		var l, r IntSet
		l.AddAll(scen.l...)
		r.AddAll(scen.r...)
		d := l.SymmetricDifference(&r)
		if d.String() != scen.want {
			t.Errorf("%v & %v = %s, want %s", scen.l, scen.r, d, scen.want)
		}
	}
}
