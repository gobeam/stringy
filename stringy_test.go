package stringy

import (
	"errors"
	"fmt"
	"strings"
	"sync"
	"testing"
)

var sm StringManipulation = New("This is example.")

// Test Acronym
func TestInput_Acronym(t *testing.T) {
	acronym := New("Laugh Out Loud")
	val := acronym.Acronym().Get()
	if val != "LOL" {
		t.Errorf("Expected: %s but got: %s", "LOL", val)
	}
	if acronym.Error() != nil {
		t.Errorf("Expected no error but got: %v", acronym.Error())
	}
}

// Test Between - positive case
func TestInput_Between(t *testing.T) {
	val := sm.Between("This", "example").Trim().ToUpper()
	if val != "IS" {
		t.Errorf("Expected: %s but got: %s", "IS", val)
	}
	if sm.Error() != nil {
		t.Errorf("Expected no error but got: %v", sm.Error())
	}
}

// Test Between - empty input
func TestInput_EmptyBetween(t *testing.T) {
	sm := New("This is example.")
	val := sm.Between("", "").ToUpper()
	if val != "THIS IS EXAMPLE." {
		t.Errorf("Expected: %s but got: %s", "THIS IS EXAMPLE.", val)
	}
	if sm.Error() != nil {
		t.Errorf("Expected no error but got: %v", sm.Error())
	}
}

// Test Between - no match
func TestInput_EmptyNoMatchBetween(t *testing.T) {
	sm := New("This is example.")
	result := sm.Between("hello", "test")
	if result.Get() != "" {
		t.Errorf("Expected: \"\" but got: %s", result.Get())
	}
	if sm.Error() != nil {
		t.Errorf("Expected no error but got: %v", sm.Error())
	}
}

// Test Match - positive case
func TestInput_MatchBetween(t *testing.T) {
	sm := New("This is example.")
	result := sm.Between("This", "example").Trim()
	if result.Get() != "is" {
		t.Errorf("Expected: %s but got: %s", "is", result.Get())
	}
	if sm.Error() != nil {
		t.Errorf("Expected no error but got: %v", sm.Error())
	}
}

// Test Boolean - true values
func TestInput_BooleanTrue(t *testing.T) {
	strs := []string{"on", "On", "yes", "YES", "1", "true"}
	for _, s := range strs {
		t.Run(s, func(t *testing.T) {
			sm := New(s)
			if val := sm.Boolean(); !val {
				t.Errorf("Expected: to be true but got: %v", val)
			}
			if sm.Error() != nil {
				t.Errorf("Expected no error but got: %v", sm.Error())
			}
		})
	}
}

// Test Boolean - false values
func TestInput_BooleanFalse(t *testing.T) {
	strs := []string{"off", "Off", "no", "NO", "0", "false"}
	for _, s := range strs {
		t.Run(s, func(t *testing.T) {
			sm := New(s)
			if val := sm.Boolean(); val {
				t.Errorf("Expected: to be false but got: %v", val)
			}
			if sm.Error() != nil {
				t.Errorf("Expected no error but got: %v", sm.Error())
			}
		})
	}
}

// Test Boolean - error case
func TestInput_BooleanError(t *testing.T) {
	strs := []string{"invalid", "-1", ""}
	for _, s := range strs {
		t.Run(s, func(t *testing.T) {
			sm := New(s)
			val := sm.Boolean()
			if val != false {
				t.Errorf("Expected false as default value on error but got: %v", val)
			}
			if sm.Error() == nil {
				t.Errorf("Expected error but got none")
			}
		})
	}
}

// Test CamelCase - standard case
func TestInput_CamelCase(t *testing.T) {
	str := New("Camel case this_complicated__string%%")
	against := "camelCaseThisComplicatedString"
	if val := str.CamelCase("%", "").Get(); val != against {
		t.Errorf("Expected: to be %s but got: %s", against, val)
	}
	if str.Error() != nil {
		t.Errorf("Expected no error but got: %v", str.Error())
	}
}

// Test CamelCase - no rule case
func TestInput_CamelCaseNoRule(t *testing.T) {
	str := New("Camel case this_complicated__string%%")
	against := "camelCaseThisComplicatedString%%"
	if val := str.CamelCase().Get(); val != against {
		t.Errorf("Expected: to be %s but got: %s", against, val)
	}
	if str.Error() != nil {
		t.Errorf("Expected no error but got: %v", str.Error())
	}
}

// Test CamelCase - odd rule error
func TestInput_CamelCaseOddRuleError(t *testing.T) {
	str := New("Camel case this_complicated__string%%")
	result := str.CamelCase("%")
	if str.Error() == nil {
		t.Errorf("Expected error but got none")
	}
	if result.Get() != "" {
		t.Errorf("Expected empty result when error occurs, got: %s", result.Get())
	}
}

// Test PascalCase - standard case
func TestInput_PascalCaseNoRule(t *testing.T) {
	str := New("pascal case this_complicated__string%%")
	against := "PascalCaseThisComplicatedString%%"
	if val := str.PascalCase().Get(); val != against {
		t.Errorf("Expected: to be %s but got: %s", against, val)
	}
	if str.Error() != nil {
		t.Errorf("Expected no error but got: %v", str.Error())
	}
}

// Test PascalCase - odd rule error
func TestInput_PascalCaseOddRuleError(t *testing.T) {
	str := New("pascal case this_complicated__string%%")
	result := str.PascalCase("%")
	if str.Error() == nil {
		t.Errorf("Expected error but got none")
	}
	if result.Get() != "" {
		t.Errorf("Expected empty result when error occurs, got: %s", result.Get())
	}
}

