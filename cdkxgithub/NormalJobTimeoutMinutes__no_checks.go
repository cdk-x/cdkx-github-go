//go:build no_runtime_type_checking

package cdkxgithub

// Building without runtime type checking enabled, so all the below just return nil

func validateNormalJobTimeoutMinutes_FromNumberParameters(value *float64) error {
	return nil
}

func validateNormalJobTimeoutMinutes_FromStringParameters(value *string) error {
	return nil
}

