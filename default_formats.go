package monday

// Default date formats by country.
// Mostly taken from http://en.wikipedia.org/wiki/Date_format_by_country
const (
    DefaultFormatEnUS = "01/02/06"  
    
    DefaultFormatEnUSFull = "Monday, January 2, 2006"    // English (United States)
    DefaultFormatEnUSLong = "January 2, 2006"            
    DefaultFormatEnUSMedium = "Jan 02, 2006"
    DefaultFormatEnUSShort = "1/2/06"
    DefaultFormatEnUSDateTime = "1/2/06 3:04 PM"

    DefaultFormatEnGBFull = "Monday, 2 January 2006"    // English (United Kingdom)
    DefaultFormatEnGBLong = "2 January 2006"            
    DefaultFormatEnGBMedium = "02 Jan 2006"
    DefaultFormatEnGBShort = "02/01/2006"
    DefaultFormatEnGBDateTime = "02/01/2006 15:04"
    
    DefaultFormatDaDKFull = "Monday den 2. January 2006"    // Danish (Denmark)
    DefaultFormatDaDKLong = "2. Jan 2006"                  
    DefaultFormatDaDKMedium = "02/01/2006"
    DefaultFormatDaDKShort = "02/01/06"
    DefaultFormatDaDKDateTime = "02/01/2006 15.04"

    DefaultFormatNlBEFull = "Monday 2 January 2006"    // Dutch (Belgium)
    DefaultFormatNlBELong = "2 January 2006"           
    DefaultFormatNlBEMedium = "02-Jan-2006"
    DefaultFormatNlBEShort = "2/01/06"
    DefaultFormatNlBEDateTime = "2/01/06 15:04"

    DefaultFormatNlNLFull = "Monday 2 January 2006"    // Dutch (Netherlands)
    DefaultFormatNlNLLong = "2 January 2006"           
    DefaultFormatNlNLMedium = "02 Jan 2006"
    DefaultFormatNlNLShort = "02-01-06"
    DefaultFormatNlNLDateTime = "02-01-06 15:04"

    DefaultFormatFiFIFull = "Monday 2. January 2006"    // Finnish (Finland)
    DefaultFormatFiFILong = "2. January 2006"           
    DefaultFormatFiFIMedium = "02.1.2006"
    DefaultFormatFiFIShort = "02.1.2006"
    DefaultFormatFiFIDateTime = "02.1.2006 15.04"

    DefaultFormatFrFRFull = "Monday 2 January 2006"    // French (France)
    DefaultFormatFrFRLong = "2 January 2006"           
    DefaultFormatFrFRMedium = "02 Jan 2006"
    DefaultFormatFrFRShort = "02/01/2006"
    DefaultFormatFrFRDateTime = "02/01/2006 15:04"

    DefaultFormatFrCAFull = "Monday 2 January 2006"    // French (Canada)
    DefaultFormatFrCALong = "2 January 2006"           
    DefaultFormatFrCAMedium = "2006-01-02"
    DefaultFormatFrCAShort = "06-01-02"
    DefaultFormatFrCADateTime = "06-01-02 15:04"

    DefaultFormatDeDEFull = "Monday, 2. January 2006"    // German (Germany)
    DefaultFormatDeDELong = "2. January 2006"            
    DefaultFormatDeDEMedium = "02.01.2006"
    DefaultFormatDeDEShort = "02.01.06"
    DefaultFormatDeDEDateTime = "02.01.06 15:04"
    
    DefaultFormatHuHUFull = "2006. January 2., Monday"    // Hungarian (Hungary)
    DefaultFormatHuHULong = "2006. January 2."
    DefaultFormatHuHUMedium = "2006.01.02."
    DefaultFormatHuHUShort = "2006.01.02."
    DefaultFormatHuHUDateTime = "2006.01.02. 15:04"

    DefaultFormatItITFull = "Monday 2 January 2006"    // Italian (Italy)
    DefaultFormatItITLong = "2 January 2006"           
    DefaultFormatItITMedium = "02/Jan/2006"
    DefaultFormatItITShort = "02/01/06"
    DefaultFormatItITDateTime = "02/01/06 15:04"

    DefaultFormatNnNOFull = "Monday 2. January 2006"    // Norwegian Nynorsk (Norway)
    DefaultFormatNnNOLong = "2. January 2006"           
    DefaultFormatNnNOMedium = "02. Jan 2006"
    DefaultFormatNnNOShort = "02.01.06"
    DefaultFormatNnNODateTime = "02.01.06 15:04"

    DefaultFormatNbNOFull = "Monday 2. January 2006"    // Norwegian Bokmål (Norway)
    DefaultFormatNbNOLong = "2. January 2006"           
    DefaultFormatNbNOMedium = "02. Jan 2006"
    DefaultFormatNbNOShort = "02.01.06"
    DefaultFormatNbNODateTime = "15:04 02.01.06"

    DefaultFormatPtPTFull = "Monday, 2 de January de 2006"    // Portuguese (Portugal)
    DefaultFormatPtPTLong = "2 de January de 2006"            
    DefaultFormatPtPTMedium = "02/01/2006"
    DefaultFormatPtPTShort = "02/01/06"
    DefaultFormatPtPTDateTime = "02/01/06, 15:04"

    DefaultFormatPtBRFull = "Monday, 2 de January de 2006"    // Portuguese (Brazil)
    DefaultFormatPtBRLong = "02 de January de 2006"            
    DefaultFormatPtBRMedium = "02/01/2006"
    DefaultFormatPtBRShort = "02/01/06"
    DefaultFormatPtBRDateTime = "02/01/06, 15:04"  

    DefaultFormatRoROFull = "Monday, 02 January 2006"    // Romanian (Romania)
    DefaultFormatRoROLong = "02 January 2006"           
    DefaultFormatRoROMedium = "02.01.2006"
    DefaultFormatRoROShort = "02.01.2006"
    DefaultFormatRoRODateTime = "02.01.06, 15:04"  

    DefaultFormatRuRUFull = "Monday, 2 January 2006 г."    // Russian (Russia)
    DefaultFormatRuRULong = "2 January 2006 г."              
    DefaultFormatRuRUMedium = "02 Jan 2006 г."
    DefaultFormatRuRUShort = "02.01.06"
    DefaultFormatRuRUDateTime = "02.01.06, 15:04"  

    DefaultFormatEsESFull = "Monday, 2 de January de 2006"    // Spanish (Spain)
    DefaultFormatEsESLong = "2 de January de 2006"            
    DefaultFormatEsESMedium = "02/01/2006"
    DefaultFormatEsESShort = "02/01/06"
    DefaultFormatEsESDateTime = "02/01/06 15:04"  

    DefaultFormatSvSEFull = "Mondayen den 2:e January 2006"   // Swedish (Sweden)
    DefaultFormatSvSELong = "2 January 2006"                  
    DefaultFormatSvSEMedium = "2 Jan 2006"
    DefaultFormatSvSEShort = "2006-01-02"
    DefaultFormatSvSEDateTime = "2006-01-02 15:04"  

    DefaultFormatTrTRFull = "2 January 2006 Monday"   // Turkish (Turkey)
    DefaultFormatTrTRLong = "2 January 2006"          
    DefaultFormatTrTRMedium = "2 Jan 2006"
    DefaultFormatTrTRShort = "2.01.2006"
    DefaultFormatTrTRDateTime = "2.01.2006 15:04"  
)

