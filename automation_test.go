package monday

import (
	"fmt"
	"testing"
)

type layoutData struct {
	layout    string
	matches   []string
	unmatches []string
	locales   []Locale
}

var (
	testingLayoutsData []layoutData = []layoutData{
		layoutData{
			layout: "Mon, 2 Jan 2006 15:4:5 -0700",
			matches: []string{
				"Вт, 2 Янв 1997 12:01:44 -0330",
				"Tue, 12 Feb 2014 22:03:40 -0300",
			},
			unmatches: []string{
				"Mon, 2 Jan 2006 15:4:5 -0700 ",
				"Mon, 2  Jan 2006 15:4:5 -0700",
				"Mon, 2 Jan 2006 15:4:5",
				"Mon, 2 Jan 2006 15:4:5 4 -0700",
			},
			locales: []Locale{
				LocaleRuRU,
				LocaleEnUS,
			},
		},
		layoutData{
			layout: "Mon  2 Jan 2006 15:4:5 -0700",
		},
		layoutData{
			layout: "Пон, 2 Янв 2006 15:4:5 -0700",
		},
		// layoutData{
		// 	layout: "Mon, 2 Jan 2006 15:4:5 -0700",
		// },
		// layoutData{
		// 	layout: "Mon, 2 Jan 2006 15:4:5 -0700",
		// },
		// layoutData{
		// 	layout: "Mon, 2 Jan 2006 15:4:5 -0700",
		// },
		// layoutData{
		// 	layout: "Mon, 2 Jan 2006 15:4:5 -0700",
		// },
	}
)

func TestLayoutValidator(t *testing.T) {
	ld := NewLocaleDetector()
	for _, ltd := range testingLayoutsData {
		ld.prepareLayout(ltd.layout)
		for i, m := range ltd.matches {
			if !ld.validateValue(ltd.layout, m) {
				t.Errorf("'%s' not matches to '%s' last error position = %d\n", m, ltd.layout, ld.errorPosition())
				fmt.Printf("!\n")
			} else {
				fmt.Printf("'%s' matches to '%s'..OK\n", m, ltd.layout)
			}
			var locale Locale = ld.detectLocale(m)
			if locale != ltd.locales[i] {
				t.Errorf("locales detect error, expected '%s', result '%s'\n", ltd.locales[i], locale)
			} else {
				fmt.Printf("detect locale for '%s': expected '%s', result '%s'\n", m, ltd.locales[i], locale)
			}
		}
		for _, u := range ltd.unmatches {
			if ld.validateValue(ltd.layout, u) {
				t.Errorf("'%s' matches to '%s'\n", u, ltd.layout)
			} else {
				fmt.Printf("'%s' not matches to '%s'..OK\n", u, ltd.layout)
			}
		}
	}
}

func TestLocaleDetector(t *testing.T) {

}

func TestParsing(t *testing.T) {

}