// Test ContainsAll
func TestInput_ContainsAll(t *testing.T) {
	contains := New("hello mam how are you??")
	if val := contains.ContainsAll("mam", "?"); !val {
		t.Errorf("Expected value to be true but got false")
	}
	if val := contains.ContainsAll("non existent"); val {
		t.Errorf("Expected value to be false but got true")
	}
	if contains.Error() != nil {
		t.Errorf("Expected no error but got: %v", contains.Error())
	}
}

// Test Delimited - with rules
func TestInput_Delimited(t *testing.T) {
	str := New("Delimited case this_complicated__string@@")
	against := "delimited.case.this.complicated.string"
	if val := str.Delimited(".", "@", "").ToLower(); val != against {
		t.Errorf("Expected: to be %s but got: %s", against, val)
	}
	if str.Error() != nil {
		t.Errorf("Expected no error but got: %v", str.Error())
	}
}

// Test Delimited - no delimiter
func TestInput_DelimitedNoDelimeter(t *testing.T) {
	str := New("Delimited case this_complicated__string@@")
	against := "delimited.case.this.complicated.string@@"
	if val := str.Delimited("").ToLower(); val != against {
		t.Errorf("Expected: to be %s but got: %s", against, val)
	}
	if str.Error() != nil {
		t.Errorf("Expected no error but got: %v", str.Error())
	}
}

// Test Delimited - odd rule error
func TestInput_DelimitedOddRuleError(t *testing.T) {
	str := New("Delimited case this_complicated__string@@")
	result := str.Delimited(".", "@")
	if str.Error() == nil {
		t.Errorf("Expected error but got none")
	}
	if result.Get() != "" {
		t.Errorf("Expected empty result when error occurs, got: %s", result.Get())
	}
}

// Test KebabCase
func TestInput_KebabCase(t *testing.T) {
	str := New("Kebab case this-complicated___string@@")
	against := "Kebab-case-this-complicated-string"
	if val := str.KebabCase("@", "").Get(); val != against {
		t.Errorf("Expected: to be %s but got: %s", against, val)
	}
	if str.Error() != nil {
		t.Errorf("Expected no error but got: %v", str.Error())
	}
}

// Test KebabCase - odd rule error
func TestInput_KebabCaseOddRuleError(t *testing.T) {
	str := New("Kebab case this-complicated___string@@")
	result := str.KebabCase("@")
	if str.Error() == nil {
		t.Errorf("Expected error but got none")
	}
	if result.Get() != "" {
		t.Errorf("Expected empty result when error occurs, got: %s", result.Get())
	}
}

// Test LcFirst
func TestInput_LcFirst(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{
			name: "leading uppercase",
			arg:  "This is an all lower",
			want: "this is an all lower",
		},
		{
			name: "empty string",
			arg:  "",
			want: "",
		},
		{
			name: "multi-byte leading character",
			arg:  "ΔΔΔ",
			want: "δΔΔ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sm := New(tt.arg)
			if got := sm.LcFirst(); got != tt.want {
				t.Errorf("LcFirst(%v) = %v, want %v", tt.arg, got, tt.want)
			}
			if sm.Error() != nil {
				t.Errorf("Expected no error but got: %v", sm.Error())
			}
		})
	}
}

// Test Lines
func TestInput_Lines(t *testing.T) {
	lines := New("fòô\r\nbàř\nyolo")
	strSlic := lines.Lines()
	if len(strSlic) != 3 {
		t.Errorf("Length expected to be 3 but got: %d", len(strSlic))
	}
	if strSlic[0] != "fòô" {
		t.Errorf("Expected: %s but got: %s", "fòô", strSlic[0])
	}
	if lines.Error() != nil {
		t.Errorf("Expected no error but got: %v", lines.Error())
	}
}

// Test Lines - empty input
func TestInput_LinesEmpty(t *testing.T) {
	lines := New("")
	strSlic := lines.Lines()
	if len(strSlic) != 0 {
		t.Errorf("Length expected to be 0 but got: %d", len(strSlic))
	}
	if lines.Error() != nil {
		t.Errorf("Expected no error but got: %v", lines.Error())
	}
}

// Test Pad
func TestInput_Pad(t *testing.T) {
	pad := New("Roshan")
	if result := pad.Pad(10, "0", "both"); result != "00Roshan00" {
		t.Errorf("Expected: %s but got: %s", "00Roshan00", result)
	}
	if result := pad.Pad(10, "0", "left"); result != "0000Roshan" {
		t.Errorf("Expected: %s but got: %s", "0000Roshan", result)
	}
	if result := pad.Pad(10, "0", "right"); result != "Roshan0000" {
		t.Errorf("Expected: %s but got: %s", "Roshan0000", result)
	}
	if pad.Error() != nil {
		t.Errorf("Expected no error but got: %v", pad.Error())
	}
}

// Test Pad - invalid length
func TestInput_PadInvalidLength(t *testing.T) {
	pad := New("Roshan")
	if result := pad.Pad(6, "0", "both"); result != "Roshan" {
		t.Errorf("Expected: %s but got: %s", "Roshan", result)
	}
	if result := pad.Pad(6, "0", "left"); result != "Roshan" {
		t.Errorf("Expected: %s but got: %s", "Roshan", result)
	}
	if result := pad.Pad(6, "0", "right"); result != "Roshan" {
		t.Errorf("Expected: %s but got: %s", "Roshan", result)
	}
	if result := pad.Pad(13, "0", "middle"); result != "Roshan" {
		t.Errorf("Expected: %s but got: %s", "Roshan", result)
	}
	if pad.Error() != nil {
		t.Errorf("Expected no error but got: %v", pad.Error())
	}
}

