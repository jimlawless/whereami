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
	fmt.Printf("%s\n", whereami.WhereAmI())
}
