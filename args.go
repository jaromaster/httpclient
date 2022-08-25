package main

import (
	"strings"
)

// take list of args and return map containing flags
func CreateFlagsFromArgs(args []string) map[string]string {
	flagsMap := make(map[string]string)

	for _, arg := range args {
		if !strings.HasPrefix(arg, "--") {
			continue
		}

		key_val := strings.SplitN(arg, "=", 2)
		if len(key_val) == 2 {
			flagsMap[key_val[0]] = key_val[1]
		} else {
			flagsMap[key_val[0]] = key_val[0]
		}
	}

	return flagsMap
}
