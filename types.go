package kpsclient

// Person represents the registration type of the queried individual.
// It expresses the identity type returned by the service.
type Person string

const (
	// PersonTC is used for citizens of the Republic of Turkey.
	PersonTC Person = "tc_vatandasi"

	// PersonYab represents foreign nationals.
	PersonYab Person = "yabanci"

	// PersonMavi represents Blue Card holders.
	PersonMavi Person = "mavi"

	// PersonEmpty is used when the registration type cannot be determined or is not returned.
	PersonEmpty Person = ""
)

// QueryRequest represents the input query data sent to the identity verification service.
// Required fields may vary depending on the service type.
//
// Required Fields:
//   - TCNo
//   - FirstName
//   - LastName
//   - BirthYear
//
// Optional Fields:
//   - BirthMonth
//   - BirthDay
//   - SerialNumber (TCKK serial number – used for new identity cards)
//
// Example Usage:
//
//	q := QueryRequest{
//	    TCNo:       "12345678901",
//	    FirstName:  "AHMET",
//	    LastName:   "YILMAZ",
//	    BirthYear:  "1990",
//	    BirthMonth: "05",
//	    BirthDay:   "12",
//	}
type QueryRequest struct {
	// TCNo is the 11-digit Turkish Republic identity number of the person to be queried.
	TCNo string `json:"tcno"`

	// FirstName should be the person's official first name (all uppercase).
	FirstName string `json:"firstname"`

	// LastName should be the person's official last name (all uppercase).
	LastName string `json:"lastname"`

	// BirthYear represents the person's year of birth (YYYY format).
	BirthYear string `json:"birthyear"`

	// BirthMonth represents the person's birth month (MM format). Optional.
	BirthMonth string `json:"birthmonth,omitempty"`

	// BirthDay represents the person's birth day (DD format). Optional.
	BirthDay string `json:"birthday,omitempty"`

	// SerialNumber is the new generation TCKK serial number. Optional in some validations.
	SerialNumber string `json:"serialnumber,omitempty"`
}

// Result represents the standard response returned from the identity verification service.
//
// Code Field:
//
//	1 = Verification successful. Record found for the individual and information is validated.
//	2 = Verification failed. No record found for the provided information or the details do not match actual information.
//	3 = Record found but the individual is deceased. DeathDate field is present.
//
// Status Field:
//
//	If Code == 1, then true,
//	Otherwise, false.
//
// Extra Field:
//
//	Contains additional information about the queried individual.
//	Example: Name, Surname, IdentityNo, Nationality, BirthDate, DeathDate, etc.
//
// Raw Field:
//
//	Contains the full raw SOAP/XML response from the service.
//	Useful for debugging, logging, or error analysis.
type Result struct {
	// Status is the logical equivalent of the verification result.
	// Returns true when Code == 1; otherwise false.
	Status bool `json:"status"`

	// Code represents the transaction result returned by the service.
	// 1: Success, 2: Failure, 3: Deceased Record.
	Code int `json:"code"`

	// Message is a human-readable short summary of the transaction result.
	// E.g.: "Verification successful", "No record found", etc.
	Message string `json:"message,omitempty"`

	// Person contains basic demographic info for the verified individual.
	// Varies for Turkish citizen, foreign national or Blue Card holder.
	Person Person `json:"person,omitempty"`

	// Extra flexibly holds additional attributes returned by the service.
	// Available fields may change depending on the service.
	// Example: IdentityNo, Name, Surname, Nationality, BirthDate, DeathDate, etc.
	Extra map[string]string `json:"extra,omitempty"`

	// Raw contains the full raw SOAP/XML payload returned by the service.
	// Useful in development and debugging phases.
	Raw string `json:"raw,omitempty"`
}
