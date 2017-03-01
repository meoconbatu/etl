package etl

import (
	"strings"
)

const testVersion = 1

func Transform(input map[int][]string) (output map[string]int) {
	output = make(map[string]int,26)
	for iKey, iValues := range input {
		for i := 0; i < len(iValues); i++ {
			oKey := strings.ToLower(iValues[i])
			output[oKey] = iKey
		}
	}
	return output
}
