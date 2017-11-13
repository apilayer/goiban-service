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
	"database/sql"
	"strconv"
)

var (
	SELECT_BY_BANK_CODE      = "SELECT 1 FROM BANK_DATA WHERE bankcode = ? and country = ?;"
	SELECT_BY_BANK_CODE_STMT *sql.Stmt
)

func ValidateBankCode(iban *Iban, intermediateResult *ValidationResult, db *sql.DB) *ValidationResult {
	if SELECT_BY_BANK_CODE_STMT == nil {
		prepareSelectByBankCodeStatement(db)
	}

	length, ok := COUNTRY_CODE_TO_BANK_CODE_LENGTH[(iban.countryCode)]

	if !ok {
		intermediateResult.CheckResults["bankCode"] = false
		intermediateResult.Messages = append(intermediateResult.Messages, "Cannot validate bank code length. No information available.")
		return intermediateResult
	}

	if len(iban.bban) < length {
		intermediateResult.CheckResults["bankCode"] = false
		intermediateResult.Valid = false

		intermediateResult.Messages = append(intermediateResult.Messages, "Bank code validation impossible; Invalid bank code length for country '"+iban.countryCode+"' (expected "+strconv.Itoa(length)+" digits)")
		return intermediateResult
	}

	bankCode := iban.bban[0:length]
	//bankCode = strings.TrimLeft(bankCode, "0")

	var res int
	err := SELECT_BY_BANK_CODE_STMT.QueryRow(bankCode, iban.countryCode).Scan(&res)

	switch {
	case err == sql.ErrNoRows:
		intermediateResult.Valid = false
		intermediateResult.Messages = append(intermediateResult.Messages, "Invalid bank code: "+bankCode)
		intermediateResult.CheckResults["bankCode"] = false
		return intermediateResult
	}
	intermediateResult.Messages = append(intermediateResult.Messages, "Bank code valid: "+bankCode)
	intermediateResult.CheckResults["bankCode"] = true

	return intermediateResult
}

func prepareSelectByBankCodeStatement(db *sql.DB) {
	var err error
	SELECT_BY_BANK_CODE_STMT, err = db.Prepare(SELECT_BY_BANK_CODE)
	if err != nil {
		panic("Couldn't prepare statement: " + SELECT_BY_BANK_CODE)
	}
}
