package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	x := "1 + 1 + 3"

	if Evaluate(x) != -1 {
		fmt.Printf("%s = %d", x, Evaluate(x))
	} else {
		fmt.Println("x Not Valid")
	}
}

func tokenizing(x string) []string {
	s := strings.Split(x, " ")
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

func isOperand(c string) bool {
	_, err := strconv.ParseFloat(c, 64)
	return err == nil
}

func value(c string) int {
	val, _ := strconv.Atoi(c)
	return val
}

//Evaluate Expression Strings
func Evaluate(x string) int {
	exp := tokenizing(x)
	len1 := len(exp)

	if len1 == 0 {
		return -1
	}

	res := value(exp[0])

	for i := 1; i < len1; i += 2 {
		opr := exp[i]
		opd := exp[i+1]

		if !isOperand(opd) {
			return -1
		}

		switch opr {
		case "+":
			res += value(opd)
			break
		case "-":
			res -= value(opd)
			break
		default:
			return -1
		}
	}

	return res
}
