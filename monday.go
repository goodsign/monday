package monday

import (
	"fmt"
	"time"
)

// internalFormatFunc is a preprocessor for default time.Format func
type internalFormatFunc func(value, layout string) string

var internalFormatFuncs = map[Locale]internalFormatFunc{
	LocaleEnUS: createCommonFormatFunc(LocaleEnUS),
	LocaleEnGB: createCommonFormatFunc(LocaleEnGB),
	LocaleDaDK: createCommonFormatFunc(LocaleDaDK),
	LocaleNlBE: createCommonFormatFunc(LocaleNlBE),
	LocaleNlNL: createCommonFormatFunc(LocaleNlNL),
	LocaleFrFR: createCommonFormatFunc(LocaleFrFR),
	LocaleFrCA: createCommonFormatFunc(LocaleFrFR),
	LocaleRuRU: createCommonFormatFuncWithGenitive(LocaleRuRU),
	LocaleFiFI: createCommonFormatFuncWithGenitive(LocaleFiFI),
	LocaleDeDE: createCommonFormatFunc(LocaleDeDE),
	LocaleHuHU: createCommonFormatFunc(LocaleHuHU),
	LocaleItIT: createCommonFormatFunc(LocaleItIT),
	LocaleNnNO: createCommonFormatFunc(LocaleNnNO),
	LocaleNbNO: createCommonFormatFunc(LocaleNbNO),
	LocalePtPT: createCommonFormatFunc(LocalePtPT),
	LocalePtBR: createCommonFormatFunc(LocalePtBR),
	LocaleRoRO: createCommonFormatFunc(LocaleRoRO),
	LocaleEsES: createCommonFormatFunc(LocaleEsES),
	LocaleCaES: createCommonFormatFunc(LocaleCaES),
	LocaleSvSE: createCommonFormatFunc(LocaleSvSE),
	LocaleTrTR: createCommonFormatFunc(LocaleTrTR),
	LocaleBgBG: createCommonFormatFunc(LocaleBgBG),
	LocaleZhCN: createCommonFormatFunc(LocaleZhCN),
	LocaleZhTW: createCommonFormatFunc(LocaleZhTW),
	LocaleZhHK: createCommonFormatFunc(LocaleZhHK),
	LocaleJaJP: createCommonFormatFunc(LocaleJaJP),
}

// internalParseFunc is a preprocessor for default time.ParseInLocation func
type internalParseFunc func(layout, value string) string

var internalParseFuncs = map[Locale]internalParseFunc{
	LocaleEnUS: createCommonParseFunc(LocaleEnUS),
	LocaleEnGB: createCommonParseFunc(LocaleEnGB),
	LocaleDaDK: createCommonParseFunc(LocaleDaDK),
	LocaleNlBE: createCommonParseFunc(LocaleNlBE),
	LocaleNlNL: createCommonParseFunc(LocaleNlNL),
	LocaleFrFR: createCommonParseFunc(LocaleFrFR),
	LocaleFrCA: createCommonParseFunc(LocaleFrFR),
	LocaleRuRU: createCommonParsetFuncWithGenitive(LocaleRuRU),
	LocaleFiFI: createCommonParsetFuncWithGenitive(LocaleFiFI),
	LocaleDeDE: createCommonParseFunc(LocaleDeDE),
	LocaleHuHU: createCommonParseFunc(LocaleHuHU),
	LocaleItIT: createCommonParseFunc(LocaleItIT),
	LocaleNnNO: createCommonParseFunc(LocaleNnNO),
	LocaleNbNO: createCommonParseFunc(LocaleNbNO),
	LocalePtPT: parseFuncPtCommon(LocalePtPT),
	LocalePtBR: parseFuncPtCommon(LocalePtBR),
	LocaleRoRO: createCommonParseFunc(LocaleRoRO),
	LocaleEsES: createCommonParseFunc(LocaleEsES),
	LocaleCaES: createCommonParseFunc(LocaleCaES),
	LocaleSvSE: createCommonParseFunc(LocaleSvSE),
	LocaleTrTR: createCommonParseFunc(LocaleTrTR),
	LocaleBgBG: createCommonParseFunc(LocaleBgBG),
	LocaleZhCN: parseFuncZhCommon(LocaleZhCN),
	LocaleZhTW: parseFuncZhCommon(LocaleZhTW),
	LocaleZhHK: parseFuncZhCommon(LocaleZhHK),
	LocaleJaJP: parseFuncJaCommon(LocaleJaJP),
}

