package stringy

import (
	"errors"
	"regexp"
	"strings"
	"unicode"
)

var selectCapitalRegexp = regexp.MustCompile(SelectCapital)

/*
 * appendPadding is a helper function to append padding to the result string.
 * It takes a string builder, the padding character, the count of padding characters,
 * and the size of the padding.
 * @param result string builder
 * @param with string padding character
 * @param padCount int count of padding characters
 * @param padSize int size of the padding
 */
func appendPadding(result *strings.Builder, with string, padCount, padSize int) {
	for i := 0; i < padSize; i++ {
		result.WriteByte(with[i%len(with)])
	}
}

/*
 * caseHelper is a helper function to split the input string into words based on the provided rules.
 * It takes an input string, a boolean indicating if the input is camel case,
 * and a variadic number of rules to split the string.
 * @param input string
 * @param isCamel bool indicates if the input is camel case
 * @param rule ...string variadic number of rules to split the string
 * @return []string slice of words
 * @return error if any error occurs
 */
func caseHelper(input string, isCamel bool, rule ...string) ([]string, error) {
	if !isCamel {
		input = selectCapitalRegexp.ReplaceAllString(input, ReplaceCapital)
	}
	input = strings.Join(strings.Fields(strings.TrimSpace(input)), " ")
	if len(rule) > 0 && len(rule)%2 != 0 {
		return nil, errors.New(OddError)
	}
	rule = append(rule, ".", " ", "_", " ", "-", " ")

	replacer := strings.NewReplacer(rule...)
	input = replacer.Replace(input)

	// word splitting for multi-byte characters
	var words []string
	var currentWord strings.Builder

	for _, r := range input {
		if unicode.IsSpace(r) {
			if currentWord.Len() > 0 {
				words = append(words, currentWord.String())
				currentWord.Reset()
			}
		} else {
			currentWord.WriteRune(r)
		}
	}

	if currentWord.Len() > 0 {
		words = append(words, currentWord.String())
	}

	return words, nil
}

/**
 * getInput is a helper function to get the input string from the input struct.
 * It checks if there is an error in the input struct and returns an empty string if there is.
 * If there is no error, it checks if the Result field is not empty and returns that.
 * If the Result field is empty, it returns the Input field.
 * @param i input struct
 * @return string
 */
func getInput(i input) (input string) {
	// If there's an error, return an empty string
	if i.err != nil {
		return ""
	}

	if i.Result != "" {
		input = i.Result
	} else {
		input = i.Input
	}
	return
}

/*
 * replaceStr is a helper function to replace the first or last occurrence of a substring in a string.
 * It takes the input string, the substring to search for, the replacement string,
 * and the type of replacement (first or last).
 * @param input string
 * @param search string substring to search for
 * @param replace string replacement string
 * @param types string type of replacement (first or last)
 * @return string the modified string
 */
func replaceStr(input, search, replace, types string) string {
	lcInput := strings.ToLower(input)
	lcSearch := strings.ToLower(search)
	if input == "" || !strings.Contains(lcInput, lcSearch) {
		return input
	}
	var start int
	if types == "last" {
		start = strings.LastIndex(lcInput, lcSearch)
	} else {
		start = strings.Index(lcInput, lcSearch)
	}
	end := start + len(search)
	return input[:start] + replace + input[end:]
}
