package monday

import (
	"testing"
	"unicode"
)

// A word must not mix letters from Latin, Cyrillic and Greek: lookalike
// letters (Latin "t" in Cyrillic "Чт") break parsing of text typed by people.
func TestKnownWordsDoNotMixAlphabets(t *testing.T) {
	tables := map[string]map[Locale]map[string]string{
		"days short":            knownDaysShort,
		"days long":             knownDaysLong,
		"months short":          knownMonthsShort,
		"months long":           knownMonthsLong,
		"months genitive short": knownMonthsGenitiveShort,
		"months genitive long":  knownMonthsGenitiveLong,
		"periods":               knownPeriods,
	}
	scripts := []*unicode.RangeTable{unicode.Latin, unicode.Cyrillic, unicode.Greek}

	for name, table := range tables {
		for locale, words := range table {
			for key, word := range words {
				found := 0
				for _, script := range scripts {
					for _, r := range word {
						if unicode.Is(script, r) {
							found++
							break
						}
					}
				}
				if found > 1 {
					t.Errorf("%s %s %q: %q mixes alphabets", locale, name, key, word)
				}
			}
		}
	}
}
