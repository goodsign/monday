package monday

import (
	"reflect"
	"testing"
)

type DateLayoutTest struct {
	name     string
	input    string
	expected []dateStringLayoutItem
}

var dateLayoutTests = []DateLayoutTest{
	{"ANSIC", "Thu Feb  4 21:00:57 2010", []dateStringLayoutItem{
		{"Thu", true, false},
		{" ", false, false},
		{"Feb", true, false},
		{"  ", false, false},
		{"4", false, true},
		{" ", false, false},
		{"21", false, true},
		{":", false, false},
		{"00", false, true},
		{":", false, false},
		{"57", false, true},
		{" ", false, false},
		{"2010", false, true},
	}},
	{"RFC1123Z", "Wed, 04 Feb 2013 21:00:57 -0800", []dateStringLayoutItem{
		{"Wed", true, false},
		{", ", false, false},
		{"04", false, true},
		{" ", false, false},
		{"Feb", true, false},
		{" ", false, false},
		{"2013", false, true},
		{" ", false, false},
		{"21", false, true},
		{":", false, false},
		{"00", false, true},
		{":", false, false},
		{"57", false, true},
		{" -", false, false},
		{"0800", false, true},
	}},
}

func TestLayoutParser(t *testing.T) {
	for i, ts := range dateLayoutTests {
		items := stringToLayoutItems(ts.input)

		if !reflect.DeepEqual(items, ts.expected) {
			t.Errorf("Test #%d (%s) failed.\n         Got: %#v\n    Expected: %#v\n", i, ts.name, items, ts.expected)
		}
	}
}
