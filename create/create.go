package create

import (
	"fmt"
	"strings"
)

var badWindowsNames = []string{"com1", "com2", "com3", "com4", "com5", "com6", "com7", "com8", "com9", "lpt1", "lpt2", "lpt3", "lpt4", "lpt5", "lpt6", "lpt7", "lpt8", "lpt9", "con", "nul", "prn"}

var badCharacters = []rune{'^', '/', '?', '<', '>', '\\', ':', '*', '|', '"', '.'}

func Project(s string) error {
	s = strings.ToLower(s)
	if err := hasIllegalFilename(s); err != nil {
		return fmt.Errorf("Given project name %s has an illegal filename: %s", s, err)
	}
	return nil
}

func hasIllegalFilename(s string) error {
	if s == "" {
		return fmt.Errorf("No project name given. Please specify a project name")
	}
	if len(s) > 255 {
		return fmt.Errorf("Project name cannot exceed 255 characters")
	}
	for _, bwn := range badWindowsNames {
		if s == bwn {
			return fmt.Errorf("Project name contains a reserved Windows operating system file name")
		}
	}
	for _, r := range badCharacters {
		if strings.ContainsRune(s, r) {
			return fmt.Errorf("Found illegal character %s in filename", string(r))
		}
	}
	return nil
}
