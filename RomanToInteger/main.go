package main

import "fmt"

func romanToInt(s string) int {
	a := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	k := 0
	for i := len(s) - 1; i != 0; i-- {
		if (s[i] == 'V' || s[i] == 'X') && (s[i-1] == 'I') {
			k += a[rune(s[i])] - 1

			if i != 1 {
				i--
			}

		} else if (s[i] == 'L' || s[i] == 'C') && (s[i-1] == 'X') {
			k += a[rune(s[i])] - 10

			if i != 1 {
				i--
			}

		} else if (s[i] == 'D' || s[i] == 'M') && (s[i-1] == 'C') {
			k += a[rune(s[i])] - 100

			if i != 1 {
				i--
			}
		} else {
			k += a[rune(s[i])]

		}

	}
	k += a[rune(s[0])]

	return k
}

func main() {
	s := "LVIII"
	fmt.Println(romanToInt(s))
}

// //func romanToInt(s string) int {
// 	var v, lv, cv int
// 	h := map[uint8]int{
// 		'I': 1,
// 		'V': 5,
// 		'X': 10,
// 		'L': 50,
// 		'C': 100,
// 		'D': 500,
// 		'M': 1000,
// 	}

//     for i := len(s) - 1; i >= 0; i-- {
// 		cv = h[s[i]]
// 		if cv < lv {
// 			v -= cv
// 		} else {
// 			v += cv
// 		}
// 		lv = cv
// 	}

// 	return v
// }
