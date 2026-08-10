//go:build no_runtime_type_checking

package cdkxgithub

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GhbWorkflow) validateAddDependencyParameters(target cdkxcore.Resource) error {
	return nil
}

func (g *jsiiProxy_GhbWorkflow) validateGetAttParameters(attr *string) error {
	return nil
}

func validateGhbWorkflow_IsConstructParameters(x interface{}) error {
	return nil
}

func validateGhbWorkflow_IsResourceParameters(x interface{}) error {
	return nil
}

func validateGhbWorkflow_OfParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewGhbWorkflowParameters(scope constructs.Construct, id *string, props *WorkflowProps) error {
	return nil
}

