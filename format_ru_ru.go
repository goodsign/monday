package monday

// ============================================================
// Format rules for "ru_RU" locale: Russian (Russia)
// ============================================================

var longDayNamesRuRU = map[string]string {
    "Sunday": "Воскресенье",
    "Monday": "Понедельник",
    "Tuesday": "Вторник",
    "Wednesday": "Среда",
    "Thursday": "Четверг",
    "Friday": "Пятница",
    "Saturday": "Суббота",
}

var shortDayNamesRuRU = map[string]string {
    "Sun": "Вс",
    "Mon": "Пн",
    "Tue": "Вт",
    "Wed": "Ср",
    "Thu": "Чт",
    "Fri": "Пт",
    "Sat": "Сб",
}

var longMonthNamesRuRU = map[string]string {
    "January": "Январь",
    "February": "Февраль",
    "March": "Март",
    "April": "Апрель",
    "May": "Май",
    "June": "Июнь",
    "July": "Июль",
    "August": "Август",
    "September": "Сентябрь",
    "October": "Октябрь",
    "November": "Ноябрь",
    "December": "Декабрь",
}

var longMonthNamesGenitiveRuRU = map[string]string {
    "January": "января",
    "February": "февраля",
    "March": "марта",
    "April": "апреля",
    "May": "мая",
    "June": "июня",
    "July": "июля",
    "August": "августа",
    "September": "сентября",
    "October": "октября",
    "November": "ноября",
    "December": "декабря",
}

var shortMonthNamesRuRU = map[string]string {
    "Jan": "Янв",
    "Feb": "Фев",
    "Mar": "Мар",
    "Apr": "Апр",
    "May": "Май",
    "Jun": "Июн",
    "Jul": "Июл",
    "Aug": "Авг",
    "Sep": "Сен",
    "Oct": "Окт",
    "Nov": "Ноя",
    "Dec": "Дек",
}

var shortMonthNamesGenitiveRuRU = map[string]string {
    "Jan": "янв",
    "Feb": "фев",
    "Mar": "мар",
    "Apr": "апр",
    "May": "мая",
    "Jun": "июн",
    "Jul": "июл",
    "Aug": "авг",
    "Sep": "сен",
    "Oct": "окт",
    "Nov": "ноя",
    "Dec": "дек",
}

func hasDigitBefore(l []dateStringLayoutItem, position int) bool {
    if position >= 2 {
        return l[position - 2].isDigit && len(l[position - 2].item) <= 2
    }
    return false
}

func formatRuRU(value, format string) (res string) {
    l := stringToLayoutItems(value, LocaleRuRU)
    knw := knownWords[LocaleRuRU]

    for i, v := range l {
        changed := false

        if v.isWord {
            tr, ok := knw[v.item]

            if ok {
                _, isShortName := shortMonthNamesRuRU[v.item]
                _, isLongName := longMonthNamesRuRU[v.item]

                if isShortName {
                    if hasDigitBefore(l, i) {
                        res = res + shortMonthNamesGenitiveRuRU[v.item]
                    } else {
                        res = res + shortMonthNamesRuRU[v.item]
                    }
                } else if isLongName {
                    if hasDigitBefore(l, i) {
                        res = res + longMonthNamesGenitiveRuRU[v.item]
                    } else {
                        res = res + longMonthNamesRuRU[v.item]
                    }
                } else {
                    res = res + tr[0]    
                }

                
                changed = true
            }
        }

        if !changed {
            res = res + v.item                
        }
    }
    return res
}