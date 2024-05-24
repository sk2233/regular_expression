/*
@author: sk
@date: 2024/4/27
*/
package main

func MatchStr(expr RegexExpr, str string) bool { // 递归进行
	chars := []rune(str)
	indexes := expr.Match(chars, []int{0}) // 获取从 0 开始的所有可能
	for _, item := range indexes {
		if item == len(chars) {
			return true
		}
	}
	return false
}