var knownDaysShort = map[Locale]map[string]string{}           // Mapping for 'Format', days of week, short form
var knownDaysLong = map[Locale]map[string]string{}            // Mapping for 'Format', days of week, long form
var knownMonthsLong = map[Locale]map[string]string{}          // Mapping for 'Format', months: long form
var knownMonthsShort = map[Locale]map[string]string{}         // Mapping for 'Format', months: short form
var knownMonthsGenitiveShort = map[Locale]map[string]string{} // Mapping for 'Format', special for names in genitive, short form
var knownMonthsGenitiveLong = map[Locale]map[string]string{}  // Mapping for 'Format', special for names in genitive, long form
var knownPeriods = map[Locale]map[string]string{}

// Reverse maps for the same

var knownDaysShortReverse = map[Locale]map[string]string{}           // Mapping for 'Format', days of week, short form
var knownDaysLongReverse = map[Locale]map[string]string{}            // Mapping for 'Format', days of week, long form
var knownMonthsLongReverse = map[Locale]map[string]string{}          // Mapping for 'Format', months: long form
var knownMonthsShortReverse = map[Locale]map[string]string{}         // Mapping for 'Format', months: short form
var knownMonthsGenitiveShortReverse = map[Locale]map[string]string{} // Mapping for 'Format', special for names in genitive, short form
var knownMonthsGenitiveLongReverse = map[Locale]map[string]string{}  // Mapping for 'Format', special for names in genitive, long form
var knownPeriodsReverse = map[Locale]map[string]string{}

func init() {
	fillKnownWords()
}

