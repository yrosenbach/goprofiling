package main

import (
	"fmt"
	"strings"
)

// HelloToWorld get a space separated string with replaces any occurance of case in-sensative 'hello' with an upper case 'world'.
func HelloToWorld(helloStr string) string {
	var result string
	tokens := strings.Split(helloStr, " ")

	for _, token := range tokens {
		if strings.ToLower(token) == "hello" {
			result += "world"
		} else {
			result += token
		}

		result += " "
	}

	return result
}

func main() {
	fmt.Println(HelloToWorld("My hellO is HELLO"))
}
