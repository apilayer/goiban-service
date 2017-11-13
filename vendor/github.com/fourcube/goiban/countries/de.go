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

package countries

import (
	"strconv"
	s "strings"
)

/*
	Describes the field width of the fixed width text file
	which can be found at
	http://www.bundesbank.de/Redaktion/DE/Standardartikel/Kerngeschaeftsfelder/Unbarer_Zahlungsverkehr/bankleitzahlen_download.html?searchArchive=0&submit=Suchen&searchIssued=0&templateQueryString=Bankleitzahlen

	Example:
	bankcodemname######                                                zip##city                               short#####                 pan##bic########z##id###cdnext####
	100000001Bundesbank                                                10591Berlin                             BBk Berlin                 20100MARKDEF110009011380U000000000
*/

type BundesbankFileEntry struct {
	Bankcode  string // 8
	M         int    // 1
	Name      string // 58
	Zip       string // 5
	City      string // 35
	ShortName string // 27
	Pan       int    // 5
	Bic       string // 11
	CheckAlgo string // 2 	enumerates some checksum algorithms
	// 		described in http://www.bundesbank.de/Redaktion/DE/Downloads/Kerngeschaeftsfelder/Unbarer_Zahlungsverkehr/pruefzifferberechnungsmethoden.pdf?__blob=publicationFile
	Id          string // 5	internal id
	Change      string // 1
	ToBeDeleted int    // 1
	NewBankCode string // 8
}

func BundesbankStringToEntry(val string) *BundesbankFileEntry {
	runeVal := []rune(val)
	m, _ := strconv.Atoi(string(runeVal[8:9]))
	pan, _ := strconv.Atoi(string(runeVal[134:139]))
	id := string(runeVal[152:157])
	toBeDeleted, _ := strconv.Atoi(string(runeVal[159:160]))

	return &BundesbankFileEntry{
		toTrimmedString(runeVal[0:8]),
		m,
		toTrimmedString(runeVal[9:67]),
		toTrimmedString(runeVal[67:72]),
		toTrimmedString(runeVal[72:107]),
		toTrimmedString(runeVal[107:134]),
		pan,
		toTrimmedString(runeVal[139:150]),
		toTrimmedString(runeVal[150:152]),
		id,
		toTrimmedString(runeVal[158:159]),
		toBeDeleted,
		toTrimmedString(runeVal[160:168]),
	}
}

/*
	Converts an array of runes to a trimmed string
*/
func toTrimmedString(runes []rune) string {
	return s.TrimSpace(string(runes))
}
