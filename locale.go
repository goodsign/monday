package monday

// Locale identifies locales supported by 'monday' package.
// Here the ICU locale identifiers are used to provide ICU compatibility if
// needed. See http://userguide.icu-project.org/locale
type Locale string

const (
    LocaleEnUS = "en_US"    // English (United States)
    LocaleEnGB = "en_GB"    // English (United Kingdom)
    LocaleDaDK = "da_DK"    // Danish (Denmark)
    LocaleNlBE = "nl_BE"    // Dutch (Belgium)
    LocaleNlNL = "nl_NL"    // Dutch (Netherlands)
    LocaleFiFI = "fi_FI"    // Finnish (Finland)
    LocaleFrFR = "fr_FR"    // French (France)
    LocaleFrCA = "fr_CA"    // French (Canada)
    LocaleDeDE = "de_DE"    // German (Germany)
    LocaleHuHU = "hu_HU"    // Hungarian (Hungary)
    LocaleItIT = "it_IT"    // Italian (Italy)
    LocaleNnNO = "nn_NO"    // Norwegian Nynorsk (Norway)
    LocaleNbNO = "nb_NO"    // Norwegian Bokmål (Norway)
    LocalePtPT = "pt_PT"    // Portuguese (Portugal)
    LocalePtBR = "pt_BR"    // Portuguese (Brazil)
    LocaleRoRO = "ro_RO"    // Romanian (Romania)
    LocaleRuRU = "ru_RU"    // Russian (Russia)
    LocaleEsES = "es_ES"    // Spanish (Spain)
    LocaleSvSE = "sv_SE"    // Swedish (Sweden)
    LocaleTrTR = "tr_TR"    // Turkish (Turkey)
)