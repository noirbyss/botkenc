package main

import (
	"fmt"
	"strings"
)

func main() {
	callback1 := "main(one_device)"
	callback2 := "main"

	menu1, callbackOne := parseCallback(callback1)
	menu2, callbackTwo := parseCallback(callback2)

	fmt.Printf("%s:\nmenu:%s\nargs:%s\n\n\n%s:menu:%s\nargs:%v", callback1, menu1, *callbackOne, callback2, menu2, callbackTwo)
}

func parseCallback(s string) (string, *string) {
	hasArgs := strings.IndexByte(s, '(')

	if hasArgs == -1 {
		return s, nil
	}

	menu := s[:hasArgs]
	args := s[hasArgs+1 : len(s)-1]
	return menu, &args
}
