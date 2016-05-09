package create

import (
	"fmt"
)

func NewEndpoint(name, endpointFolder string) error {
	if endpointFolder == "" {
		endpointFolder = "v1"
	}
	fmt.Println(name, endpointFolder)
	return nil
}
