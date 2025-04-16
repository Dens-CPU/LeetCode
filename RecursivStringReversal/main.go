package main

import "fmt"

func RecursivString(str string) string {
	fmt.Println("Начало работы функции")
	var recusriv func(str string)
	newstr := ""
	recusriv = func(str string) {
		newstr += string(str[len(str)-1])
		str = str[:len(str)-1]
		if len(str) != 0 {
			recusriv(str)
		}
	}
	recusriv(str)
	return newstr
}

//Чистый вариант

func RecursivStringNew(str string) string {
	if len(str) == 0 {
		return ""
	}
	return RecursivStringNew(str[1:]) + string(str[0])
}
func main() {
	str := "abc"
	fmt.Println(RecursivString(str))
	fmt.Println(RecursivStringNew(str))

}
