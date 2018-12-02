package tests

import (
	"testing"

	orderedset "github.com/enixdark/dag-block/lib/utils"
)

func TestOrderedSet(t *testing.T) {
	s := orderedset.NewOrderedSet()

	if !s.Empty() {
		t.Fatalf("New set expected to be empty but it is not")
	}

	if s.Size() != 0 {
		t.Fatalf("New set expected to have 0 elements but got %d", s.Size())
	}
}

func TestOrderedSet_Add(t *testing.T) {
	s := orderedset.NewOrderedSet()
	s.Add("e", "f", "g", "c", "d", "x", "b", "a")
	s.Add("b") //overwrite
	structValue := customType{"svalue"}
	s.Add(structValue)
	s.Add(&structValue)
	s.Add(true)

	actualOutput := s.Values()
	expectedOutput := []interface{}{"e", "f", "g", "c", "d", "x", "b", "a", structValue, &structValue, true}
	if !sameElements(actualOutput, expectedOutput) {
		t.Errorf("Got %v expected %v", actualOutput, expectedOutput)
	}
}

func TestOrderedSet_Remove(t *testing.T) {
	s := orderedset.NewOrderedSet()
	s.Add("e", "f", "g", "c", "d", "x", "b", "a")
	s.Add("b") //overwrite
	structValue := customType{"svalue"}
	s.Add(structValue)
	s.Add(&structValue)
	s.Add(true)

	s.Remove("f", "g", &structValue, true)

	actualOutput := s.Values()
	expectedOutput := []interface{}{"e", "c", "d", "x", "b", "a", structValue}
	if !sameElements(actualOutput, expectedOutput) {
		t.Errorf("Got %v expected %v", actualOutput, expectedOutput)
	}

	// already removed, doesn't fails.
	s.Remove("f", "g", &structValue, true)
	actualOutput = s.Values()
	expectedOutput = []interface{}{"e", "c", "d", "x", "b", "a", structValue}
	if !sameElements(actualOutput, expectedOutput) {
		t.Errorf("Got %v expected %v", actualOutput, expectedOutput)
	}
}

func TestOrderedSet_Contains(t *testing.T) {
	s := orderedset.NewOrderedSet()
	s.Add("e", "f", "g", "c", "d", "x", "b", "a")
	s.Add("b") //overwrite
	structValue := customType{"svalue"}
	s.Add(structValue)
	s.Add(&structValue)
	s.Add(true)

	table := []struct {
		input          []interface{}
		expectedOutput bool
	}{
		{[]interface{}{"c", "d", &structValue}, true},
		{[]interface{}{"c", "d", nil}, false},
		{[]interface{}{true}, true},
		{[]interface{}{structValue}, true},
	}

	for _, test := range table {
		actualOutput := s.Contains(test.input...)
		if actualOutput != test.expectedOutput {
			t.Errorf("Got %v expected %v", actualOutput, test.expectedOutput)
		}
	}
}

func TestOrderedSet_Empty(t *testing.T) {
	s := orderedset.NewOrderedSet()

	if empty := s.Empty(); !empty {
		t.Errorf("Got %v expected %v", empty, true)
	}

	s.Add("e", "f", "g", "c", "d", "x", "b", "a")
	if empty := s.Empty(); empty {
		t.Errorf("Got %v expected %v", empty, false)
	}

	s.Remove("e", "f", "g", "c", "d", "x", "b", "a")
	if empty := s.Empty(); !empty {
		t.Errorf("Got %v expected %v", empty, true)
	}
}

func TestOrderedSet_Values(t *testing.T) {
	s := orderedset.NewOrderedSet()
	s.Add("e", "f", "g", "c", "d", "x", "b", "a")
	s.Add("b") //overwrite
	structValue := customType{"svalue"}
	s.Add(structValue)
	s.Add(&structValue)
	s.Add(true)

	actualOutput := s.Values()
	expectedOutput := []interface{}{"e", "f", "g", "c", "d", "x", "b", "a", structValue, &structValue, true}
	if !sameElements(actualOutput, expectedOutput) {
		t.Errorf("Got %v expected %v", actualOutput, expectedOutput)
	}
}

func TestOrderedSet_Size(t *testing.T) {
	s := orderedset.NewOrderedSet()

	if size := s.Size(); size != 0 {
		t.Errorf("Got %v expected %v", size, 0)
	}

	s.Add("e", "f", "g", "c", "d", "x", "b", "a")
	s.Add("e", "f", "g", "c", "d", "x", "b", "a", "z") // overwrite
	if size := s.Size(); size != 9 {
		t.Errorf("Got %v expected %v", size, 9)
	}

	s.Remove("e", "f", "g", "c", "d", "x", "b", "a", "z")
	if size := s.Size(); size != 0 {
		t.Errorf("Got %v expected %v", size, 0)
	}
}

func TestOrderedSet_String(t *testing.T) {
	s := orderedset.NewOrderedSet()

	s.Add("foo", "bar")
	expected := "[foo bar]"
	result := s.String()
	if expected != result {
		t.Fatalf("OrderedSet_Stringer expected to be %q but got %q", expected, result)
	}
}
