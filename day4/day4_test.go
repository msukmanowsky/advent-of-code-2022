package day4

import (
	"log"
	"testing"
)

type assignmentTest struct {
	a    assignment
	b    assignment
	want bool
}

var assignmentTests = []assignmentTest{
	{
		a:    assignment{Min: 1, Max: 10},
		b:    assignment{Min: 1, Max: 10},
		want: true,
	},
	{
		a:    assignment{Min: 2, Max: 8},
		b:    assignment{Min: 3, Max: 7},
		want: true,
	},
	{
		a:    assignment{Min: 3, Max: 7},
		b:    assignment{Min: 2, Max: 8},
		want: false,
	},
}

type overlapTest struct {
	a    assignment
	b    assignment
	want bool
}

var overlapTests = []overlapTest{
	{
		a:    assignment{Min: 2, Max: 4},
		b:    assignment{Min: 6, Max: 8},
		want: false,
	},
	{
		a:    assignment{Min: 2, Max: 3},
		b:    assignment{Min: 4, Max: 5},
		want: false,
	},
	{
		a:    assignment{Min: 5, Max: 7},
		b:    assignment{Min: 7, Max: 9},
		want: true,
	},
	{
		a:    assignment{Min: 2, Max: 6},
		b:    assignment{Min: 4, Max: 8},
		want: true,
	},
}

func Test_assignmentContained(t *testing.T) {
	for _, test := range assignmentTests {
		if ans := test.a.contains(test.b); ans != test.want {
			log.Printf("expected %+v.contains(%+v) = %+v (got: %+v)", test.a, test.b, test.want, ans)
			t.Fail()
		}
	}
}

func Test_assignmentOverlaps(t *testing.T) {
	for _, test := range overlapTests {
		if ans := test.a.overlaps(test.b); ans != test.want {
			log.Printf("expected %+v.overlaps(%+v = %+v got(: %+v)", test.a, test.b, test.want, ans)
			t.Fail()
		}
	}
}
