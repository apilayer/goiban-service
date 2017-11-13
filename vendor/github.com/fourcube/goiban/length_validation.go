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
	COUNTRY_CODE_TO_BANK_CODE_LENGTH = map[string]int{
		"DE": 8,
		"BE": 3,
		"NL": 4,
		"LU": 3,
		"CH": 5,
		"AT": 5,
		"LI": 5,
	}

	COUNTRY_CODE_TO_LENGTH_MAP = map[string]int{
		"LV": 21,
		"LU": 20,
		"LT": 20,
		"HR": 21,
		"RO": 24,
		"DZ": 24,
		"VG": 24,
		"HU": 28,
		"MG": 27,
		"DO": 28,
		"ME": 22,
		"MK": 19,
		"ML": 28,
		"DE": 22,
		"MC": 27,
		"MD": 24,
		"DK": 18,
		"IE": 22,
		"AT": 20,
		"MU": 30,
		"IL": 23,
		"MZ": 25,
		"IR": 26,
		"AZ": 28,
		"IS": 26,
		"IT": 27,
		"MR": 27,
		"BA": 20,
		"MT": 31,
		"PT": 25,
		"AD": 24,
		"UA": 29,
		"ES": 24,
		"AE": 23,
		"NL": 18,
		"PS": 29,
		"EG": 27,
		"AL": 28,
		"EE": 20,
		"AO": 25,
		"GE": 22,
		"BR": 29,
		"GA": 27,
		"GB": 22,
		"TN": 24,
		"TR": 26,
		"NO": 15,
		"BF": 27,
		"FR": 27,
		"BG": 22,
		"BH": 22,
		"BI": 16,
		"FO": 18,
		"BE": 16,
		"BJ": 28,
		"FI": 18,
		"CZ": 24,
		"CY": 28,
		"SE": 24,
		"CV": 25,
		"SI": 19,
		"KW": 30,
		"SK": 24,
		"SN": 28,
		"KZ": 20,
		"SM": 27,
		"CI": 28,
		"PL": 28,
		"RS": 22,
		"GT": 28,
		"CG": 27,
		"LB": 28,
		"CH": 21,
		"GR": 27,
		"PK": 24,
		"LI": 21,
		"CR": 21,
		"GL": 18,
		"CM": 27,
		"GI": 23,
		"SA": 24,
	}
)

/*
	Returns the allowed length of an IBAN code for a certain country.
			or -1 if the country code could not be looked up.
*/
func getAllowedLength(countryCode string) int {
	var length int
	var ok bool

	length, ok = COUNTRY_CODE_TO_LENGTH_MAP[countryCode]
	if ok {
		return length
	} else {
		return -1
	}
}
