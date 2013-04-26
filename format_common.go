package monday

import "strings"

func findInString(where string, what string, foundIndex *int, trimRight *int) (found bool) {
	ind := strings.Index(where, what)
	if ind != -1 {
		*foundIndex = ind
		*trimRight = len(where) - ind - len(what)
		return true
	}

	return false
}

// commonFormatFunc is used for languages which don't have changed forms of month names dependent
// on their position (after day or standalone)
func commonFormatFunc(value, format string,
	knownDaysShort, knownDaysLong, knownMonthsShort, knownMonthsLong map[string]string) (res string) {
	l := stringToLayoutItems(value)
	f := stringToLayoutItems(format)

	for i, v := range l {
		var knw map[string]string

		// number of symbols before replaced term
		foundIndex := 0
		trimRight := 0

		switch {
		case findInString(f[i].item, "Monday", &foundIndex, &trimRight):
			knw = knownDaysLong
		case findInString(f[i].item, "Mon", &foundIndex, &trimRight):
			knw = knownDaysShort
		case findInString(f[i].item, "January", &foundIndex, &trimRight):
			knw = knownMonthsLong
		case findInString(f[i].item, "Jan", &foundIndex, &trimRight):
			knw = knownMonthsShort
		}

		if knw != nil {
			trimmedItem := v.item[foundIndex : len(v.item)-trimRight]
			tr, ok := knw[trimmedItem]

			if ok {
				res = res + v.item[:foundIndex] + tr + v.item[len(v.item)-trimRight:]
			} else {
				res = res + v.item
			}
		} else {
			res = res + v.item
		}
	}
	return res
}

func hasDigitBefore(l []dateStringLayoutItem, position int) bool {
	if position >= 2 {
		return l[position-2].isDigit && len(l[position-2].item) <= 2
	}
	return false
}

// commonGenitiveFormatFunc is used for languages with genitive forms of names, like Russian.
func commonGenitiveFormatFunc(value, format string,
	knownDaysShort, knownDaysLong, knownMonthsShort, knownMonthsLong,
	knownMonthsGenShort, knownMonthsGenLong map[string]string) (res string) {
	l := stringToLayoutItems(value)
	f := stringToLayoutItems(format)

	for i, v := range l {

		var knw map[string]string
		switch f[i].item {
		case "Mon":
			knw = knownDaysShort
		case "Monday":
			knw = knownDaysLong
		case "Jan":
			if hasDigitBefore(l, i) {
				knw = knownMonthsGenShort
			} else {
				knw = knownMonthsShort
			}
		case "January":
			if hasDigitBefore(l, i) {
				knw = knownMonthsGenLong
			} else {
				knw = knownMonthsLong
			}
		}

		if knw != nil {
			tr, _ := knw[v.item]
			res = res + tr
		} else {
			res = res + v.item
		}
	}
	return res
}

func createCommonFormatFunc(locale Locale) internalFormatFunc {
	return func(value, layout string) (res string) {
		return commonFormatFunc(value, layout,
			knownDaysShort[locale], knownDaysLong[locale], knownMonthsShort[locale], knownMonthsLong[locale])
	}
}

func createCommonFormatFuncWithGenitive(locale Locale) internalFormatFunc {
	return func(value, layout string) (res string) {
		return commonGenitiveFormatFunc(value, layout,
			knownDaysShort[locale], knownDaysLong[locale], knownMonthsShort[locale], knownMonthsLong[locale],
			knownMonthsGenitiveShort[locale], knownMonthsGenitiveLong[locale])
	}
}

func createCommonParseFunc(locale Locale) internalParseFunc {
	return func(layout, value string) string {
		return commonFormatFunc(value, layout,
			knownDaysShortReverse[locale], knownDaysLongReverse[locale],
			knownMonthsShortReverse[locale], knownMonthsLongReverse[locale])
	}
}

func createCommonParsetFuncWithGenitive(locale Locale) internalParseFunc {
	return func(layout, value string) (res string) {
		return commonGenitiveFormatFunc(value, layout,
			knownDaysShortReverse[locale], knownDaysLongReverse[locale],
			knownMonthsShortReverse[locale], knownMonthsLongReverse[locale],
			knownMonthsGenitiveShortReverse[locale], knownMonthsGenitiveLongReverse[locale])
	}
}
