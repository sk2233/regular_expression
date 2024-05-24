/*
@author: sk
@date: 2024/4/27
*/
package main

import "fmt"

var (
	exprs []rune
	index = 0
)

func Match(c rune) bool {
	if index >= len(exprs) || exprs[index] != c {
		return false
	}
	index++
	return true
}

func MustMatch(c rune) {
	if exprs[index] != c {
		panic(fmt.Sprintf("%c not %c", exprs[index], c))
	}
	index++
}

func IsEnd() bool {
	return index >= len(exprs)
}

func Get() rune {
	return exprs[index]
}

func ReadChar() rune {
	index++
	return exprs[index-1]
}

// abc+c*(av)*|a 由于都是单个字符可以直接进行解析，无需尽量 token
func ParseRegex(expr string) RegexExpr {
	// 先初始化
	exprs = []rune(expr)
	// 再解析
	return ParseOr()
}

func ParseOr() RegexExpr {
	expr1 := ParseConcat()
	for Match('|') {
		expr2 := ParseConcat()
		expr1 = NewOr(expr1, expr2)
	}
	return expr1
}

func ParseConcat() RegexExpr {
	expr1 := ParseRepeat()
	if !IsEnd() && Get() != ')' && Get() != '|' {
		expr2 := ParseConcat()
		expr1 = NewConcat(expr1, expr2)
	}
	return expr1
}

func ParseRepeat() RegexExpr {
	expr := ParseLiteral()
	if Match('*') {
		return NewRepeat0(expr)
	}
	if Match('+') {
		return NewRepeat1(expr)
	}
	return expr
}

func ParseLiteral() RegexExpr {
	char := ReadChar()
	if char >= 'a' && char <= 'z' {
		return NewLiteral(char)
	}
	if char == '.' {
		return NewLiteral(char)
	}
	if char == '(' {
		expr := ParseOr()
		MustMatch(')')
		return expr
	}
	panic(fmt.Sprintf("err char of %c", char))
}