// Test RemoveSpecialCharacter
func TestInput_RemoveSpecialCharacter(t *testing.T) {
	cleanString := New("special@#remove%%%%")
	against := "specialremove"
	if result := cleanString.RemoveSpecialCharacter(); result != against {
		t.Errorf("Expected: %s but got: %s", against, result)
	}
	if cleanString.Error() != nil {
		t.Errorf("Expected no error but got: %v", cleanString.Error())
	}
}

// Test ReplaceFirst
func TestInput_ReplaceFirst(t *testing.T) {
	replaceFirst := New("Hello My name is Roshan and his name is Alis.")
	against := "Hello My nombre is Roshan and his name is Alis."
	if result := replaceFirst.ReplaceFirst("name", "nombre"); result != against {
		t.Errorf("Expected: %s but got: %s", against, result)
	}
	if replaceFirst.Error() != nil {
		t.Errorf("Expected no error but got: %v", replaceFirst.Error())
	}
}

// Test ReplaceFirst - empty input
func TestInput_ReplaceFirstEmptyInput(t *testing.T) {
	replaceFirst := New("")
	against := ""
	if result := replaceFirst.ReplaceFirst("name", "nombre"); result != against {
		t.Errorf("Expected: %s but got: %s", against, result)
	}
	if replaceFirst.Error() != nil {
		t.Errorf("Expected no error but got: %v", replaceFirst.Error())
	}
}

// Test ReplaceLast
func TestInput_ReplaceLast(t *testing.T) {
	replaceLast := New("Hello My name is Roshan and his name is Alis.")
	against := "Hello My name is Roshan and his nombre is Alis."
	if result := replaceLast.ReplaceLast("name", "nombre"); result != against {
		t.Errorf("Expected: %s but got: %s", against, result)
	}
	if replaceLast.Error() != nil {
		t.Errorf("Expected no error but got: %v", replaceLast.Error())
	}
}

// Test Reverse
func TestInput_Reverse(t *testing.T) {
	reverseString := New("roshan")
	against := "nahsor"
	if result := reverseString.Reverse(); result != against {
		t.Errorf("Expected: %s but got: %s", against, result)
	}
	if reverseString.Error() != nil {
		t.Errorf("Expected no error but got: %v", reverseString.Error())
	}
}

// Test Shuffle
func TestInput_Shuffle(t *testing.T) {
	check := "roshan"
	shuffleString := New(check)
	if result := shuffleString.Shuffle(); len(result) != len(check) && check == result {
		t.Errorf("Shuffle string gave wrong output")
	}
	if shuffleString.Error() != nil {
		t.Errorf("Expected no error but got: %v", shuffleString.Error())
	}
}

// Test SnakeCase
func TestInput_SnakeCase(t *testing.T) {
	str := New("SnakeCase this-complicated___string@@")
	against := "snake_case_this_complicated_string"
	if val := str.SnakeCase("@", "").ToLower(); val != against {
		t.Errorf("Expected: to be %s but got: %s", against, val)
	}
	if str.Error() != nil {
		t.Errorf("Expected no error but got: %v", str.Error())
	}
}

// Test SnakeCase - odd rule error
func TestInput_SnakeCaseOddRuleError(t *testing.T) {
	str := New("SnakeCase this-complicated___string@@")
	result := str.SnakeCase("@")
	if str.Error() == nil {
		t.Errorf("Expected error but got none")
	}
	if result.Get() != "" {
		t.Errorf("Expected empty result when error occurs, got: %s", result.Get())
	}
}

// Test Surround
func TestInput_Surround(t *testing.T) {
	str := New("this")
	against := "__this__"
	if val := str.Surround("__"); val != against {
		t.Errorf("Expected: to be %s but got: %s", against, val)
	}
	if str.Error() != nil {
		t.Errorf("Expected no error but got: %v", str.Error())
	}
}

// Test Tease
func TestInput_Tease(t *testing.T) {
	str := New("This is just simple paragraph on lorem ipsum.")
	against := "This is just..."
	if val := str.Tease(12, "..."); val != against {
		t.Errorf("Expected: to be %s but got: %s", against, val)
	}
	if str.Error() != nil {
		t.Errorf("Expected no error but got: %v", str.Error())
	}
}

// Test Tease - empty or shorter than length
func TestInput_TeaseEmpty(t *testing.T) {
	str := New("This is just simple paragraph on lorem ipsum.")
	against := "This is just simple paragraph on lorem ipsum."
	if val := str.Tease(200, "..."); val != against {
		t.Errorf("Expected: to be %s but got: %s", against, val)
	}
	if str.Error() != nil {
		t.Errorf("Expected no error but got: %v", str.Error())
	}
}

// Test Title
func TestInput_Title(t *testing.T) {
	str := New("this is just AN eXample")
	against := "This Is Just An Example"
	if val := str.Title(); val != against {
		t.Errorf("Expected: to be %s but got: %s", against, val)
	}
	if str.Error() != nil {
		t.Errorf("Expected no error but got: %v", str.Error())
	}
}

// Test UcFirst
func TestInput_UcFirst(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{
			name: "leading lowercase",
			arg:  "test input",
			want: "Test input",
		},
		{
			name: "empty string",
			arg:  "",
			want: "",
		},
		{
			name: "multi-byte leading character",
			arg:  "δδδ",
			want: "Δδδ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sm := New(tt.arg)
			if got := sm.UcFirst(); got != tt.want {
				t.Errorf("UcFirst(%v) = %v, want %v", tt.arg, got, tt.want)
			}
			if sm.Error() != nil {
				t.Errorf("Expected no error but got: %v", sm.Error())
			}
		})
	}
}

