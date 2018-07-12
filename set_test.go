package monday

import (
	"strconv"
	"testing"
)

func TestSet_Add(t *testing.T) {
	s := newSet()
	s.Add(LocaleEnUS)
	s.Add(LocaleEnGB)
	s.Add(LocaleEnGB) // duplicate
	if !s.Has(LocaleEnGB) {
		t.Error("Add: added item not available in the set")
	}
	if !s.Has(LocaleEnUS, LocaleEnGB) {
		t.Error("Add: added items are not availabile in the set.")
	}
}

func TestSet_RaceAdd(t *testing.T) {
	// Create two sets. Add concurrently items to each of them. Remove from the
	// other one.
	// "go test -race" should detect this if the library is not thread-safe.
	s := newSet()
	u := newSet()

	go func() {
		for i := 0; i < 1000; i++ {
			item := "item" + strconv.Itoa(i)
			go func(i int) {
				s.Add(Locale(item))
				u.Add(Locale(item))
			}(i)
		}
	}()

	for i := 0; i < 1000; i++ {
		item := "item" + strconv.Itoa(i)
		go func(i int) {
			s.Add(Locale(item))
			u.Add(Locale(item))
		}(i)
	}
}
