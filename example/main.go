package main

import (
	"fmt"

	"github.com/gobeam/stringy"
)

func main() {

	strBetween := stringy.New("HelloMyName")
	fmt.Println(strBetween.Between("hello", "name").ToUpper())

	teaseString := stringy.New("Hello My name is Roshan. I am full stack developer")
	fmt.Println(teaseString.Tease(20, "..."))

	replaceFirst := stringy.New("Hello My name is Roshan and his name is Alis.")
	fmt.Println(replaceFirst.ReplaceFirst("name", "nombre"))

	replaceLast := stringy.New("Hello My name is Roshan and his name is Alis.")
	fmt.Println(replaceLast.ReplaceLast("name", "nombre"))

	snakeCase := stringy.New("ThisIsOne___messed up string. Can we Really Snake Case It?")
	fmt.Println(snakeCase.SnakeCase("?", "").Get())
	fmt.Println(snakeCase.SnakeCase("?", "").ToUpper())
	fmt.Println(snakeCase.SnakeCase("?", "").ToLower())

	camelCase := stringy.New("ThisIsOne___messed up string. Can we Really camel-case It ?##")
	fmt.Println(camelCase.CamelCase("?", "", "#", ""))

	delimiterString := stringy.New("ThisIsOne___messed up string. Can we Really delimeter-case It?")
	fmt.Println(delimiterString.Delimited("?").Get())

	contains := stringy.New("hello mam how are you??")
	fmt.Println(contains.ContainsAll("mams", "?"))

	lines := stringy.New("fòô\r\nbàř\nyolo123")
	fmt.Println(lines.Lines())

	reverse := stringy.New("This is only test")
	fmt.Println(reverse.Reverse())

	pad := stringy.New("Roshan")
	fmt.Println(pad.Pad(10, "0", "both"))
	fmt.Println(pad.Pad(10, "0", "left"))
	fmt.Println(pad.Pad(10, "0", "right"))

	shuffleString := stringy.New("roshan")
	fmt.Println(shuffleString.Shuffle())

	cleanString := stringy.New("special@#remove%%%%")
	fmt.Println(cleanString.RemoveSpecialCharacter())

	boolString := stringy.New("off")
	fmt.Println(boolString.Boolean())

	surroundStr := stringy.New("__")
	fmt.Println(surroundStr.Surround("-"))

	str := stringy.New("hello__man how-Are you??")
	result := str.CamelCase("?", "")
	fmt.Println(result) // HelloManHowAreYou

	snakeStr := str.SnakeCase("?", "")
	fmt.Println(snakeStr.ToLower()) // hello_man_how_are_you

	kebabStr := str.KebabCase("?", "")
	fmt.Println(kebabStr.ToUpper()) // HELLO-MAN-HOW-ARE-YOU

	fcn := stringy.New("4111 1111 1111 1111")
	first := fcn.First(4)
	fmt.Println(first) // 4111

	lcn := stringy.New("4111 1111 1111 1348")
	last := lcn.Last(4)
	fmt.Println(last) // 1348

	ufo := stringy.New("known flying object")
	fmt.Println(ufo.Prefix("un")) // unknown flying object

	pun := stringy.New("this really is a cliff")
	fmt.Println(pun.Suffix("hanger")) // this really is a cliffhanger

	acronym := stringy.New("Laugh Out Loud")
	// fmt.Println(acronym.Acronym().Get())     // LOL
	fmt.Println(acronym.Acronym().ToLower()) // lol
}
