package main

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/joomcode/errorx"
	utilities "github.com/pericles-tpt/utilities"
)

const (
	minPkgIdentLength = 1
	maxPkgIdentLength = 3
	minFncIdentLength = 1
	maxFncIdentLength = 4

	includePackageIdent = true
	includeFuncIdent    = false
)

// PrependErrorWithRuntimeInfo, prepends to an error, a shorthand descriptor for the name of the package
//
// and function where the error occurred. To make it easier to identify the error location
func PrependErrorWithRuntimeInfo(err error, msg string, args ...interface{}) error {
	pc, _, _, ok := runtime.Caller(1)
	if ok {
		fn := runtime.FuncForPC(pc)
		if fn != nil {
			runtimeIdent := getRuntimeIdent(fn)

			if err == nil && msg == "" {
				return fmt.Errorf("%sA CALL TO `PrependErrorWithRuntimeInfo` IS INVALID, `err` == nil && msg == \"\"", runtimeIdent)
			}

			if msg != "" {
				return errorx.Decorate(err, "%s%s", runtimeIdent, fmt.Sprintf(msg, args...))
			}
			return errorx.Decorate(err, "%s", runtimeIdent)
		}
	}
	return errorx.Decorate(err, msg, args...)
}

// getRuntimeIdent, constructs a runtime identifier based on local flags
func getRuntimeIdent(fn *runtime.Func) string {
	var (
		ident       string
		identParts  []string
		fnNameParts = strings.Split(fn.Name(), ".")
	)
	if includePackageIdent {
		identParts = append(identParts, shortenPackageName(fnNameParts[0]))
	}
	if includeFuncIdent {
		identParts = append(identParts, shortenFunctionName(fnNameParts[1]))
	}
	if len(identParts) > 0 {
		ident = fmt.Sprintf("[%s] ", strings.Join(identParts, "_"))
	}
	return ident
}

// shortenPackageName, generates a shorthand string
//
// identifier for a package
//
// e.g. main -> mn, handlers -> hnd
func shortenPackageName(s string) string {
	var (
		ret           = s
		withoutVowels = utilities.RemoveCharsIn(s, "aeiou")
	)
	if (len(withoutVowels) + 1) > minPkgIdentLength {
		ret = string(ret[0]) + withoutVowels
	}

	if len(ret) > maxPkgIdentLength {
		ret = ret[:maxPkgIdentLength]
	}
	return strings.ToLower(ret)
}

// shortenFunctionName, generates a short string
//
// identifier for a function
//
// e.g. shortenFunctionName -> sfn, Open -> open
func shortenFunctionName(s string) string {
	var (
		ret              = s
		upperCaseLetters = utilities.GetUpperChars(s[1:])
	)
	if (len(upperCaseLetters) + 1) > minFncIdentLength {
		ret = string(ret[0]) + upperCaseLetters
	}

	if len(ret) > maxFncIdentLength {
		ret = ret[:maxFncIdentLength]
	}
	return strings.ToLower(ret)
}
