// +build !prod

package uassert

import "fmt"

func True(cond bool, a ...interface{}) {
	inAssertionf(cond, fmt.Sprint(a...))
}

func False(cond bool, a ...interface{}) {
	inAssertionf(!cond, fmt.Sprint(a...))
}

func inAssertionf(cond bool, format string, a ...interface{}) {
	if cond {
		return;
	}
	fmt.Println("----- Assertion Failed -----");
	if len(a) == 0 {
		panic(format);
	} else {
		panic(fmt.Sprintf(format, a...))
	}
}

func Truef(cond bool, format string, a ...interface{}) {
	inAssertionf(cond, format, a...)
}

func Falsef(cond bool, format string, a ...interface{}) {
	inAssertionf(!cond, format, a...)
}
