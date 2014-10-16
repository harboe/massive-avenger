package model

import ()

type Bank struct {
	Address Address `json:"address"`
	Account string  `json:"account"`
	// The SWIFT code consists of 8 or 11 characters. When 8-digits code is given, it refers to the primary office.
	// * First 4 characters - bank code (only letters)
	// * Next 2 characters - ISO 3166-1 alpha-2 country code (only letters)
	// * Next 2 characters - location code (letters and digits) (passive participant will have "1" in the second character)
	// * Last 3 characters - branch code, optional ('XXX' for primary office) (letters and digits)
	SwiftCode string `json:"swift"`
}
