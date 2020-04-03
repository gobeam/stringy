package string_manipulation

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

type input struct {
	Input  string
	Result string
	Error  error
}

func caseHelper(input, key string,)(result string) {
	matchSpecial := regexp.MustCompile("[^-._A-Za-z0-9]*")
	input = matchSpecial.ReplaceAllString(input,"")
	matchWord := regexp.MustCompile("[-._]*[^A-Za-z0-9]")
	input = matchWord.ReplaceAllString(input,fmt.Sprintf("%s",key))
	for _, word := range strings.Fields(strings.TrimSpace(input)) {
		result += ucfirst(word)
	}
	return
}

func (i *input) KebabCase() StringManipulation {
	input := getInput(*i)
	i.Result = caseHelper(input, "-")
	return i
}

func ucfirst(val string) string {
	for i, v := range val {
		return string(unicode.ToUpper(v)) + val[i+1:]
	}
	return ""
}

func (i *input) UcFirst() string {
	input := getInput(*i)
	return ucfirst(input)
}

func (i *input) LcFirst() string {
	input := getInput(*i)
	for i, v := range input {
		return string(unicode.ToUpper(v)) + input[i+1:]
	}
	return input
}

func (i *input) SnakeCase() StringManipulation {
	input := getInput(*i)
	matchFirstCap := regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap := regexp.MustCompile("([a-z0-9])([A-Z])")
	snake := matchFirstCap.ReplaceAllString(input, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	i.Result = strings.Join(strings.Fields(strings.TrimSpace(snake)), "_")
	return i
}

func (i *input) CamelCase() string {
	input := getInput(*i)
	matchSpecial := regexp.MustCompile("[^-._A-Za-z0-9]*")
	input = matchSpecial.ReplaceAllString(input,"")
	matchWord := regexp.MustCompile("[-._]*[^A-Za-z0-9]")
	input = matchWord.ReplaceAllString(input," ")
	var result string
	for _, word := range strings.Fields(strings.TrimSpace(input)) {
			result += ucfirst(word)
	}
	return result

}

func (i *input) Get() string {
	return getInput(*i)
}

func (i *input) Slugify() string {
	input := getInput(*i)
	return strings.ReplaceAll(input, " ", "-")
}

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

func (i *input) ReplaceFirst(search, replace string) string {
	input := getInput(*i)
	return replaceStr(input, search, replace, First)
}

func (i *input) ReplaceLast(search, replace string) string {
	input := getInput(*i)
	return replaceStr(input, search, replace, Last)
}

func (i *input) Tease(length int, indicator string) string {
	input := getInput(*i)
	if input == "" || len(input) < length {
		return input
	}
	return input[:length] + indicator
}

type StringManipulation interface {
	Between(start, end string) StringManipulation
	Get() string
	ToUpper() string
	UcFirst() string
	LcFirst() string
	CamelCase() string
	SnakeCase() StringManipulation
	KebabCase() StringManipulation
	Slugify() string
	ReplaceFirst(search, replace string) string
	ReplaceLast(search, replace string) string
	ToLower() string
	Tease(length int, indicator string) string
}

func New(val string) StringManipulation {
	return &input{Input: val}
}

func getInput(i input) (input string) {
	if i.Result != "" {
		input = i.Result
	} else {
		input = i.Input
	}
	return
}

func (i *input) ToUpper() string {
	input := getInput(*i)
	return strings.ToUpper(input)
}

func (i *input) ToLower() (result string) {
	input := getInput(*i)
	return strings.ToLower(input)
}

func (i *input) Between(start, end string) StringManipulation {
	if start == "" && end == "" || i.Input == "" {
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
	i.Result = i.Input[startIndex:endIndex]
	return i
}
