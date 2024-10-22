package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func FindRune(str string, r rune, start int, step int) int {
	rarr := []rune(str)
	i := start
	for i > -1 && i < len(str) && rarr[i] != r {
		i += step
	}
	if i >= len(str) {
		i = -1
	}
	return i
}

func Calc(expression string) (float64, error) {
	expression = strings.Replace(expression, " ", "", -1)
	rarr := []rune(expression)
	// preprocess
	bracket_depth := 0
	l := 0
	// r := 0;
	new_exp := ""
	for i := 0; i < len(expression); i++ {
		if rarr[i] == '(' {
			if bracket_depth == 0 {
				l = i
			}
			bracket_depth++
		} else if rarr[i] == ')' {
			if bracket_depth == 1 {
				num, ok := Calc(expression[l+1 : i])
				if ok != nil {
					return num, ok
				}
				new_exp += strconv.FormatFloat(num, 'f', -1, 64)
				continue
			}
			if bracket_depth == 0 {
				return 0, errors.New("bad brackets seq")
			}
			bracket_depth--
		} else if bracket_depth == 0 {
			new_exp += string(rarr[i])
		}
	}
	fmt.Println(new_exp)
	return 0, nil
}

func main() {
	// fmt.Print(FindRune("()", '(', 1, -1))
	fmt.Print(Calc("1 + 3 + 4(1)"))
}
