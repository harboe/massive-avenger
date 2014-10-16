package model

import (
	"fmt"
	"testing"
)

var addressTestCases = map[string]Address{
	"Dennis Harboe\nÅlekistevej 203\n2720 Vanløse\nDenmark": Address{
		Name:    "Dennis Harboe",
		Street:  "Ålekistevej 203",
		Zip:     2720,
		City:    "Vanløse",
		Country: "Denmark",
	},
	"Dennis Harboe\nÅlekistevej 203\n2720 Vanløse\nCopenhagen\nDenmark": Address{
		Name:    "Dennis Harboe",
		Street:  "Ålekistevej 203",
		Zip:     2720,
		City:    "Vanløse",
		Country: "Denmark",
		State:   "Copenhagen",
	},
}

func Test_addrees_string(t *testing.T) {

	for expected, obj := range addressTestCases {

		if expected != fmt.Sprintf("%s", obj) {
			t.Errorf("Expected: \"%s\" got \"%s\".", expected, obj)
		}
	}

}
