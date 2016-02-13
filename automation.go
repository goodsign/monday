package monday

import (
	"gopkg.in/fatih/set.v0"
	"regexp"
	"strings"
	"time"
)

var wordsRx = regexp.MustCompile("(\\p{L}+)")

type LocaleDetector struct {
	localeMap  map[string]*set.Set
	lastLocale Locale
}

func (this *LocaleDetector) addWords(words []string, v Locale) {
	for _, w := range words {
		l := strings.ToLower(w)
		if _, ok := this.localeMap[w]; !ok {
			this.localeMap[w] = set.New(v)
			if l != w {
				this.localeMap[l] = set.New(v)
			}
		} else {
			this.localeMap[w].Add(v)
			if l != w {
				this.localeMap[l].Add(v)
			}
		}
	}
}

func NewLocaleDetector() *LocaleDetector {
	this := &LocaleDetector{localeMap: make(map[string]*set.Set), lastLocale: LocaleEnGB}
	for _, v := range ListLocales() {
		days := GetShortDays(v)
		this.addWords(days, v)
		days = GetLongDays(v)
		this.addWords(days, v)
		months := GetShortMonths(v)
		this.addWords(months, v)
		months = GetLongMonths(v)
		this.addWords(months, v)
	}
	return this
}

func (this *LocaleDetector) Parse(layout, value string) (time.Time, error) {
	// t, e := ParseInLocation(layout, value, time.UTC, this.lastLocale)
	// if e == nil {
	// 	return t, nil
	// }
	this.lastLocale = this.detectLocale(value)
	return ParseInLocation(layout, value, time.UTC, this.lastLocale)
}

func (this *LocaleDetector) detectLocale(value string) Locale {
	var localesMap map[Locale]int = make(map[Locale]int)
	for _, v := range wordsRx.FindAllStringSubmatchIndex(value, -1) {
		word := strings.ToLower(value[v[0]:v[1]])
		// fmt.Printf("--word:'%s'\n", word)
		if localesSet, ok := this.localeMap[word]; ok {
			localesSet.Each(func(i interface{}) bool {
				if loc, ok := i.(Locale); ok {
					//      fmt.Printf("\tinc %s\n", loc)
					if _, ok := localesMap[loc]; !ok {
						localesMap[loc] = 1
					} else {
						localesMap[loc]++
					}
				}
				return true
			})
		}
	}
	var result Locale = LocaleEnUS
	var frequency int = 0
	for key, counter := range localesMap {
		if counter > frequency {
			frequency = counter
			result = key
		}
	}
	return result
}
