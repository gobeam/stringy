# Golang String manipulation helper package
[![Build][Build-Status-Image]][Build-Status-Url] [![Go Report Card](https://goreportcard.com/badge/github.com/gobeam/stringy?branch=master&kill_cache=1)](https://goreportcard.com/report/github.com/gobeam/Stringy) [![GoDoc][godoc-image]][godoc-url]
[![Coverage Status](https://coveralls.io/repos/github/gobeam/stringy/badge.svg)](https://coveralls.io/github/gobeam/stringy)

Convert string to camel case, snake case, kebab case / slugify, custom delimiter, pad string, tease string and many other functionality with help of by Stringy package. You can convert camelcase to snakecase or kebabcase, or snakecase to camelcase and kebabcase and vice versa. This package was inspired from PHP [danielstjules/Stringy](https://github.com/danielstjules/Stringy).

* [Why?](#why)
* [Installation](#installation)
* [Functions](#functions)
* [Running the tests](#running-the-tests)
* [Contributing](#contributing)
* [License](#license)

<table>
    <tr>
        <td><a href="#betweenstart-end-string-stringmanipulation">Between</a></td>
        <td><a href="#boolean-bool">Boolean</a></td>
        <td><a href="#camelcaserule-string-string">CamelCase</a></td>
    </tr>
    <tr>
        <td><a href="#containsallcheck-string-bool">ContainsAll</a></td>
        <td><a href="#delimiteddelimiter-string-rule-string-stringmanipulation">Delimited</a></td>
        <td><a href="#get-string">Get</a></td>
    </tr>
    <tr>
        <td><a href="#kebabcaserule-string-stringmanipulation">KebabCase</a></td>
        <td><a href="#lcfirst-string">LcFirst</a></td>
        <td><a href="#lines-string">Lines</a></td>
    </tr>
    <tr>
        <td><a href="#padlength-int-with-padtype-string-string">Pad</a></td>
        <td><a href="#removespecialcharacter-string">RemoveSpecialCharacter</a></td>
        <td><a href="#replacefirstsearch-replace-string-string">ReplaceFirst</a></td>
    </tr>
    <tr>
        <td><a href="#replacelastsearch-replace-string-string">ReplaceLast</a></td>
        <td><a href="#reverse-string">Reverse</a></td>
        <td><a href="#shuffle-string">Shuffle</a></td>
    </tr>
    <tr>
        <td><a href="#surroundwith-string-string">Surround</a></td>
        <td><a href="#snakecaserule-string-stringmanipulation">SnakeCase</a></td>
        <td><a href="#teaselength-int-indicator-string-string">Tease</a></td>
    </tr>
    <tr>
        <td><a href="#tolower-string">ToLower</a></td>
        <td><a href="#toupper-string">ToUpper</a></td>
        <td><a href="#ucfirst-string">UcFirst</a></td>
    </tr>
     <tr>
         <td><a href="#firstlength-int-string">First</a></td>
         <td><a href="#lastlength-int-string">Last</a></td>
         <td><a href="#prefixstring-string">Prefix</a></td>
     </tr>
     <tr>
         <td><a href="#suffixstring-string">Suffix</a></td>
         <td></td>
         <td></td>
     </tr>
</table>


## Why?

Golang has very rich strings core package despite some extra helper function are not available and this stringy package is here to fill that void. Plus there are other some packages in golang, that have same functionality but for some extreme cases they fail to provide correct output. This package cross flexibility is it's main advantage. You can convert to camelcase  to snakecase or kebabcase or vice versa. 

```go
package main

import (
	"fmt"
	"github.com/gobeam/stringy"
    )

func main() {
 str := stringy.New("hello__man how-Are you??")
 result := str.CamelCase("?", "")
 fmt.Println(result) // HelloManHowAreYou

 snakeStr := str.SnakeCase("?", "")
 fmt.Println(snakeStr.ToLower()) // hello_man_how_are_you

 kebabStr := str.KebabCase("?", "")
 fmt.Println(kebabStr.ToUpper()) // HELLO-MAN-HOW-ARE-YOU
}
```

## Installation

``` bash
$ go get -u -v github.com/gobeam/stringy
```

or with dep

``` bash
$ dep ensure -add github.com/gobeam/stringy
```


## Functions

#### Between(start, end string) StringManipulation

Between takes two string params start and end which and returns value which is in middle of start and end part of input. You can chain to upper which with make result all uppercase or ToLower which will make result all lower case or Get which will return result as it is.

```go
  strBetween := stringy.New("HelloMyName")
  fmt.Println(strBetween.Between("hello", "name").ToUpper()) // MY
```

#### Boolean() bool

Boolean func returns boolean value of string value like on, off, 0, 1, yes, no returns boolean value of string input. You can chain this function on other function which returns implemented StringManipulation interface.

```go
  boolString := stringy.New("off")
  fmt.Println(boolString.Boolean()) // false
```

#### CamelCase(rule ...string) string

CamelCase is variadic function which takes one Param rule i.e slice of strings and it returns input type string in camel case form and rule helps to omit character you want to omit from string. By default special characters like "_", "-","."," " are treated like word separator and treated accordingly by default and you dont have to worry about it.

```go
  camelCase := stringy.New("ThisIsOne___messed up string. Can we Really camel-case It ?##")
  fmt.Println(camelCase.CamelCase("?", "", "#", "")) // ThisIsOneMessedUpStringCanWeReallyCamelCaseIt
```
look how it omitted ?## from string. If you dont want to omit anything and since it returns plain strings and you cant actually cap all or lower case all camelcase string its not required.

```go
  camelCase := stringy.New("ThisIsOne___messed up string. Can we Really camel-case It ?##")
  fmt.Println(camelCase.CamelCase()) // ThisIsOneMessedUpStringCanWeReallyCamelCaseIt?##
```

#### ContainsAll(check ...string) bool

ContainsAll is variadic function which takes slice of strings as param and checks if they are present in input and returns boolean value accordingly.

```go
  contains := stringy.New("hello mam how are you??")
  fmt.Println(contains.ContainsAll("mam", "?")) // true
```

#### Delimited(delimiter string, rule ...string) StringManipulation

Delimited is variadic function that takes two params delimiter and slice of strings named rule. It joins the string by passed delimeter. Rule param helps to omit character you want to omit from string. By default special characters like "_", "-","."," " are treated like word separator and treated accordingly by default and you dont have to worry about it. If you don't want to omit any character pass empty string.

```go
  delimiterString := stringy.New("ThisIsOne___messed up string. Can we Really delimeter-case It?")
  fmt.Println(delimiterString.Delimited("?").Get())
```
You can chain to upper which with make result all uppercase or ToLower which will make result all lower case or Get which will return result as it is.


#### First(length int) string

First returns first n characters from provided input. It removes all spaces in string before doing so.

```go
  fcn := stringy.New("4111 1111 1111 1111")
  first := fcn.First(4)
  fmt.Println(first) // 4111
```


#### Get() string

Get simply returns result and can be chained on function which returns StringManipulation interface view above examples


#### KebabCase(rule ...string) StringManipulation

KebabCase/slugify is variadic function that takes one Param slice of strings named rule and it returns passed string in kebab case or slugify form. Rule param helps to omit character you want to omit from string. By default special characters like "_", "-","."," " are treated like word separator and treated accordingly by default and you don't have to worry about it. If you don't want to omit any character pass nothing.

```go
  str := stringy.New("hello__man how-Are you??")
  kebabStr := str.KebabCase("?","")
  fmt.Println(kebabStr.ToUpper()) // HELLO-MAN-HOW-ARE-YOU
  fmt.Println(kebabStr.Get()) // hello-man-how-Are-you
```
You can chain to upper which with make result all uppercase or ToLower which will make result all lower case or Get which will return result as it is.


#### Last(length int) string

Last returns last n characters from provided input. It removes all spaces in string before doing so.

```go
  lcn := stringy.New("4111 1111 1111 1348")
  last := lcn.Last(4)
  fmt.Println(last) // 1348
```


#### LcFirst() string

LcFirst simply returns result by lower casing first letter of string and it can be chained on function which return StringManipulation interface

```go
  contains := stringy.New("Hello roshan")
  fmt.Println(contains.LcFirst()) // hello roshan
```


#### Lines() []string

Lines returns slice of strings by removing white space characters

```go
  lines := stringy.New("fòô\r\nbàř\nyolo123")
  fmt.Println(lines.Lines()) // [fòô bàř yolo123]
```


#### Pad(length int, with, padType string) string

Pad takes three param length i.e total length to be after padding, with i.e  what to pad with and pad type which can be ("both" or "left" or "right") it return string after padding upto length by with param and on padType type it can be chained on function which return StringManipulation interface

```go
  pad := stringy.New("Roshan")
  fmt.Println(pad.Pad(0, "0", "both"))  // 00Roshan00
  fmt.Println(pad.Pad(0, "0", "left"))  // 0000Roshan
  fmt.Println(pad.Pad(0, "0", "right")) // Roshan0000
```


#### RemoveSpecialCharacter() string

RemoveSpecialCharacter removes all special characters and returns the string nit can be chained on function which return StringManipulation interface

```go
  cleanString := stringy.New("special@#remove%%%%")
  fmt.Println(cleanString.RemoveSpecialCharacter()) // specialremove
```


#### ReplaceFirst(search, replace string) string

ReplaceFirst takes two param search and replace. It returns string by searching search sub string and replacing it with replace substring on first occurrence it can be chained on function which return StringManipulation interface.

```go
  replaceFirst := stringy.New("Hello My name is Roshan and his name is Alis.")
  fmt.Println(replaceFirst.ReplaceFirst("name", "nombre")) // Hello My nombre is Roshan and his name is Alis.
```


#### ReplaceLast(search, replace string) string

ReplaceLast takes two param search and replace it return string by searching search sub string and replacing it with replace substring on last occurrence it can be chained on function which return StringManipulation interface

```go
  replaceLast := stringy.New("Hello My name is Roshan and his name is Alis.")
  fmt.Println(replaceLast.ReplaceLast("name", "nombre")) // Hello My name is Roshan and his nombre is Alis.
```


#### Reverse() string

Reverse function reverses the passed strings it can be chained on function which return StringManipulation interface.

```go
  reverse := stringy.New("This is only test")
  fmt.Println(reverse.Reverse()) // tset ylno si sihT
```


#### Shuffle() string

Shuffle shuffles the given string randomly it can be chained on function which return StringManipulation interface.

```go
  shuffleString := stringy.New("roshan")
  fmt.Println(shuffleString.Shuffle()) // nhasro
```


#### Surround(with string) string

Surround takes one param with which is used to surround user input and it can be chained on function which return StringManipulation interface.

```go
  surroundStr := stringy.New("__")
  fmt.Println(surroundStr.Surround("-")) // -__-
```


#### SnakeCase(rule ...string) StringManipulation

SnakeCase is variadic function that takes one Param slice of strings named rule and it returns passed string in snake case form. Rule param helps to omit character you want to omit from string. By default special characters like "_", "-","."," " are treated like word separator and treated accordingly by default and you don't have to worry about it. If you don't want to omit any character pass nothing.

```go
  snakeCase := stringy.New("ThisIsOne___messed up string. Can we Really Snake Case It?")
  fmt.Println(snakeCase.SnakeCase("?", "").Get()) // This_Is_One_messed_up_string_Can_we_Really_Snake_Case_It
  fmt.Println(snakeCase.SnakeCase("?", "").ToUpper()) // THIS_IS_ONE_MESSED_UP_STRING_CAN_WE_REALLY_SNAKE_CASE_IT
```
You can chain to upper which with make result all uppercase or ToLower which will make result all lower case or Get which will return result as it is.


#### Tease(length int, indicator string) string

Tease takes two params length and indicator and it shortens given string on passed length and adds indicator on end it can be chained on function which return StringManipulation interface.

```go
  teaseString := stringy.New("Hello My name is Roshan. I am full stack developer")
  fmt.Println(teaseString.Tease(20, "...")) // Hello My name is Ros...
```


#### ToLower() string

ToLower makes all string of user input to lowercase and it can be chained on function which return StringManipulation interface.

```go
  snakeCase := stringy.New("ThisIsOne___messed up string. Can we Really Snake Case It?")
  fmt.Println(snakeCase.SnakeCase("?", "").ToLower()) // this_is_one_messed_up_string_can_we_really_snake_case_it
```


#### ToUpper() string

ToUpper makes all string of user input to uppercase and it can be chained on function which return StringManipulation interface.

```go
  snakeCase := stringy.New("ThisIsOne___messed up string. Can we Really Snake Case It?")
  fmt.Println(snakeCase.SnakeCase("?", "").ToUpper()) // THIS_IS_ONE_MESSED_UP_STRING_CAN_WE_REALLY_SNAKE_CASE_IT
```


#### UcFirst() string

LcFirst simply returns result by lower casing first letter of string and it can be chained on function which return StringManipulation interface.

```go
  contains := stringy.New("hello roshan")
  fmt.Println(contains.UcFirst()) // Hello roshan
```


#### Prefix(string) string

Prefix makes sure string has been prefixed with a given string and avoids adding it again if it has.

```go
  ufo := stringy.New("known flying object")
  fmt.Println(ufo.Prefix("un")) // unknown flying object
```


#### Suffix(string) string

Suffix makes sure string has been suffixed with a given string and avoids adding it again if it has.

```go
  pun := stringy.New("this really is a cliff")
  fmt.Println(pun.Suffix("hanger")) // this really is a cliffhanger
```


## Running the tests

``` bash
$ go test
```


## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate. - see `CONTRIBUTING.md` for details.


## License

Released under the MIT License - see `LICENSE.txt` for details.


[Build-Status-Url]: https://travis-ci.com/gobeam/stringy
[Build-Status-Image]: https://travis-ci.com/gobeam/stringy.svg?branch=master
[godoc-url]: https://pkg.go.dev/github.com/gobeam/stringy?tab=doc
[godoc-image]: https://godoc.org/github.com/gobeam/stringy?status.svg

