/*
@author: sk
@date: 2024/4/27
*/
package main

type RegexExpr interface {
	// 返回匹配的所有下标
	Match(chars []rune, index []int) []int
}

type Literal struct { // 单字符   a b . 等
	Char rune
}

func (l *Literal) Match(chars []rune, indexes []int) []int {
	res := make([]int, 0)
	for _, item := range indexes {
		if l.match(chars, item) {
			res = append(res, item+1)
		}
	}
	return res
}

func (l *Literal) match(chars []rune, index int) bool {
	if index >= len(chars) {
		return false
	}
	if l.Char == '.' && chars[index] >= 'a' && chars[index] <= 'z' {
		return true
	}
	return l.Char == chars[index]
}

func NewLiteral(char rune) *Literal {
	return &Literal{Char: char}
}

type Or struct { // ｜   a|b
	Expr1, Expr2 RegexExpr
}

// 可能虽然可能但是位置不对，这里无法判断进行回溯 所有需要计算所有结果，也是唯一会使结果增加的情况
func (o *Or) Match(chars []rune, indexes []int) []int {
	if len(indexes) == 0 {
		return make([]int, 0)
	}
	res1 := o.Expr1.Match(chars, indexes)
	res2 := o.Expr2.Match(chars, indexes)
	return append(res1, res2...)
}

func NewOr(expr1 RegexExpr, expr2 RegexExpr) *Or {
	return &Or{Expr1: expr1, Expr2: expr2}
}

type Concat struct { // 分组 abc -> Concat(a,Concat(b,c))
	Expr1, Expr2 RegexExpr
}

func (c *Concat) Match(chars []rune, indexes []int) []int {
	if len(indexes) == 0 {
		return make([]int, 0)
	}
	res := c.Expr1.Match(chars, indexes)
	res = c.Expr2.Match(chars, res)
	return res
}

func NewConcat(expr1 RegexExpr, expr2 RegexExpr) *Concat {
	return &Concat{Expr1: expr1, Expr2: expr2}
}

type Repeat0 struct { // a* 可以出现零次
	Expr RegexExpr
}

func (r *Repeat0) Match(chars []rune, indexes []int) []int {
	if len(indexes) == 0 {
		return make([]int, 0)
	}
	return matchRepeat(r.Expr, chars, indexes, false)
}

func matchRepeat(expr RegexExpr, chars []rune, indexes []int, lastOnce bool) []int {
	if lastOnce {
		indexes = expr.Match(chars, indexes)
		if len(indexes) == 0 {
			return make([]int, 0)
		}
	}
	res := make([]int, 0)
	for len(indexes) > 0 {
		res = append(res, indexes...)
		indexes = expr.Match(chars, indexes)
	}
	return res
}

func NewRepeat0(expr RegexExpr) *Repeat0 {
	return &Repeat0{Expr: expr}
}

type Repeat1 struct { // a+ 至少出现一次
	Expr RegexExpr
}

func (r *Repeat1) Match(chars []rune, indexes []int) []int {
	if len(indexes) == 0 {
		return make([]int, 0)
	}
	return matchRepeat(r.Expr, chars, indexes, true)
}

func NewRepeat1(expr RegexExpr) *Repeat1 {
	return &Repeat1{Expr: expr}
}
