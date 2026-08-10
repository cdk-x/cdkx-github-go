//go:build no_runtime_type_checking

package cdkxgithub

// Building without runtime type checking enabled, so all the below just return nil

func validateServiceContainerPorts_FromNumberParameters(value *float64) error {
	return nil
}

func validateServiceContainerPorts_FromStringParameters(value *string) error {
	return nil
}