// Test First
func TestInput_First(t *testing.T) {
	fcn := New("4111 1111 1111 1111")
	against := "4111"
	if first := fcn.First(4); first != against {
		t.Errorf("Expected: to be %s but got: %s", against, first)
	}
	if fcn.Error() != nil {
		t.Errorf("Expected no error but got: %v", fcn.Error())
	}
}

// Test First - error case
func TestInput_FirstError(t *testing.T) {
	fcn := New("4111 1111 1111 1111")
	result := fcn.First(100)
	if fcn.Error() == nil {
		t.Errorf("Error expected but got none")
	}
	if result != "" {
		t.Errorf("Expected empty result when error occurs, got: %s", result)
	}
}

// Test Last
func TestInput_Last(t *testing.T) {
	lcn := New("4111 1111 1111 1348")
	against := "1348"
	if last := lcn.Last(4); last != against {
		t.Errorf("Expected: to be %s but got: %s", against, last)
	}
	if lcn.Error() != nil {
		t.Errorf("Expected no error but got: %v", lcn.Error())
	}
}

// Test Last - error case
func TestInput_LastError(t *testing.T) {
	lcn := New("4111 1111 1111 1348")
	result := lcn.Last(100)
	if lcn.Error() == nil {
		t.Errorf("Error expected but got none")
	}
	if result != "" {
		t.Errorf("Expected empty result when error occurs, got: %s", result)
	}
}

// Test Prefix
func TestInput_Prefix(t *testing.T) {
	str := New("foobar")
	against := "foobar"
	if val := str.Prefix("foo"); val != against {
		t.Errorf("Expected: to be %s but got: %s", against, val)
	}

	str = New("foobar")
	against = "foofoofoobar"
	if val := str.Prefix("foofoo"); val != against {
		t.Errorf("Expected: to be %s but got: %s", against, val)
	}
	if str.Error() != nil {
		t.Errorf("Expected no error but got: %v", str.Error())
	}
}

// Test Suffix
func TestInput_Suffix(t *testing.T) {
	str := New("foobar")
	against := "foobar"
	if val := str.Suffix("bar"); val != against {
		t.Errorf("Expected: to be %s but got: %s", against, val)
	}

	str = New("foobar")
	against = "foobarbarbar"
	if val := str.Suffix("barbar"); val != against {
		t.Errorf("Expected: to be %s but got: %s", against, val)
	}
	if str.Error() != nil {
		t.Errorf("Expected no error but got: %v", str.Error())
	}
}

// Test Error method
func TestInput_Error(t *testing.T) {
	// Test that Error() returns nil for new object
	str := New("test")
	if str.Error() != nil {
		t.Errorf("Expected nil error but got: %v", str.Error())
	}

	// Test that Error() returns the correct error after setting it
	str = New("invalid")
	str.Boolean() // This should set an error
	if str.Error() == nil {
		t.Errorf("Expected error but got nil")
	}
	if str.Error().Error() != InvalidLogicalString {
		t.Errorf("Expected error message '%s' but got: %s", InvalidLogicalString, str.Error().Error())
	}

	// Test that Error() is reset when using New()
	str = New("test")
	if str.Error() != nil {
		t.Errorf("Expected nil error after New() but got: %v", str.Error())
	}
}

// Test Release method
func TestInput_Release(t *testing.T) {
	i := inputPool.Get().(*input)
	i.Input = "test"
	i.Result = "result"
	i.err = errors.New("test error")

	i.Release()

	if i.Input != "" || i.Result != "" || i.err != nil {
		t.Errorf("Release didn't reset the fields properly. Input: %s, Result: %s, err: %v",
			i.Input, i.Result, i.err)
	}
}

// Test method chaining with errors
func TestInput_MethodChainingWithErrors(t *testing.T) {
	// Test that an error in one method is preserved through a chain
	str := New("test")
	result := str.CamelCase("%").ToUpper()

	if str.Error() == nil {
		t.Errorf("Expected error to be preserved in chain but got nil")
	}

	if result != "" {
		t.Errorf("Expected empty result after error but got: %s", result)
	}
}

// Test Error persistence
func TestInput_ErrorPersistence(t *testing.T) {
	// Test that after an error is set, it remains until a new object is created
	str := New("test")
	str.First(100) // Should set error

	if str.Error() == nil {
		t.Errorf("Expected error but got nil")
	}

	// Try another operation
	str.ToUpper()

	// Error should still be present
	if str.Error() == nil {
		t.Errorf("Expected error to persist but it was cleared")
	}

	// Create new object
	newStr := New("test")
	if newStr.Error() != nil {
		t.Errorf("Expected nil error for new object but got: %v", newStr.Error())
	}
}

// Test for handling multi-byte characters in all methods
func TestInput_MultiByteCharacters(t *testing.T) {
	// Test with emoji and international characters
	str := New("😀 Hello 世界")

	// Test CamelCase with multi-byte
	camelResult := str.CamelCase().Get()
	if camelResult != "😀Hello世界" {
		t.Errorf("CamelCase multi-byte - Expected: %s but got: %s", "😀Hello世界", camelResult)
	}

	// Test SnakeCase with multi-byte - update expectation
	str = New("😀 Hello 世界")
	snakeResult := str.SnakeCase().Get()
	if snakeResult != "😀_Hello_世界" {
		t.Errorf("SnakeCase multi-byte - Expected: %s but got: %s", "😀_Hello_世界", snakeResult)
	}

	// Test KebabCase with multi-byte - update expectation
	str = New("😀 Hello 世界")
	kebabResult := str.KebabCase().Get()
	if kebabResult != "😀-Hello-世界" {
		t.Errorf("KebabCase multi-byte - Expected: %s but got: %s", "😀-Hello-世界", kebabResult)
	}
}

