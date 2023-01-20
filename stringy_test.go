package stringy

import (
	"testing"
)

var sm StringManipulation = New("This is example.")

func TestInput_Acronym(t *testing.T) {
	acronym := New("Laugh Out Loud")
	val := acronym.Acronym().Get()
	if val != "LOL" {
		t.Errorf("Expected: %s but got: %s", "IS", val)
	}
}

func TestInput_Between(t *testing.T) {
	val := sm.Between("This", "example").ToUpper()
	if val != "IS" {
		t.Errorf("Expected: %s but got: %s", "IS", val)
	}
}

func TestInput_EmptyBetween(t *testing.T) {
	sm := New("This is example.")
	val := sm.Between("", "").ToUpper()
	if val != "THIS IS EXAMPLE." {
		t.Errorf("Expected: %s but got: %s", "THIS IS EXAMPLE.", val)
	}
}

func TestInput_EmptyNoMatchBetween(t *testing.T) {
	sm := New("This is example.")
	val := sm.Between("hello", "test").ToUpper()
	if val != "THIS IS EXAMPLE." {
		t.Errorf("Expected: %s but got: %s", "THIS IS EXAMPLE.", val)
	}
}

func TestInput_BooleanTrue(t *testing.T) {
	strs := []string{"on", "On", "yes", "YES", "1", "true"}
	for _, s := range strs {
		t.Run(s, func(t *testing.T) {
			if val := New(s).Boolean(); !val {
				t.Errorf("Expected: to be true but got: %v", val)
			}
		})
	}
}

func TestInput_BooleanFalse(t *testing.T) {
	strs := []string{"off", "Off", "no", "NO", "0", "false"}
	for _, s := range strs {
		t.Run(s, func(t *testing.T) {
			if val := New(s).Boolean(); val {
				t.Errorf("Expected: to be false but got: %v", val)
			}
		})
	}
}

func TestInput_BooleanError(t *testing.T) {
	strs := []string{"invalid", "-1", ""}
	for _, s := range strs {
		t.Run(s, func(t *testing.T) {
			defer func() {
				if err := recover(); err == nil {
					t.Errorf("Error expected")
				}
			}()
			val := New(s).Boolean()
			t.Errorf("Expected: to panic but got: %v", val)
		})
	}

}

func TestInput_CamelCase(t *testing.T) {
	str := New("Camel case this_complicated__string%%")
	val := str.CamelCase("%", "")
	if val != "camelCaseThisComplicatedString" {
		t.Errorf("Expected: to be %s but got: %s", "camelCaseThisComplicatedString", val)
	}
}

func TestInput_CamelCaseNoRule(t *testing.T) {
	str := New("Camel case this_complicated__string%%")
	val := str.CamelCase()
	if val != "camelCaseThisComplicatedString%%" {
		t.Errorf("Expected: to be %s but got: %s", "camelCaseThisComplicatedString", val)
	}
}

func TestInput_CamelCaseOddRuleError(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Errorf("Error expected")
		}
	}()
	str := New("Camel case this_complicated__string%%")
	val := str.CamelCase("%")

	if val != "camelCaseThisComplicatedString%%" {
		t.Errorf("Expected: to be %s but got: %s", "camelCaseThisComplicatedString", val)
	}
}

func TestInput_ContainsAll(t *testing.T) {
	contains := New("hello mam how are you??")
	if val := contains.ContainsAll("mam", "?"); !val {
		t.Errorf("Expected value to be true but got false")
	}
	if val := contains.ContainsAll("non existent"); val {
		t.Errorf("Expected value to be false but got true")
	}
}

func TestInput_Delimited(t *testing.T) {
	str := New("Delimited case this_complicated__string@@")
	against := "delimited.case.this.complicated.string"
	if val := str.Delimited(".", "@", "").ToLower(); val != against {
		t.Errorf("Expected: to be %s but got: %s", against, val)
	}
}

func TestInput_DelimitedNoDelimeter(t *testing.T) {
	str := New("Delimited case this_complicated__string@@")
	against := "delimited.case.this.complicated.string@@"
	if val := str.Delimited("").ToLower(); val != against {
		t.Errorf("Expected: to be %s but got: %s", against, val)
	}
}

func TestInput_KebabCase(t *testing.T) {
	str := New("Kebab case this-complicated___string@@")
	against := "Kebab-case-this-complicated-string"
	if val := str.KebabCase("@", "").Get(); val != against {
		t.Errorf("Expected: to be %s but got: %s", against, val)
	}
}

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
			if got := New(tt.arg).LcFirst(); got != tt.want {
				t.Errorf("LcFirst(%v) = %v, want %v", tt.arg, got, tt.want)
			}
		})
	}
}

