package main

import (
	"fmt"
	. "github.com/gobeam/string-manipulation"
)


func main() {
	//str := New("HelloMyName")
	//fmt.Println(str.Between("hello","name").ToUpper())

	//teaseString := New("Hello My name is Roshan.")
	//fmt.Println(teaseString.Tease(100,"..."))

	//replaceFirst := New("Hello My name is Roshan and his name is Alis.")
	//fmt.Println(replaceFirst.ReplaceFirst("name", "nau"))
	//
	//replaceLast := New("Hello My name is Roshan and his name is Alis.")
	//fmt.Println(replaceLast.ReplaceLast("name", "nau"))

	//snakeCase := New("hey man how are you")
	//fmt.Println(snakeCase.ToSnake().ToLower())
	//camelCase := New("any__Yoko    _._-_po122ΩΩΩß##s_si")
	//fmt.Println(camelCase.CamelCase())

	kebabCase := New("any__Yoko    _._-_po122ΩΩΩß##s_si")
	fmt.Println(kebabCase.KebabCase())

	//fmt.Println(strcase.ToKebab("any__Yoko    _._-_po122ΩΩΩß##s_si"))

	//matchFirstCap := regexp.MustCompile("[-._][^a-z0-9]*")
	//rslt := matchFirstCap.ReplaceAllString("Any__       _._-_pos_si"," ")
	//fmt.Println(rslt)
}