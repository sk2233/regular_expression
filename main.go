/*
@author: sk
@date: 2024/4/27
*/
package main

import "fmt"

// https://rcoh.svbtle.com/no-magic-regular-expressions

func main() {
	regex := ParseRegex("(a*b)|(b+a.)")
	fmt.Println(MatchStr(regex, "aabb"))
}