func fillKnownWords() {

	// En_US: English (United States)
	fillKnownDaysLong(longDayNamesEnUS, LocaleEnUS)
	fillKnownDaysShort(shortDayNamesEnUS, LocaleEnUS)
	fillKnownMonthsLong(longMonthNamesEnUS, LocaleEnUS)
	fillKnownMonthsShort(shortMonthNamesEnUS, LocaleEnUS)

	// En_GB: English (United Kingdom)
	fillKnownDaysLong(longDayNamesEnUS, LocaleEnGB)
	fillKnownDaysShort(shortDayNamesEnUS, LocaleEnGB)
	fillKnownMonthsLong(longMonthNamesEnUS, LocaleEnGB)
	fillKnownMonthsShort(shortMonthNamesEnUS, LocaleEnGB)

	// Da_DK: Danish (Denmark)
	fillKnownDaysLong(longDayNamesDaDK, LocaleDaDK)
	fillKnownDaysShort(shortDayNamesDaDK, LocaleDaDK)
	fillKnownMonthsLong(longMonthNamesDaDK, LocaleDaDK)
	fillKnownMonthsShort(shortMonthNamesDaDK, LocaleDaDK)

	// Nl_BE: Dutch (Belgium)
	fillKnownDaysLong(longDayNamesNlBE, LocaleNlBE)
	fillKnownDaysShort(shortDayNamesNlBE, LocaleNlBE)
	fillKnownMonthsLong(longMonthNamesNlBE, LocaleNlBE)
	fillKnownMonthsShort(shortMonthNamesNlBE, LocaleNlBE)

	// Nl_NL: Dutch (Netherlands)
	fillKnownDaysLong(longDayNamesNlBE, LocaleNlNL)
	fillKnownDaysShort(shortDayNamesNlBE, LocaleNlNL)
	fillKnownMonthsLong(longMonthNamesNlBE, LocaleNlNL)
	fillKnownMonthsShort(shortMonthNamesNlBE, LocaleNlNL)

	// Fi_FI: Finnish (Finland)
	fillKnownDaysLong(longDayNamesFiFI, LocaleFiFI)
	fillKnownDaysShort(shortDayNamesFiFI, LocaleFiFI)
	fillKnownMonthsLong(longMonthNamesFiFI, LocaleFiFI)
	fillKnownMonthsShort(shortMonthNamesFiFI, LocaleFiFI)
	fillKnownMonthsGenitiveLong(longMonthNamesGenitiveFiFI, LocaleFiFI)
	fillKnownMonthsGenitiveShort(shortMonthNamesFiFI, LocaleFiFI)

	// Fr_FR: French (France)
	fillKnownDaysLong(longDayNamesFrFR, LocaleFrFR)
	fillKnownDaysShort(shortDayNamesFrFR, LocaleFrFR)
	fillKnownMonthsLong(longMonthNamesFrFR, LocaleFrFR)
	fillKnownMonthsShort(shortMonthNamesFrFR, LocaleFrFR)

	// Fr_CA: French (France)
	fillKnownDaysLong(longDayNamesFrFR, LocaleFrCA)
	fillKnownDaysShort(shortDayNamesFrFR, LocaleFrCA)
	fillKnownMonthsLong(longMonthNamesFrFR, LocaleFrCA)
	fillKnownMonthsShort(shortMonthNamesFrFR, LocaleFrCA)

	// De_DE: German (Germany)
	fillKnownDaysLong(longDayNamesDeDE, LocaleDeDE)
	fillKnownDaysShort(shortDayNamesDeDE, LocaleDeDE)
	fillKnownMonthsLong(longMonthNamesDeDE, LocaleDeDE)
	fillKnownMonthsShort(shortMonthNamesDeDE, LocaleDeDE)

	// Hu_HU: Hungarian (Hungary)
	fillKnownDaysLong(longDayNamesHuHU, LocaleHuHU)
	fillKnownDaysShort(shortDayNamesHuHU, LocaleHuHU)
	fillKnownMonthsLong(longMonthNamesHuHU, LocaleHuHU)
	fillKnownMonthsShort(shortMonthNamesHuHU, LocaleHuHU)

	// It_IT: Italian (Italy)
	fillKnownDaysLong(longDayNamesItIT, LocaleItIT)
	fillKnownDaysShort(shortDayNamesItIT, LocaleItIT)
	fillKnownMonthsLong(longMonthNamesItIT, LocaleItIT)
	fillKnownMonthsShort(shortMonthNamesItIT, LocaleItIT)

	// Nn_NO: Norwegian Nynorsk (Norway)
	fillKnownDaysLong(longDayNamesNnNO, LocaleNnNO)
	fillKnownDaysShort(shortDayNamesNnNO, LocaleNnNO)
	fillKnownMonthsLong(longMonthNamesNnNO, LocaleNnNO)
	fillKnownMonthsShort(shortMonthNamesNnNO, LocaleNnNO)

	// Nb_NO: Norwegian Bokmål (Norway)
	fillKnownDaysLong(longDayNamesNbNO, LocaleNbNO)
	fillKnownDaysShort(shortDayNamesNbNO, LocaleNbNO)
	fillKnownMonthsLong(longMonthNamesNbNO, LocaleNbNO)
	fillKnownMonthsShort(shortMonthNamesNbNO, LocaleNbNO)

	// Pt_PT: Portuguese (Portugal)
	fillKnownDaysLong(longDayNamesPtPT, LocalePtPT)
	fillKnownDaysShort(shortDayNamesPtPT, LocalePtPT)
	fillKnownMonthsLong(longMonthNamesPtPT, LocalePtPT)
	fillKnownMonthsShort(shortMonthNamesPtPT, LocalePtPT)

	// Pt_BR: Portuguese (Brazil)
	fillKnownDaysLong(longDayNamesPtBR, LocalePtBR)
	fillKnownDaysShort(shortDayNamesPtBR, LocalePtBR)
	fillKnownMonthsLong(longMonthNamesPtBR, LocalePtBR)
	fillKnownMonthsShort(shortMonthNamesPtBR, LocalePtBR)

	// Ro_RO: Portuguese (Brazil)
	fillKnownDaysLong(longDayNamesRoRO, LocaleRoRO)
	fillKnownDaysShort(shortDayNamesRoRO, LocaleRoRO)
	fillKnownMonthsLong(longMonthNamesRoRO, LocaleRoRO)
	fillKnownMonthsShort(shortMonthNamesRoRO, LocaleRoRO)

	// Ru_RU: Russian (Russia)
	fillKnownDaysLong(longDayNamesRuRU, LocaleRuRU)
	fillKnownDaysShort(shortDayNamesRuRU, LocaleRuRU)
	fillKnownMonthsLong(longMonthNamesRuRU, LocaleRuRU)
	fillKnownMonthsShort(shortMonthNamesRuRU, LocaleRuRU)
	fillKnownMonthsGenitiveLong(longMonthNamesGenitiveRuRU, LocaleRuRU)
	fillKnownMonthsGenitiveShort(shortMonthNamesGenitiveRuRU, LocaleRuRU)

	// Es_ES: Spanish (Spain)
	fillKnownDaysLong(longDayNamesEsES, LocaleEsES)
	fillKnownDaysShort(shortDayNamesEsES, LocaleEsES)
	fillKnownMonthsLong(longMonthNamesEsES, LocaleEsES)
	fillKnownMonthsShort(shortMonthNamesEsES, LocaleEsES)

	// Ca_ES: Catalan (Spain)
	fillKnownDaysLong(longDayNamesCaES, LocaleCaES)
	fillKnownDaysShort(shortDayNamesCaES, LocaleCaES)
	fillKnownMonthsLong(longMonthNamesCaES, LocaleCaES)
	fillKnownMonthsShort(shortMonthNamesCaES, LocaleCaES)

	// Sv_SE: Swedish (Sweden)
	fillKnownDaysLong(longDayNamesSvSE, LocaleSvSE)
	fillKnownDaysShort(shortDayNamesSvSE, LocaleSvSE)
	fillKnownMonthsLong(longMonthNamesSvSE, LocaleSvSE)
	fillKnownMonthsShort(shortMonthNamesSvSE, LocaleSvSE)

	// Tr_TR: Turkish (Turkey)
	fillKnownDaysLong(longDayNamesTrTR, LocaleTrTR)
	fillKnownDaysShort(shortDayNamesTrTR, LocaleTrTR)
	fillKnownMonthsLong(longMonthNamesTrTR, LocaleTrTR)
	fillKnownMonthsShort(shortMonthNamesTrTR, LocaleTrTR)

	// Bg_BG: Bulgarian (Bulgaria)
	fillKnownDaysLong(longDayNamesBgBG, LocaleBgBG)
	fillKnownDaysShort(shortDayNamesBgBG, LocaleBgBG)
	fillKnownMonthsLong(longMonthNamesBgBG, LocaleBgBG)
	fillKnownMonthsShort(shortMonthNamesBgBG, LocaleBgBG)

	// Zh_CN: Chinese (Mainland)
	fillKnownDaysLong(longDayNamesZhCN, LocaleZhCN)
	fillKnownDaysShort(shortDayNamesZhCN, LocaleZhCN)
	fillKnownMonthsLong(longMonthNamesZhCN, LocaleZhCN)
	fillKnownMonthsShort(shortMonthNamesZhCN, LocaleZhCN)

	// Zh_TW: Chinese (Taiwan)
	fillKnownDaysLong(longDayNamesZhTW, LocaleZhTW)
	fillKnownDaysShort(shortDayNamesZhTW, LocaleZhTW)
	fillKnownMonthsLong(longMonthNamesZhTW, LocaleZhTW)
	fillKnownMonthsShort(shortMonthNamesZhTW, LocaleZhTW)

	// Zh_HK: Chinese (Hong Kong)
	fillKnownDaysLong(longDayNamesZhHK, LocaleZhHK)
	fillKnownDaysShort(shortDayNamesZhHK, LocaleZhHK)
	fillKnownMonthsLong(longMonthNamesZhHK, LocaleZhHK)
	fillKnownMonthsShort(shortMonthNamesZhHK, LocaleZhHK)

	fillKnownDaysLong(longDayNamesJaJP, LocaleJaJP)
	fillKnownDaysShort(shortDayNamesJaJP, LocaleJaJP)
	fillKnownMonthsLong(longMonthNamesJaJP, LocaleJaJP)
	fillKnownMonthsShort(shortMonthNamesJaJP, LocaleJaJP)
	fillKnownPeriods(periodsJaJP, LocaleJaJP)

}

