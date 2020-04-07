package stringy

// const below are used in packages
const (
	First                = "first"
	Last                 = "last"
	Left                 = "left"
	Right                = "right"
	Both                 = "both"
	OddError             = "odd number rule provided please provide in even count"
	SelectCapital        = "([a-z])([A-Z])"
	ReplaceCapital       = "$1 $2"
	LengthError          = "passed length cannot be greater than input length"
	InvalidLogicalString = "invalid string value to test boolean value"
)

// False is slice of array for false logical representation in string
var False = []string{"off", "no", "0", "false"}

// True is slice of array for true logical representation in string
var True = []string{"on", "yes", "1", "True"}
