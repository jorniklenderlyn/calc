package main

import (
	"errors"
	// "fmt"
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

func ProcessNum(n string) (float64, error) {
	rarr := []rune(n)
	cur_num := ""
	var ret float64 = 0
	var sign rune = 0
	for _, el := range rarr {
		if el == '*' || el == '/' {
			fnum, ok := strconv.ParseFloat(cur_num, 64)
			if ok != nil {
				return 0, errors.New("bad num format1")
			}
			if ret == 0 {
				ret = fnum
			} else {
				if sign == '*' {
					ret *= fnum
				} else {
					ret /= fnum
				}
			}
			sign = el
			cur_num = ""
		} else {
			cur_num += string(el)
		}
	}
	if cur_num != "" {
		fnum, ok := strconv.ParseFloat(cur_num, 64)
		if ok != nil {
			return 0, errors.New("bad num format2")
		}
		if sign == '*' {
			ret *= fnum
		} else if sign == 0 {
			ret = fnum
		} else {
			ret /= fnum
		}
	}
	return ret, nil
}

func Calc(expression string) (float64, error) {
	// fmt.Println("S", expression)
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
				// continue
			}
			if bracket_depth == 0 {
				return 0, errors.New("bad brackets seq1")
			}
			bracket_depth--
		} else if bracket_depth == 0 {
			new_exp += string(rarr[i])
		}
	}
	if bracket_depth != 0 {
		// fmt.Println(bracket_depth)
		return 0, errors.New("bad brackets seq2")
	}
	rarr = []rune(new_exp)
	// narr := []int{}
	// cur_num := ""
	// ret := 0
	// cur_num := 1
	// a := 1
	// for i := 0; i < len(new_exp); i++ {
	// 	if rarr[i] == '+' || rarr[i] == '-' {
	// 		strconv.
	// 		if rarr == '-' {
	// 			a = -1
	// 		} else {
	// 			a = 1
	// 		}
	// 	}
	// }
	// if cur_num != "" {
	// 	ret +
	// }
	signs := []rune{}
	for _, el := range rarr {
		if el == '+' || el == '-' {
			signs = append(signs, el)
		}
	}
	new_exp = strings.Replace(new_exp, "-", "+", -1)
	nums := strings.Split(new_exp, "+")
	fnums := []float64{}
	for _, n := range nums {
		fnum, ok := ProcessNum(n)
		if ok != nil {
			return 0, ok
		}
		fnums = append(fnums, fnum)
	}
	ret := fnums[0]
	for i := 1; i < len(fnums); i++ {
		if signs[i-1] == '+' {
			ret += fnums[i]
		} else {
			ret -= fnums[i]
		}
	}
	// fmt.Println(signs)
	// fmt.Println(fnums)
	// fmt.Println(new_exp)
	return ret, nil
}

// func main() {
// fmt.Print(FindRune("()", '(', 1, -1))
// fmt.Print(Calc("(1 + 3 + 4*2) / 12"))
// fmt.Print(ProcessNum("1"))
// }
