// Copyright 2016 - by Jim Lawless
// License: MIT / X11
// See: http://www.mailsend-online.com/license2016.php
//
// This code may not conform to popular Go coding idioms

package main

import (
	"fmt"
	"github.com/jimlawless/whereami"
)

func main() {
	trace()
	fmt.Printf("%d\n", fact(4))
}

// reference the proper stack entry index to obtain the
// information about trace()'s caller
func trace() {
	fmt.Printf("%s\n", whereami.WhereAmI(2))
}

func fact(num int) int {
	trace()
	if num == 1 {
		return 1
	} else {
		return num * fact(num-1)
	}
}
