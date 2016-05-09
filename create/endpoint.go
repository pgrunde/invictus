package create

import (
	"fmt"
)

func NewEndpoint(name, endpointFolder string) error {
	fmt.Println(name, endpointFolder)
	return nil
}
