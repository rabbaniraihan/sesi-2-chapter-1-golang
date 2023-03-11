package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("Nilai i =", i)
	}

	for j := 0; j < 11; j++ {
		if j == 5 {
			fmt.Println("character U+0421 '\u0421' starts at byte position 0")
			fmt.Println("character U+0410 '\u0410' starts at byte position 2")
			fmt.Println("character U+0428 '\u0428' starts at byte position 4")
			fmt.Println("character U+0410 '\u0410' starts at byte position 6")
			fmt.Println("character U+0420 '\u0420' starts at byte position 8")
			fmt.Println("character U+0412 '\u0412' starts at byte position 10")
			fmt.Println("character U+041E '\u041e' starts at byte position 12")

		} else {
			fmt.Println("Nilai j =", j)
		}
	}
}
