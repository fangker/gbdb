package utils

import (
	"strings"
	"os"
)

var ENV_DIR = strings.Split(os.Environ()[13], "=")[1]