func TestInput_Lines(t *testing.T) {
	lines := New("fòô\r\nbàř\nyolo")
	strSlic := lines.Lines()
	if len(strSlic) != 3 {
		t.Errorf("Length expected to be 3 but got: %d", len(strSlic))
	}
	if strSlic[0] != "fòô" {
		t.Errorf("Expected: %s but got: %s", "fòô", strSlic[0])
	}
}

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
}

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
}

func TestInput_RemoveSpecialCharacter(t *testing.T) {
	cleanString := New("special@#remove%%%%")
	against := "specialremove"
	if result := cleanString.RemoveSpecialCharacter(); result != against {
		t.Errorf("Expected: %s but got: %s", against, result)
	}
}

func TestInput_ReplaceFirst(t *testing.T) {
	replaceFirst := New("Hello My name is Roshan and his name is Alis.")
	against := "Hello My nombre is Roshan and his name is Alis."
	if result := replaceFirst.ReplaceFirst("name", "nombre"); result != against {
		t.Errorf("Expected: %s but got: %s", against, result)
	}
}

func TestInput_ReplaceFirstEmptyInput(t *testing.T) {
	replaceFirst := New("")
	against := ""
	if result := replaceFirst.ReplaceFirst("name", "nombre"); result != against {
		t.Errorf("Expected: %s but got: %s", against, result)
	}
}

func TestInput_ReplaceLast(t *testing.T) {
	replaceLast := New("Hello My name is Roshan and his name is Alis.")
	against := "Hello My name is Roshan and his nombre is Alis."
	if result := replaceLast.ReplaceLast("name", "nombre"); result != against {
		t.Errorf("Expected: %s but got: %s", against, result)
	}
}

func TestInput_Reverse(t *testing.T) {
	reverseString := New("roshan")
	against := "nahsor"
	if result := reverseString.Reverse(); result != against {
		t.Errorf("Expected: %s but got: %s", against, result)
	}
}

func TestInput_Shuffle(t *testing.T) {
	check := "roshan"
	shuffleString := New(check)
	if result := shuffleString.Shuffle(); len(result) != len(check) && check == result {
		t.Errorf("Shuffle string gave wrong output")
	}
}

func TestInput_SnakeCase(t *testing.T) {
	str := New("SnakeCase this-complicated___string@@")
	against := "snake_case_this_complicated_string"
	if val := str.SnakeCase("@", "").ToLower(); val != against {
		t.Errorf("Expected: to be %s but got: %s", against, val)
	}
}

func TestInput_Surround(t *testing.T) {
	str := New("this")
	against := "__this__"
	if val := str.Surround("__"); val != against {
		t.Errorf("Expected: to be %s but got: %s", against, val)
	}
}

func TestInput_Tease(t *testing.T) {
	str := New("This is just simple paragraph on lorem ipsum.")
	against := "This is just..."
	if val := str.Tease(12, "..."); val != against {
		t.Errorf("Expected: to be %s but got: %s", against, val)
	}
}

func TestInput_TeaseEmpty(t *testing.T) {
	str := New("This is just simple paragraph on lorem ipsum.")
	against := "This is just simple paragraph on lorem ipsum."
	if val := str.Tease(200, "..."); val != against {
		t.Errorf("Expected: to be %s but got: %s", against, val)
	}
}

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
			if got := New(tt.arg).UcFirst(); got != tt.want {
				t.Errorf("UcFirst(%v) = %v, want %v", tt.arg, got, tt.want)
			}
		})
	}
}

func TestInput_First(t *testing.T) {
	fcn := New("4111 1111 1111 1111")
	against := "4111"
	if first := fcn.First(4); first != against {
		t.Errorf("Expected: to be %s but got: %s", against, first)
	}
}

func TestInput_FirstError(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Errorf("Error expected but got none")
		}
	}()
	fcn := New("4111 1111 1111 1111")
	fcn.First(100)
}

func TestInput_LastError(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Errorf("Error expected but got none")
		}
	}()
	fcn := New("4111 1111 1111 1111")
	fcn.Last(100)
}

func TestInput_Last(t *testing.T) {
	lcn := New("4111 1111 1111 1348")
	against := "1348"
	if last := lcn.Last(4); last != against {
		t.Errorf("Expected: to be %s but got: %s", against, last)
	}
}

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
}

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
}
