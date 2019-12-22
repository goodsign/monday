package monday

import (
	"fmt"
	"testing"
	"time"
)

type FormatTest struct {
	locale   Locale
	date     time.Time
	layout   string
	expected string
}

var formatTests = []FormatTest{
	{LocaleEnUS, time.Date(2013, 9, 3, 0, 0, 0, 0, time.UTC), "Mon Jan 2 2006", "Tue Sep 3 2013"},
	{LocaleEnUS, time.Date(2013, 9, 4, 0, 0, 0, 0, time.UTC), "Monday Jan 2 2006", "Wednesday Sep 4 2013"},
	{LocaleEnUS, time.Date(2013, 10, 3, 0, 0, 0, 0, time.UTC), "Monday January 02 2006", "Thursday October 03 2013"},
	{LocaleEnUS, time.Date(2013, 11, 3, 0, 0, 0, 0, time.UTC), "Monday. 2 January 2006", "Sunday. 3 November 2013"},
	{LocaleEnUS, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2006. 2 January. Monday", "2013. 13 May. Monday"},
	{LocaleEnUS, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2 Jan 2006", "13 May 2013"},
	{LocaleEnUS, time.Date(0, 5, 1, 0, 0, 0, 0, time.UTC), "January", "May"},
	{LocaleEnUS, time.Date(0, 5, 13, 0, 0, 0, 0, time.UTC), "2 January", "13 May"},

	{LocaleEnGB, time.Date(2013, 9, 3, 0, 0, 0, 0, time.UTC), "Mon Jan 2 2006", "Tue Sep 3 2013"},
	{LocaleEnGB, time.Date(2013, 9, 4, 0, 0, 0, 0, time.UTC), "Monday Jan 2 2006", "Wednesday Sep 4 2013"},
	{LocaleEnGB, time.Date(2013, 10, 3, 0, 0, 0, 0, time.UTC), "Monday January 02 2006", "Thursday October 03 2013"},
	{LocaleEnGB, time.Date(2013, 11, 3, 0, 0, 0, 0, time.UTC), "Monday. 2 January 2006", "Sunday. 3 November 2013"},
	{LocaleEnGB, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2006. 2 January. Monday", "2013. 13 May. Monday"},
	{LocaleEnGB, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2 Jan 2006", "13 May 2013"},
	{LocaleEnGB, time.Date(0, 5, 1, 0, 0, 0, 0, time.UTC), "January", "May"},
	{LocaleEnGB, time.Date(0, 5, 13, 0, 0, 0, 0, time.UTC), "2 January", "13 May"},

	{LocaleDaDK, time.Date(2013, 9, 3, 0, 0, 0, 0, time.UTC), "Mon Jan 2 2006", "tir sep 3 2013"},
	{LocaleDaDK, time.Date(2013, 9, 4, 0, 0, 0, 0, time.UTC), "Monday Jan 2 2006", "onsdag sep 4 2013"},
	{LocaleDaDK, time.Date(2013, 10, 3, 0, 0, 0, 0, time.UTC), "Monday January 02 2006", "torsdag oktober 03 2013"},
	{LocaleDaDK, time.Date(2013, 11, 3, 0, 0, 0, 0, time.UTC), "Monday. 2 January 2006", "søndag. 3 november 2013"},
	{LocaleDaDK, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2006. 2 January. Monday", "2013. 13 maj. mandag"},
	{LocaleDaDK, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2 Jan 2006", "13 maj 2013"},
	{LocaleDaDK, time.Date(0, 5, 1, 0, 0, 0, 0, time.UTC), "January", "maj"},
	{LocaleDaDK, time.Date(0, 5, 13, 0, 0, 0, 0, time.UTC), "2 January", "13 maj"},

	{LocaleNlBE, time.Date(2013, 9, 3, 0, 0, 0, 0, time.UTC), "Mon Jan 2 2006", "di sep 3 2013"},
	{LocaleNlBE, time.Date(2013, 9, 4, 0, 0, 0, 0, time.UTC), "Monday Jan 2 2006", "woensdag sep 4 2013"},
	{LocaleNlBE, time.Date(2013, 10, 3, 0, 0, 0, 0, time.UTC), "Monday January 02 2006", "donderdag oktober 03 2013"},
	{LocaleNlBE, time.Date(2013, 11, 3, 0, 0, 0, 0, time.UTC), "Monday. 2 January 2006", "zondag. 3 november 2013"},
	{LocaleNlBE, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2006. 2 January. Monday", "2013. 13 mei. maandag"},
	{LocaleNlBE, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2 Jan 2006", "13 mei 2013"},
	{LocaleNlBE, time.Date(0, 5, 1, 0, 0, 0, 0, time.UTC), "January", "mei"},
	{LocaleNlBE, time.Date(0, 5, 13, 0, 0, 0, 0, time.UTC), "2 January", "13 mei"},

	{LocaleNlNL, time.Date(2013, 9, 3, 0, 0, 0, 0, time.UTC), "Mon Jan 2 2006", "di sep 3 2013"},
	{LocaleNlNL, time.Date(2013, 9, 4, 0, 0, 0, 0, time.UTC), "Monday Jan 2 2006", "woensdag sep 4 2013"},
	{LocaleNlNL, time.Date(2013, 10, 3, 0, 0, 0, 0, time.UTC), "Monday January 02 2006", "donderdag oktober 03 2013"},
	{LocaleNlNL, time.Date(2013, 11, 3, 0, 0, 0, 0, time.UTC), "Monday. 2 January 2006", "zondag. 3 november 2013"},
	{LocaleNlNL, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2006. 2 January. Monday", "2013. 13 mei. maandag"},
	{LocaleNlNL, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2 Jan 2006", "13 mei 2013"},
	{LocaleNlNL, time.Date(0, 5, 1, 0, 0, 0, 0, time.UTC), "January", "mei"},
	{LocaleNlNL, time.Date(0, 5, 13, 0, 0, 0, 0, time.UTC), "2 January", "13 mei"},

	{LocaleFiFI, time.Date(2013, 9, 3, 0, 0, 0, 0, time.UTC), "Mon Jan 2 2006", "ti syys 3 2013"},
	{LocaleFiFI, time.Date(2013, 9, 4, 0, 0, 0, 0, time.UTC), "Monday Jan 2 2006", "keskiviikko syys 4 2013"},
	{LocaleFiFI, time.Date(2013, 10, 3, 0, 0, 0, 0, time.UTC), "Monday January 02 2006", "torstai lokakuu 03 2013"},
	{LocaleFiFI, time.Date(2013, 11, 3, 0, 0, 0, 0, time.UTC), "Monday. 2 January 2006", "sunnuntai. 3 marraskuuta 2013"},
	{LocaleFiFI, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2006. 2 January. Monday", "2013. 13 toukokuuta. maanantai"},
	{LocaleFiFI, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2 Jan 2006", "13 touko 2013"},
	{LocaleFiFI, time.Date(0, 5, 1, 0, 0, 0, 0, time.UTC), "January", "toukokuu"},
	{LocaleFiFI, time.Date(0, 5, 13, 0, 0, 0, 0, time.UTC), "2 January", "13 toukokuuta"},

	{LocaleFrFR, time.Date(2013, 9, 3, 0, 0, 0, 0, time.UTC), "Mon Jan 2 2006", "mar sept 3 2013"},
	{LocaleFrFR, time.Date(2013, 9, 4, 0, 0, 0, 0, time.UTC), "Monday Jan 2 2006", "mercredi sept 4 2013"},
	{LocaleFrFR, time.Date(2013, 10, 3, 0, 0, 0, 0, time.UTC), "Monday January 02 2006", "jeudi octobre 03 2013"},
	{LocaleFrFR, time.Date(2013, 11, 3, 0, 0, 0, 0, time.UTC), "Monday. 2 January 2006", "dimanche. 3 novembre 2013"},
	{LocaleFrFR, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2006. 2 January. Monday", "2013. 13 mai. lundi"},
	{LocaleFrFR, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2 Jan 2006", "13 mai 2013"},
	{LocaleFrFR, time.Date(0, 5, 1, 0, 0, 0, 0, time.UTC), "January", "mai"},
	{LocaleFrFR, time.Date(0, 5, 13, 0, 0, 0, 0, time.UTC), "2 January", "13 mai"},

	{LocaleFrCA, time.Date(2013, 9, 3, 0, 0, 0, 0, time.UTC), "Mon Jan 2 2006", "mar sept 3 2013"},
	{LocaleFrCA, time.Date(2013, 9, 4, 0, 0, 0, 0, time.UTC), "Monday Jan 2 2006", "mercredi sept 4 2013"},
	{LocaleFrCA, time.Date(2013, 10, 3, 0, 0, 0, 0, time.UTC), "Monday January 02 2006", "jeudi octobre 03 2013"},
	{LocaleFrCA, time.Date(2013, 11, 3, 0, 0, 0, 0, time.UTC), "Monday. 2 January 2006", "dimanche. 3 novembre 2013"},
	{LocaleFrCA, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2006. 2 January. Monday", "2013. 13 mai. lundi"},
	{LocaleFrCA, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2 Jan 2006", "13 mai 2013"},
	{LocaleFrCA, time.Date(0, 5, 1, 0, 0, 0, 0, time.UTC), "January", "mai"},
	{LocaleFrCA, time.Date(0, 5, 13, 0, 0, 0, 0, time.UTC), "2 January", "13 mai"},

	{LocaleFrGP, time.Date(2013, 9, 3, 0, 0, 0, 0, time.UTC), "Mon Jan 2 2006", "mar sept 3 2013"},
	{LocaleFrGP, time.Date(2013, 9, 4, 0, 0, 0, 0, time.UTC), "Monday Jan 2 2006", "mercredi sept 4 2013"},
	{LocaleFrGP, time.Date(2013, 10, 3, 0, 0, 0, 0, time.UTC), "Monday January 02 2006", "jeudi octobre 03 2013"},
	{LocaleFrGP, time.Date(2013, 11, 3, 0, 0, 0, 0, time.UTC), "Monday. 2 January 2006", "dimanche. 3 novembre 2013"},
	{LocaleFrGP, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2006. 2 January. Monday", "2013. 13 mai. lundi"},
	{LocaleFrGP, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2 Jan 2006", "13 mai 2013"},
	{LocaleFrGP, time.Date(0, 5, 1, 0, 0, 0, 0, time.UTC), "January", "mai"},
	{LocaleFrGP, time.Date(0, 5, 13, 0, 0, 0, 0, time.UTC), "2 January", "13 mai"},

	{LocaleFrLU, time.Date(2013, 9, 3, 0, 0, 0, 0, time.UTC), "Mon Jan 2 2006", "mar sept 3 2013"},
	{LocaleFrLU, time.Date(2013, 9, 4, 0, 0, 0, 0, time.UTC), "Monday Jan 2 2006", "mercredi sept 4 2013"},
	{LocaleFrLU, time.Date(2013, 10, 3, 0, 0, 0, 0, time.UTC), "Monday January 02 2006", "jeudi octobre 03 2013"},
	{LocaleFrLU, time.Date(2013, 11, 3, 0, 0, 0, 0, time.UTC), "Monday. 2 January 2006", "dimanche. 3 novembre 2013"},
	{LocaleFrLU, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2006. 2 January. Monday", "2013. 13 mai. lundi"},
	{LocaleFrLU, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2 Jan 2006", "13 mai 2013"},
	{LocaleFrLU, time.Date(0, 5, 1, 0, 0, 0, 0, time.UTC), "January", "mai"},
	{LocaleFrLU, time.Date(0, 5, 13, 0, 0, 0, 0, time.UTC), "2 January", "13 mai"},

	{LocaleFrMQ, time.Date(2013, 9, 3, 0, 0, 0, 0, time.UTC), "Mon Jan 2 2006", "mar sept 3 2013"},
	{LocaleFrMQ, time.Date(2013, 9, 4, 0, 0, 0, 0, time.UTC), "Monday Jan 2 2006", "mercredi sept 4 2013"},
	{LocaleFrMQ, time.Date(2013, 10, 3, 0, 0, 0, 0, time.UTC), "Monday January 02 2006", "jeudi octobre 03 2013"},
	{LocaleFrMQ, time.Date(2013, 11, 3, 0, 0, 0, 0, time.UTC), "Monday. 2 January 2006", "dimanche. 3 novembre 2013"},
	{LocaleFrMQ, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2006. 2 January. Monday", "2013. 13 mai. lundi"},
	{LocaleFrMQ, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2 Jan 2006", "13 mai 2013"},
	{LocaleFrMQ, time.Date(0, 5, 1, 0, 0, 0, 0, time.UTC), "January", "mai"},
	{LocaleFrMQ, time.Date(0, 5, 13, 0, 0, 0, 0, time.UTC), "2 January", "13 mai"},

	{LocaleFrGF, time.Date(2013, 9, 3, 0, 0, 0, 0, time.UTC), "Mon Jan 2 2006", "mar sept 3 2013"},
	{LocaleFrGF, time.Date(2013, 9, 4, 0, 0, 0, 0, time.UTC), "Monday Jan 2 2006", "mercredi sept 4 2013"},
	{LocaleFrGF, time.Date(2013, 10, 3, 0, 0, 0, 0, time.UTC), "Monday January 02 2006", "jeudi octobre 03 2013"},
	{LocaleFrGF, time.Date(2013, 11, 3, 0, 0, 0, 0, time.UTC), "Monday. 2 January 2006", "dimanche. 3 novembre 2013"},
	{LocaleFrGF, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2006. 2 January. Monday", "2013. 13 mai. lundi"},
	{LocaleFrGF, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2 Jan 2006", "13 mai 2013"},
	{LocaleFrGF, time.Date(0, 5, 1, 0, 0, 0, 0, time.UTC), "January", "mai"},
	{LocaleFrGF, time.Date(0, 5, 13, 0, 0, 0, 0, time.UTC), "2 January", "13 mai"},

	{LocaleFrRE, time.Date(2013, 9, 3, 0, 0, 0, 0, time.UTC), "Mon Jan 2 2006", "mar sept 3 2013"},
	{LocaleFrRE, time.Date(2013, 9, 4, 0, 0, 0, 0, time.UTC), "Monday Jan 2 2006", "mercredi sept 4 2013"},
	{LocaleFrRE, time.Date(2013, 10, 3, 0, 0, 0, 0, time.UTC), "Monday January 02 2006", "jeudi octobre 03 2013"},
	{LocaleFrRE, time.Date(2013, 11, 3, 0, 0, 0, 0, time.UTC), "Monday. 2 January 2006", "dimanche. 3 novembre 2013"},
	{LocaleFrRE, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2006. 2 January. Monday", "2013. 13 mai. lundi"},
	{LocaleFrRE, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2 Jan 2006", "13 mai 2013"},
	{LocaleFrRE, time.Date(0, 5, 1, 0, 0, 0, 0, time.UTC), "January", "mai"},
	{LocaleFrRE, time.Date(0, 5, 13, 0, 0, 0, 0, time.UTC), "2 January", "13 mai"},

	{LocaleDeDE, time.Date(2013, 9, 3, 0, 0, 0, 0, time.UTC), "Mon Jan 2 2006", "Di Sep 3 2013"},
	{LocaleDeDE, time.Date(2013, 9, 4, 0, 0, 0, 0, time.UTC), "Monday Jan 2 2006", "Mittwoch Sep 4 2013"},
	{LocaleDeDE, time.Date(2013, 10, 3, 0, 0, 0, 0, time.UTC), "Monday January 02 2006", "Donnerstag Oktober 03 2013"},
	{LocaleDeDE, time.Date(2013, 11, 3, 0, 0, 0, 0, time.UTC), "Monday. 2 January 2006", "Sonntag. 3 November 2013"},
	{LocaleDeDE, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2006. 2 January. Monday", "2013. 13 Mai. Montag"},
	{LocaleDeDE, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2 Jan 2006", "13 Mai 2013"},
	{LocaleDeDE, time.Date(0, 5, 1, 0, 0, 0, 0, time.UTC), "January", "Mai"},
	{LocaleDeDE, time.Date(0, 5, 13, 0, 0, 0, 0, time.UTC), "2 January", "13 Mai"},

	{LocaleHuHU, time.Date(2013, 9, 3, 0, 0, 0, 0, time.UTC), "Mon Jan 2 2006", "K szept 3 2013"},
	{LocaleHuHU, time.Date(2013, 9, 4, 0, 0, 0, 0, time.UTC), "Monday Jan 2 2006", "szerda szept 4 2013"},
	{LocaleHuHU, time.Date(2013, 10, 3, 0, 0, 0, 0, time.UTC), "Monday January 02 2006", "csütörtök október 03 2013"},
	{LocaleHuHU, time.Date(2013, 11, 3, 0, 0, 0, 0, time.UTC), "Monday. 2 January 2006", "vasárnap. 3 november 2013"},
	{LocaleHuHU, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2006. 2 January. Monday", "2013. 13 május. hétfő"},
	{LocaleHuHU, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2 Jan 2006", "13 máj 2013"},
	{LocaleHuHU, time.Date(0, 5, 1, 0, 0, 0, 0, time.UTC), "January", "május"},
	{LocaleHuHU, time.Date(0, 5, 13, 0, 0, 0, 0, time.UTC), "2 January", "13 május"},

	{LocaleItIT, time.Date(2013, 9, 3, 0, 0, 0, 0, time.UTC), "Mon Jan 2 2006", "mar set 3 2013"},
	{LocaleItIT, time.Date(2013, 9, 4, 0, 0, 0, 0, time.UTC), "Monday Jan 2 2006", "mercoledì set 4 2013"},
	{LocaleItIT, time.Date(2013, 10, 3, 0, 0, 0, 0, time.UTC), "Monday January 02 2006", "giovedì ottobre 03 2013"},
	{LocaleItIT, time.Date(2013, 11, 3, 0, 0, 0, 0, time.UTC), "Monday. 2 January 2006", "domenica. 3 novembre 2013"},
	{LocaleItIT, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2006. 2 January. Monday", "2013. 13 maggio. lunedì"},
	{LocaleItIT, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2 Jan 2006", "13 mag 2013"},
	{LocaleItIT, time.Date(0, 5, 1, 0, 0, 0, 0, time.UTC), "January", "maggio"},
	{LocaleItIT, time.Date(0, 5, 13, 0, 0, 0, 0, time.UTC), "2 January", "13 maggio"},

	{LocaleNnNO, time.Date(2013, 9, 3, 0, 0, 0, 0, time.UTC), "Mon Jan 2 2006", "ty sep 3 2013"},
	{LocaleNnNO, time.Date(2013, 9, 4, 0, 0, 0, 0, time.UTC), "Monday Jan 2 2006", "onsdag sep 4 2013"},
	{LocaleNnNO, time.Date(2013, 10, 3, 0, 0, 0, 0, time.UTC), "Monday January 02 2006", "torsdag oktober 03 2013"},
	{LocaleNnNO, time.Date(2013, 11, 3, 0, 0, 0, 0, time.UTC), "Monday. 2 January 2006", "søndag. 3 november 2013"},
	{LocaleNnNO, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2006. 2 January. Monday", "2013. 13 mai. måndag"},
	{LocaleNnNO, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2 Jan 2006", "13 mai 2013"},
	{LocaleNnNO, time.Date(0, 5, 1, 0, 0, 0, 0, time.UTC), "January", "mai"},
	{LocaleNnNO, time.Date(0, 5, 13, 0, 0, 0, 0, time.UTC), "2 January", "13 mai"},

	{LocaleNbNO, time.Date(2013, 9, 3, 0, 0, 0, 0, time.UTC), "Mon Jan 2 2006", "ti sep 3 2013"},
	{LocaleNbNO, time.Date(2013, 9, 4, 0, 0, 0, 0, time.UTC), "Monday Jan 2 2006", "onsdag sep 4 2013"},
	{LocaleNbNO, time.Date(2013, 10, 3, 0, 0, 0, 0, time.UTC), "Monday January 02 2006", "torsdag oktober 03 2013"},
	{LocaleNbNO, time.Date(2013, 11, 3, 0, 0, 0, 0, time.UTC), "Monday. 2 January 2006", "søndag. 3 november 2013"},
	{LocaleNbNO, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2006. 2 January. Monday", "2013. 13 mai. mandag"},
	{LocaleNbNO, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2 Jan 2006", "13 mai 2013"},
	{LocaleNbNO, time.Date(0, 5, 1, 0, 0, 0, 0, time.UTC), "January", "mai"},
	{LocaleNbNO, time.Date(0, 5, 13, 0, 0, 0, 0, time.UTC), "2 January", "13 mai"},

	{LocalePlPL, time.Date(2013, 9, 3, 0, 0, 0, 0, time.UTC), "Mon Jan 2 2006", "Wto Wrz 3 2013"},
	{LocalePlPL, time.Date(2013, 9, 4, 0, 0, 0, 0, time.UTC), "Monday Jan 2 2006", "Środa Wrz 4 2013"},
	{LocalePlPL, time.Date(2013, 10, 3, 0, 0, 0, 0, time.UTC), "Monday January 02 2006", "Czwartek Październik 03 2013"},
	{LocalePlPL, time.Date(2013, 11, 3, 0, 0, 0, 0, time.UTC), "Monday. 2 January 2006", "Niedziela. 3 Listopad 2013"},
	{LocalePlPL, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2006. 2 January. Monday", "2013. 13 Maj. Poniedziałek"},
	{LocalePlPL, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2 Jan 2006", "13 Maj 2013"},
	{LocalePlPL, time.Date(0, 5, 1, 0, 0, 0, 0, time.UTC), "January", "Maj"},
	{LocalePlPL, time.Date(0, 5, 13, 0, 0, 0, 0, time.UTC), "2 January", "13 Maj"},

	{LocalePtPT, time.Date(2013, 9, 3, 0, 0, 0, 0, time.UTC), "Mon Jan 2 2006", "ter Set 3 2013"},
	{LocalePtPT, time.Date(2013, 9, 4, 0, 0, 0, 0, time.UTC), "Monday Jan 2 2006", "Quarta-feira Set 4 2013"},
	{LocalePtPT, time.Date(2013, 10, 3, 0, 0, 0, 0, time.UTC), "Monday January 02 2006", "Quinta-feira Outubro 03 2013"},
	{LocalePtPT, time.Date(2013, 11, 3, 0, 0, 0, 0, time.UTC), "Monday. 2 January 2006", "Domingo. 3 Novembro 2013"},
	{LocalePtPT, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2006. 2 January. Monday", "2013. 13 Maio. Segunda-feira"},
	{LocalePtPT, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2 Jan 2006", "13 Mai 2013"},
	{LocalePtPT, time.Date(0, 5, 1, 0, 0, 0, 0, time.UTC), "January", "Maio"},
	{LocalePtPT, time.Date(0, 5, 13, 0, 0, 0, 0, time.UTC), "2 January", "13 Maio"},

	{LocalePtBR, time.Date(2013, 9, 3, 0, 0, 0, 0, time.UTC), "Mon Jan 2 2006", "ter set 3 2013"},
	{LocalePtBR, time.Date(2013, 9, 4, 0, 0, 0, 0, time.UTC), "Monday Jan 2 2006", "quarta-feira set 4 2013"},
	{LocalePtBR, time.Date(2013, 10, 3, 0, 0, 0, 0, time.UTC), "Monday January 02 2006", "quinta-feira outubro 03 2013"},
	{LocalePtBR, time.Date(2013, 11, 3, 0, 0, 0, 0, time.UTC), "Monday. 2 January 2006", "domingo. 3 novembro 2013"},
	{LocalePtBR, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2006. 2 January. Monday", "2013. 13 maio. segunda-feira"},
	{LocalePtBR, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2 Jan 2006", "13 mai 2013"},
	{LocalePtBR, time.Date(0, 5, 1, 0, 0, 0, 0, time.UTC), "January", "maio"},
	{LocalePtBR, time.Date(0, 5, 13, 0, 0, 0, 0, time.UTC), "2 January", "13 maio"},

	{LocaleRoRO, time.Date(2013, 9, 3, 0, 0, 0, 0, time.UTC), "Mon Jan 2 2006", "Ma sept 3 2013"},
	{LocaleRoRO, time.Date(2013, 9, 4, 0, 0, 0, 0, time.UTC), "Monday Jan 2 2006", "miercuri sept 4 2013"},
	{LocaleRoRO, time.Date(2013, 10, 3, 0, 0, 0, 0, time.UTC), "Monday January 02 2006", "joi octombrie 03 2013"},
	{LocaleRoRO, time.Date(2013, 11, 3, 0, 0, 0, 0, time.UTC), "Monday. 2 January 2006", "duminică. 3 noiembrie 2013"},
	{LocaleRoRO, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2006. 2 January. Monday", "2013. 13 mai. luni"},
	{LocaleRoRO, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2 Jan 2006", "13 mai 2013"},
	{LocaleRoRO, time.Date(0, 5, 1, 0, 0, 0, 0, time.UTC), "January", "mai"},
	{LocaleRoRO, time.Date(0, 5, 13, 0, 0, 0, 0, time.UTC), "2 January", "13 mai"},

	{LocaleRuRU, time.Date(2013, 9, 3, 0, 0, 0, 0, time.UTC), "Mon Jan 2 2006", "Вт Сен 3 2013"},
	{LocaleRuRU, time.Date(2013, 9, 4, 0, 0, 0, 0, time.UTC), "Monday Jan 2 2006", "Среда Сен 4 2013"},
	{LocaleRuRU, time.Date(2013, 10, 3, 0, 0, 0, 0, time.UTC), "Monday January 02 2006", "Четверг Октябрь 03 2013"},
	{LocaleRuRU, time.Date(2013, 11, 3, 0, 0, 0, 0, time.UTC), "Monday. 2 January 2006", "Воскресенье. 3 ноября 2013"},
	{LocaleRuRU, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2006. 2 January. Monday", "2013. 13 мая. Понедельник"},
	{LocaleRuRU, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2 Jan 2006", "13 мая 2013"},
	{LocaleRuRU, time.Date(0, 5, 1, 0, 0, 0, 0, time.UTC), "January", "Май"},
	{LocaleRuRU, time.Date(0, 5, 13, 0, 0, 0, 0, time.UTC), "2 January", "13 мая"},

	{LocaleEsES, time.Date(2013, 9, 3, 0, 0, 0, 0, time.UTC), "Mon Jan 2 2006", "mar sep 3 2013"},
	{LocaleEsES, time.Date(2013, 9, 4, 0, 0, 0, 0, time.UTC), "Monday Jan 2 2006", "miércoles sep 4 2013"},
	{LocaleEsES, time.Date(2013, 10, 3, 0, 0, 0, 0, time.UTC), "Monday January 02 2006", "jueves octubre 03 2013"},
	{LocaleEsES, time.Date(2013, 11, 3, 0, 0, 0, 0, time.UTC), "Monday. 2 January 2006", "domingo. 3 noviembre 2013"},
	{LocaleEsES, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2006. 2 January. Monday", "2013. 13 mayo. lunes"},
	{LocaleEsES, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2 Jan 2006", "13 may 2013"},
	{LocaleEsES, time.Date(0, 5, 1, 0, 0, 0, 0, time.UTC), "January", "mayo"},
	{LocaleEsES, time.Date(0, 5, 13, 0, 0, 0, 0, time.UTC), "2 January", "13 mayo"},

	{LocaleCaES, time.Date(2013, 9, 3, 0, 0, 0, 0, time.UTC), "Mon Jan 2 2006", "dt set 3 2013"},
	{LocaleCaES, time.Date(2013, 9, 4, 0, 0, 0, 0, time.UTC), "Monday Jan 2 2006", "dimecres set 4 2013"},
	{LocaleCaES, time.Date(2013, 10, 3, 0, 0, 0, 0, time.UTC), "Monday January 02 2006", "dijous octubre 03 2013"},
	{LocaleCaES, time.Date(2013, 11, 3, 0, 0, 0, 0, time.UTC), "Monday. 2 January 2006", "diumenge. 3 novembre 2013"},
	{LocaleCaES, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2006. 2 January. Monday", "2013. 13 maig. dilluns"},
	{LocaleCaES, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2 Jan 2006", "13 maig 2013"},
	{LocaleCaES, time.Date(0, 5, 1, 0, 0, 0, 0, time.UTC), "January", "maig"},
	{LocaleCaES, time.Date(0, 5, 13, 0, 0, 0, 0, time.UTC), "2 January", "13 maig"},

	{LocaleSvSE, time.Date(2013, 9, 3, 0, 0, 0, 0, time.UTC), "Mon Jan 2 2006", "tis sep 3 2013"},
	{LocaleSvSE, time.Date(2013, 9, 4, 0, 0, 0, 0, time.UTC), "Monday Jan 2 2006", "onsdag sep 4 2013"},
	{LocaleSvSE, time.Date(2013, 10, 3, 0, 0, 0, 0, time.UTC), "Monday January 02 2006", "torsdag oktober 03 2013"},
	{LocaleSvSE, time.Date(2013, 11, 3, 0, 0, 0, 0, time.UTC), "Monday. 2 January 2006", "söndag. 3 november 2013"},
	{LocaleSvSE, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2006. 2 January. Monday", "2013. 13 maj. måndag"},
	{LocaleSvSE, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2 Jan 2006", "13 maj 2013"},
	{LocaleSvSE, time.Date(0, 5, 1, 0, 0, 0, 0, time.UTC), "January", "maj"},
	{LocaleSvSE, time.Date(0, 5, 13, 0, 0, 0, 0, time.UTC), "2 January", "13 maj"},

	{LocaleTrTR, time.Date(2013, 9, 3, 0, 0, 0, 0, time.UTC), "Mon Jan 2 2006", "Sal Eyl 3 2013"},
	{LocaleTrTR, time.Date(2013, 9, 4, 0, 0, 0, 0, time.UTC), "Monday Jan 2 2006", "Çarşamba Eyl 4 2013"},
	{LocaleTrTR, time.Date(2013, 10, 3, 0, 0, 0, 0, time.UTC), "Monday January 02 2006", "Perşembe Ekim 03 2013"},
	{LocaleTrTR, time.Date(2013, 11, 3, 0, 0, 0, 0, time.UTC), "Monday. 2 January 2006", "Pazar. 3 Kasım 2013"},
	{LocaleTrTR, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2006. 2 January. Monday", "2013. 13 Mayıs. Pazartesi"},
	{LocaleTrTR, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2 Jan 2006", "13 May 2013"},
	{LocaleTrTR, time.Date(0, 5, 1, 0, 0, 0, 0, time.UTC), "January", "Mayıs"},
	{LocaleTrTR, time.Date(0, 5, 13, 0, 0, 0, 0, time.UTC), "2 January", "13 Mayıs"},

	{LocaleUkUA, time.Date(2013, 9, 3, 0, 0, 0, 0, time.UTC), "Mon Jan 2 2006", "Вт Вер 3 2013"},
	{LocaleUkUA, time.Date(2013, 9, 4, 0, 0, 0, 0, time.UTC), "Monday Jan 2 2006", "Середа Вер 4 2013"},
	{LocaleUkUA, time.Date(2013, 10, 3, 0, 0, 0, 0, time.UTC), "Monday January 02 2006", "Четвер Жовтень 03 2013"},
	{LocaleUkUA, time.Date(2013, 11, 3, 0, 0, 0, 0, time.UTC), "Monday. 2 January 2006", "Неділя. 3 листопада 2013"},
	{LocaleUkUA, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2006. 2 January. Monday", "2013. 13 травня. Понеділок"},
	{LocaleUkUA, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2 Jan 2006", "13 тра 2013"},
	{LocaleUkUA, time.Date(0, 5, 1, 0, 0, 0, 0, time.UTC), "January", "Травень"},
	{LocaleUkUA, time.Date(0, 5, 13, 0, 0, 0, 0, time.UTC), "2 January", "13 травня"},

	{LocaleBgBG, time.Date(2013, 9, 3, 0, 0, 0, 0, time.UTC), "Mon Jan 2 2006", "Вт Сеп 3 2013"},
	{LocaleBgBG, time.Date(2013, 9, 4, 0, 0, 0, 0, time.UTC), "Monday Jan 2 2006", "Сряда Сеп 4 2013"},
	{LocaleBgBG, time.Date(2013, 10, 3, 0, 0, 0, 0, time.UTC), "Monday January 02 2006", "Четвъртък Октомври 03 2013"},
	{LocaleBgBG, time.Date(2013, 11, 3, 0, 0, 0, 0, time.UTC), "Monday. 2 January 2006", "Неделя. 3 Ноември 2013"},
	{LocaleBgBG, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2006. 2 January. Monday", "2013. 13 Май. Понеделник"},
	{LocaleBgBG, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2 Jan 2006", "13 Май 2013"},
	{LocaleBgBG, time.Date(0, 5, 1, 0, 0, 0, 0, time.UTC), "January", "Май"},
	{LocaleBgBG, time.Date(0, 5, 13, 0, 0, 0, 0, time.UTC), "2 January", "13 Май"},

	{LocaleZhCN, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2006-01-2", "2013-05-13"},
	{LocaleZhCN, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2006/1/2 Monday", "2013/5/13 星期一"},
	{LocaleZhCN, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2006年1月2日 Monday", "2013年5月13日 星期一"},
	{LocaleZhCN, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2006年1月2日", "2013年5月13日"},
	{LocaleZhCN, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2006年 Jan 2日", "2013年 5 13日"},
	{LocaleZhCN, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2006年 January 2日", "2013年 5 月 13日"},
	{LocaleZhCN, time.Date(0, 5, 1, 0, 0, 0, 0, time.UTC), "January", "5 月"},

	{LocaleZhTW, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2006-01-2", "2013-05-13"},
	{LocaleZhTW, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2006/1/2 Monday", "2013/5/13 星期一"},
	{LocaleZhTW, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2006年1月2日 Monday", "2013年5月13日 星期一"},
	{LocaleZhTW, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2006年1月2日", "2013年5月13日"},
	{LocaleZhTW, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2006年 Jan 2日", "2013年 5 13日"},
	{LocaleZhTW, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2006年 January 2日", "2013年 5 月 13日"},
	{LocaleZhTW, time.Date(0, 5, 1, 0, 0, 0, 0, time.UTC), "January", "5 月"},

	{LocaleZhHK, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2006-01-2", "2013-05-13"},
	{LocaleZhHK, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2006/1/2 Monday", "2013/5/13 星期一"},
	{LocaleZhHK, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2006年1月2日 Monday", "2013年5月13日 星期一"},
	{LocaleZhHK, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2006年1月2日", "2013年5月13日"},
	{LocaleZhHK, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2006年 Jan 2日", "2013年 5 13日"},
	{LocaleZhHK, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2006年 January 2日", "2013年 5 月 13日"},
	{LocaleZhHK, time.Date(0, 5, 1, 0, 0, 0, 0, time.UTC), "January", "5 月"},

	{LocaleKoKR, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2006/01/2", "2013/05/13"},
	{LocaleKoKR, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2006/1/2 Monday", "2013/5/13 월요일"},
	{LocaleKoKR, time.Date(2013, 5, 13, 10, 30, 0, 0, time.UTC), "2006/1/2 Monday 3:04pm", "2013/5/13 월요일 10:30오전"},
	{LocaleKoKR, time.Date(2013, 5, 13, 23, 30, 0, 0, time.UTC), "2006/1/2 Monday 3:04PM", "2013/5/13 월요일 11:30오후"},
	{LocaleKoKR, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2006 Jan 2 Monday", "2013 5월 13 월요일"},
	{LocaleKoKR, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2006 Jan 2", "2013 5월 13"},
	{LocaleKoKR, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2006 January 2", "2013 5월 13"},
	{LocaleKoKR, time.Date(0, 5, 1, 0, 0, 0, 0, time.UTC), "January", "5월"},

	{LocaleJaJP, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2006/01/2", "2013/05/13"},
	{LocaleJaJP, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2006/1/2 Monday", "2013/5/13 月曜日"},
	{LocaleJaJP, time.Date(2013, 5, 13, 10, 30, 0, 0, time.UTC), "2006/1/2 Monday 3:04pm", "2013/5/13 月曜日 10:30午前"},
	{LocaleJaJP, time.Date(2013, 5, 13, 23, 30, 0, 0, time.UTC), "2006/1/2 Monday 3:04PM", "2013/5/13 月曜日 11:30午後"},
	{LocaleJaJP, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2006 Jan 2 Monday", "2013 5月 13 月曜日"},
	{LocaleJaJP, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2006 Jan 2", "2013 5月 13"},
	{LocaleJaJP, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2006 January 2", "2013 5月 13"},
	{LocaleJaJP, time.Date(0, 5, 1, 0, 0, 0, 0, time.UTC), "January", "5月"},

	{LocaleElGR, time.Date(2013, 9, 3, 0, 0, 0, 0, time.UTC), "Mon Jan 2 2006", "Τρι Σεπ 3 2013"},
	{LocaleElGR, time.Date(2013, 9, 4, 0, 0, 0, 0, time.UTC), "Monday Jan 2 2006", "Τετάρτη Σεπ 4 2013"},
	{LocaleElGR, time.Date(2013, 10, 3, 0, 0, 0, 0, time.UTC), "Monday January 02 2006", "Πέμπτη Οκτώβριος 03 2013"},
	{LocaleElGR, time.Date(2013, 11, 3, 0, 0, 0, 0, time.UTC), "Monday. 2 January 2006", "Κυριακή. 3 Νοεμβρίου 2013"},
	{LocaleElGR, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2006. 2 January. Monday", "2013. 13 Μαΐου. Δευτέρα"},
	{LocaleElGR, time.Date(2015, 5, 23, 22, 7, 0, 0, time.UTC), "2 Jan 2006 3:04pm", "23 Μαϊ 2015 10:07μμ"},
	{LocaleElGR, time.Date(0, 5, 1, 0, 0, 0, 0, time.UTC), "January", "Μάιος"},
	{LocaleElGR, time.Date(0, 5, 13, 0, 0, 0, 0, time.UTC), "2 January", "13 Μαΐου"},

	{LocaleIdID, time.Date(2013, 9, 3, 0, 0, 0, 0, time.UTC), "Mon Jan 2 2006", "Sel Sep 3 2013"},
	{LocaleIdID, time.Date(2013, 9, 4, 0, 0, 0, 0, time.UTC), "Monday Jan 2 2006", "Rabu Sep 4 2013"},
	{LocaleIdID, time.Date(2013, 10, 3, 0, 0, 0, 0, time.UTC), "Monday January 02 2006", "Kamis Oktober 03 2013"},
	{LocaleIdID, time.Date(2013, 11, 3, 0, 0, 0, 0, time.UTC), "Monday. 2 January 2006", "Minggu. 3 November 2013"},
	{LocaleIdID, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2006. 2 January. Monday", "2013. 13 Mei. Senin"},
	{LocaleIdID, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2 Jan 2006", "13 Mei 2013"},
	{LocaleIdID, time.Date(0, 5, 1, 0, 0, 0, 0, time.UTC), "January", "Mei"},
	{LocaleIdID, time.Date(0, 5, 13, 0, 0, 0, 0, time.UTC), "2 January", "13 Mei"},

	{LocaleCsCZ, time.Date(2013, 9, 3, 0, 0, 0, 0, time.UTC), "Mon Jan 2 2006", "út zář 3 2013"},
	{LocaleCsCZ, time.Date(2013, 9, 4, 0, 0, 0, 0, time.UTC), "Monday Jan 2 2006", "středa zář 4 2013"},
	{LocaleCsCZ, time.Date(2013, 10, 3, 0, 0, 0, 0, time.UTC), "Monday January 02 2006", "čtvrtek říjen 03 2013"},
	{LocaleCsCZ, time.Date(2013, 11, 3, 0, 0, 0, 0, time.UTC), "Monday. 2 January 2006", "neděle. 3 listopad 2013"},
	{LocaleCsCZ, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2006. 2 January. Monday", "2013. 13 květen. pondělí"},
	{LocaleCsCZ, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2 Jan 2006", "13 kvě 2013"},
	{LocaleCsCZ, time.Date(0, 5, 1, 0, 0, 0, 0, time.UTC), "January", "květen"},
	{LocaleCsCZ, time.Date(0, 5, 13, 0, 0, 0, 0, time.UTC), "2 January", "13 květen"},

	{LocaleSlSI, time.Date(2013, 9, 3, 0, 0, 0, 0, time.UTC), "Mon Jan 2 2006", "tor sep 3 2013"},
	{LocaleSlSI, time.Date(2013, 9, 4, 0, 0, 0, 0, time.UTC), "Monday Jan 2 2006", "sreda sep 4 2013"},
	{LocaleSlSI, time.Date(2013, 10, 3, 0, 0, 0, 0, time.UTC), "Monday January 02 2006", "četrtek oktober 03 2013"},
	{LocaleSlSI, time.Date(2013, 11, 3, 0, 0, 0, 0, time.UTC), "Monday. 2 January 2006", "nedelja. 3 november 2013"},
	{LocaleSlSI, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2006. 2 January. Monday", "2013. 13 maj. ponedeljek"},
	{LocaleSlSI, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2 Jan 2006", "13 maj 2013"},
	{LocaleSlSI, time.Date(0, 5, 1, 0, 0, 0, 0, time.UTC), "January", "maj"},
	{LocaleSlSI, time.Date(0, 5, 13, 0, 0, 0, 0, time.UTC), "2 January", "13 maj"},
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

func BenchmarkFormat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, ts := range formatTests {
			txt := Format(ts.date, ts.layout, ts.locale)

			if txt != ts.expected {
				b.Errorf("failed")
				continue
			}
		}
	}
}

func TestBadLocale(t *testing.T) {
	txt := Format(time.Date(2013, 9, 3, 0, 0, 0, 0, time.UTC), "Mon Jan 2 2006", "aa_AA")
	if txt != "Tue Sep 3 2013" {
		t.Error("Failed to test with bad locale. ", txt)
	}

	_, err := ParseInLocation("Mon January 2006", "Sun April 2013", time.UTC, "aa_AA")
	if err.Error() != "unsupported locale: aa_AA" {
		t.Error("Failed to test with unsupported locale.")
	}
}

func TestGetShortDays(t *testing.T) {
	locales := ListLocales()
	for _, locale := range locales {
		days := GetShortDays(locale)
		if days == nil || len(days) != 7 {
			t.Error("Not expected result. ", days)
		}

		// according to https://www.timeanddate.com/calendar/days/monday.html
		// only Canada, USA and Japan use Sunday as first day of the week
		switch locale {
		case LocaleEnUS:
			if days[0] != shortDayNamesEnUS["Sun"] {
				t.Error("first day of week in US should be Sunday", days)
			}
		case LocaleJaJP:
			if days[0] != shortDayNamesJaJP["Sun"] {
				t.Error("first day of week in JP should be Sunday", days)
			}

		default:
			if days[0] != knownDaysShort[locale]["Mon"] {
				t.Error("first day of week should be Monday", days, locale)
			}
		}
	}
}

func TestGetLongDays(t *testing.T) {
	locales := ListLocales()
	for _, locale := range locales {
		days := GetLongDays(locale)
		if days == nil || len(days) != 7 {
			t.Error("Not expected result. ", days)
		}

		// according to https://www.timeanddate.com/calendar/days/monday.html
		// only Canada, USA and Japan use Sunday as first day of the week
		switch locale {
		case LocaleEnUS:
			if days[0] != longDayNamesEnUS["Sunday"] {
				t.Error("first day of week in US should be Sunday", days)
			}
		case LocaleJaJP:
			if days[0] != longDayNamesJaJP["Sunday"] {
				t.Error("first day of week in JP should be Sunday", days)
			}

		default:
			if days[0] != knownDaysLong[locale]["Monday"] {
				t.Error("first day of week should be Monday", days, locale)
			}
		}
	}
}

func ExampleFormat() {
	t := time.Date(2013, 4, 25, 0, 0, 0, 0, time.UTC)

	locales := ListLocales()
	for _, loc := range locales {
		fmt.Printf("Locale: %s\n", loc)

		// Full date format
		fmt.Printf("    Full: %s\n", Format(t, FullFormatsByLocale[loc], loc))
		// Long date format
		fmt.Printf("    Long: %s\n", Format(t, LongFormatsByLocale[loc], loc))
		// Medium date format
		fmt.Printf("    Medium: %s\n", Format(t, MediumFormatsByLocale[loc], loc))
		// Short date format
		fmt.Printf("    Short: %s\n", Format(t, ShortFormatsByLocale[loc], loc))
		// DateTime format
		fmt.Printf("    DateTime: %s\n", Format(t, DateTimeFormatsByLocale[loc], loc))
	}
}

func ExampleParseInLocation() {
	layout := "2 January 2006 15:04:05 MST"

	fmt.Println(ParseInLocation(layout, "12 April 2013 00:00:00 MST", time.UTC, LocaleEnUS))
	fmt.Println(ParseInLocation(layout, "12 апреля 2013 00:00:00 MST", time.UTC, LocaleRuRU))
	fmt.Println(ParseInLocation(layout, "12 квітня 2013 00:00:00 MST", time.UTC, LocaleUkUA))
}

func ExampleParse() {
	layout := "2 January 2006 15:04:05 MST"

	fmt.Println(Parse(layout, "12 April 2013 00:00:00 MST", LocaleEnUS))
	fmt.Println(Parse(layout, "12 апреля 2013 00:00:00 MST", LocaleRuRU))
	fmt.Println(Parse(layout, "12 квітня 2013 00:00:00 MST", LocaleUkUA))
}
