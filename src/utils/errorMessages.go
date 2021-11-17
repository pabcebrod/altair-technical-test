package utils

import "errors"

// Error codes returned by failures.
var (
	NO_PERSON_ID = errors.New("There is no person with that ID in the team")

)