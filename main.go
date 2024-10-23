package main

import (
	"errors"
	"strconv"
	"strings"
)

func ProcessNum(n string) (float64, error) {
	if strings.Replace(n, " ", "", -1) == "" {
		return 0, errors.New("bad num format1")
	}
	if n[len(n)-1] == '*' || n[len(n)-1] == '/' {
		return 0, errors.New("bad num format1")
	}
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
	expression = strings.Replace(expression, " ", "", -1)
	rarr := []rune(expression)
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
		return 0, errors.New("bad brackets seq2")
	}
	rarr = []rune(new_exp)
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
	return ret, nil
}
