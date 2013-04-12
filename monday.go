package monday

import (
    "time"
)

// internalFormatFunc is a preprocessor for default time.Format func
type internalFormatFunc func (value, layout string) string

var internalFormatFuncs = map[Locale]internalFormatFunc {
    LocaleEnUS: createCommonFormatFunc(LocaleEnUS),
    LocaleRuRU: formatRuRU,
}

// internalParseFunc is a preprocessor for default time.ParseInLocation func
type internalParseFunc func (layout, value string) string

var internalParseFuncs = map[Locale]internalParseFunc {
    LocaleEnUS: createCommonParseFunc(LocaleEnUS),
    LocaleRuRU: createCommonParseFunc(LocaleRuRU),
}

var knownWords = map[Locale]map[string][]string {}        // Mapping for 'Format'
var knownWordsReverse = map[Locale]map[string][]string {} // Mapping for 'Parse'

func init() {
    fillKnownWords()
}

func fillKnownWords() {

    // En_US: English (United States)
    fillKnownWordsForLocale(longDayNamesEnUS, LocaleEnUS)
    fillKnownWordsForLocale(shortDayNamesEnUS, LocaleEnUS)
    fillKnownWordsForLocale(longMonthNamesEnUS, LocaleEnUS)
    fillKnownWordsForLocale(shortMonthNamesEnUS, LocaleEnUS)

    // Ru_RU: Russian (Russia) 
    fillKnownWordsForLocale(longDayNamesRuRU, LocaleRuRU)
    fillKnownWordsForLocale(shortDayNamesRuRU, LocaleRuRU)
    fillKnownWordsForLocale(longMonthNamesRuRU, LocaleRuRU)
    fillKnownWordsForLocale(longMonthNamesGenitiveRuRU, LocaleRuRU)
    fillKnownWordsForLocale(shortMonthNamesRuRU, LocaleRuRU)
    fillKnownWordsForLocale(shortMonthNamesGenitiveRuRU, LocaleRuRU)
}

func fillKnownWordsForFormat(src map[string]string, locale Locale) {
    loc, ok := knownWords[locale]

    if !ok {
        loc = make(map[string][]string)
        knownWords[locale] = loc
    }

    for k, v := range src {
        loc[k] = append(loc[k], v)
    }
}

func fillKnownWordsForParse(src map[string]string, locale Locale) {
    loc, ok := knownWordsReverse[locale]

    if !ok {
        loc = make(map[string][]string)
        knownWordsReverse[locale] = loc
    }

    for k, v := range src {
        loc[v] = append(loc[v], k)
    }
}

func fillKnownWordsForLocale(src map[string]string, locale Locale) {
    fillKnownWordsForParse(src, locale)
    fillKnownWordsForFormat(src, locale)    
}

// Format is the standard time.Format wrapper, that replaces known standard 'time' package
// identifiers for months and days to their equivalents in the specified language.
//
// Values of variables 'longDayNames', 'shortDayNames', 'longMonthNames', 'shortMonthNames'
// from file 'time/format.go' (revision 'go1') are chosen as the 'known' words.
//
// Some languages have specific behavior, e.g. in Russian language 
// month names have different suffix when they are presented stand-alone (i.e. in a list or something) 
// and yet another one when they are part of a formatted date. 
// So, even though March is "Март" in Russian, correctly formatted today's date would be: "7 Марта 2007"
// Thus, some transformations for some languages may be a bit more complex than just plain replacements.
func Format(dt time.Time, layout string, locale Locale) string {
    intFunc := internalFormatFuncs[locale]
    fm := dt.Format(layout)
    return intFunc(fm, layout)
}

// ParseInLocation is the standard time.ParseInLocation wrapper, which replaces
// known month/day translations for a specified locale back to English before
// calling time.ParseInLocation. So, you can parse localized dates with this wrapper.
func ParseInLocation(layout, value string, loc *time.Location, locale Locale) (time.Time, error) {
   intFunc := internalParseFuncs[locale]
   pl := intFunc(layout, value)
   return time.ParseInLocation(layout, pl, loc)
}