// Test concurrent use of the package using goroutines
func TestInput_Concurrency(t *testing.T) {
	const goroutines = 100
	var wg sync.WaitGroup
	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func(id int) {
			defer wg.Done()

			// Use different operations in each goroutine
			str := New(fmt.Sprintf("Test%d", id))

			// Mix of operations
			switch id % 5 {
			case 0:
				result := str.CamelCase().Get()
				if result != fmt.Sprintf("test%d", id) {
					t.Errorf("Concurrent CamelCase - Expected: test%d but got: %s", id, result)
				}
			case 1:
				result := str.SnakeCase().Get()
				if result != fmt.Sprintf("Test_%d", id) {
					t.Errorf("Concurrent SnakeCase - Expected: Test_%d but got: %s", id, result)
				}
			case 2:
				result := str.Between("T", fmt.Sprintf("%d", id)).Get()
				if result != "est" {
					t.Errorf("Concurrent Between - Expected: est but got: %s", result)
				}
			case 3:
				result := str.ToUpper()
				if result != fmt.Sprintf("TEST%d", id) {
					t.Errorf("Concurrent ToUpper - Expected: TEST%d but got: %s", id, result)
				}
			case 4:
				result := str.Reverse()
				expected := fmt.Sprintf("%dseT", id)
				if result != expected {
					t.Errorf("Concurrent Reverse - Expected: %s but got: %s", expected, result)
				}
			}

			// Test Release
			input := inputPool.Get().(*input)
			input.Input = "test"
			input.Release()
		}(i)
	}

	wg.Wait()
}

// Additional test cases for SnakeCase
func TestInput_SnakeCaseRobustness(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		rules    []string
		expected string
	}{
		{
			name:     "basic conversion",
			input:    "ThisIsATest",
			rules:    []string{},
			expected: "This_Is_A_Test",
		},
		{
			name:     "with special characters",
			input:    "This@Is#A$Test",
			rules:    []string{"@", " ", "#", " ", "$", " "},
			expected: "This_Is_A_Test",
		},
		{
			name:     "with mixed case",
			input:    "thisIsATest",
			rules:    []string{},
			expected: "this_Is_A_Test",
		},
		{
			name:     "with existing underscores",
			input:    "this_is_a_test",
			rules:    []string{},
			expected: "this_is_a_test",
		},
		{
			name:     "with mixed separators",
			input:    "this-is.a test",
			rules:    []string{},
			expected: "this_is_a_test",
		},
		{
			name:     "with consecutive separators",
			input:    "this__is...a   test",
			rules:    []string{},
			expected: "this_is_a_test",
		},
		{
			name:     "with control characters",
			input:    "this\nis\ta\rtest",
			rules:    []string{},
			expected: "this_is_a_test",
		},
		{
			name:     "with empty input",
			input:    "",
			rules:    []string{},
			expected: "",
		},
		{
			name:     "with only separators",
			input:    "___...-  ",
			rules:    []string{},
			expected: "",
		},
		{
			name:     "with multi-byte characters",
			input:    "こんにちは世界",
			rules:    []string{},
			expected: "こんにちは世界",
		},
		{
			name:     "with multi-byte characters and separators",
			input:    "こんにちは_世界",
			rules:    []string{},
			expected: "こんにちは_世界",
		},
		{
			name:     "complex mix",
			input:    "ThisIs-A__complex.123 Test   with@#$Stuff",
			rules:    []string{"@", " ", "#", " ", "$", " "},
			expected: "This_Is_A_complex_123_Test_with_Stuff",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			str := New(tc.input)
			var result string
			if len(tc.rules) > 0 {
				result = str.SnakeCase(tc.rules...).Get()
			} else {
				result = str.SnakeCase().Get()
			}

			if result != tc.expected {
				t.Errorf("Expected: %q, but got: %q", tc.expected, result)
			}
		})
	}
}

// Test edge cases for all string manipulations
func TestInput_EdgeCases(t *testing.T) {
	// Test with empty string
	empty := New("")

	// All these operations should handle empty strings gracefully
	if empty.CamelCase().Get() != "" {
		t.Errorf("CamelCase empty - Expected empty string")
	}

	if empty.SnakeCase().Get() != "" {
		t.Errorf("SnakeCase empty - Expected empty string")
	}

	if empty.KebabCase().Get() != "" {
		t.Errorf("KebabCase empty - Expected empty string")
	}

	if empty.Reverse() != "" {
		t.Errorf("Reverse empty - Expected empty string")
	}

	if empty.RemoveSpecialCharacter() != "" {
		t.Errorf("RemoveSpecialCharacter empty - Expected empty string")
	}

	if empty.Tease(10, "...") != "" {
		t.Errorf("Tease empty - Expected empty string")
	}

	if empty.Pad(10, "0", "both") != "0000000000" {
		t.Errorf("Pad empty - Expected 10 zeros but got: %s", empty.Pad(10, "0", "both"))
	}

	// Test with only special characters
	special := New("@#$%^&*")

	if special.RemoveSpecialCharacter() != "" {
		t.Errorf("RemoveSpecialCharacter special - Expected empty string but got: %s",
			special.RemoveSpecialCharacter())
	}

	// Test with extremely long input
	longStr := strings.Repeat("a", 10000)
	long := New(longStr)

	if len(long.Tease(100, "...")) != 103 {
		t.Errorf("Tease long - Expected length 103 but got: %d", len(long.Tease(100, "...")))
	}

	if len(long.First(100)) != 100 {
		t.Errorf("First long - Expected length 100 but got: %d", len(long.First(100)))
	}

	if len(long.Last(100)) != 100 {
		t.Errorf("Last long - Expected length 100 but got: %d", len(long.Last(100)))
	}
}