// 'Full' date formats for all supported locales
var FullFormatsByLocale = map[Locale]string {
    LocaleEnUS : DefaultFormatEnUSFull,
    LocaleEnGB : DefaultFormatEnGBFull,
    LocaleDaDK : DefaultFormatDaDKFull,
    LocaleNlBE : DefaultFormatNlBEFull,
    LocaleNlNL : DefaultFormatNlNLFull,
    LocaleFiFI : DefaultFormatFiFIFull,
    LocaleFrFR : DefaultFormatFrFRFull,
    LocaleFrCA : DefaultFormatFrCAFull,
    LocaleDeDE : DefaultFormatDeDEFull,
    LocaleHuHU : DefaultFormatHuHUFull,
    LocaleItIT : DefaultFormatItITFull,
    LocaleNnNO : DefaultFormatNnNOFull,
    LocaleNbNO : DefaultFormatNbNOFull,
    LocalePtPT : DefaultFormatPtPTFull,
    LocalePtBR : DefaultFormatPtBRFull,
    LocaleRoRO : DefaultFormatRoROFull,
    LocaleRuRU : DefaultFormatRuRUFull,
    LocaleEsES : DefaultFormatEsESFull,
    LocaleSvSE : DefaultFormatSvSEFull,
    LocaleTrTR : DefaultFormatTrTRFull,
}

// 'Long' date formats for all supported locales
var LongFormatsByLocale = map[Locale]string {
    LocaleEnUS : DefaultFormatEnUSLong,
    LocaleEnGB : DefaultFormatEnGBLong,
    LocaleDaDK : DefaultFormatDaDKLong,
    LocaleNlBE : DefaultFormatNlBELong,
    LocaleNlNL : DefaultFormatNlNLLong,
    LocaleFiFI : DefaultFormatFiFILong,
    LocaleFrFR : DefaultFormatFrFRLong,
    LocaleFrCA : DefaultFormatFrCALong,
    LocaleDeDE : DefaultFormatDeDELong,
    LocaleHuHU : DefaultFormatHuHULong,
    LocaleItIT : DefaultFormatItITLong,
    LocaleNnNO : DefaultFormatNnNOLong,
    LocaleNbNO : DefaultFormatNbNOLong,
    LocalePtPT : DefaultFormatPtPTLong,
    LocalePtBR : DefaultFormatPtBRLong,
    LocaleRoRO : DefaultFormatRoROLong,
    LocaleRuRU : DefaultFormatRuRULong,
    LocaleEsES : DefaultFormatEsESLong,
    LocaleSvSE : DefaultFormatSvSELong,
    LocaleTrTR : DefaultFormatTrTRLong,
}

