package monday

import "strings"

// ============================================================
// Format rules for "th_TH" locale: Thai (Thailand)
// ============================================================

var longDayNamesThTH = map[string]string{
	"Sunday":    "วันอาทิตย์",
	"Monday":    "วันจันทร์",
	"Tuesday":   "วันอังคาร",
	"Wednesday": "วันพุธ",
	"Thursday":  "วันพฤหัสบดี",
	"Friday":    "วันศุกร์",
	"Saturday":  "วันเสาร์",
}

var shortDayNamesThTH = map[string]string{
	"Sun": "อา.",
	"Mon": "จ.",
	"Tue": "อ.",
	"Wed": "พ.",
	"Thu": "พฤ.",
	"Fri": "ศ.",
	"Sat": "ส.",
}

var longMonthNamesThTH = map[string]string{
	"January":   "มกราคม",
	"February":  "กุมภาพันธ์",
	"March":     "มีนาคม",
	"April":     "เมษายน",
	"May":       "พฤษภาคม",
	"June":      "มิถุนายน",
	"July":      "กรกฎาคม",
	"August":    "สิงหาคม",
	"September": "กันยายน",
	"October":   "ตุลาคม",
	"November":  "พฤศจิกายน",
	"December":  "ธันวาคม",
}

var shortMonthNamesThTH = map[string]string{
	"Jan": "ม.ค.",
	"Feb": "ก.พ.",
	"Mar": "มี.ค.",
	"Apr": "เม.ย.",
	"May": "พ.ค.",
	"Jun": "มิ.ย.",
	"Jul": "ก.ค.",
	"Aug": "ส.ค.",
	"Sep": "ก.ย.",
	"Oct": "ต.ค.",
	"Nov": "พ.ย.",
	"Dec": "ธ.ค.",
}

func parseFuncThCommon(locale Locale) internalParseFunc {
	return func(layout, value string) string {
		// This special case is needed because th_TH... contains month and day names
		// that consist of dots, and special character. Example: "February" = "กุมภาพันธ์", "Feb" = "ก.พ."
		//
		// This means that probably default time package layout IDs like 'January' or 'Jan'
		// shouldn't be used in th_TH. But this is a time-compatible package, so someone
		// might actually use those and we need to replace those before doing standard procedures.
		for k, v := range knownMonthsShortReverse[locale] {
			value = strings.Replace(value, k, v, -1)
		}
		for k, v := range knownDaysShortReverse[locale] {
			value = strings.Replace(value, k, v, -1)
		}
		for k, v := range knownMonthsLongReverse[locale] {
			value = strings.Replace(value, k, v, -1)
		}
		for k, v := range knownDaysLongReverse[locale] {
			value = strings.Replace(value, k, v, -1)
		}

		return commonFormatFunc(value, layout,
			knownDaysShortReverse[locale], knownDaysLongReverse[locale],
			knownMonthsShortReverse[locale], knownMonthsLongReverse[locale], knownPeriods[locale])
	}
}