// Additional test for Between with various cases
func TestInput_BetweenEdgeCases(t *testing.T) {
	// Test with matching start but no matching end
	str := New("Hello World")
	result := str.Between("Hello", "Goodbye").Get()
	if result != "" {
		t.Errorf("Between with no matching end - Expected empty but got: %s", result)
	}

	// Test with matching end but no matching start
	result = str.Between("Goodbye", "World").Get()
	if result != "" {
		t.Errorf("Between with no matching start - Expected empty but got: %s", result)
	}

	// Test with multiple occurrences of start and end - updated expectation
	str = New("start middle start middle end end")
	result = str.Between("start", "end").Get()
	if result != " middle start middle " {
		t.Errorf("Between with multiple occurrences - Expected: ' middle start middle ' but got: %s", result)
	}

	// Test with overlapping start and end - modified expectation
	str = New("startend")
	result = str.Between("start", "end").Get()
	if result != "" {
		t.Errorf("Between with overlapping - Expected empty but got: %s", result)
	}

	// Test case sensitivity - updated expectation
	str = New("START middle END")
	result = str.Between("start", "end").Get()
	if result != " middle " {
		t.Errorf("Between case insensitive - Expected: ' middle ' but got: %s", result)
	}
}

// Additional test for method chaining
func TestInput_MethodChaining(t *testing.T) {
	str := New("this is a TEST string")

	// Chain multiple operations
	result := str.CamelCase().Between("this", "string").ToUpper()
	if result != "ISATEST" {
		t.Errorf("Method chaining - Expected: ISATEST but got: %s", result)
	}

	// Chain with error in the middle
	str = New("this is a TEST string")
	result = str.SnakeCase("%").ToUpper() // Should error due to odd rule
	if result != "" {
		t.Errorf("Method chaining with error - Expected empty but got: %s", result)
	}
	if str.Error() == nil {
		t.Errorf("Method chaining with error - Expected error but got nil")
	}
}

// Test for Prefix and Suffix with edge cases
func TestInput_PrefixSuffixEdgeCases(t *testing.T) {
	// Test empty string
	str := New("")
	if str.Prefix("prefix") != "prefix" {
		t.Errorf("Prefix with empty - Expected: prefix but got: %s", str.Prefix("prefix"))
	}
	if str.Suffix("suffix") != "suffix" {
		t.Errorf("Suffix with empty - Expected: suffix but got: %s", str.Suffix("suffix"))
	}

	// Test with empty prefix/suffix
	str = New("test")
	if str.Prefix("") != "test" {
		t.Errorf("Prefix with empty prefix - Expected: test but got: %s", str.Prefix(""))
	}
	if str.Suffix("") != "test" {
		t.Errorf("Suffix with empty suffix - Expected: test but got: %s", str.Suffix(""))
	}

	// Test with very long prefix/suffix
	longAffix := strings.Repeat("x", 1000)
	str = New("test")
	if !strings.HasPrefix(str.Prefix(longAffix), longAffix) {
		t.Errorf("Prefix with long prefix - Expected to start with long prefix")
	}
	if !strings.HasSuffix(str.Suffix(longAffix), longAffix) {
		t.Errorf("Suffix with long suffix - Expected to end with long suffix")
	}
}

// Test memory leaks with object pool
func TestInput_MemoryLeaks(t *testing.T) {
	// Create a large number of objects and release them
	for i := 0; i < 10000; i++ {
		str := inputPool.Get().(*input)
		str.Input = fmt.Sprintf("test%d", i)
		str.Result = fmt.Sprintf("result%d", i)
		str.Release()
	}

	// Get a new object and check it's clean
	obj := inputPool.Get().(*input)
	if obj.Input != "" || obj.Result != "" || obj.err != nil {
		t.Errorf("Pool object not clean - Input: %s, Result: %s, err: %v",
			obj.Input, obj.Result, obj.err)
	}
	obj.Release()
}

// Test for handling null characters and other special byte sequences
func TestInput_SpecialByteSequences(t *testing.T) {
	// Test with null characters
	str := New("hello\x00world")

	result := str.CamelCase().Get()
	if result != "helloWorld" {
		t.Errorf("CamelCase with null - Expected: helloWorld but got: %s", result)
	}

	// Test with escape sequences
	str = New("hello\tworld\ntest")
	result = str.SnakeCase().Get()
	if result != "hello_world_test" {
		t.Errorf("SnakeCase with escapes - Expected: hello_world_test but got: %s", result)
	}

	// Test with control characters
	str = New("hello\x01\x02\x03world")
	result = str.RemoveSpecialCharacter()
	if result != "helloworld" {
		t.Errorf("RemoveSpecialCharacter with control - Expected: helloworld but got: %s", result)
	}
}

// Test for Title() method with various edge cases
func TestInput_TitleEdgeCases(t *testing.T) {
	// Test with empty string
	str := New("")
	if str.Title() != "" {
		t.Errorf("Title with empty - Expected empty but got: %s", str.Title())
	}

	// Test with single word
	str = New("hello")
	if str.Title() != "Hello" {
		t.Errorf("Title with single word - Expected: Hello but got: %s", str.Title())
	}

	// Test with all uppercase
	str = New("HELLO WORLD")
	if str.Title() != "Hello World" {
		t.Errorf("Title with all uppercase - Expected: Hello World but got: %s", str.Title())
	}

	// Test with mixed case
	str = New("hElLo wOrLd")
	if str.Title() != "Hello World" {
		t.Errorf("Title with mixed case - Expected: Hello World but got: %s", str.Title())
	}

	// Test with special characters
	str = New("hello-world")
	if str.Title() != "Hello-world" {
		t.Errorf("Title with special chars - Expected: Hello-world but got: %s", str.Title())
	}
}