// 'Medium' date formats for all supported locales
var MediumFormatsByLocale = map[Locale]string {
    LocaleEnUS : DefaultFormatEnUSMedium,
    LocaleEnGB : DefaultFormatEnGBMedium,
    LocaleDaDK : DefaultFormatDaDKMedium,
    LocaleNlBE : DefaultFormatNlBEMedium,
    LocaleNlNL : DefaultFormatNlNLMedium,
    LocaleFiFI : DefaultFormatFiFIMedium,
    LocaleFrFR : DefaultFormatFrFRMedium,
    LocaleFrCA : DefaultFormatFrCAMedium,
    LocaleDeDE : DefaultFormatDeDEMedium,
    LocaleHuHU : DefaultFormatHuHUMedium,
    LocaleItIT : DefaultFormatItITMedium,
    LocaleNnNO : DefaultFormatNnNOMedium,
    LocaleNbNO : DefaultFormatNbNOMedium,
    LocalePtPT : DefaultFormatPtPTMedium,
    LocalePtBR : DefaultFormatPtBRMedium,
    LocaleRoRO : DefaultFormatRoROMedium,
    LocaleRuRU : DefaultFormatRuRUMedium,
    LocaleEsES : DefaultFormatEsESMedium,
    LocaleSvSE : DefaultFormatSvSEMedium,
    LocaleTrTR : DefaultFormatTrTRMedium,
}

// 'Short' date formats for all supported locales
var ShortFormatsByLocale = map[Locale]string {
    LocaleEnUS : DefaultFormatEnUSShort,
    LocaleEnGB : DefaultFormatEnGBShort,
    LocaleDaDK : DefaultFormatDaDKShort,
    LocaleNlBE : DefaultFormatNlBEShort,
    LocaleNlNL : DefaultFormatNlNLShort,
    LocaleFiFI : DefaultFormatFiFIShort,
    LocaleFrFR : DefaultFormatFrFRShort,
    LocaleFrCA : DefaultFormatFrCAShort,
    LocaleDeDE : DefaultFormatDeDEShort,
    LocaleHuHU : DefaultFormatHuHUShort,
    LocaleItIT : DefaultFormatItITShort,
    LocaleNnNO : DefaultFormatNnNOShort,
    LocaleNbNO : DefaultFormatNbNOShort,
    LocalePtPT : DefaultFormatPtPTShort,
    LocalePtBR : DefaultFormatPtBRShort,
    LocaleRoRO : DefaultFormatRoROShort,
    LocaleRuRU : DefaultFormatRuRUShort,
    LocaleEsES : DefaultFormatEsESShort,
    LocaleSvSE : DefaultFormatSvSEShort,
    LocaleTrTR : DefaultFormatTrTRShort,
}

// 'DateTime' date formats for all supported locales
var DateTimeFormatsByLocale = map[Locale]string {
    LocaleEnUS : DefaultFormatEnUSDateTime,
    LocaleEnGB : DefaultFormatEnGBDateTime,
    LocaleDaDK : DefaultFormatDaDKDateTime,
    LocaleNlBE : DefaultFormatNlBEDateTime,
    LocaleNlNL : DefaultFormatNlNLDateTime,
    LocaleFiFI : DefaultFormatFiFIDateTime,
    LocaleFrFR : DefaultFormatFrFRDateTime,
    LocaleFrCA : DefaultFormatFrCADateTime,
    LocaleDeDE : DefaultFormatDeDEDateTime,
    LocaleHuHU : DefaultFormatHuHUDateTime,
    LocaleItIT : DefaultFormatItITDateTime,
    LocaleNnNO : DefaultFormatNnNODateTime,
    LocaleNbNO : DefaultFormatNbNODateTime,
    LocalePtPT : DefaultFormatPtPTDateTime,
    LocalePtBR : DefaultFormatPtBRDateTime,
    LocaleRoRO : DefaultFormatRoRODateTime,
    LocaleRuRU : DefaultFormatRuRUDateTime,
    LocaleEsES : DefaultFormatEsESDateTime,
    LocaleSvSE : DefaultFormatSvSEDateTime,
    LocaleTrTR : DefaultFormatTrTRDateTime,
}