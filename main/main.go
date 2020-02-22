package main

import "fmt"

func alphanumeric(str string) bool {
	if len(str) < 1 {
		return false
	} else {
		for _, ch := range str {
			if !(ch >= '0' && ch <= '9' ) && !(ch >= 'a' && ch <= 'z') && !(ch >= 'A' && ch <= 'Z') {
				return false
			}
		}
	}

	return true
}

func main(){
	fmt.Println(alphanumeric("asdlkjAKLkla123123"))
	fmt.Println(alphanumeric("asdlkjAKLkla 123"))
}
