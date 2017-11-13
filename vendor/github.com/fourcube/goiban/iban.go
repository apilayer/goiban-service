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

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

// describes the structure of an IBAN
type Iban struct {
	countryCode string
	checkDigit  string
	bban        string
	original    string
	bic         string
}

var (
	successIndicator = big.NewInt(1)
	ibanMod          = big.NewInt(97)
)

func (i *Iban) GetCountryCode() string {
	return i.countryCode
}

/**
Returns a pointer to an Iban instance or nil on structural errors.
*/
func ParseToIban(val string) *Iban {
	// Init empty Iban object
	cc := ExtractCountryCode(val)
	checkDigit := extractCheckDigit(val)
	bbanResult, bbanOk := extractBBAN(val)

	if len(cc) == 0 || len(checkDigit) == 0 || !bbanOk {
		return nil
	}

	iban := &Iban{cc, checkDigit, bbanResult.Data, val, ""}
	return iban
}

// CalculateIBAN returns a parser result with an IBAN
func CalculateIBAN(countryCode string, bankCode string, account string) *ParserResult {
	var iban string

	cc := ExtractCountryCode(countryCode)
	if cc == "" {
		return NewParserResult(false, "Invalid country code.", "")
	}

	bankCode = padBankCode(bankCode, countryCode)
	account = padAccountNumber(bankCode, account, countryCode)

	iban = strings.ToUpper(bankCode + account + countryCode + "00")
	allowedLength := getAllowedLength(countryCode)
	if allowedLength > 0 && len(iban) != allowedLength {
		return NewParserResult(false, "Invalid IBAN length. Check your country, bank code and account number for additional digits.", "")
	}

	// If we don't have any info on allowed length, what shall we do?

	iban = toNumericString(iban)
	intBuf := big.NewInt(0)
	intBuf, ok := intBuf.SetString(iban, 10)
	if !ok {
		return NewParserResult(false, "Could not generate check digits.", "")
	}

	checkdigits := big.NewInt(98)
	checkdigits.Sub(checkdigits, intBuf.Mod(intBuf, ibanMod))

	iban = strings.ToUpper(fmt.Sprintf("%s%02d%s%s", countryCode, checkdigits, bankCode, account))

	finalValidation, success := extractBBAN(iban)
	if success {
		return NewParserResult(true, "", iban)
	}

	return NewParserResult(false, finalValidation.Message, "")
}

func padBankCode(bankCode string, countryCode string) string {
	if length := COUNTRY_CODE_TO_BANK_CODE_LENGTH[countryCode]; length > 0 {
		for len(bankCode) < length {
			bankCode = "0" + bankCode
		}
	}

	return bankCode
}

func padAccountNumber(bankCode string, accountNumber string, countryCode string) string {
	bankCodeLength := len(bankCode)
	totalLength := COUNTRY_CODE_TO_LENGTH_MAP[countryCode]
	expectedLength := totalLength - bankCodeLength - 4 // subtract country code and check digits

	if expectedLength < 1 {
		return accountNumber
	}

	for len(accountNumber) < expectedLength {
		accountNumber = "0" + accountNumber
	}

	return accountNumber
}

/*
	Returns a pointer to a goiban.ValidationResult.
*/
func (iban *Iban) Validate() *ValidationResult {
	var ok bool
	validateableString := toNumericString(iban.bban) + countryCodeToNumericString(iban.countryCode) + iban.checkDigit

	intBuf := big.NewInt(0)

	intBuf, ok = intBuf.SetString(validateableString, 10)
	if !ok {
		return NewValidationResult(false, "Could not parse IBAN number.", iban.original)
	}
	result := intBuf.Mod(intBuf, ibanMod)

	if result.Cmp(successIndicator) == 0 {
		return NewValidationResult(true, "", iban.original)
	}

	return NewValidationResult(false, "Validation failed.", iban.original)
}

/*
	Returns true if the string val can be parsed to an Iban Struct.
*/
func IsParseable(val string) *ParserResult {
	// Init empty Iban object
	cc := ExtractCountryCode(val)
	if cc == "" {
		return NewParserResult(false, "Invalid country code.", "")
	}

	checkDigit := extractCheckDigit(val)
	if checkDigit == "" {
		return NewParserResult(false, "Invalid / no check digits found.", "")
	}

	bbanResult, _ := extractBBAN(val)

	return bbanResult
}

/*
	Returns a country code from a string value representing an IBAN.

	Can return an empty string if value is invalid.
*/
func ExtractCountryCode(val string) string {
	// has to be at least two digits long
	if len(val) < 2 {
		return ""
	}

	possibleCountryCode := strings.ToUpper(val[0:2])

	if !isValidChar(possibleCountryCode[0]) ||
		!isValidChar(possibleCountryCode[1]) {
		return ""
	}

	return possibleCountryCode
}

/*
	Returns two check digits from a string value representing an IBAN.

	Can return an empty string if the value is invalid.
*/
func extractCheckDigit(val string) string {
	// starts at position 2 and is 2 digits long
	if len(val) < 4 {
		return ""
	}

	possibleCheckDigit := strings.ToUpper(val[2:4])

	if !isValidNum(possibleCheckDigit[0]) ||
		!isValidNum(possibleCheckDigit[1]) {
		return ""
	}

	return possibleCheckDigit
}

/*
	Returns a BBAN from a string value representing an IBAN.
*/
func extractBBAN(val string) (*ParserResult, bool) {
	// replace all spaces
	val = strings.Replace(val, " ", "", -1)
	// starts at position 4 in the string
	if len(val) < 5 || len(val) > 34 {
		return NewParserResult(false, "Invalid BBAN length.", ""), false
	}

	countryCode := ExtractCountryCode(val)
	allowedLength := getAllowedLength(countryCode)
	if allowedLength > 0 && (len(val) != allowedLength) {
		return NewParserResult(false, "IBAN length invalid. Expected length for "+countryCode+" is "+strconv.Itoa(allowedLength)+".", ""), false
	}

	bban := strings.ToUpper(val[4:len(val)])

	// only alphanumeric chars may be used
	for _, ch := range bban {
		if !isValidNum(uint8(ch)) && !isValidChar(uint8(ch)) {
			return NewParserResult(false, "Invalid characters in BBAN: "+string(ch), ""), false
		}
	}

	return NewParserResult(true, "", bban), true
}

/*
	Returns a numeric representation of a two-character country code

	e.g. DE -> 1314

	The char value is diminished by 55 and it's integer representation is concatenated onto a string.
*/
func countryCodeToNumericString(countryCode string) string {
	if len(countryCode) > 2 || len(countryCode) < 2 {
		return ""
	}
	return toNumericString(countryCode)
}
