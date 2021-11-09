package stringy

import (
	"errors"
	"math"
	"math/rand"
	"regexp"
	"strings"
	"time"
	"unicode"
)

// input is struct that holds input from user and result
type input struct {
	Input  string
	Result string
}

// StringManipulation is an interface that holds all abstract methods to manipulate strings
type StringManipulation interface {
	Between(start, end string) StringManipulation
	Boolean() bool
	CamelCase(rule ...string) string
	ContainsAll(check ...string) bool
	Delimited(delimiter string, rule ...string) StringManipulation
	First(length int) string
	Get() string
	KebabCase(rule ...string) StringManipulation
	Last(length int) string
	LcFirst() string
	Lines() []string
	Pad(length int, with, padType string) string
	RemoveSpecialCharacter() string
	ReplaceFirst(search, replace string) string
	ReplaceLast(search, replace string) string
	Reverse() string
	Shuffle() string
	Surround(with string) string
	SnakeCase(rule ...string) StringManipulation
	Tease(length int, indicator string) string
	ToLower() string
	ToUpper() string
	UcFirst() string
	Prefix(with string) string
	Suffix(with string) string
}

// New func returns pointer to input struct
func New(val string) StringManipulation {
	return &input{Input: val}
}

// Between takes two string params start and end which and returns
// value which is in middle of start and end part of input. You can
// chain to upper which with make result all upercase or ToLower which
// will make result all lower case or Get which will return result as it is
func (i *input) Between(start, end string) StringManipulation {
	if (start == "" && end == "") || i.Input == "" {
		return i
	}

	input := strings.ToLower(i.Input)
	lcStart := strings.ToLower(start)
	lcEnd := strings.ToLower(end)
	var startIndex, endIndex int

	if len(start) > 0 && strings.Contains(input, lcStart) {
		startIndex = len(start)
	}
	if len(end) > 0 && strings.Contains(input, lcEnd) {
		endIndex = strings.Index(input, lcEnd)
	} else if len(input) > 0 {
		endIndex = len(input)
	}
	i.Result = strings.TrimSpace(i.Input[startIndex:endIndex])
	return i
}

// Boolean func returns boolean value of string value like on, off, 0, 1, yes, no
// returns boolean value of string input. You can chain this function on other function
// which returns implemented StringManipulation interface
func (i *input) Boolean() bool {
	input := getInput(*i)
	inputLower := strings.ToLower(input)
	off := contains(False, inputLower)
	if off {
		return false
	}
	on := contains(True, inputLower)
	if on {
		return true
	}
	panic(errors.New(InvalidLogicalString))
}

// CamelCase is variadic function which takes one Param rule i.e slice of strings and it returns
// input type string in camel case form and rule helps to omit character you want to omit from string.
// By default special characters like "_", "-","."," " are l\treated like word separator and treated
// accordingly by default and you dont have to worry about it
// Example input: hello user
// Result : HelloUser
func (i *input) CamelCase(rule ...string) string {
	input := getInput(*i)
	// removing excess space
	wordArray := caseHelper(input, true, rule...)
	for i, word := range wordArray {
		wordArray[i] = ucfirst(word)
	}
	return strings.Join(wordArray, "")
}

// ContainsAll is variadic function which takes slice of strings as param and checks if they are present
// in input and returns boolean value accordingly
func (i *input) ContainsAll(check ...string) bool {
	input := getInput(*i)
	for _, item := range check {
		if !strings.Contains(input, item) {
			return false
		}
	}
	return true
}

// Delimited is variadic function that takes two params delimiter and slice of strings i.e rule. It joins
// the string by passed delimeter. Rule param helps to omit character you want to omit from string. By
// default special characters like "_", "-","."," " are l\treated like word separator and treated accordingly
// by default and you dont have to worry about it.
func (i *input) Delimited(delimiter string, rule ...string) StringManipulation {
	input := getInput(*i)
	if strings.TrimSpace(delimiter) == "" {
		delimiter = "."
	}
	wordArray := caseHelper(input, false, rule...)
	i.Result = strings.Join(wordArray, delimiter)
	return i
}

// First returns first n characters from provided input. It removes all spaces in string before doing so.
func (i *input) First(length int) string {
	input := getInput(*i)
	input = strings.ReplaceAll(input, " ", "")
	if len(input) < length {
		panic(errors.New(LengthError))
	}
	return input[0:length]
}

// Get simply returns result and can be chained on function which
// returns StringManipulation interface
func (i *input) Get() string {
	return getInput(*i)
}

// KebabCase is variadic function that takes one Param slice of strings named rule and it returns passed string
// in kebab case form. Rule param helps to omit character you want to omit from string. By default special characters
// like "_", "-","."," " are l\treated like word separator and treated accordingly by default and you dont have to worry
// about it. If you don't want to omit any character pass nothing.
// Example input: hello user
// Result : hello-user
func (i *input) KebabCase(rule ...string) StringManipulation {
	input := getInput(*i)
	wordArray := caseHelper(input, false, rule...)
	i.Result = strings.Join(wordArray, "-")
	return i
}

// Last returns last n characters from provided input. It removes all spaces in string before doing so.
func (i *input) Last(length int) string {
	input := getInput(*i)
	input = strings.ReplaceAll(input, " ", "")
	inputLen := len(input)
	if len(input) < length {
		panic(errors.New(LengthError))
	}
	start := inputLen - length
	return input[start:inputLen]
}

