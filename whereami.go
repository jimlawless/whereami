// Copyright 2016 - by Jim Lawless
// License: MIT / X11
// See: http://www.mailsend-online.com/license2016.php
//
// This code may not conform to popular Go coding idioms

package whereami

import (
	"fmt"
	"runtime"
	"strings"
)

type WhereAmIDataOutput struct {
	File       string
	Function   string
	LineNumber int
}

type WhereAmIDataInput struct {
	DepthList          []int
	FullFunctionOutput bool
	FullFilePath       bool
}

// return a string containing the file name, function name
// and the line number of a specified entry on the call stack
func WhereAmI(depthList ...int) string {
	data := WhereAmIData(WhereAmIDataInput{
		DepthList:          depthList,
		FullFilePath:       false,
		FullFunctionOutput: true,
	})
	return fmt.Sprintf("File: %s  Function: %s Line: %d", data.File, data.Function, data.LineNumber)
}

// return  file name, function name and the line number
// of a specified entry on the call stack
func WhereAmIData(in WhereAmIDataInput) WhereAmIDataOutput {
	var depth int
	if in.DepthList == nil {
		depth = 1
	} else {
		depth = in.DepthList[0]
	}
	function, file, line, _ := runtime.Caller(depth)

	out := WhereAmIDataOutput{}

	if in.FullFunctionOutput {
		out.Function = runtime.FuncForPC(function).Name()
	} else {
		out.Function = chopPath(runtime.FuncForPC(function).Name())
	}

	if in.FullFilePath {
		out.File = file
	} else {
		out.File = chopPath(file)
	}

	out.LineNumber = line

	return out
}

// return the source filename after the last slash
func chopPath(original string) string {
	i := strings.LastIndex(original, "/")
	if i == -1 {
		return original
	} else {
		return original[i+1:]
	}
}
