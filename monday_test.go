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

	{LocaleJaJP, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2006/01/2", "2013/05/13"},
	{LocaleJaJP, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2006/1/2 Monday", "2013/5/13 月曜日"},
	{LocaleJaJP, time.Date(2013, 5, 13, 10, 30, 0, 0, time.UTC), "2006/1/2 Monday 3:04pm", "2013/5/13 月曜日 10:30午前"},
	{LocaleJaJP, time.Date(2013, 5, 13, 23, 30, 0, 0, time.UTC), "2006/1/2 Monday 3:04PM", "2013/5/13 月曜日 11:30午後"},
	{LocaleJaJP, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2006 Jan 2 Monday", "2013 5月 13 月曜日"},
	{LocaleJaJP, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2006 Jan 2", "2013 5月 13"},
	{LocaleJaJP, time.Date(2013, 5, 13, 0, 0, 0, 0, time.UTC), "2006 January 2", "2013 5月 13"},
	{LocaleJaJP, time.Date(0, 5, 1, 0, 0, 0, 0, time.UTC), "January", "5月"},
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

func TestBadLocale(t *testing.T) {
	txt := Format(time.Date(2013, 9, 3, 0, 0, 0, 0, time.UTC), "Mon Jan 2 2006", "aa_AA")
	if txt != "Tue Sep 3 2013" {
		t.Error("Failed to test with bad locale. ", txt)
	}
}

func TestGetShortDays(t *testing.T) {
	locales := ListLocales()
	for _, locale := range locales {
		if days := GetShortDays(locale); days == nil {
			t.Error("Not expected result. ", days)
		}
	}
}

func TestGetLongDays(t *testing.T) {
	locales := ListLocales()
	for _, locale := range locales {
		if days := GetLongDays(locale); days == nil {
			t.Error("Not expected result. ", days)
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
}