// LcFirst simply returns result by lower casing first letter of string and it can be chained on
// function which return StringManipulation interface
func (i *input) LcFirst() string {
	input := getInput(*i)
	for _, v := range input {
		return string(unicode.ToLower(v)) + input[len(string(v)):]
	}
	return ""
}

// Lines returns slice of strings by removing white space characters
func (i *input) Lines() []string {
	input := getInput(*i)
	matchWord := regexp.MustCompile(`[\s]*[\W]\pN`)
	result := matchWord.ReplaceAllString(input, " ")
	return strings.Fields(strings.TrimSpace(result))
}

// Pad takes three param length i.e total length to be after padding, with i.e  what to pad
// with and pad type which can be ("both" or "left" or "right") it return string after padding
// upto length by with param and on padType type it can be chained on function which return
// StringManipulation interface
func (i *input) Pad(length int, with, padType string) string {
	input := getInput(*i)
	inputLength := len(input)
	padLength := len(with)
	if inputLength >= length {
		return input
	}
	switch padType {
	case Right:
		var count = 1 + ((length - padLength) / padLength)
		var result = input + strings.Repeat(with, count)
		return result[:length]
	case Left:
		var count = 1 + ((length - padLength) / padLength)
		var result = strings.Repeat(with, count) + input
		return result[(len(result) - length):]
	case Both:
		length := (float64(length - inputLength)) / float64(2)
		repeat := math.Ceil(length / float64(padLength))
		return strings.Repeat(with, int(repeat))[:int(math.Floor(float64(length)))] + input + strings.Repeat(with, int(repeat))[:int(math.Ceil(float64(length)))]
	default:
		return input
	}
}

// RemoveSpecialCharacter removes all special characters and returns the string
// it can be chained on function which return StringManipulation interface
func (i *input) RemoveSpecialCharacter() string {
	input := getInput(*i)
	var result strings.Builder
	for i := 0; i < len(input); i++ {
		b := input[i]
		if ('a' <= b && b <= 'z') ||
			('A' <= b && b <= 'Z') ||
			('0' <= b && b <= '9') ||
			b == ' ' {
			result.WriteByte(b)
		}
	}
	return result.String()
}

// ReplaceFirst takes two param search and replace. It returns string by searching search
// sub string and replacing it with replace substring on first occurrence it can be chained
// on function which return StringManipulation interface.
func (i *input) ReplaceFirst(search, replace string) string {
	input := getInput(*i)
	return replaceStr(input, search, replace, First)
}

// ReplaceLast takes two param search and replace
// it return string by searching search sub string and replacing it
// with replace substring on last occurrence
// it can be chained on function which return StringManipulation interface
func (i *input) ReplaceLast(search, replace string) string {
	input := getInput(*i)
	return replaceStr(input, search, replace, Last)
}

// Reverse reverses the passed strings
// it can be chained on function which return StringManipulation interface
func (i *input) Reverse() string {
	input := getInput(*i)
	r := []rune(input)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// Shuffle shuffles the given string randomly
// it can be chained on function which return StringManipulation interface
func (i *input) Shuffle() string {
	input := getInput(*i)
	rand.Seed(time.Now().Unix())

	inRune := []rune(input)
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})
	return string(inRune)
}

// SnakeCase is variadic function that takes one Param slice of strings named rule
// and it returns passed string in snake case form. Rule param helps to omit character
// you want to omit from string. By default special characters like "_", "-","."," " are treated
// like word separator and treated accordingly by default and you don't have to worry about it.
// If you don't want to omit any character pass nothing.
func (i *input) SnakeCase(rule ...string) StringManipulation {
	input := getInput(*i)
	wordArray := caseHelper(input, false, rule...)
	i.Result = strings.Join(wordArray, "_")
	return i
}

// Surround takes one param with which is used to surround user input and it
// can be chained on function which return StringManipulation interface.
func (i *input) Surround(with string) string {
	input := getInput(*i)
	return with + input + with
}

// Tease takes two params length and indicator and it shortens given string
// on passed length and adds indicator on end
// it can be chained on function which return StringManipulation interface
func (i *input) Tease(length int, indicator string) string {
	input := getInput(*i)
	if input == "" || len(input) < length {
		return input
	}
	return input[:length] + indicator
}

// ToLower makes all string of user input to lowercase
// it can be chained on function which return StringManipulation interface
func (i *input) ToLower() (result string) {
	input := getInput(*i)
	return strings.ToLower(input)
}

// ToUpper makes all string of user input to uppercase
// it can be chained on function which return StringManipulation interface
func (i *input) ToUpper() string {
	input := getInput(*i)
	return strings.ToUpper(input)
}

// UcFirst makes first word of user input to uppercase
// it can be chained on function which return StringManipulation interface
func (i *input) UcFirst() string {
	input := getInput(*i)
	return ucfirst(input)
}

// Prefix makes sure that string is prefixed with a given string
func (i *input) Prefix(with string) string {
	input := getInput(*i)
	if strings.HasPrefix(input, with) {
		return input
	}

	return with + input
}

// Suffix makes sure that string is suffixed with a given string
func (i *input) Suffix(with string) string {
	input := getInput(*i)
	if strings.HasSuffix(input, with) {
		return input
	}

	return input + with
}
