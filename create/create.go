package create

import (
	"fmt"
	"strings"
)

var badWindowsNames = []string{"com1", "com2", "com3", "com4", "com5", "com6", "com7", "com8", "com9", "lpt1", "lpt2", "lpt3", "lpt4", "lpt5", "lpt6", "lpt7", "lpt8", "lpt9", "con", "nul", "prn"}

var badCharacters = []rune{'/', '?', '<', '>', '\\', ':', '*', '|', '"', '.'}

func Project(s string) error {
	s = strings.ToLower(s)
	if hasIllegalFilename(s) {
		return fmt.Errorf("Given project name %s has an illegal filename", s)
	}
	return nil
}

// / ? < > \ : * | " . com1, com2, com3, com4, com5, com6, com7, com8, com9, lpt1, lpt2, lpt3, lpt4, lpt5, lpt6, lpt7, lpt8, lpt9, con, nul, and prn
// must be <= 255 characters long
func hasIllegalFilename(s string) bool {
	if len(s) > 255 {
		return true
	}
	for _, bwn := range badWindowsNames {
		if s == bwn {
			return true
		}
	}
	for _, r := range badCharacters {
		if strings.ContainsRune(s, r) {
			return true
		}
	}
	return false
}
