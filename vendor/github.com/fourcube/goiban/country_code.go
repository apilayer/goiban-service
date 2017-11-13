/*
The MIT License (MIT)

Copyright (c) 2014 Chris Grieger

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package goiban

var (
	COUNTRY_TO_CC_MAP = map[string]string{
		"Albania":                   "AL",
		"Algeria":                   "DZ",
		"Andorra":                   "AD",
		"Angola":                    "AO",
		"Austria":                   "AT",
		"Azerbaijan":                "AZ",
		"Bahrain":                   "BH",
		"Belgium":                   "BE",
		"Benin":                     "BJ",
		"Bosnia and Herzegovina":    "BA",
		"Brazil":                    "BR",
		"British Virgin Islands":    "VG",
		"Bulgaria":                  "BG",
		"Burkina Faso":              "BF",
		"Burundi":                   "BI",
		"Cameroon":                  "CM",
		"Cape Verde":                "CV",
		"Central African Republic":  "FR",
		"Congo":                     "CG",
		"Costa Rica":                "CR",
		"Croatia":                   "HR",
		"Cyprus":                    "CY",
		"Czech Republic":            "CZ",
		"Denmark":                   "DK",
		"Dominican Republic":        "DO",
		"Egypt":                     "EG",
		"Estonia":                   "EE",
		"Faroe Islands":             "FO",
		"Finland":                   "FI",
		"France":                    "FR",
		"French Guiana":             "FR",
		"French Polynesia":          "FR",
		"Gabon":                     "GA",
		"Georgia":                   "GE",
		"Germany":                   "DE",
		"Gibraltar":                 "GI",
		"Greece":                    "GR",
		"Greenland":                 "GL",
		"Guadeloupe":                "FR",
		"Guatemala":                 "GT",
		"Guernsey":                  "GB",
		"Hungary":                   "HU",
		"Iceland":                   "IS",
		"Iran":                      "IR",
		"Ireland":                   "IE",
		"Isle of Man":               "GB",
		"Israel":                    "IL",
		"Italy":                     "IT",
		"Ivory Coast":               "CI",
		"Jersey":                    "GB",
		"Kazakhstan":                "KZ",
		"Kuwait":                    "KW",
		"Latvia":                    "LV",
		"Lebanon":                   "LB",
		"Liechtenstein":             "LI",
		"Lithuania":                 "LT",
		"Luxembourg":                "LU",
		"Macedonia":                 "MK",
		"Madagascar":                "MG",
		"Mali":                      "ML",
		"Malta":                     "MT",
		"Martinique":                "FR",
		"Mauritania":                "MR",
		"Mauritius":                 "MU",
		"Moldova":                   "MD",
		"Monaco":                    "MC",
		"Montenegro":                "ME",
		"Mozambique":                "MZ",
		"Netherlands":               "NL",
		"New Caledonia":             "FR",
		"Norway":                    "NO",
		"Pakistan":                  "PK",
		"Palestine, State of":       "PS",
		"Poland":                    "PL",
		"Portugal":                  "PT",
		"Romania":                   "RO",
		"Réunion":                   "FR",
		"Saint-Pierre and Miquelon": "FR",
		"San Marino":                "SM",
		"Sao Tome and Principe":     "PT",
		"Saudi Arabia":              "SA",
		"Senegal":                   "SN",
		"Serbia":                    "RS",
		"Slovakia":                  "SK",
		"Slovenia":                  "SI",
		"Spain":                     "ES",
		"Sweden":                    "SE",
		"Switzerland":               "CH",
		"Tunisia":                   "TN",
		"Turkey":                    "TR",
		"Ukraine":                   "UA",
		"United Arab Emirates":      "AE",
		"United Kingdom":            "GB",
		"Wallis and Futuna":         "FR",
	}
)

/*
	Returns the 2-digit country code for a country name.
	or an empty string if the country code could not be looked up.
*/
func getCountryCode(countryName string) string {
	var countryCode string
	var ok bool

	countryCode, ok = COUNTRY_TO_CC_MAP[countryName]
	if ok {
		return countryCode
	}

	return ""
}