func fill(src map[string]string, dest map[Locale]map[string]string, locale Locale) {
	loc, ok := dest[locale]

	if !ok {
		loc = make(map[string]string)
		dest[locale] = loc
	}

	for k, v := range src {
		loc[k] = v
	}
}

func fillReverse(src map[string]string, dest map[Locale]map[string]string, locale Locale) {
	loc, ok := dest[locale]

	if !ok {
		loc = make(map[string]string)
		dest[locale] = loc
	}

	for k, v := range src {
		loc[v] = k
	}
}

func fillKnownMonthsGenitiveShort(src map[string]string, locale Locale) {
	fillReverse(src, knownMonthsGenitiveShortReverse, locale)
	fill(src, knownMonthsGenitiveShort, locale)
}

func fillKnownMonthsGenitiveLong(src map[string]string, locale Locale) {
	fillReverse(src, knownMonthsGenitiveLongReverse, locale)
	fill(src, knownMonthsGenitiveLong, locale)
}

func fillKnownDaysShort(src map[string]string, locale Locale) {
	fillReverse(src, knownDaysShortReverse, locale)
	fill(src, knownDaysShort, locale)
}

func fillKnownDaysLong(src map[string]string, locale Locale) {
	fillReverse(src, knownDaysLongReverse, locale)
	fill(src, knownDaysLong, locale)
}

