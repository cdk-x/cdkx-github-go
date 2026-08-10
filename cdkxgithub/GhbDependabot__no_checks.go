//go:build no_runtime_type_checking

package cdkxgithub

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GhbDependabot) validateAddDependencyParameters(target cdkxcore.Resource) error {
	return nil
}

func (g *jsiiProxy_GhbDependabot) validateGetAttParameters(attr *string) error {
	return nil
}

func validateGhbDependabot_IsConstructParameters(x interface{}) error {
	return nil
}

func validateGhbDependabot_IsResourceParameters(x interface{}) error {
	return nil
}

func validateGhbDependabot_OfParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewGhbDependabotParameters(scope constructs.Construct, id *string, props *DependabotProps) error {
	return nil
}

