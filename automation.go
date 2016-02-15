package monday

import (
	"errors"
	"fmt"
	"gopkg.in/fatih/set.v0"
	"regexp"
	"strconv"
	"strings"
	"text/scanner"
	"time"
)

var (
	wordsRx             = regexp.MustCompile("(\\p{L}+)")
	debugLayoutDef bool = false
)

type InvalidTypeError struct {
	error
}

type InvalidLengthError struct {
	error
}

func NewInvalidTypeError() InvalidTypeError {
	return InvalidTypeError{error: errors.New("invalid type for token")}
}

func NewInvalidLengthError() InvalidLengthError {
	return InvalidLengthError{error: errors.New("invalid length for token")}
}

type layoutSpanI interface {
	scanInt(s *scanner.Scanner) (int, error)
	scanString(s *scanner.Scanner) (string, error)
	isString() bool
	isDelimiter() bool
}

type lengthLimitSpan struct {
	minLength int
	maxLength int
}

func (this lengthLimitSpan) scanInt(s *scanner.Scanner) (int, error) {
	return -1, NewInvalidTypeError()
}

func (this lengthLimitSpan) scanString(s *scanner.Scanner) (string, error) {
	return "", NewInvalidTypeError()
}

func (this lengthLimitSpan) isString() bool    { return false }
func (this lengthLimitSpan) isDelimiter() bool { return false }

func initLengthLimitSpan(min, max int) lengthLimitSpan {
	return lengthLimitSpan{
		minLength: min,
		maxLength: max,
	}
}

type limitedStringSpan struct {
	lengthLimitSpan
}

func initLimitedStringSpan(minLength, maxLength int) limitedStringSpan {
	return limitedStringSpan{lengthLimitSpan: initLengthLimitSpan(minLength, maxLength)}
}

func (this limitedStringSpan) scanString(s *scanner.Scanner) (string, error) {
	tok := s.Scan()
	if tok != scanner.EOF && tok == -2 {
		return s.TokenText(), nil
	}
	return "", NewInvalidTypeError()
}

func (this limitedStringSpan) isString() bool { return true }
func (this limitedStringSpan) String() string {
	return fmt.Sprintf("[limitedStringSpan:%v]", this.lengthLimitSpan)
}

type rangeIntSpan struct {
	lengthLimitSpan
	min int
	max int
}

func initRangeIntSpan(minValue, maxValue, minLength, maxLength int) rangeIntSpan {
	return rangeIntSpan{
		lengthLimitSpan: initLengthLimitSpan(minLength, maxLength),
		min:             minValue,
		max:             maxValue,
	}
}

func (this rangeIntSpan) scanInt(s *scanner.Scanner) (int, error) {
	var tok = s.Scan()
	var negative bool = false
	if tok == 45 {
		negative = true
		if debugLayoutDef {
			fmt.Printf("scan negative:'%s'\n", s.TokenText())
		}
		tok = s.Scan()
	}
	if tok == -3 {
		str := s.TokenText()
		i, err := strconv.Atoi(str)
		if err != nil {
			return 0, err
		}
		if negative {
			i = i * -1
		}
		return i, nil
	} else {
		if debugLayoutDef {
			fmt.Printf("invalid tok: %s '%s'\n", tok, s.TokenText())
		}
	}
	return 0, NewInvalidTypeError()
}

func (this rangeIntSpan) String() string {
	return fmt.Sprintf("[rangeIntSpan:%v]", this.lengthLimitSpan)
}

type delimiterSpan struct {
	lengthLimitSpan
	character string
}

func initDelimiterSpan(character string, minLength, maxLength int) delimiterSpan {
	return delimiterSpan{
		lengthLimitSpan: initLengthLimitSpan(minLength, maxLength),
		character:       character,
	}
}

func (this delimiterSpan) scanString(s *scanner.Scanner) (string, error) {
	tok := s.Scan()
	if tok != scanner.EOF && tok != -2 && tok != 45 && tok != -3 {
		return s.TokenText(), nil
	} else {
		if debugLayoutDef {
			fmt.Printf("expected tok:=!(-2,-3,45), received:%d ('%s')\n", tok, s.TokenText())
		}
	}
	return "", NewInvalidTypeError()
}

func (this delimiterSpan) isString() bool    { return false }
func (this delimiterSpan) isDelimiter() bool { return true }
func (this delimiterSpan) String() string {
	return fmt.Sprintf("[delimiterSpan '%s':%v]", this.character, this.lengthLimitSpan)
}

type layoutDef struct {
	spans         []layoutSpanI
	errorPosition int
}

func (this *layoutDef) validate(value string) bool {
	s := &scanner.Scanner{}
	s.Init(strings.NewReader(value))
	s.Whitespace = 0
	for _, span := range this.spans {
		if span.isString() || span.isDelimiter() {
			if _, err := span.scanString(s); err != nil {
				this.errorPosition = s.Pos().Offset
				if debugLayoutDef {
					fmt.Printf("error at pos: %d: %s (span=%+v) - expected string or delimiter\n", s.Pos().Offset, err.Error(), span)
				}
				return false
			}
		} else if _, err := span.scanInt(s); err != nil {
			if debugLayoutDef {
				fmt.Printf("error at pos: %d: %s (span=%+v) - expected integer\n", s.Pos().Offset, err.Error(), span)
			}
			this.errorPosition = s.Pos().Offset
			return false
		}
	}
	this.errorPosition = s.Pos().Offset
	return s.Pos().Offset == len(value)
}

type LocaleDetector struct {
	localeMap         map[string]*set.Set
	lastLocale        Locale
	layoutsMap        map[string]layoutDef
	lastErrorPosition int
}

func (this *LocaleDetector) prepareLayout(layout string) layoutDef {
	s := scanner.Scanner{}
	s.Init(strings.NewReader(layout))
	s.Whitespace = 0
	result := make([]layoutSpanI, 0)
	var tok rune
	// var pos int = 0
	var span layoutSpanI
	var sign bool = false
	//	var neg bool = false
	for tok != scanner.EOF {
		tok = s.Scan()
		switch tok {
		case -2: // text
			span = initLimitedStringSpan(1, -1)
		case -3: // digit
			span = initRangeIntSpan(-1, -1, 1, -1)
			if sign {
				sign = false
			}
		case 45: // negative sign
			sign = true
			// neg = s.TokenText() == "-"
			continue
		case scanner.EOF:
			continue
		default: // fixed character
			span = initDelimiterSpan(s.TokenText(), 1, 1)
		}
		result = append(result, span)
		// length := s.Pos().Offset - pos
		// pos = s.Pos().Offset
		// fmt.Printf("tok'%s' [%d %d] length=%d\n", s.TokenText(), pos, s.Pos().Offset, length)

	}
	if debugLayoutDef {
		fmt.Printf("layout:'%s'\n", layout)
		fmt.Printf("layout:%v\n", result)
	}
	ld := layoutDef{spans: result}
	this.layoutsMap[layout] = ld
	return ld
}

func (this *LocaleDetector) validateValue(layout string, value string) bool {
	l, ok := this.layoutsMap[layout]
	if !ok {
		l = this.prepareLayout(layout)
	}
	result := l.validate(value)
	this.lastErrorPosition = l.errorPosition
	return result
}

func (this *LocaleDetector) errorPosition() int { return this.lastErrorPosition }

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
	this := &LocaleDetector{localeMap: make(map[string]*set.Set), lastLocale: LocaleEnGB, layoutsMap: make(map[string]layoutDef)}
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
