package main

import (
	"fmt"
	. "github.com/gobeam/Stringy"
)

func main() {

	strBetween := New("HelloMyName")
	fmt.Println(strBetween.Between("hello", "name").ToUpper())

	teaseString := New("Hello My name is Roshan. I am full stack developer")
	fmt.Println(teaseString.Tease(20, "..."))

	replaceFirst := New("Hello My name is Roshan and his name is Alis.")
	fmt.Println(replaceFirst.ReplaceFirst("name", "nombre"))

	replaceLast := New("Hello My name is Roshan and his name is Alis.")
	fmt.Println(replaceLast.ReplaceLast("name", "nombre"))

	snakeCase := New("ThisIsOne___messed up string. Can we Really Snake Case It?")
	fmt.Println(snakeCase.SnakeCase("?", "").Get())
	fmt.Println(snakeCase.SnakeCase("?", "").ToUpper())
	fmt.Println(snakeCase.SnakeCase("?", "").ToLower())

	camelCase := New("ThisIsOne___messed up string. Can we Really camel-case It ?##")
	fmt.Println(camelCase.CamelCase("?", "", "#", ""))

	delimiterString := New("ThisIsOne___messed up string. Can we Really delimeter-case It?")
	fmt.Println(delimiterString.Delimited("?").Get())

	contains := New("hello mam how are you??")
	fmt.Println(contains.ContainsAll("mams", "?"))

	lines := New("fòô\r\nbàř\nyolo123")
	fmt.Println(lines.Lines())

	reverse := New("This is only test")
	fmt.Println(reverse.Reverse())

	pad := New("Roshan")
	fmt.Println(pad.Pad(10, "0", "both"))
	fmt.Println(pad.Pad(10, "0", "left"))
	fmt.Println(pad.Pad(10, "0", "right"))

	shuffleString := New("roshan")
	fmt.Println(shuffleString.Shuffle())

	cleanString := New("special@#remove%%%%")
	fmt.Println(cleanString.RemoveSpecialCharacter())

	boolString := New("off")
	fmt.Println(boolString.Boolean())

	surroundStr := New("__")
	fmt.Println(surroundStr.Surround("-"))

	str := New("hello__man how-Are you??")
	result := str.CamelCase("?", "")
	fmt.Println(result) // HelloManHowAreYou

	snakeStr := str.SnakeCase("?", "")
	fmt.Println(snakeStr.ToLower()) // hello_man_how_are_you

	kebabStr := str.KebabCase("?", "")
	fmt.Println(kebabStr.ToUpper()) // HELLO-MAN-HOW-ARE-YOU

}
