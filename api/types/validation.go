package types

import (
	"fmt"
	"strings"
)

type ValidatedRequest interface {
	Validate() error
}

func ErrorMsg(errorList []string) error {
	if len(errorList) == 0 {
		return nil
	}
	return fmt.Errorf("validation failed: " + strings.Join(errorList, ","))
}
