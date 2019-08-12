package ulog

import (
	"os"
	"strings"
	"runtime"
	"path"
)

var values = os.Getenv("DEBUG")

func checkScopeEnable() bool {
	if values == "" || values == "*" {
		return true
	}
	tags := strings.Split(values, " ")
	_, fileName, _, _ := runtime.Caller(1)
	pkgDirStrings := strings.Split(path.Dir(fileName), "/")
	pkgDirName := pkgDirStrings[len(pkgDirStrings)-1]
	for _, t := range tags {
		if t == pkgDirName {
			return true;
		}
	}
	return false
}
