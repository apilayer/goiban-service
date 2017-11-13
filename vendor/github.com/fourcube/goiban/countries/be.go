/*
The MIT License (MIT)

Copyright (c) 2015 Chris Grieger

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
	"fmt"
	"strconv"
	"strings"
)

type BelgiumFileEntry struct {
	Bankcode string
	Name     string
	Bic      string
}

// Return a slice of BankEntries because BICs can map to multiple
// Bankcodes
func BelgiumRowToEntry(row []string) []BelgiumFileEntry {
	var entries []BelgiumFileEntry
	// Those are NAP / NAV / VRIJ entries
	if len(row[2]) < 8 {
		return entries
	}

	lowerBankcodeBound, _ := strconv.Atoi(row[0])
	upperBankcodeBound, _ := strconv.Atoi(row[1])

	for lowerBankcodeBound <= upperBankcodeBound {
		entries = append(entries, BelgiumFileEntry{
			Bankcode: strings.TrimSpace(fmt.Sprintf("%03d", lowerBankcodeBound)),
			Name:     strings.TrimSpace(row[3]),
			Bic:      strings.TrimSpace(row[2]),
		})

		lowerBankcodeBound++
	}

	return entries
}
