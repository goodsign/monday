package monday

import (
    "fmt"
    "time"
    "testing"
)

type FormatTest struct {
    locale Locale
    date time.Time
    layout string
    expected string
}

var formatTests = []FormatTest {
    { LocaleEnUS, time.Date(2013, 9, 3, 0, 0, 0, 0, time.UTC), "Mon Jan 2 2006", "Tue Sep 3 2013" },
    { LocaleEnUS, time.Date(2013, 9, 4, 0, 0, 0, 0, time.UTC), "Monday Jan 2 2006", "Wednesday Sep 4 2013" },
    { LocaleEnUS, time.Date(2013, 10, 3, 0, 0, 0, 0, time.UTC), "Monday January 02 2006", "Thursday October 03 2013" },
    { LocaleEnUS, time.Date(2013, 11, 3, 0, 0, 0, 0, time.UTC), "Monday. 2 January 2006", "Sunday. 3 November 2013" },
    { LocaleEnUS, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2006. 2 January. Monday", "2013. 13 May. Monday" },
    { LocaleEnUS, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2 Jan 2006", "13 May 2013" },
    { LocaleEnUS, time.Date(0, 5, 1, 0, 0, 0, 0, time.UTC), "January", "May" },
    { LocaleEnUS, time.Date(0, 5, 13, 0, 0, 0, 0, time.UTC), "2 January", "13 May" },

    { LocaleRuRU, time.Date(2013, 9, 3, 0, 0, 0, 0, time.UTC), "Mon Jan 2 2006", "Вт Сен 3 2013" },
    { LocaleRuRU, time.Date(2013, 9, 4, 0, 0, 0, 0, time.UTC), "Monday Jan 2 2006", "Среда Сен 4 2013" },
    { LocaleRuRU, time.Date(2013, 10, 3, 0, 0, 0, 0, time.UTC), "Monday January 02 2006", "Четверг Октябрь 03 2013" },
    { LocaleRuRU, time.Date(2013, 11, 3, 0, 0, 0, 0, time.UTC), "Monday. 2 January 2006", "Воскресенье. 3 ноября 2013" },
    { LocaleRuRU, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2006. 2 January. Monday", "2013. 13 мая. Понедельник" },
    { LocaleRuRU, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2 Jan 2006", "13 мая 2013" },
    { LocaleRuRU, time.Date(0, 5, 1, 0, 0, 0, 0, time.UTC), "January", "Май" },
    { LocaleRuRU, time.Date(0, 5, 13, 0, 0, 0, 0, time.UTC), "2 January", "13 мая" },
}

func TestFormat(t *testing.T) {

    for i, ts := range formatTests {
        txt := Format(ts.date, ts.layout, ts.locale)

        if txt != ts.expected {
            t.Errorf("Test #%d (%s: %s) => Format failed.\n         Got: %s\n    Expected: %s\n", i, ts.locale, ts.layout, txt, ts.expected)
            continue
        }

        reverseDate, err := ParseInLocation(ts.layout, txt, time.UTC, ts.locale)

        if err != nil {
            t.Errorf("Test #%d (%s: %s) => Reverse parse from '%s' error: %s", i, ts.locale, ts.layout, txt, err)
            continue
        }

        if reverseDate != ts.date {
            t.Errorf("Test #%d (%s: %s) => Reverse parse from '%s' failed.\n         Got: %s\n    Expected: %s\n",
                     i, ts.locale, ts.layout, txt, reverseDate.Format(time.RFC850), ts.date.Format(time.RFC850))
            continue
        }
    }
}

func ExampleFormatUsage() {
    t := time.Date(2013, 4, 12, 0, 0, 0, 0, time.UTC)
    layout := "2 January 2006 15:04:05 MST"

    fmt.Println(Format(t, layout, LocaleEnUS))
    fmt.Println(Format(t, layout, LocaleRuRU))
}

func ExampleParseUsage() {
    layout := "2 January 2006 15:04:05 MST"

    fmt.Println(ParseInLocation(layout, "12 April 2013 00:00:00 MST", time.UTC, LocaleEnUS))
    fmt.Println(ParseInLocation(layout, "12 апреля 2013 00:00:00 MST", time.UTC, LocaleRuRU))
}