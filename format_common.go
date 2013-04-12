package monday

func commonFormatFunc(value, layout string, locale Locale, knw map[string][]string) (res string) {
    l := stringToLayoutItems(value, locale)
    
    for _, v := range l {
        changed := false

        if v.isWord {
            tr, ok := knw[v.item]

            if ok && len(tr) > 0 {
                res = res + tr[0]
                changed = true
            }
        }

        if !changed {
            res = res + v.item                
        }
    }
    return res
}

func createCommonFormatFunc(locale Locale) internalFormatFunc {
    return func (value, layout string) (res string) {
        return commonFormatFunc(value, layout, locale, knownWords[locale])
    }
}

func createCommonParseFunc(locale Locale) internalParseFunc {
    return func (layout, value string) string {
        return commonFormatFunc(value, layout, locale, knownWordsReverse[locale])
    }
}