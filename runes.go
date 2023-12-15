package main

import (
	"fmt"
	"unicode/utf8"
	// "unicode/utf8"
)

/*
go uses UTF-8 encoding. so we have runes which is an alias for int32 datatype
so in go we don't have char which is 7 bit ASCII encoding.
so the length of string != no of characters in the string
here length refers to the no of byte the string consumes

for eg->const s = "สวัสดี" this thai string has 4 chars and
string is of length 18, and as these characters are multiByte long
so length !=4


	So const s = "สวัสดี" this thai string is 6 runes long and each rune is 3 byte long
	which makes length eq to 18
*/

func examineRune(r rune) {
	//each of these characters to which we are matching the condition, we call this each
	//char a rune literal and can only be used inside single quotation
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}

}

func Main14() {

	const s = "สวัสดี" // instead of []char it's []byte
	fmt.Println(len(s))

	// for i := 0; i < len(s); i++ {
	// 	fmt.Printf("%T ", s[i])  //prints uint8
	// 	fmt.Printf("%x\n", s[i]) //hexadecimal for each byte
	// }
	// fmt.Println()

	// fmt.Println("Rune count:", utf8.RuneCountInString(s))

	// for idx, runeValue := range s {
	// 	fmt.Printf("%#U starts at %d\n", runeValue, idx) //this %#U is a format specifier for runeValues
	// }

	for i, w := 0, 0; i < len(s); i += w {
		runeV, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeV, i)
		w = width
		examineRune(runeV)
	}

}
