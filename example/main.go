package main

import (
	"fmt"
	"reflect"

	"github.com/gobeam/stringy"
)

func main() {
	fmt.Println("=== Stringy Library Functionality Demo with Assertions ===")

	// Between
	fmt.Println("== Between ==")
	strBetween := stringy.New("HelloMyName")
	betweenResult := strBetween.Between("hello", "name").ToUpper()
	fmt.Println(betweenResult)
	assertStringEquals("Between with ToUpper", betweenResult, "MY")

	// Tease
	fmt.Println("\n== Tease ==")
	teaseString := stringy.New("Hello My name is Roshan. I am full stack developer")
	teaseResult := teaseString.Tease(20, "...")
	fmt.Println(teaseResult)
	assertStringEquals("Tease", teaseResult, "Hello My name is Ros...")

	// ReplaceFirst
	fmt.Println("\n== ReplaceFirst ==")
	replaceFirst := stringy.New("Hello My name is Roshan and his name is Alis.")
	replaceFirstResult := replaceFirst.ReplaceFirst("name", "nombre")
	fmt.Println(replaceFirstResult)
	assertStringEquals("ReplaceFirst", replaceFirstResult, "Hello My nombre is Roshan and his name is Alis.")

	// ReplaceLast
	fmt.Println("\n== ReplaceLast ==")
	replaceLast := stringy.New("Hello My name is Roshan and his name is Alis.")
	replaceLastResult := replaceLast.ReplaceLast("name", "nombre")
	fmt.Println(replaceLastResult)
	assertStringEquals("ReplaceLast", replaceLastResult, "Hello My name is Roshan and his nombre is Alis.")

	// SnakeCase
	fmt.Println("\n== SnakeCase ==")
	snakeCase := stringy.New("ThisIsOne___messed up string. Can we Really Snake Case It?")
	snakeCaseResult := snakeCase.SnakeCase("?", "").Get()
	snakeCaseUpperResult := snakeCase.SnakeCase("?", "").ToUpper()
	snakeCaseLowerResult := snakeCase.SnakeCase("?", "").ToLower()
	fmt.Println(snakeCaseResult)
	fmt.Println(snakeCaseUpperResult)
	fmt.Println(snakeCaseLowerResult)
	assertStringEquals("SnakeCase", snakeCaseResult, "This_Is_One_messed_up_string_Can_we_Really_Snake_Case_It")
	assertStringEquals("SnakeCase ToUpper", snakeCaseUpperResult, "THIS_IS_ONE_MESSED_UP_STRING_CAN_WE_REALLY_SNAKE_CASE_IT")
	assertStringEquals("SnakeCase ToLower", snakeCaseLowerResult, "this_is_one_messed_up_string_can_we_really_snake_case_it")

	// CamelCase
	fmt.Println("\n== CamelCase ==")
	camelCase := stringy.New("ThisIsOne___messed up string. Can we Really camel-case It ?##")
	camelCaseResult := camelCase.CamelCase("?", "", "#", "").Get()
	fmt.Println(camelCaseResult)
	assertStringEquals("CamelCase", camelCaseResult, "thisIsOneMessedUpStringCanWeReallyCamelCaseIt")

	// Delimited
	fmt.Println("\n== Delimited ==")
	delimiterString := stringy.New("ThisIsOne___messed up string. Can we Really delimeter-case It")
	delimitedResult := delimiterString.Delimited(".").Get()
	fmt.Println(delimitedResult)
	assertStringEquals("Delimited", delimitedResult, "This.Is.One.messed.up.string.Can.we.Really.delimeter.case.It")

	// ContainsAll
	fmt.Println("\n== ContainsAll ==")
	contains := stringy.New("hello mam how are you??")
	containsResult := contains.ContainsAll("mam", "?")
	fmt.Println(containsResult)
	assertTrue("ContainsAll true case", containsResult)
	containsFalseResult := contains.ContainsAll("xyz")
	assertFalse("ContainsAll false case", containsFalseResult)

	// Lines
	fmt.Println("\n== Lines ==")
	linesString := stringy.New("fòô\r\nbàř\nyolo123")
	linesResult := linesString.Lines()
	fmt.Println(linesResult)
	assertSliceEquals("Lines", linesResult, []string{"fòô", "bàř", "yolo123"})

	// Reverse
	fmt.Println("\n== Reverse ==")
	reverse := stringy.New("This is only test")
	reverseResult := reverse.Reverse()
	fmt.Println(reverseResult)
	assertStringEquals("Reverse", reverseResult, "tset ylno si sihT")

	// Pad
	fmt.Println("\n== Pad ==")
	pad := stringy.New("Roshan")
	padBothResult := pad.Pad(10, "0", "both")
	padLeftResult := pad.Pad(10, "0", "left")
	padRightResult := pad.Pad(10, "0", "right")
	fmt.Println(padBothResult)
	fmt.Println(padLeftResult)
	fmt.Println(padRightResult)
	assertStringEquals("Pad both", padBothResult, "00Roshan00")
	assertStringEquals("Pad left", padLeftResult, "0000Roshan")
	assertStringEquals("Pad right", padRightResult, "Roshan0000")

	// Shuffle - can't assert exact result as it's random
	fmt.Println("\n== Shuffle ==")
	shuffleString := stringy.New("roshan")
	shuffleResult := shuffleString.Shuffle()
	fmt.Println(shuffleResult)
	assertTrue("Shuffle length", len(shuffleResult) == len("roshan"))

	// RemoveSpecialCharacter
	fmt.Println("\n== RemoveSpecialCharacter ==")
	cleanString := stringy.New("special@#remove%%%%")
	cleanResult := cleanString.RemoveSpecialCharacter()
	fmt.Println(cleanResult)
	assertStringEquals("RemoveSpecialCharacter", cleanResult, "specialremove")

	// Boolean
	fmt.Println("\n== Boolean ==")
	boolString := stringy.New("off")
	boolResult := boolString.Boolean()
	fmt.Println(boolResult)
	assertFalse("Boolean false", boolResult)
	boolTrueString := stringy.New("on")
	boolTrueResult := boolTrueString.Boolean()
	assertTrue("Boolean true", boolTrueResult)

	// Surround
	fmt.Println("\n== Surround ==")
	surroundStr := stringy.New("__")
	surroundResult := surroundStr.Surround("-")
	fmt.Println(surroundResult)
	assertStringEquals("Surround", surroundResult, "-__-")

	// More CamelCase and SnakeCase examples
	fmt.Println("\n== More Case Conversion Examples ==")
	str := stringy.New("hello__man how-Are you??")
	caseResult := str.CamelCase("?", "").Get()
	fmt.Println(caseResult)
	assertStringEquals("CamelCase complex", caseResult, "helloManHowAreYou")

	snakeStr := str.SnakeCase("?", "")
	snakeStrResult := snakeStr.ToLower()
	fmt.Println(snakeStrResult)
	assertStringEquals("SnakeCase with ToLower", snakeStrResult, "hello_man_how_are_you")

	kebabStr := str.KebabCase("?", "")
	kebabStrResult := kebabStr.ToUpper()
	fmt.Println(kebabStrResult)
	assertStringEquals("KebabCase with ToUpper", kebabStrResult, "HELLO-MAN-HOW-ARE-YOU")

	// First and Last
	fmt.Println("\n== First and Last ==")
	fcn := stringy.New("4111 1111 1111 1111")
	firstResult := fcn.First(4)
	fmt.Println(firstResult)
	assertStringEquals("First", firstResult, "4111")

	lcn := stringy.New("4111 1111 1111 1348")
	lastResult := lcn.Last(4)
	fmt.Println(lastResult)
	assertStringEquals("Last", lastResult, "1348")

	// Prefix and Suffix
	fmt.Println("\n== Prefix and Suffix ==")
	ufo := stringy.New("known flying object")
	prefixResult := ufo.Prefix("un")
	fmt.Println(prefixResult)
	assertStringEquals("Prefix", prefixResult, "unknown flying object")

	pun := stringy.New("this really is a cliff")
	suffixResult := pun.Suffix("hanger")
	fmt.Println(suffixResult)
	assertStringEquals("Suffix", suffixResult, "this really is a cliffhanger")

	// Acronym
	fmt.Println("\n== Acronym ==")
	acronym := stringy.New("Laugh Out Loud")
	acronymResult := acronym.Acronym().ToLower()
	fmt.Println(acronymResult)
	assertStringEquals("Acronym with ToLower", acronymResult, "lol")

	// Title
	fmt.Println("\n== Title ==")
	title := stringy.New("this is just AN eXample")
	titleResult := title.Title()
	fmt.Println(titleResult)
	assertStringEquals("Title", titleResult, "This Is Just An Example")

	// Substring
	fmt.Println("\n== Substring ==")
	subStr := stringy.New("Hello World")
	subStrResult := subStr.Substring(0, 5).Get()
	fmt.Println(subStrResult)
	assertStringEquals("Substring", subStrResult, "Hello")

	// For multi-byte characters
	subStrMB := stringy.New("Hello 世界")
	subStrMBResult := subStrMB.Substring(6, 8).Get()
	fmt.Println(subStrMBResult)
	assertStringEquals("Substring with multi-byte", subStrMBResult, "世界")

	// Empty result - start == end
	subStrEmptyResult := subStr.Substring(2, 2).Get()
	assertStringEquals("Substring empty (start == end)", subStrEmptyResult, "")

	// Contains
	fmt.Println("\n== Contains ==")
	containsStr := stringy.New("Hello World")
	containsTrue := containsStr.Contains("World")
	containsFalse := containsStr.Contains("Universe")
	fmt.Println("Contains 'World':", containsTrue)
	fmt.Println("Contains 'Universe':", containsFalse)
	assertTrue("Contains true case", containsTrue)
	assertFalse("Contains false case", containsFalse)

	// ReplaceAll
	fmt.Println("\n== ReplaceAll ==")
	replaceAllStr := stringy.New("Hello World World")
	replaceAllResult := replaceAllStr.ReplaceAll("World", "Universe").Get()
	fmt.Println(replaceAllResult)
	assertStringEquals("ReplaceAll", replaceAllResult, "Hello Universe Universe")

	// Trim
	fmt.Println("\n== Trim ==")
	trimStr := stringy.New("  Hello World  ")
	trimResult := trimStr.Trim().Get()
	fmt.Println(trimResult)
	assertStringEquals("Trim whitespace", trimResult, "Hello World")

	specialTrimStr := stringy.New("!!!Hello World!!!")
	specialTrimResult := specialTrimStr.Trim("!").Get()
	fmt.Println(specialTrimResult)
	assertStringEquals("Trim specific chars", specialTrimResult, "Hello World")

	// IsEmpty
	fmt.Println("\n== IsEmpty ==")
	emptyStr := stringy.New("")
	isEmptyResult := emptyStr.IsEmpty()
	fmt.Println("'' is empty:", isEmptyResult)
	assertTrue("IsEmpty with empty string", isEmptyResult)

	nonEmptyStr := stringy.New("Hello")
	isNotEmptyResult := nonEmptyStr.IsEmpty()
	fmt.Println("'Hello' is empty:", isNotEmptyResult)
	assertFalse("IsEmpty with non-empty string", isNotEmptyResult)

	whitespaceStr := stringy.New("   \t\n")
	isWhitespaceEmptyResult := whitespaceStr.IsEmpty()
	fmt.Println("Whitespace is empty:", isWhitespaceEmptyResult)
	assertTrue("IsEmpty with whitespace", isWhitespaceEmptyResult)

	// WordCount
	fmt.Println("\n== WordCount ==")
	wordCountStr := stringy.New("This is a test")
	wordCount := wordCountStr.WordCount()
	fmt.Println("Word count:", wordCount)
	assertTrue("WordCount", wordCount == 4)

	// TruncateWords
	fmt.Println("\n== TruncateWords ==")
	truncateStr := stringy.New("This is a long sentence that needs to be truncated")
	truncateResult := truncateStr.TruncateWords(4, "...").Get()
	fmt.Println(truncateResult)
	assertStringEquals("TruncateWords", truncateResult, "This is a long...")

	// SlugifyWithCount
	fmt.Println("\n== SlugifyWithCount ==")
	slugifyStr := stringy.New("This is a blog post title")
	slugifyResult := slugifyStr.SlugifyWithCount(1).Get()
	fmt.Println(slugifyResult)
	assertStringEquals("SlugifyWithCount", slugifyResult, "this-is-a-blog-post-title-1")

	fmt.Println("\n=== All assertions passed! ===")
}

// Assertion helper functions
func assertStringEquals(name, actual, expected string) {
	if actual != expected {
		fmt.Printf("❌ ASSERTION FAILED for %s:\nExpected: %q\nActual:   %q\n",
			name, expected, actual)
		panic("Assertion failed")
	} else {
		fmt.Printf("✅ %s: Passed\n", name)
	}
}

func assertTrue(name string, condition bool) {
	if !condition {
		fmt.Printf("❌ ASSERTION FAILED for %s: Expected true, got false\n", name)
		panic("Assertion failed")
	} else {
		fmt.Printf("✅ %s: Passed\n", name)
	}
}

func assertFalse(name string, condition bool) {
	if condition {
		fmt.Printf("❌ ASSERTION FAILED for %s: Expected false, got true\n", name)
		panic("Assertion failed")
	} else {
		fmt.Printf("✅ %s: Passed\n", name)
	}
}

func assertSliceEquals(name string, actual, expected []string) {
	if !reflect.DeepEqual(actual, expected) {
		fmt.Printf("❌ ASSERTION FAILED for %s:\nExpected: %v\nActual:   %v\n",
			name, expected, actual)
		panic("Assertion failed")
	} else {
		fmt.Printf("✅ %s: Passed\n", name)
	}
}
