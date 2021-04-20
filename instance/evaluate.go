package instance

import (
	"fmt"
	newStack "github.com/DataStructure/stack"
)

func EvaluateExpression(expression string) {
	op := []string{"+", "-", "*", "/"}

	// 运算符
	OPTR := newStack.InitStack()
	// 运算数
	OPND := newStack.InitStack()

	for _, v := range expression {
		ch := string(v)
		s, _ := OPTR.Top()
		for ch != "#" || s.(string) != "#" {
			if ch != "" {
				fmt.Println(op)
			}
		}

	}
	fmt.Println(OPTR, OPND)
}