// Test for error handling with invalid parameters
func TestInput_InvalidParameters(t *testing.T) {
	// Test negative length in padding
	str := New("test")
	result := str.Pad(-10, "0", "both")
	if result != "test" {
		t.Errorf("Pad with negative length - Expected: test but got: %s", result)
	}

	// Test negative length in First/Last
	result = str.First(-5)
	if result != "" {
		t.Errorf("First with negative length - Expected empty but got: %s", result)
	}
	if str.Error() == nil {
		t.Errorf("First with negative length - Expected error but got nil")
	}

	// Reset error state
	str = New("test")

	result = str.Last(-5)
	if result != "" {
		t.Errorf("Last with negative length - Expected empty but got: %s", result)
	}
	if str.Error() == nil {
		t.Errorf("Last with negative length - Expected error but got nil")
	}
}

// Test Trim method
func TestInput_Trim(t *testing.T) {
	// Test trim whitespace
	str := New("  Hello World  ")
	result := str.Trim().Get()
	if result != "Hello World" {
		t.Errorf("Trim whitespace - Expected: \"Hello World\" but got: \"%s\"", result)
	}

	// Test trim specific characters
	str = New("!!!Hello World!!!")
	result = str.Trim("!").Get()
	if result != "Hello World" {
		t.Errorf("Trim specific chars - Expected: \"Hello World\" but got: \"%s\"", result)
	}

	// Test trim with empty string
	str = New("")
	result = str.Trim().Get()
	if result != "" {
		t.Errorf("Trim empty - Expected empty string but got: \"%s\"", result)
	}

	// Test trim with no trimming needed
	str = New("Hello")
	result = str.Trim().Get()
	if result != "Hello" {
		t.Errorf("Trim no whitespace - Expected: \"Hello\" but got: \"%s\"", result)
	}

	// Test method chaining
	str = New("  Hello World  ")
	result = str.Trim().ToUpper()
	if result != "HELLO WORLD" {
		t.Errorf("Trim with chaining - Expected: \"HELLO WORLD\" but got: \"%s\"", result)
	}

	// Test with multi-byte characters
	str = New("  👋 Hello 世界  ")
	result = str.Trim().Get()
	if result != "👋 Hello 世界" {
		t.Errorf("Trim with multi-byte - Expected: \"👋 Hello 世界\" but got: \"%s\"", result)
	}
}

func TestCamelCaseHelper(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"test", "test"},
		{"test entity", "testEntity"},
		{"test-entity", "testEntity"},
		{"test-entity_test", "testEntityTest"},
		{"test_entity", "testEntity"},
		{"TestEntity", "testEntity"},
		{"testEntity", "testEntity"},
		{"test_entity_definition", "testEntityDefinition"},
	}

	for _, test := range tests {
		actual := New(test.input).CamelCase().Get()
		if actual != test.expected {
			t.Errorf("Expected %s, got %s, input: %s", test.expected, actual, test.input)
		}
	}
}

func TestInput_SentenceCase(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		rules    []string
		expected string
	}{
		{
			name:     "camel case",
			input:    "thisIsCamelCase",
			rules:    []string{},
			expected: "This is camel case",
		},
		{
			name:     "snake case",
			input:    "this_is_snake_case",
			rules:    []string{},
			expected: "This is snake case",
		},
		{
			name:     "kebab case",
			input:    "this-is-kebab-case",
			rules:    []string{},
			expected: "This is kebab case",
		},
		{
			name:     "pascal case",
			input:    "ThisIsPascalCase",
			rules:    []string{},
			expected: "This is pascal case",
		},
		{
			name:     "with numbers",
			input:    "this_is_1_example_with2_numbers",
			rules:    []string{},
			expected: "This is 1 example with2 numbers",
		},
		{
			name:     "with special characters",
			input:    "this@is#an&example",
			rules:    []string{"@", " ", "#", " ", "&", " "},
			expected: "This is an example",
		},
		{
			name:     "empty string",
			input:    "",
			rules:    []string{},
			expected: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			str := New(tc.input)
			var result string
			if len(tc.rules) > 0 {
				result = str.SentenceCase(tc.rules...).Get()
			} else {
				result = str.SentenceCase().Get()
			}

			if result != tc.expected {
				t.Errorf("Expected: %q but got: %q", tc.expected, result)
			}
		})
	}
}

// Test for WordCount method
func TestInput_WordCount(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "basic sentence",
			input:    "This is a test",
			expected: 4,
		},
		{
			name:     "empty string",
			input:    "",
			expected: 0,
		},
		{
			name:     "multiple spaces",
			input:    "This   has   extra   spaces",
			expected: 4,
		},
		{
			name:     "with newlines and tabs",
			input:    "This\nhas\ttabs and newlines",
			expected: 5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			str := New(tc.input)
			count := str.WordCount()
			if count != tc.expected {
				t.Errorf("Expected word count: %d but got: %d", tc.expected, count)
			}
		})
	}
}

