//go:build !no_runtime_type_checking

package cdkxgithub

import (
	"fmt"
)

func validateShell_FromStringParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