func fillKnownMonthsShort(src map[string]string, locale Locale) {
	fillReverse(src, knownMonthsShortReverse, locale)
	fill(src, knownMonthsShort, locale)
}

func fillKnownMonthsLong(src map[string]string, locale Locale) {
	fillReverse(src, knownMonthsLongReverse, locale)
	fill(src, knownMonthsLong, locale)
}

func fillKnownPeriods(src map[string]string, locale Locale) {
	fillReverse(src, knownPeriodsReverse, locale)
	fill(src, knownPeriods, locale)
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
// So, even though March is "Март" in Russian, correctly formatted today's date would be: "7 марта 2007".
// Thus, some transformations for some languages may be a bit more complex than just plain replacements.
func Format(dt time.Time, layout string, locale Locale) string {
	fm := dt.Format(layout)
	intFunc, ok := internalFormatFuncs[locale]
	if !ok {
		return fm
	}
	return intFunc(fm, layout)
}

// ParseInLocation is the standard time.ParseInLocation wrapper, which replaces
// known month/day translations for a specified locale back to English before
// calling time.ParseInLocation. So, you can parse localized dates with this wrapper.
func ParseInLocation(layout, value string, loc *time.Location, locale Locale) (time.Time, error) {
	intFunc, ok := internalParseFuncs[locale]
	if ok {
		value = intFunc(layout, value)
	} else {
		return time.Now(), fmt.Errorf("monday: coudln't find parse func for locale %v", locale)
	}

	return time.ParseInLocation(layout, value, loc)
}

func GetShortDays(locale Locale) (arr []string) {
	days, ok := knownDaysShort[locale]
	if !ok {
		return
	}
	for _, day := range days {
		arr = append(arr, day)
	}
	return
}

func GetLongDays(locale Locale) (arr []string) {
	days, ok := knownDaysLong[locale]
	if !ok {
		return
	}
	for _, day := range days {
		arr = append(arr, day)
	}
	return
}