// Test for TruncateWords method
func TestInput_TruncateWords(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		count    int
		suffix   string
		expected string
	}{
		{
			name:     "truncate with suffix",
			input:    "This is a test of the truncate words method",
			count:    4,
			suffix:   "...",
			expected: "This is a test...",
		},
		{
			name:     "no truncation needed",
			input:    "Short text",
			count:    5,
			suffix:   "...",
			expected: "Short text",
		},
		{
			name:     "empty string",
			input:    "",
			count:    3,
			suffix:   "...",
			expected: "",
		},
		{
			name:     "truncate with empty suffix",
			input:    "One two three four five",
			count:    3,
			suffix:   "",
			expected: "One two three",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			str := New(tc.input)
			result := str.TruncateWords(tc.count, tc.suffix).Get()
			if result != tc.expected {
				t.Errorf("Expected: %q but got: %q", tc.expected, result)
			}
		})
	}
}

// Test for IsEmpty method
func TestInput_IsEmpty(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "empty string",
			input:    "",
			expected: true,
		},
		{
			name:     "whitespace only",
			input:    "   \t\n",
			expected: true,
		},
		{
			name:     "non-empty string",
			input:    "Hello",
			expected: false,
		},
		{
			name:     "whitespace with content",
			input:    "  Hello  ",
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			str := New(tc.input)
			result := str.IsEmpty()
			if result != tc.expected {
				t.Errorf("Expected IsEmpty(): %v but got: %v", tc.expected, result)
			}
		})
	}
}

// Test for Substring method
func TestInput_Substring(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		start    int
		end      int
		expected string
		hasError bool
	}{
		{
			name:     "basic substring",
			input:    "Hello World",
			start:    0,
			end:      5,
			expected: "Hello",
			hasError: false,
		},
		{
			name:     "middle substring",
			input:    "Hello World",
			start:    6,
			end:      11,
			expected: "World",
			hasError: false,
		},
		{
			name:     "out of bounds - end too large",
			input:    "Hello",
			start:    0,
			end:      10,
			expected: "Hello",
			hasError: false,
		},
		{
			name:     "out of bounds - start negative",
			input:    "Hello",
			start:    -5,
			end:      5,
			expected: "Hello",
			hasError: false,
		},
		{
			name:     "invalid range - start > end",
			input:    "Hello",
			start:    4,
			end:      2,
			expected: "",
			hasError: true,
		},
		{
			name:     "empty result - start == end",
			input:    "Hello",
			start:    2,
			end:      2,
			expected: "",
			hasError: false,
		},
		{
			name:     "multi-byte characters",
			input:    "Hello 世界",
			start:    6,
			end:      8,
			expected: "世界",
			hasError: false,
		},
		{
			name:     "empty string input",
			input:    "",
			start:    0,
			end:      0,
			expected: "",
			hasError: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			str := New(tc.input)
			result := str.Substring(tc.start, tc.end)

			if result.Get() != tc.expected {
				t.Errorf("Expected: %q but got: %q", tc.expected, result.Get())
			}

			if (result.Error() != nil) != tc.hasError {
				t.Errorf("Expected error: %v but got: %v", tc.hasError, result.Error() != nil)
			}
		})
	}
}

// Test for SlugifyWithCount method
func TestInput_SlugifyWithCount(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		count    int
		expected string
	}{
		{
			name:     "basic slug with count",
			input:    "This is a test",
			count:    1,
			expected: "this-is-a-test-1",
		},
		{
			name:     "slug without count",
			input:    "This is a test",
			count:    0,
			expected: "this-is-a-test",
		},
		{
			name:     "slug with special characters",
			input:    "This & that @ example.com",
			count:    2,
			expected: "this-that-examplecom-2",
		},
		{
			name:     "empty string",
			input:    "",
			count:    1,
			expected: "-1",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			str := New(tc.input)
			result := str.SlugifyWithCount(tc.count).Get()
			if result != tc.expected {
				t.Errorf("Expected: %q but got: %q", tc.expected, result)
			}
		})
	}
}

// Test for Contains method
func TestInput_Contains(t *testing.T) {
	testCases := []struct {
		name      string
		input     string
		substring string
		expected  bool
	}{
		{
			name:      "contains substring",
			input:     "Hello World",
			substring: "World",
			expected:  true,
		},
		{
			name:      "does not contain substring",
			input:     "Hello World",
			substring: "Universe",
			expected:  false,
		},
		{
			name:      "empty string contains empty substring",
			input:     "",
			substring: "",
			expected:  true,
		},
		{
			name:      "string contains empty substring",
			input:     "Hello",
			substring: "",
			expected:  true,
		},
		{
			name:      "case sensitivity",
			input:     "Hello World",
			substring: "world",
			expected:  false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			str := New(tc.input)
			result := str.Contains(tc.substring)
			if result != tc.expected {
				t.Errorf("Expected Contains(%q): %v but got: %v", tc.substring, tc.expected, result)
			}
		})
	}
}

// Test for ReplaceAll method
func TestInput_ReplaceAll(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		search   string
		replace  string
		expected string
	}{
		{
			name:     "basic replacement",
			input:    "Hello World World",
			search:   "World",
			replace:  "Universe",
			expected: "Hello Universe Universe",
		},
		{
			name:     "no matches",
			input:    "Hello World",
			search:   "Universe",
			replace:  "Galaxy",
			expected: "Hello World",
		},
		{
			name:     "empty search string",
			input:    "Hello",
			search:   "",
			replace:  "x",
			expected: "xHxexlxlxox",
		},
		{
			name:     "empty replace string",
			input:    "Hello World",
			search:   "l",
			replace:  "",
			expected: "Heo Word",
		},
		{
			name:     "empty input",
			input:    "",
			search:   "test",
			replace:  "replace",
			expected: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			str := New(tc.input)
			result := str.ReplaceAll(tc.search, tc.replace).Get()
			if result != tc.expected {
				t.Errorf("Expected: %q but got: %q", tc.expected, result)
			}
		})
	}
}
