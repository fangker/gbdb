// +build prod

package uassert

func True(cond bool, a ...interface{}) {}

func False(cond bool, a ...interface{}) {}

func Truef(cond bool, format string, a ...interface{}) {}

func Falsef(cond bool, format string, a ...interface{}) {}